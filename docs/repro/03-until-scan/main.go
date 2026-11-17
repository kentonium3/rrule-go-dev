//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kentonium3/rrule-go"
)

func main() {
	var pos []string
	for i := 2; i <= 366; i++ {
		pos = append(pos, strconv.Itoa(i))
	}
	for _, c := range []struct{ label, rule string }{
		{"BYSETPOS=2", "FREQ=DAILY;BYSETPOS=2;UNTIL=20260201T000000Z"},
		{"BYSETPOS=2..366", "FREQ=DAILY;BYSETPOS=" + strings.Join(pos, ",") + ";UNTIL=20260201T000000Z"},
	} {
		r, err := rrule.StrToRRule("DTSTART:20260105T100000Z\nRRULE:" + c.rule)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		start := time.Now()
		all := r.All()
		fmt.Printf("FREQ=DAILY;%s;UNTIL=20260201T000000Z  %d occurrences in %v\n", c.label, len(all), time.Since(start).Round(time.Millisecond))
	}
}
