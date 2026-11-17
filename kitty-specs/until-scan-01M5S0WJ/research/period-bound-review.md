# Independent UNTIL bound design challenge

Read AGENTS.md, docs/CONTEXT.md, issue 03, mission spec/research and iterator code. No repository edits.

## Confirmed real-world counterexample

A guard that stops once the nominal period calendar date exceeds UNTIL's date in DTSTART's zone changes existing output, even for ordinary DST transitions.

Baseline command: `GOCACHE=/tmp/rrule-review-gocache go run /tmp/rrule-until-dates.go` (Go 1.25.14).

- America/Sao_Paulo, DAILY, DTSTART 2018-11-03 00:00 -03, UNTIL 2018-11-03 23:59:59 -03: baseline yields Nov 3 midnight twice. time.Date(2018,11,4,0,...) normalizes to Nov 3 23:00.
- Pacific/Apia, DAILY, DTSTART 2011-12-29 00:00 -10, UNTIL 2011-12-29 23:59:59 -10: baseline yields Dec 29 midnight twice. time.Date(2011,12,30,0,...) normalizes to Dec 29 midnight.

Mechanism: firstyday.AddDate(0,0,i).Date() normalizes the period's intended midnight before its date is extracted; candidate reconstruction then uses that prior date and the timeset clock. A nominal next-day guard removes the duplicate. This is a real IANA historical transition, not a crafted TZif. It should be a golden compatibility test under the current exact-output requirement. Even if duplicate output is considered an existing impossible result, that classification is a scope decision rather than a proof that the guard preserves behavior.

## Assessment

A raw cursor bound also fails yearly/monthly expanded periods as already documented. An extra calendar day addresses these two examples but is not a general proof: firstyday itself is normalized and arbitrary time.Location data has no practical offset bound exposed by Go 1.16 APIs. A fixed one/two-day margin needs an explicit supported-zone assumption; it cannot claim preservation of every accepted location.

The smallest unconditionally provable fast path is UTC (and demonstrably fixed zones if their constancy can be established) with the original behavior for other zones. That narrows FR-001 and requires user authorization, so do not silently select it.

A broader timezone-aware design might use ZoneBounds (Go >=1.19), but I have not established a complete, small implementation/proof: merely inspecting the current interval does not bound future transitions or offsets used by time.Date's normalization. Do not claim a minimum-version bump alone solves correctness. Enumerating transitions across an adequately justified domain and accounting for firstyday/AddDate normalization needs further design review.

Recommendation: present the concrete ordinary-DST counterexample and ask the user to choose an explicit zone-domain assumption / narrower fast path or authorize further timezone-bound design. Do not introduce a speculative multi-century padding bound or claim date-only equivalence.
