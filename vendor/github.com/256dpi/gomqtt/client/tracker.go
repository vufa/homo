package client

import (
	"sync"
	"time"
)

// A Tracker keeps track of keep alive intervals.
type Tracker struct {
	last    time.Time
	pings   uint8
	timeout time.Duration
	mutex   sync.RWMutex
}

// NewTracker returns a new tracker.
func NewTracker(timeout time.Duration) *Tracker {
	return &Tracker{
		last:    time.Now(),
		timeout: timeout,
	}
}

// Reset will reset the tracker.
func (t *Tracker) Reset() {
	// acquire mutex
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// reset timestamp
	t.last = time.Now()
}

// Window returns the time until a new ping should be sent.
func (t *Tracker) Window() time.Duration {
	// acquire mutex
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	return t.timeout - time.Since(t.last)
}

// Ping marks a ping.
func (t *Tracker) Ping() {
	// acquire mutex
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// increment
	t.pings++
}

// Pong marks a pong.
func (t *Tracker) Pong() {
	// acquire mutex
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// decrement
	t.pings--
}

// Pending returns if pings are pending.
func (t *Tracker) Pending() bool {
	// acquire mutex
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	return t.pings > 0
}
