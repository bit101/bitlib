// Package bltime has time related functions
package bltime

import (
	"fmt"
	"time"
)

type timer struct {
	startTime time.Time
}

var timers = make(map[any]*timer)

// StartTimer starts a timer. Can also reset an existing timer.
func StartTimer(key any) error {
	t, ok := timers[key]
	if !ok {
		t = &timer{}
		timers[key] = t
	}
	t.startTime = time.Now()
	return nil
}

// TimerElapsed returns the current duration of the timer but does not stop it.
func TimerElapsed(key any) (time.Duration, error) {
	t, ok := timers[key]
	if !ok {
		return 0, fmt.Errorf("no timer for key %v", key)
	}
	return time.Now().Sub(t.startTime), nil
}

// PrintTimerElapsed prints the current duration of the timer but does not stop it.
func PrintTimerElapsed(key any) {
	t, ok := timers[key]
	if !ok {
		fmt.Printf("no timer for key %v\n", key)
	}
	duration := time.Now().Sub(t.startTime)
	fmt.Println(duration)
}
