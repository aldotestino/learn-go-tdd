package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {

	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 100
		counter := NewCounter()

		// syncronise concurrent processes
		// WaitGroup waits for a collection of goroutines to finish
		var wg sync.WaitGroup
		wg.Add(wantedCount) // sets the number of goroutines to wait for

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done() // each goroutines call done when finis∂hed
			}(&wg)
		}

		wg.Wait() // blocks until all the goroutines have finished

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d want %d", got.Value(), want)
	}
}
