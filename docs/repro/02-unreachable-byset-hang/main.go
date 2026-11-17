//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"time"

	"github.com/kentonium3/rrule-go"
)

func main() {
	// DTSTART is at an even hour, minute and second, so with INTERVAL=2
	// the iterator only ever visits even values and BY*=1 is unreachable.
	for _, rule := range []string{
		"FREQ=HOURLY;INTERVAL=2;BYHOUR=2",     // reachable: returns at once
		"FREQ=HOURLY;INTERVAL=2;BYHOUR=1",     // unreachable
		"FREQ=MINUTELY;INTERVAL=2;BYMINUTE=1", // unreachable
		"FREQ=SECONDLY;INTERVAL=2;BYSECOND=1", // unreachable
		"FREQ=HOURLY;INTERVAL=2;BYHOUR=1;UNTIL=20260201T000000Z",
	} {
		r, err := rrule.StrToRRule("DTSTART:20260105T100000Z\nRRULE:" + rule)
		if err != nil {
			fmt.Printf("%-55s error: %v\n", rule, err)
			continue
		}
		done := make(chan time.Time, 1)
		go func() { done <- r.After(time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC), false) }()
		select {
		case t := <-done:
			fmt.Printf("%-55s After() = %v\n", rule, t)
		case <-time.After(10 * time.Second):
			fmt.Printf("%-55s After() did not return within 10s\n", rule)
		}
	}
}
