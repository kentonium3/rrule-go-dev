# Discovery: unreachable sub-daily BY-lists

Scope confirmed by Kent on 2026-10-25 in decision 01M5RZ9H1HV8HPSBD3VE3YF6WE.

## Evidence and decisions

- Source brief: `docs/issues/02-unreachable-byset-hang.md`, private issue 5; reproductions in `docs/repro/02-unreachable-byset-hang/main.go`.
- Baseline: merged master `1a62a24b9460201d16e43b693c5ea9cfc3700198`. The first mission bounds arithmetic at MAXYEAR, but the inner sub-daily loops in `rrule.go` still search repeatedly for unreachable clock filters. This can require scanning the entire remaining calendar horizon. The prior description of literal infinite iteration applies to the original upstream baseline; the current mission requires prompt termination on the patched baseline too.
- Trigger: accepted interval and clock filters select disjoint clock residues (for example, an even start hour and interval 2 with BYHOUR=1). Ordinary small values suffice; likely accidental with user-authored filters and trivial to trigger deliberately. An evaluation consumes CPU without returning promptly.
- Keep all exported signatures and correct sequences. `DTStart` and `Until` rebuild rule state, so constructor-only rejection cannot be the sole safeguard. Decide the smallest iterator-time solution during planning.
- Confirmed regression threshold: one second for each reported unreachable fixture. Tests must clean up timed-out execution rather than leave spinning goroutines in the suite.
- General never-matching calendar filters and UNTIL scanning are outside this mission. No public release or Vikunja changes.

## Planning questions

- Prefer a bounded clock-cycle search or equivalent reachability proof; verify combined hour/minute/second filters, interval carries and filtered-day shortcuts before selecting the implementation.
- Preserve existing buffered occurrences and evaluate fresh iterators after setters in both reachable-to-unreachable and reverse directions.
- Select timeout-safe red/green tests and exact reachable controls, including interval/calendar boundaries and mixed BY-lists.

## Confirmed planning outcome — 2026-10-25

Decision: bounded existing inner searches, normal iterator exhaustion, constructors unchanged. Kent confirmed engineering alignment in decision 01M5RZHPBXMKMK5YBHQB6R8BTR.

Evidence: rrule.go generate advances hourly/minutely/secondly clock values through advanceClock, whose returned clock is total modulo unitsPerDay. Existing predicates depend only on those normalized clock fields. A modular increment visits at most M states before repeating, so M unsuccessful attempts suffice. Filtered-day shortcuts add interval multiples and normalize on the first step. Check a candidate before consuming the final permitted attempt.

Alternatives: constructor-only rejection cannot cover mutable DTSTART without API/state complexity; gcd shortens the bound but adds unnecessary code; UNTIL or wall-time cutoffs do not prove unreachability. Keep Go 1.16 because no new language/runtime feature is needed. No dependencies change, so supply-chain review does not apply.

Open questions: none before task authoring. Implementation must measure the one-second operation threshold and independently verify timeout cleanup, carry behavior and final-attempt matches. No runtime tests have yet been claimed for this mission.
