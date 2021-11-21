package timer

import (
	"errors"
	"time"
)

type Interface interface {
	Start() <-chan time.Time
	Stop()
	Pause()
	Resume()
}

type timer struct {
	ticker  *time.Ticker
	done    chan bool
	out     chan time.Time
	started bool
	paused  bool
	dur     time.Duration
}

func (t *timer) Start() <-chan time.Time {
	if t.paused {
		t.Resume()
	}

	if !t.started {
		t.ticker = time.NewTicker(t.dur)
		t.done = make(chan bool, 1)

		go func() {
			for {
				select {
				case <-t.done:
					close(t.out)
					return
				case res := <-t.ticker.C:
					if !t.paused {
						t.out <- res
					}
				}
			}
		}()

		t.started = true
		t.paused = false
	}

	return t.out
}

func (t *timer) Stop() {
	t.ticker.Stop()
	t.done <- true
	t.started = false
	t.paused = false
}

func (t *timer) Pause() {
	t.paused = true
}

func (t *timer) Resume() {
	t.paused = false
}

func NewTimer(dur time.Duration) Interface {
	if dur <= 0 {
		panic(errors.New("non-positive interval for NewTimer"))
	}

	return &timer{
		dur: dur,
		out: make(chan time.Time, 1),
	}
}

func NewDefaultTimer() Interface {
	return NewTimer(1 * time.Second)
}
