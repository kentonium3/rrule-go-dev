# 06: Floating UNTIL values are read as UTC

GitHub issue: kentonium3/rrule-go#6 (public). From teambition/rrule-go#67. Priority P2.

**Normative rule (Kent, 2026-10-25): adhere to RFC 5545 in all cases.**

## Problem

`str.go` parses a DATE-TIME with no `Z` and no `TZID` (a floating local time) as UTC, and accepts
UNTIL forms that RFC 5545 forbids. Measured on v1.8.3 (`docs/repro/05-07-backlog/out-v1.8.3.txt`):

1. **A floating rule round-trips as UTC.** `DTSTART:20201228T090000` + `UNTIL=20201230T090000`
   comes back from `String()` with `Z` on both, which is a different rule.
2. **A forbidden UNTIL form is silently reinterpreted.** With a TZID DTSTART (Sydney 09:00 daily),
   `UNTIL=20201230T080000` (floating) is read as 08:00 UTC and gives 3 occurrences. RFC 5545
   requires a UTC UNTIL here.

## RFC 5545 requirements

- **§3.3.10, UNTIL:** UNTIL must have the same value type as DTSTART (DATE or DATE-TIME).
  - A floating DTSTART needs a floating UNTIL.
  - A UTC or TZID DTSTART needs a UTC UNTIL.
- **§3.3.5:** `TZID` must not be applied to a DATE-TIME given in UTC (`TZID=…:…Z`).
- **§3.3.5, FORM #1:** a floating value is bound to no time zone. It must not become UTC on output.

## Acceptance

- **Reject what RFC 5545 forbids.** `StrToRRule`, `StrToRRuleSet` and `StrToROption*` return an
  error, naming the rule part, for:
  - a TZID or UTC DTSTART with a floating UNTIL;
  - a floating DTSTART with a UTC UNTIL;
  - a DATE DTSTART with a DATE-TIME UNTIL, or the reverse;
  - `TZID=…` applied to a value ending in `Z`.

  python-dateutil (2.7+) also raises for a TZ-aware DTSTART with a non-UTC UNTIL. This is a
  deliberate behaviour change for inputs that parse today: list every newly rejected form in the
  spec and in the release notes.
- **Keep accepting every conformant form** with identical occurrences and `String()` output. Above
  all that means `DTSTART;TZID=…` + `UNTIL=…Z`, the most common form.
- **Floating stays floating.** A floating DTSTART and UNTIL round-trip through `String()` without
  `Z`, and their occurrences are unchanged. Go has no floating time type, so the parser must
  remember that a value was floating. Choose the smallest design that works. An exported addition
  needs a justification in the plan and makes the release v1.9.0.
- **A RRULE with no DTSTART line** (`StrToROption("FREQ=…;UNTIL=…")`, the path Vikunja uses) has
  nothing to compare UNTIL with. It keeps today's behaviour: a floating UNTIL is read in the `loc`
  argument. The plan confirms that this conforms to the RFC.
- `NewRRule` with an `Until` in any location keeps today's `String()` output (UTC with `Z`).
- The order of DTSTART and UNTIL in the input doesn't matter. DTSTART may be a separate line or
  inside the RRULE. `StrToRRuleSet` checks each RRULE/EXRULE line against the shared DTSTART.

## Prior art: teambition PR #68 (do not trust it)

https://github.com/teambition/rrule-go/pull/68, open since 2024-08-15. Probed on 2026-10-25 by
applying its head `d898b57a07` to v1.8.2:

- **It rejects the RFC-required form.** `DTSTART;TZID=Australia/Sydney:…` + `UNTIL=20201229T210000Z`
  fails with "DTSTART and UNTIL time format mismatch". Its check compares "ends in Z" between
  DTSTART and UNTIL, but a TZID DTSTART has no Z and must have a UTC UNTIL. Do not adopt that check.
- For a TZID DTSTART it writes UNTIL **without** `Z`, which RFC 5545 forbids. That also changes
  `String()` for callers that pass a local `Until` to `NewRRule`.
- It does not fix the floating round-trip.
- **Reusable after verification:**
  - its `TZID=…:…Z` rejection in `StrToDtStart`, which is RFC-correct;
  - the idea of parsing a floating UNTIL in DTSTART's location. Under the RFC rule that only
    applies when DTSTART is floating, and then the location is floating too.

  Note that it only works when DTSTART was parsed first, and its `Location() != nil` check is
  always true.
