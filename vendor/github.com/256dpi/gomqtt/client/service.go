package client

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/256dpi/gomqtt/client/future"
	"github.com/256dpi/gomqtt/packet"
	"github.com/256dpi/gomqtt/session"
	"github.com/256dpi/gomqtt/topic"

	"github.com/jpillora/backoff"
	"gopkg.in/tomb.v2"
)

type command struct {
	publish       bool
	subscribe     bool
	unsubscribe   bool
	future        *future.Future
	message       *packet.Message
	subscriptions []packet.Subscription
	topics        []string
}

// Service is an abstraction for Client that provides a stable interface to the
// application, while it automatically connects and reconnects clients in the
// background. Errors are not returned but emitted using the ErrorCallback.
// All methods return Futures that get completed once the acknowledgements are
// received. Once the services is stopped all waiting futures get canceled.
//
// Note: If clean session is false and there are packets in the store, messages
// might get completed after starting without triggering any futures to complete.
type Service struct {
	// The session used by the client to store unacknowledged packets.
	Session Session

	// The OnlineCallback is called when the service is connected.
	//
	// Note: Execution of the service is resumed after the callback returns.
	// This means that waiting on a future inside the callback will deadlock the
	// service.
	OnlineCallback func(resumed bool)

	// The MessageCallback is called when a message is received. If an error is
	// returned the underlying client will be prevented from acknowledging the
	// specified message and closed immediately. The errors is logged and a
	// reconnect attempt initiated.
	//
	// Note: Execution of the service is resumed after the callback returns.
	// This means that waiting on a future inside the callback will deadlock the
	// service.
	MessageCallback func(*packet.Message) error

	// The ErrorCallback is called when an error occurred.
	//
	// Note: Execution of the service is resumed after the callback returns.
	// This means that waiting on a future inside the callback will deadlock the
	// service.
	ErrorCallback func(error)

	// The OfflineCallback is called when the service is disconnected.
	//
	// Note: Execution of the service is resumed after the callback returns.
	// This means that waiting on a future inside the callback will deadlock the
	// service.
	OfflineCallback func()

	// The logger that is used to log write low level information like packets
	// that have ben successfully sent and received, details about the automatic
	// keep alive handler, reconnection and occurring errors.
	Logger func(msg string)

	// The minimum delay between reconnects.
	//
	// Note: The value must be changed before calling Start.
	MinReconnectDelay time.Duration

	// The maximum delay between reconnects.
	//
	// Note: The value must be changed before calling Start.
	MaxReconnectDelay time.Duration

	// The allowed timeout until a connection attempt is canceled.
	ConnectTimeout time.Duration

	// The allowed timeout until a connection is forcefully closed.
	DisconnectTimeout time.Duration

	// The allowed timeout until a subscribe action is forcefully closed during
	// reconnect.
	ResubscribeTimeout time.Duration

	// Whether to resubscribe all subscriptions after reconnecting. Can be
	// disabled if the broker supports persistent sessions and the client is
	// configured to request one.
	ResubscribeAllSubscriptions bool

	config        *Config
	started       bool
	backoff       *backoff.Backoff
	subscriptions *topic.Tree
	commandQueue  chan *command
	futureStore   *future.Store
	mutex         sync.Mutex
	tomb          *tomb.Tomb
}

// NewService allocates and returns a new service. The optional parameter queueSize
// specifies how many Subscribe, Unsubscribe and Publish commands can be queued
// up before actually sending them on the wire. The default queueSize is 100.
func NewService(queueSize ...int) *Service {
	var qs = 100
	if len(queueSize) > 0 {
		qs = queueSize[0]
	}

	return &Service{
		Session:                     session.NewMemorySession(),
		MinReconnectDelay:           50 * time.Millisecond,
		MaxReconnectDelay:           10 * time.Second,
		ConnectTimeout:              5 * time.Second,
		DisconnectTimeout:           10 * time.Second,
		ResubscribeTimeout:          5 * time.Second,
		ResubscribeAllSubscriptions: true,
		subscriptions:               topic.NewStandardTree(),
		commandQueue:                make(chan *command, qs),
		futureStore:                 future.NewStore(),
	}
}

// Start will start the service with the specified configuration. From now on
// the service will automatically reconnect on any error until Stop is called.
// It returns false if the service was already started.
func (s *Service) Start(config *Config) bool {
	// check config
	if config == nil {
		panic("missing config")
	}

	// acquire mutex
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// return if already started
	if s.started {
		return false
	}

	// set state
	s.started = true

	// save config
	s.config = config

	// initialize backoff
	s.backoff = &backoff.Backoff{
		Min:    s.MinReconnectDelay,
		Max:    s.MaxReconnectDelay,
		Factor: 2,
	}

	// mark future store as protected
	s.futureStore.Protect(true)

	// create new tomb
	s.tomb = new(tomb.Tomb)

	// start supervisor
	s.tomb.Go(s.supervisor)

	return true
}

// Publish will send a Publish packet containing the passed parameters. It will
// return a PublishFuture that gets completed once the quality of service flow
// has been completed.
func (s *Service) Publish(topic string, payload []byte, qos packet.QOS, retain bool) GenericFuture {
	return s.PublishMessage(&packet.Message{
		Topic:   topic,
		Payload: payload,
		QOS:     qos,
		Retain:  retain,
	})
}

// PublishMessage will send a Publish packet containing the passed message. It will
// return a PublishFuture that gets completed once the quality of service flow
// has been completed.
func (s *Service) PublishMessage(msg *packet.Message) GenericFuture {
	// acquire mutex
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// allocate future
	f := future.New()

	// queue publish
	s.commandQueue <- &command{
		publish: true,
		future:  f,
		message: msg,
	}

	return f
}

// Subscribe will send a Subscribe packet containing one topic to subscribe. It
// will return a SubscribeFuture that gets completed once the acknowledgements
// have been received.
func (s *Service) Subscribe(topic string, qos packet.QOS) SubscribeFuture {
	return s.SubscribeMultiple([]packet.Subscription{
		{Topic: topic, QOS: qos},
	})
}

// SubscribeMultiple will send a Subscribe packet containing multiple topics to
// subscribe. It will return a SubscribeFuture that gets completed once the
// acknowledgements have been received.
func (s *Service) SubscribeMultiple(subscriptions []packet.Subscription) SubscribeFuture {
	// acquire mutex
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// save subscription
	for _, v := range subscriptions {
		s.subscriptions.Set(v.Topic, v)
	}

	// allocate future
	f := future.New()

	// queue subscribe
	s.commandQueue <- &command{
		subscribe:     true,
		future:        f,
		subscriptions: subscriptions,
	}

	return &subscribeFuture{f}
}

// Unsubscribe will send a Unsubscribe packet containing one topic to unsubscribe.
// It will return a SubscribeFuture that gets completed once the acknowledgements
// have been received.
func (s *Service) Unsubscribe(topic string) GenericFuture {
	return s.UnsubscribeMultiple([]string{topic})
}

// UnsubscribeMultiple will send a Unsubscribe packet containing multiple
// topics to unsubscribe. It will return a SubscribeFuture that gets completed
// once the acknowledgements have been received.
func (s *Service) UnsubscribeMultiple(topics []string) GenericFuture {
	// acquire mutex
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// remove subscription
	for _, v := range topics {
		s.subscriptions.Empty(v)
	}

	// allocate future
	f := future.New()

	// queue unsubscribe
	s.commandQueue <- &command{
		unsubscribe: true,
		future:      f,
		topics:      topics,
	}

	return f
}

// Stop will disconnect the client if online and cancel all futures if requested.
// After the service is stopped in can be started again. It returns false if the
// service was not running.
//
// Note: You should clear the futures on the last stop before exiting to ensure
// that all goroutines return that wait on futures.
func (s *Service) Stop(clearFutures bool) bool {
	// acquire mutex
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// return if service not started
	if !s.started {
		return false
	}

	// set state
	s.started = false

	// kill and wait
	s.tomb.Kill(nil)
	_ = s.tomb.Wait()

	// clear futures if requested
	if clearFutures {
		s.futureStore.Protect(false)
		s.futureStore.Clear()
	}

	return true
}

// the supervised reconnect loop
func (s *Service) supervisor() error {
	// prepare flag
	first := true

	for {
		// delay if not first
		if !first {
			// get backoff duration
			d := s.backoff.Duration()
			s.log(fmt.Sprintf("Delay Reconnect: %v", d))

			// sleep but return on Stop
			select {
			case <-time.After(d):
			case <-s.tomb.Dying():
				return tomb.ErrDying
			}
		}

		s.log("Next Reconnect")

		// clear flag
		first = false

		// prepare the kill channel
		kill := make(chan struct{})

		// try once to get a client
		client, resumed := s.connect(kill)
		if client == nil {
			continue
		}

		// resubscribe
		if s.ResubscribeAllSubscriptions {
			if !s.resubscribe(client) {
				_ = client.Close()
				continue
			}
		}

		// run callback
		if s.OnlineCallback != nil {
			s.OnlineCallback(resumed)
		}

		// run dispatcher on client
		dying := s.dispatcher(client, kill)

		// ensure client is closed
		_ = client.Close()

		// run callback
		if s.OfflineCallback != nil {
			s.OfflineCallback()
		}

		// return goroutine if dying
		if dying {
			return tomb.ErrDying
		}
	}
}

// will try to connect one client to the broker
func (s *Service) connect(kill chan struct{}) (*Client, bool) {
	// prepare new client
	client := New()
	client.Session = s.Session
	client.Logger = s.Logger
	client.futureStore = s.futureStore

	// set callback
	client.Callback = func(msg *packet.Message, err error) error {
		if err != nil {
			s.err("Callback", err)
			close(kill)
			return nil
		}

		// call the handler
		if s.MessageCallback != nil {
			return s.MessageCallback(msg)
		}

		return nil
	}

	// attempt to connect
	connectFuture, err := client.Connect(s.config)
	if err != nil {
		_ = client.Close()
		s.err("Connect", err)
		return nil, false
	}

	// await future
	err = connectFuture.Wait(s.ConnectTimeout)
	if err != nil {
		_ = client.Close()
		s.err("Connect", err)
		return nil, false
	}

	return client, connectFuture.SessionPresent()
}

func (s *Service) resubscribe(client *Client) bool {
	// get all subscriptions and return if empty
	items := s.subscriptions.All()
	if len(items) == 0 {
		return true
	}

	// prepare subscriptions
	subs := make([]packet.Subscription, 0, len(items))
	for _, v := range items {
		subs = append(subs, v.(packet.Subscription))
	}

	// sort subscriptions
	sort.Slice(subs, func(i, j int) bool {
		return subs[i].Topic < subs[j].Topic
	})

	// resubscribe all subscriptions
	subscribeFuture, err := client.SubscribeMultiple(subs)
	if err != nil {
		s.err("Resubscribe", err)
		return false
	}

	// await future
	err = subscribeFuture.Wait(s.ResubscribeTimeout)
	if err != nil {
		s.err("Resubscribe", err)
		return false
	}

	return true
}

// reads from the queues and calls the current client
func (s *Service) dispatcher(client *Client, kill chan struct{}) bool {
	for {
		select {
		case cmd := <-s.commandQueue:
			// handle subscribe command
			if cmd.subscribe {
				// perform subscribe
				f2, err := client.SubscribeMultiple(cmd.subscriptions)
				if err != nil {
					s.err("Subscribe", err)
					cmd.future.Cancel(nil)
					return false
				}

				// attach future
				f2.(*subscribeFuture).Future.Attach(cmd.future)
			}

			// handle unsubscribe command
			if cmd.unsubscribe {
				// perform unsubscribe
				f2, err := client.UnsubscribeMultiple(cmd.topics)
				if err != nil {
					s.err("Unsubscribe", err)
					cmd.future.Cancel(nil)
					return false
				}

				// attach future
				f2.(*future.Future).Attach(cmd.future)
			}

			// handle publish command
			if cmd.publish {
				// perform publish
				f2, err := client.PublishMessage(cmd.message)
				if err != nil {
					s.err("Publish", err)
					cmd.future.Cancel(nil)
					return false
				}

				// attach future
				f2.(*future.Future).Attach(cmd.future)
			}
		case <-s.tomb.Dying():
			// disconnect client on Stop
			err := client.Disconnect(s.DisconnectTimeout)
			if err != nil {
				s.err("Disconnect", err)
			}

			return true
		case <-kill:
			return false
		}
	}
}

func (s *Service) err(sys string, err error) {
	s.log(fmt.Sprintf("%s Error: %s", sys, err.Error()))

	if s.ErrorCallback != nil {
		s.ErrorCallback(err)
	}
}

func (s *Service) log(str string) {
	if s.Logger != nil {
		s.Logger(str)
	}
}
