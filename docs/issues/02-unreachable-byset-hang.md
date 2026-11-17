# 02: Unreachable sub-daily BY-list never returns

GitHub issue: #5 (this repo). Mission 2.

## Goal

HOURLY, MINUTELY and SECONDLY rules whose BYHOUR, BYMINUTE or BYSECOND values cannot be reached from DTSTART at the given INTERVAL must not loop forever. Either the rule fails at construction, or iteration finishes with no occurrences in bounded time.

## Acceptance

- Each unreachable case in `docs/repro/02-unreachable-byset-hang/main.go` returns an error from `StrToRRule`/`NewRRule`, or `After()`/`All()` return within a test timeout of 1 s. Make it a red-first test that guards itself with a timeout and doesn't hang the suite.
- Reachable cases are unchanged; `FREQ=HOURLY;INTERVAL=2;BYHOUR=2` still returns at once.
- **A plan decision:** construction-time rejection depends on DTSTART, and `DTStart()`/`Until()` can change DTSTART after construction. PR #65's approach changed those setters' signatures to return an error. That breaks the API, which is not allowed (AGENTS.md rule 3). Find a design that keeps the signatures. One option is to detect unreachability lazily when iteration starts, and yield nothing.

## The report as sent to the maintainer (2026-10-23)


**Describe the bug**
teambition/rrule-go version: v1.8.2 (also reproduced on master e74d1634)

For `FREQ=HOURLY`, `MINUTELY` and `SECONDLY`, if no value in `BYHOUR` / `BYMINUTE` / `BYSECOND` can be reached from DTSTART in steps of `INTERVAL`, the iterator loops forever. For example, `FREQ=HOURLY;INTERVAL=2;BYHOUR=1` with DTSTART at an even hour only ever visits even hours. `After()`, `All()` and so on never return; the calling goroutine spins in a tight loop. Setting `UNTIL` does not help.

The cause appears to be the inner `for` loop that advances the hour (and the equivalent minute and second loops) ([rrule.go#L737-L748](https://github.com/teambition/rrule-go/blob/v1.8.2/rrule.go#L737-L748)). It only exits when the BY-list matches. MAXYEAR and UNTIL are checked outside this loop, so they are never reached. This is different from rules that simply match nothing, which do eventually stop at MAXYEAR.

**To Reproduce**
```go
package main

import (
	"fmt"
	"time"

	"github.com/teambition/rrule-go"
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
```

Output:
```
FREQ=HOURLY;INTERVAL=2;BYHOUR=2                         After() = 2026-01-06 02:00:00 +0000 UTC
FREQ=HOURLY;INTERVAL=2;BYHOUR=1                         After() did not return within 10s
FREQ=MINUTELY;INTERVAL=2;BYMINUTE=1                     After() did not return within 10s
FREQ=SECONDLY;INTERVAL=2;BYSECOND=1                     After() did not return within 10s
FREQ=HOURLY;INTERVAL=2;BYHOUR=1;UNTIL=20260201T000000Z  After() did not return within 10s
```

(The three unreachable cases were also left running for 60 seconds with the same result.)

**Expected behavior**
`NewRRule` / `StrToRRule` should return an error for these rules, or iteration should finish with no occurrences. It should not loop forever.

python-dateutil handles this case in `rrule.__construct_byset`, which raises `ValueError("Invalid rrule byxxx generates an empty set.")` when the BY-list cannot be reached from DTSTART at the given interval. PR #65 ported that approach (`constructByset`, called from `buildRRule` for HOURLY/MINUTELY/SECONDLY) but was closed shortly after it was opened, so it may be a useful starting point. As written, it also changes the signatures of `DTStart()` and `Until()` to return an error, because changing DTSTART can make a previously valid rule invalid.

**Additional context**
- rrule-go v1.8.2 and master (e74d1634)
- Go 1.25.14 and Go 1.27.0, linux/amd64 (same output)
- Found while validating user-supplied rules in an application. A rule like this, submitted once, leaves a goroutine spinning until the process restarts.
- Reproduced 2026-10-23

## Private fix validation — 2026-10-25

The iterator now bounds the existing hourly, minutely and secondly clock searches
at 24, 1440 and 86400 failed candidates. A clock residue repeats within that
many interval steps, so exhaustion proves the clock filters unreachable. The
last candidate is checked, buffered occurrences survive, and constructors and
DTStart/Until signatures and acceptance remain unchanged.

The trigger is an accepted interval whose clock residues cannot satisfy the
applicable BY-lists together. Ordinary interval-2 rules suffice; accidental
user input is plausible and deliberate triggering is trivial. Original upstream
can spin indefinitely. The earlier overflow fix makes the patched baseline
scan to MAXYEAR instead: the hourly reproduction measured 475 ms here, while
minute/second reproductions exceeded the four-second subprocess watchdog.
Historical upstream hourly also exceeded that watchdog.

Public-path tests cover both constructors, After/All/Iterator, exact reachable
controls (hourly first result 2026-01-06T02:00:00Z), setters, combined/mixed
filters, missing filters, local time, filtered days, year carry, final-cycle
matches, buffered results and MAXYEAR. Measured fixed original fixture All
operations were about 30 microseconds hourly/minutely and 1.4 ms secondly;
operation time is measured inside each child, separately from the four-second
startup/teardown watchdog. CommandContext kills timed-out children and
CombinedOutput waits for them. Logs and the full verification record are in
/tmp/rrule-hang-WP01-handoff.md during the private mission.

This change does not address general never-matching calendar scans or change
UNTIL checks for those scans; see issue 03. Implementation: gpt-6.1-sol on
2026-10-25; independent different-model review is recorded by the mission.
