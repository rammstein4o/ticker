ticker [![GoDoc](https://godoc.org/github.com/rammstein4o/ticker?status.svg)](https://godoc.org/github.com/rammstein4o/ticker)
=====

This package implements simple interface that uses `time.Ticker` underneath, but unlike `time.Ticker` it can be paused/resumed.


## Sample Use

```Go
package main

import (
	"log"
	"time"

	tk "github.com/rammstein4o/ticker"
)

func main() {
	ticker := tk.NewDefaultTicker()
	ch := ticker.Start()

	go func() {
		for t := range ch {
			log.Println(t)
		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Pause()

	time.Sleep(5 * time.Second)
	ticker.Resume()

	time.Sleep(10 * time.Second)
	ticker.Stop()
}
```
 
## License
 
The MIT License
