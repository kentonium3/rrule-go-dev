//go:build ignore
// +build ignore

// Reproductions for public issues kentonium3/rrule-go#5, #6 and #7.
// Run from the repo root: go run ./docs/repro/05-07-backlog/main.go
package main

import (
	"fmt"
	"time"

	"github.com/kentonium3/rrule-go"
)

func list(ts []time.Time) []string {
	out := []string{}
	for _, t := range ts {
		out = append(out, t.Format("2006-01-02 15:04:05 MST"))
	}
	return out
}

func main() {
	syd, err := time.LoadLocation("Australia/Sydney")
	if err != nil {
		panic(err)
	}

	fmt.Println("## #5 nonexistent and repeated local times (Sydney: 2022-10-02 02:00->03:00, 2023-04-02 03:00->02:00)")
	for _, c := range []struct {
		name string
		o    rrule.ROption
	}{
		{"spring HOURLY COUNT=6", rrule.ROption{Freq: rrule.HOURLY, Count: 6, Dtstart: time.Date(2022, 10, 2, 0, 0, 0, 0, syd)}},
		{"fall HOURLY COUNT=6", rrule.ROption{Freq: rrule.HOURLY, Count: 6, Dtstart: time.Date(2023, 4, 2, 0, 0, 0, 0, syd)}},
		{"spring MINUTELY BYMINUTE=0 COUNT=4", rrule.ROption{Freq: rrule.MINUTELY, Count: 4, Byminute: []int{0}, Dtstart: time.Date(2022, 10, 2, 0, 0, 0, 0, syd)}},
		{"spring DAILY BYHOUR=2,3 COUNT=4", rrule.ROption{Freq: rrule.DAILY, Count: 4, Byhour: []int{2, 3}, Dtstart: time.Date(2022, 10, 1, 0, 0, 0, 0, syd)}},
		{"spring DAILY 02:30 COUNT=3", rrule.ROption{Freq: rrule.DAILY, Count: 3, Dtstart: time.Date(2022, 10, 1, 2, 30, 0, 0, syd)}},
		{"fall DAILY 02:30 COUNT=3", rrule.ROption{Freq: rrule.DAILY, Count: 3, Dtstart: time.Date(2023, 4, 1, 2, 30, 0, 0, syd)}},
		{"control: fall HOURLY BYHOUR=9 COUNT=2", rrule.ROption{Freq: rrule.HOURLY, Count: 2, Byhour: []int{9}, Dtstart: time.Date(2023, 4, 1, 0, 0, 0, 0, syd)}},
		{"control: fall DAILY COUNT=3", rrule.ROption{Freq: rrule.DAILY, Count: 3, Dtstart: time.Date(2023, 4, 1, 9, 0, 0, 0, syd)}},
	} {
		r, err := rrule.NewRRule(c.o)
		if err != nil {
			fmt.Printf("%-40s error: %v\n", c.name, err)
			continue
		}
		fmt.Printf("%-40s %v\n", c.name, list(r.All()))
	}

	fmt.Println("\n## #6 floating UNTIL")
	for _, s := range []string{
		"DTSTART:20201228T090000\nRRULE:FREQ=DAILY;UNTIL=20201230T090000",
		"DTSTART;TZID=Australia/Sydney:20201228T090000\nRRULE:FREQ=DAILY;UNTIL=20201230T080000",
		"control: DTSTART;TZID=Australia/Sydney:20201228T090000\nRRULE:FREQ=DAILY;UNTIL=20201229T210000Z",
	} {
		label := ""
		if len(s) > 9 && s[:9] == "control: " {
			label, s = "control ", s[9:]
		}
		r, err := rrule.StrToRRule(s)
		if err != nil {
			fmt.Printf("%s%q error: %v\n", label, s, err)
			continue
		}
		fmt.Printf("%sin  %q\n%s  n=%d %v\n%s  out %q\n", label, s, label, len(r.All()), list(r.All()), label, r.String())
	}

	fmt.Println("\n## #7 repeated BY values")
	start := time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC)
	for _, sec := range []string{"0,30", "0,30,0"} {
		set, err := rrule.StrToRRuleSet("DTSTART:20080101T090000Z\nRRULE:FREQ=MINUTELY;COUNT=12;BYSECOND=" + sec)
		if err != nil {
			panic(err)
		}
		d := set.Between(start, end, true)
		fmt.Printf("RRuleSet BYSECOND=%-7s n=%d last=%s\n", sec, len(d), d[len(d)-1].Format(time.RFC3339))
	}
	r, err := rrule.StrToRRule("DTSTART:20080101T090000Z\nRRULE:FREQ=DAILY;COUNT=6;BYHOUR=9,9,10")
	if err != nil {
		panic(err)
	}
	fmt.Printf("RRule BYHOUR=9,9,10 COUNT=6: %v\n", list(r.All()))
	fmt.Printf("RRule String: %q\n", r.String())
}
