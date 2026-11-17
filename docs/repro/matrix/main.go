//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/kentonium3/rrule-go"
)

func try(rule string, n int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%-55s PANIC: %v\n", rule, r)
		}
	}()
	r, err := rrule.StrToRRule("DTSTART:20260105T100000Z\nRRULE:" + rule)
	if err != nil {
		fmt.Printf("%-55s parse error: %v\n", rule, err)
		return
	}
	var ms0, ms1 runtime.MemStats
	runtime.ReadMemStats(&ms0)
	t0 := time.Now()
	next := r.Iterator()
	var got []string
	for i := 0; i < n; i++ {
		v, ok := next()
		if !ok {
			break
		}
		got = append(got, v.Format("2006-01-02"))
	}
	runtime.ReadMemStats(&ms1)
	fmt.Printf("%-55s ok: %v (%.2fs, %d MB allocated)\n", rule, got, time.Since(t0).Seconds(), (ms1.TotalAlloc-ms0.TotalAlloc)>>20)
}

func main() {
	fmt.Println(runtime.Version(), runtime.GOOS+"/"+runtime.GOARCH)
	for _, rule := range []string{
		"FREQ=MONTHLY;BYDAY=5MO",
		"FREQ=MONTHLY;BYDAY=-5MO",
		"FREQ=MONTHLY;BYDAY=-6MO",
		"FREQ=MONTHLY;BYDAY=7MO",
		"FREQ=MONTHLY;BYDAY=8MO",
		"FREQ=MONTHLY;BYDAY=10MO",
		"FREQ=MONTHLY;BYDAY=53MO",
		"FREQ=MONTHLY;BYDAY=-53MO",
		"FREQ=YEARLY;BYMONTH=12;BYDAY=8MO",
		"FREQ=YEARLY;BYMONTH=1;BYDAY=-6MO",
		"FREQ=YEARLY;BYDAY=53MO",
		"FREQ=YEARLY;BYDAY=-53MO",
		"FREQ=YEARLY;BYEASTER=0",
		"FREQ=YEARLY;BYEASTER=-100",
		"FREQ=YEARLY;BYEASTER=300",
		"FREQ=MONTHLY;INTERVAL=9223372036854775807",
		"FREQ=YEARLY;INTERVAL=9223372036854775807",
		"FREQ=DAILY;INTERVAL=9223372036854775807",
	} {
		try(rule, 3)
	}
}
