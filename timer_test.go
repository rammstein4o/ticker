package timer

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_timer_Start(t *testing.T) {
	assert := assert.New(t)

	var wg sync.WaitGroup
	wg.Add(1)

	timer := NewDefaultTimer()
	c := timer.Start()

	count := 0
	go func() {
		for range c {
			count++
		}
		wg.Done()
	}()

	time.Sleep(5 * time.Second)
	timer.Stop()

	wg.Wait()

	assert.Equal(5, count)
}

func Test_timer_Pause(t *testing.T) {
	assert := assert.New(t)

	var wg sync.WaitGroup
	wg.Add(1)

	timer := NewDefaultTimer()
	c := timer.Start()

	count := 0
	go func() {
		for range c {
			count++
		}
		wg.Done()
	}()

	time.Sleep(2 * time.Second)
	timer.Pause()

	time.Sleep(10 * time.Second)
	timer.Start()

	time.Sleep(2 * time.Second)
	timer.Stop()

	wg.Wait()

	// We check for number between 2 and 5 as Pause and Start calls sometimes
	// match with the ticker tick's which prevents us from getting the result
	assert.GreaterOrEqual(count, 2)
	assert.LessOrEqual(count, 5)
}
