package osuapi

import (
	"time"
)

var every time.Duration
var requestsAvailable chan struct{}
var routStarted bool

// RateLimit allows you to set the maximum number of requests to do in a
// minute to the osu! API.
//
// Please note that this function is NOT thread safe. It should be executed
// only at the start of your program, and never after it.
//
// The reason for this is that creating a Mutex for a channel is just
// absolutely ridiculous.
func RateLimit(maxRequests int) {
	if maxRequests == 0 {
		requestsAvailable = nil
	}
	every = 60000 * time.Millisecond / time.Duration(maxRequests)
	requestsAvailable = make(chan struct{}, maxRequests)
	for {
		var b bool
		select {
		case requestsAvailable <- struct{}{}:
			// nothing, just keep on moving
		default:
			b = true
		}
		if b {
			break
		}
	}
	if !routStarted {
		go requestIncreaser()
	}
	routStarted = true
}
func requestIncreaser() {
	for {
		time.Sleep(every)
		requestsAvailable <- struct{}{}
	}
}
