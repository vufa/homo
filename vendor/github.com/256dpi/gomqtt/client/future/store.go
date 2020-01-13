package future

import (
	"sync"
	"time"

	"github.com/256dpi/gomqtt/packet"
)

// A Store is used to store futures.
type Store struct {
	protected bool
	store     map[packet.ID]*Future
	mutex     sync.RWMutex
}

// NewStore will create a new Store.
func NewStore() *Store {
	return &Store{
		store: make(map[packet.ID]*Future),
	}
}

// Put will save a future to the store.
func (s *Store) Put(id packet.ID, future *Future) {
	// acquire mutex
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// set future
	s.store[id] = future
}

// Get will retrieve a future from the store.
func (s *Store) Get(id packet.ID) *Future {
	// acquire mutex
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.store[id]
}

// Delete will remove a future from the store.
func (s *Store) Delete(id packet.ID) {
	// acquire mutex
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// delete future
	delete(s.store, id)
}

// All will return a slice with all stored futures.
func (s *Store) All() []*Future {
	// acquire mutex
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// collect futures
	all := make([]*Future, 0, len(s.store))
	for _, savedFuture := range s.store {
		all = append(all, savedFuture)
	}

	return all
}

// Protect will set the protection attribute and if true prevents the store from
// being cleared.
func (s *Store) Protect(value bool) {
	// acquire mutex
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// set flag
	s.protected = value
}

// Clear will cancel all stored futures and remove them if the store is unprotected.
func (s *Store) Clear() {
	// acquire mutex
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// check flag
	if s.protected {
		return
	}

	// cancel all futures
	for _, savedFuture := range s.store {
		savedFuture.Cancel(nil)
	}

	// reset store
	s.store = make(map[packet.ID]*Future)
}

// Await will wait until all futures have completed or cancelled, or the timeout
// has been reached.
func (s *Store) Await(timeout time.Duration) error {
	// prepare deadline
	deadline := time.Now().Add(timeout)

	for {
		// get futures
		s.mutex.RLock()
		futures := s.All()
		s.mutex.RUnlock()

		// return if no futures are left
		if len(futures) == 0 {
			return nil
		}

		// wait for next future to complete
		err := futures[0].Wait(deadline.Sub(time.Now()))
		if err != nil {
			return err
		}
	}
}
