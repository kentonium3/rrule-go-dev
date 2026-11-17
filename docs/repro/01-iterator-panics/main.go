//go:build ignore
// +build ignore

package main

import (
	"fmt"

	"github.com/kentonium3/rrule-go"
)

func main() {
	for _, rule := range []string{
		"FREQ=MONTHLY;BYDAY=5MO",                    // fine
		"FREQ=MONTHLY;BYDAY=-6MO",                   // 1. BYDAY ordinal
		"FREQ=MONTHLY;BYDAY=8MO",                    // 1.
		"FREQ=MONTHLY;BYDAY=53MO",                   // 1.
		"FREQ=YEARLY;BYMONTH=1;BYDAY=-6MO",          // 1.
		"FREQ=YEARLY;BYMONTH=12;BYDAY=8MO",          // 1.
		"FREQ=YEARLY;BYEASTER=-100",                 // 2. BYEASTER
		"FREQ=YEARLY;BYEASTER=300",                  // 2.
		"FREQ=MONTHLY;INTERVAL=9223372036854775807", // 3. INTERVAL
	} {
		fmt.Printf("%-45s %s\n", rule, firstTwo(rule))
	}
}

func firstTwo(rule string) (result string) {
	defer func() {
		if r := recover(); r != nil {
			result = fmt.Sprint("panic: ", r)
		}
	}()
	r, err := rrule.StrToRRule("DTSTART:20260105T100000Z\nRRULE:" + rule)
	if err != nil {
		return "error: " + err.Error()
	}
	next := r.Iterator()
	for i := 0; i < 2; i++ {
		t, ok := next()
		if !ok {
			break
		}
		result += t.Format("2006-01-02 ")
	}
	return result
}
