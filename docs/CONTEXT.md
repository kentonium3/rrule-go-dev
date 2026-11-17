# Context

## Why this copy exists

Vikunja (go-vikunja/vikunja) is moving task recurrence to RFC 5545 RRULE strings evaluated by rrule-go v1.8.2, in PR #3071. While hardening that PR against user-supplied rules, the following were found in rrule-go:

- **Panics** on rules that `StrToRRule`/`NewRRule` accept:
  - a BYDAY ordinal outside the month;
  - a BYEASTER offset outside the year mask;
  - a very large INTERVAL.
- **A hang** that never returns: HOURLY, MINUTELY and SECONDLY rules whose BY-list can't be reached at the INTERVAL.
- **A design limit:** UNTIL is only checked when an occurrence is emitted, so a rule that never matches scans to the year 9999.

A panic in a goroutine without `recover` kills a Go service. The hang pins a CPU forever. Any service that iterates user-supplied RRULE strings is exposed.

## Upstream state (checked 2026-10-23)

- **Activity:** the last change to Go code was `7bd6c830` on 2023-01-13, in v1.8.2. Master's later commits only touch the README and CI. Issue #70 (2026-09-13) has had no maintainer response.
- **Private report:** sent to the maintainer's commit email on 2026-10-23, following the repo's "Security issue disclosure" template, with a follow-up on 2026-10-25. No reply so far. Public issues are planned for on or after 2026-11-13.
- **Existing coverage:** dedup across all 71 upstream issues and PRs found nothing covering these bugs. Related history:
  - **PR #65** (closed unmerged, never reviewed). It ports python-dateutil's `__construct_byset` empty-byset check, which is exactly the fix class for the hang in issue 02. It also changes `DTStart()`/`Until()` to return an error. Do NOT copy that: it breaks the API (AGENTS.md rule 3).
  - **PR #31** (merged 2019). It added `validateBounds`, including the ±53 BYDAY ordinal limit for every frequency. That limit is too wide for MONTHLY, and that gap is behind panic 1.
  - **#24 / PR #27.** These added the default UNTIL of DTSTART + about 292 years. It does not bound never-matching scans (issue 03).
  - **#50.** The maintainer says `StrToRRule` returns an error for an invalid rule. That supports "reject at construction" as the expected behaviour.

## How Vikunja guards today (stay consistent with these)

Vikunja checks every rule before handing it to rrule-go, and recovers from panics around evaluation. The guards live in `pkg/models/task_repeat.go`: `checkRRuleBounds`, `checkRRuleWriteLimits` and `boundedForEvaluation`.
- BY* domains: BYSECOND 0..59, BYMINUTE 0..59, BYHOUR 0..23, BYMONTHDAY ±1..31, BYYEARDAY ±1..366, BYWEEKNO ±1..53, BYMONTH 1..12, BYSETPOS ±1..366, and BYEASTER -80..257. Easter falls between day 81 and day 116, so -80..257 are the offsets that stay inside every year's mask.
- BYDAY ordinals: at most ±5 for MONTHLY, or for YEARLY with BYMONTH. No ordinal below MONTHLY.
- INTERVAL: at most a 100-year span per frequency.
- At most 10,000 occurrences per period, 10 BYSETPOS values and a COUNT of 10,000.

Keep these guards after the library is fixed, as defence in depth. The library fix must reject or safely handle at least everything Vikunja rejects. Vikunja's limits are stricter in places, such as COUNT and BYSETPOS, for CPU-cost reasons a library need not adopt.

## Versions tested

- **Builds:** v1.8.2 and master `e74d1634`; the `rrule.go` blob is identical in both.
- **Go and OS:** go1.25.14 and go1.27.0 on linux/amd64. The outputs were identical.
- **Index values:** the panic indices depend on DTSTART and the year. The repro programs fix DTSTART at `20260105T100000Z`.

## Extra observations, not in the emailed report

- `FREQ=MONTHLY;BYDAY=7MO` (no UNTIL) allocates about 280 MB while scanning to MAXYEAR. `+6MO` and `+7MO` never match, so they don't panic, but they do scan.
- `FREQ=YEARLY;BYDAY=±53MO` without BYMONTH is fine. WEEKLY ignores ordinals.
- `FREQ=HOURLY;INTERVAL=9223372036854775807` returns 24 occurrences, an overflow artifact.
- `FREQ=DAILY;INTERVAL=9223372036854775807`: the occurrences after DTSTART run backwards (2026-12-31, then 2026-12-30).
- `BYEASTER=257` silently drops occurrences in some years: the index is beyond `yearlen` but still inside the mask.
- Evidence: `docs/repro/edge-cases/` and `docs/repro/matrix/`.

## Update 2026-10-25: public fork

The upstream author confirmed the project is unmaintained, and the private-disclosure plan above was superseded. The fixes are public as `kentonium3/rrule-go` v1.8.3 (teambition/rrule-go#72 points to it). The next mission works teambition's remaining valid open issues, ported to the public repo as #5–#7; see `docs/runbooks/teambition-backlog.md`. Teambition #58 is a usage question, not a bug.
