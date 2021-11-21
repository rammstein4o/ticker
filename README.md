ticker [![GoDoc](https://godoc.org/github.com/rammstein4o/ticker?status.svg)](https://godoc.org/github.com/rammstein4o/ticker)
=====

This package implements simple interface that uses `time.Ticker` underneath, but unlike `time.Ticker` it can be paused/resumed.


## Sample Use

```Go
package main

import (
    tk "github.com/rammstein4o/ticker"
    "log"
)

func main() {
	ticker := tk.NewDefaultTicker()
	ch := ticker.Start()

	go func() {
		for t := range ch {
			log.Println(t)
		}
	}()

	time.Sleep(20 * time.Second)
	ticker.Stop()
}
```
 
## License
 
The MIT License (MIT)
