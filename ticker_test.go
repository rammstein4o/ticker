package ticker

import (
	"log"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Example() {
	ticker := NewDefaultTicker()
	ch := ticker.Start()

	go func() {
		for t := range ch {
			log.Println(t)
		}
	}()

	time.Sleep(20 * time.Second)
	ticker.Stop()
}

func TestTimer(t *testing.T) {
	assert := assert.New(t)

	var wg sync.WaitGroup
	wg.Add(1)

	ticker := NewDefaultTicker()
	ch := ticker.Start()

	count := 0
	go func() {
		for range ch {
			count++
		}
		wg.Done()
	}()

	time.Sleep(5 * time.Second)
	ticker.Stop()

	wg.Wait()

	assert.Equal(5, count)
}

func TestTimer_PauseAndResume(t *testing.T) {
	assert := assert.New(t)

	var wg sync.WaitGroup
	wg.Add(1)

	ticker := NewDefaultTicker()
	ch := ticker.Start()

	count := 0
	go func() {
		for range ch {
			count++
		}
		wg.Done()
	}()

	time.Sleep(2 * time.Second)
	ticker.Pause()

	assert.True(ticker.IsRunning())
	assert.True(ticker.IsPaused())

	time.Sleep(5 * time.Second)
	ticker.Resume()

	assert.True(ticker.IsRunning())
	assert.False(ticker.IsPaused())

	time.Sleep(2 * time.Second)
	ticker.Stop()

	assert.False(ticker.IsRunning())
	assert.False(ticker.IsPaused())

	wg.Wait()
}
