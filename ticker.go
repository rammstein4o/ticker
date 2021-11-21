// This package implements simple interface that uses time.Ticker underneath,
// but unlike time.Ticker it can be paused/resumed.
package ticker

import (
	"errors"
	"time"
)

type Interface interface {
	// Start starts the internal time.Ticker and returns output channel
	Start() <-chan time.Time
	// Pause pauses sending of ticks on the output channel
	Pause()
	// Resume resumes sending of ticks on the output channel
	Resume()
	// Stop turns off the internal ticker and closes the output channel.
	// After Stop, no more ticks will be sent.
	Stop()
	// IsRunning returns true if the ticker is running
	IsRunning() bool
	// IsPaused returns true if the ticker is paused
	IsPaused() bool
}

type ticker struct {
	done    chan bool
	out     chan time.Time
	running bool
	paused  bool
	ticker  *time.Ticker
}

func (t *ticker) Start() <-chan time.Time {
	if !t.running {
		go func() {
			for {
				select {
				case <-t.done:
					close(t.out)
					close(t.done)
					return
				case res := <-t.ticker.C:
					if !t.paused {
						t.out <- res
					}
				}
			}
		}()
		t.running = true
	}

	t.Resume()
	return t.out
}

func (t *ticker) Pause() {
	t.paused = true
}

func (t *ticker) Resume() {
	t.paused = false
}

func (t *ticker) Stop() {
	t.ticker.Stop()
	t.done <- true
	t.running = false
}

func (t *ticker) IsRunning() bool {
	return t.running
}

func (t *ticker) IsPaused() bool {
	return t.paused
}

// NewTicker returns a new ticker instance containing a channel that will send
// the time on the channel after each tick. The period of the ticks is
// specified by the duration argument.
// The duration dur must be greater than zero; if not, NewTicker will
// panic. Stop the ticker to release associated resources.
func NewTicker(dur time.Duration) Interface {
	if dur <= 0 {
		panic(errors.New("non-positive interval for NewTicker"))
	}

	return &ticker{
		done:    make(chan bool, 1),
		out:     make(chan time.Time, 1),
		running: false,
		paused:  false,
		ticker:  time.NewTicker(dur),
	}
}

// NewDefaultTicker is the same as NewTicker, but with pre-defined ticks period of one second.
func NewDefaultTicker() Interface {
	return NewTicker(1 * time.Second)
}
