# 01: Iterator panics on rules that NewRRule/StrToRRule accept

GitHub issue: #4 (this repo). Mission 1.

## Goal

Every rule that `StrToRRule` or `NewRRule` accepts must iterate without panicking. Either the rule is rejected with an error at construction, or the parts that can't match yield no occurrences, as `BYDAY=6MO` and `7MO` already do.

## Acceptance

- Every rule in `docs/repro/01-iterator-panics/main.go` either returns an error from `StrToRRule`/`NewRRule` or iterates without panicking. Turn the program's list into a table-driven test that is red on the unfixed code.
- A fuzz test over the BY* domains, BYDAY ordinals, BYEASTER and INTERVAL (`FuzzIterate` or similar) runs for at least 60 s with no panic. Commit the seed corpus.
- Rules that iterate correctly today give byte-identical occurrences. Add golden tests for nearby valid cases: `BYDAY=5MO`, `-5MO`, `BYEASTER=-80` and `257`, and the largest INTERVAL that stays valid.
- The exported API is unchanged.

## The report as sent to the maintainer (2026-10-23)

----------------------------------------------------------------------
Issue 1: Iterator panics with "index out of range" for rules that NewRRule/StrToRRule accept (BYDAY ordinal, BYEASTER, large INTERVAL)
----------------------------------------------------------------------

**Describe the bug**
teambition/rrule-go version: v1.8.2 (also reproduced on master e74d1634; `rrule.go` is unchanged between them)

Some rules pass `StrToRRule` / `NewRRule` validation without error but then panic with `index out of range` as soon as they are iterated. I found three separate cases while validating user-supplied rules in an application. In each one, the input is accepted by `validateBounds` and later used as a slice index without a bounds check:

1. **BYDAY ordinal outside the month.** For `FREQ=MONTHLY`, and for `FREQ=YEARLY` with `BYMONTH`, an ordinal larger than the number of weeks in the month panics: `+8` and up, or `-6` and below. `validateBounds` allows ±1..53 for every frequency. In `rebuild()` the index `first + (n-1)*7` / `last + (n+1)*7` is used to read `wdaymask[i]` before the `first <= i && i <= last` check ([rrule.go#L466-L476](https://github.com/teambition/rrule-go/blob/v1.8.2/rrule.go#L466-L476)). (`+6` and `+7` do not panic. They never match, so iteration runs to MAXYEAR.)
2. **BYEASTER offset outside the year.** `eastermask` has `yearlen+7` entries and `eastermask[eyday+offset]` is not bounds-checked ([rrule.go#L481-L487](https://github.com/teambition/rrule-go/blob/v1.8.2/rrule.go#L481-L487)). With the current mask size, offsets from -80 to 257 are safe for every year. -81 and 258 already panic in some years (checked for 1600 to 2500), and values further out, such as -100 or 300, panic immediately.
3. **INTERVAL large enough to overflow.** `validateBounds` only rejects `INTERVAL < 0`. For `FREQ=MONTHLY`, `iterator.month += time.Month(r.interval)` overflows ([rrule.go#L705](https://github.com/teambition/rrule-go/blob/v1.8.2/rrule.go#L705)) and the next `rebuild()` panics. Other frequencies do not panic but produce wrong output: with `FREQ=DAILY;INTERVAL=9223372036854775807` the occurrences after DTSTART are 2026-12-31, then 2026-12-30.

**To Reproduce**
```go
package main

import (
	"fmt"

	"github.com/teambition/rrule-go"
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
```

Output:
```
FREQ=MONTHLY;BYDAY=5MO                        2026-03-30 2026-06-29
FREQ=MONTHLY;BYDAY=-6MO                       panic: runtime error: index out of range [-5]
FREQ=MONTHLY;BYDAY=8MO                        panic: runtime error: index out of range [383] with length 382
FREQ=MONTHLY;BYDAY=53MO                       panic: runtime error: index out of range [395] with length 382
FREQ=YEARLY;BYMONTH=1;BYDAY=-6MO              panic: runtime error: index out of range [-5]
FREQ=YEARLY;BYMONTH=12;BYDAY=8MO              panic: runtime error: index out of range [383] with length 382
FREQ=YEARLY;BYEASTER=-100                     panic: runtime error: index out of range [-6]
FREQ=YEARLY;BYEASTER=300                      panic: runtime error: index out of range [394] with length 372
FREQ=MONTHLY;INTERVAL=9223372036854775807     panic: runtime error: index out of range [9223372036854775807] with length 13
```

None of these rules returns an error from `StrToRRule`. Every panic happens inside the iterator. The exact index varies with DTSTART.

**Expected behavior**
Iterating a rule that `NewRRule` / `StrToRRule` accepted should never panic. Either the rule is rejected with an error at construction, or iteration yields no occurrences for the parts that cannot match (as `BYDAY=6MO` and `7MO` already do).

Possible fixes, in case they help:

1. **BYDAY ordinal:** skip an ordinal whose starting index already falls outside `[first, last]` before reading `wdaymask[i]`. This keeps today's accepted inputs and makes out-of-range ordinals match nothing. Alternatively, reject `|n| > 5` in `validateBounds` when `Freq` is MONTHLY, or YEARLY with `BYMONTH`. That is stricter than the RFC 5545 grammar, but it also avoids the long scan to MAXYEAR that `6MO`/`7MO` trigger today.
2. **BYEASTER:** skip offsets where `eyday+offset` falls outside `eastermask`, or add a bound for `Byeaster` in `validateBounds`. BYEASTER is not part of RFC 5545, so there is no standard limit to follow.
3. **INTERVAL:** add an upper bound to `validateBounds`. Any limit that keeps the year/month/day/hour/minute/second arithmetic from overflowing would do; real-world intervals are far smaller.

**Additional context**
- rrule-go v1.8.2 and master (e74d1634)
- Go 1.25.14 and Go 1.27.0, linux/amd64 (same output)
- Reproduced 2026-10-23

----------------------------------------------------------------------
Issue 2: Iterator never returns for HOURLY/MINUTELY/SECONDLY rules when BYHOUR/BYMINUTE/BYSECOND cannot be reached at the given INTERVAL
----------------------------------------------------------------------

## Fix notes (2026-10-25)

The repair retains constructor acceptance and skips impossible mask candidates.
Ordinal starting indices are checked against the selected month/year before either
weekday-mask access, then checked again after weekday adjustment. Impossible
monthly ordinals yield no dates; mixed lists retain their valid entries. YEARLY
ordinals without BYMONTH and WEEKLY's ignored ordinals keep their semantics.

Easter offsets are compared against the existing mask bounds before addition,
including machine-int extremes. Offsets outside that mask are skipped for the
current year; an offset such as -100 can still match a later year. Existing
in-mask offsets and the seven-day weekly tail remain supported.

INTERVAL has no new arbitrary cap. Advancement compares each step with the
remaining supported civil calendar before addition or weekly multiplication.
Subdaily arithmetic uses bounded int64 clock units, including accumulated day
carries on every inner BY-filter iteration. UTC civil dates provide the horizon
without a multi-century time.Duration or DST elapsed-time drift; emitted dates
retain the original local wall-clock behavior. A step beyond MAXYEAR exhausts
the iterator after already-buffered occurrences have been returned. In particular,
2147483647 seconds/minutes still advance correctly on 32-bit systems when the
next date is representable.

**Known Easter limit retained:** YEARLY BYEASTER=257 with DTSTART
1943-01-01T10:00:00Z and UNTIL 1943-12-31T23:59:59Z returns no occurrence.
Easter 1943 is April 25; adding 257 days gives January 7, 1944, inside the
seven-day mask tail but outside the YEARLY emitted dayset. The nearby 1944
fixture (Easter April 9) correctly yields December 22, 1944. These are explicit
regression cases. This repair does not redesign cross-year Easter semantics.
Boundary fixtures also include 1818 (March 22 Easter, offset -81 outside the
mask and -80 yielding January 1) and 1943 (offset 258 beyond the mask).

The crash conditions are accepted ordinal/Easter selections or large intervals
reaching unchecked arithmetic/indexing, followed by real iterator evaluation.
They are realistic for services evaluating user-supplied RRULEs; no frequency
estimate across deployed services is claimed. The deterministic safety suite
exercises both NewRRule and StrToRRule, exact positive controls, all frequencies,
integer widths, buffered exhaustion and DST. Fuzz coverage is delivered separately
by WP02. General never-matching UNTIL scans and the separate unreachable-subdaily
BY-list problem remain outside this repair's scope.
