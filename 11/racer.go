package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

// Racer return the fastest url with a default timeout
func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

// ConfigurableRacer return the fastest url given a timeout
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	// select waits on multiple channel
	// will be executed the case with the first channel to send a value
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout): // `time.After` returns a `chan`
		return "", fmt.Errorf("time out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	// using `chan struct{}` instead of `chan` bool to use less memory
	// `chan struct{}` is the smallest data type
	// not using `var ch chan struct{}` because it would be `nil`
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func oldRacer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
