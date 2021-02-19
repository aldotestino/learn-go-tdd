package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Sleeper is a dependency (using a real Sleeper in main and a spy sleeper in tests)
type Sleeper interface {
	Sleep()
}

// SpySleeper implements the Sleeper interface (dependency)
// this is the fake sleeper
type SpySleeper struct {
	Calls int
}

// Sleep the implementation of the Sleeper interface for the fake sleeper
func (s *SpySleeper) Sleep() {
	s.Calls++
}

// DefaultSleeper implemements the Sleeper interface (dependency)
// this is the real sleeper
type DefaultSleeper struct{}

// Sleep the implementation of the Sleeper interface for the real sleeper
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const write = "write"
const sleep = "sleep"

// CountdownOperationSpy controls the call stack
// implements both io.Writer and Sleeper
type CountdownOperationSpy struct {
	Calls []string
}

// Sleep appends a sleep operation to the call stack
// implements the Sleep method of the Sleeper interface
func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// Write appends a write operation to the call stack
// implements the Write method of the io.Writer interface
func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

// ConfigurableSleeper accepts what we need for configuration and testing
// implements the Sleeper interface
type ConfigurableSleeper struct {
	duration time.Duration
	// this is a function we have to pass down when creating an instance
	sleep func(time.Duration)
}

// Sleep make the call to the sleep function passed down
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// SpyTime helps us to cotrol the effective time passed
// implements the Sleeper interface
type SpyTime struct {
	durationSlept time.Duration
}

// Sleep modify the Sleep function to be used in our test
func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

// Countdown takes a io.Writer
// for the arguments we don't need * because we are using interfaces
// func Countdown(out *bytes.Buffer, sleeper *SpySleeper)
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	// creating an instance of ConfigurableSleeper passing down the
	// real time.Sleep function
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
