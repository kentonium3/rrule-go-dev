# Independent WP01 review — APPROVE

Date: 2026-10-25. Reviewer: gpt-6-astra (Reviewer Renata). Implementer: gpt-6.1-sol. Mission: until-scan-01M5S0WJ, WP01. Reviewed implementation: 4c9dba12904a9ef396c0eebae756bc9cb852d653. No blocking findings or required changes.

Review scope is the user's approved UTC-pointer-only optimization; every other location retains the original algorithm. Source and lifecycle were not edited. This report is advisory to the root orchestrator, which owns verdict capture/transitions.

## Context and provenance

Read full generated prompt 883fec01d1d744549cfafe4ada2281e4, spk-charter-profile-load and spk-run-review-wp, resolved reviewer-renata, loaded review charter context, read AGENTS.md then docs/CONTEXT.md then issue03, mission spec/plan/tasks/research/data-model/quickstart, research/period-bound-design.md and period-bound-review.md, and implementation handoff/logs. Code-review checklist lookup returned an exit-0 authored-section-absent advisory. An initial profile lookup for `reviewer` returned a normal unknown-profile error listing reviewer-renata; corrected to that activated profile. No actual Spec Kitty fault or workaround occurred. Every Spec Kitty invocation used SPEC_KITTY_NO_UPGRADE_CHECK=1; no upgrades.

Lane HEAD during validation was cb31718c68aba92b3bf99e74b16a3ef840fcb4f3, a mission status metadata merge. `git diff 4c9dba... HEAD -- rrule.go rrule_until_test.go go.mod go.sum` is empty. Thus tests exercised the exact reviewed implementation/test bytes. Working tree clean.

Independently reconstructed root Go files, go.mod and go.sum from merged baseline 886f07adf97d59c5348bd4e64c71ca616227304b into /tmp/rrule-until-independent-baseline, then added the exact test file from 4c9dba. Baseline rrule.go and lane pre-code b9e8e1097ebd9ccb2dab50cee18a230bfdd86aa0 both resolve to blob b3ea1cbbbee3c12558c2cb4f77589aa1920e92e1. Existing /tmp/rrule-until-baseline/rrule.go also byte-matches this baseline, independently verified.

## Correctness and requirement assessment

- FR-001 / NFR-001: the guard executes before filtering only within the empty-remain loop. UTC Jan 1 for YEARLY, month day 1 for MONTHLY, otherwise current civil day are lower bounds on candidate dates. WEEKLY calcDaySet starts at the cursor, including initial partial weeks, and extends forward. UTC firstyday/AddDate and timeset reconstruction cannot normalize to earlier dates. Strict After preserves equality. Both explicit short-horizon fixtures are exercised with NewRRule and StrToRRule and Iterator/All/After, alongside positive controls. Fixed ordinary operations measured 1.353–192.784 microseconds across the recorded harness, well under 100ms.
- FR-002: independently running exact committed TestUntilGolden on unfixed baseline passed, so expected sequences are actual baseline behavior. It covers all seven frequencies with before/equal/after endpoints; expanded next-year earlier-month and next-month earlier-day candidates; partial cross-year Sunday-start weeks; COUNT with buffered BYSETPOS results; empty dayset; differently offset UNTIL; New York gap/fold; Sao Paulo and Apia duplicate normalization; and fixed offsets including locations named UTC. Guard location comparison is pointer identity, with no heuristic. Per-candidate checks, buffered-result draining, existing reachability/overflow guards are unchanged. No exported API changed.
- FR-003: TestUntilSetters exercises fresh public iteration after shortened/extended UNTIL, DTSTART date/location switches and reset-to-default value. TestUntilEmptyStopsAtHorizon initializes the existing production iterator and invokes generate, observing actual exhaustion date after default, shortened, extended, DTSTART and zero-UNTIL resets. This is a production-path deterministic regression, not a synthetic output fixture. Independent baseline execution failed at 9999-12-31 versus expected 2318-04-18; final fixed targeted race passed. No default horizon or active-iterator guarantee changed.
- NFR-002: inspected genuine red log: all six BYSETPOS=2 paths exceed 100ms (141.995–152.012ms); all six heavy runs hit 3s watchdog and were killed/collected. CommandContext CombinedOutput waits for each child. Ordinary golden rows do not spawn children. The deterministic baseline failure supplements timing evidence.
- C-001/C-002/C-003: only rrule.go, rrule_until_test.go and issue03 changed in implementation commit. No dependency/minimum Go/API changes, license or README changes, arbitrary caps, new state, timezone machinery or unrelated repairs. Different-model review completed. No publication/push/master merge performed by reviewer. SC-001/002/003 supported by these checks.

## Commands and observed results

Toolchain: go1.25.14 linux/amd64. GOCACHE=/tmp/rrule-until-gocache for Go commands.

Independently executed:
- Lane `go test -race -run Until -count=1 ./...`: PASS, 19.382s, exact final implementation/test bytes.
- Independent baseline `go test -run '^TestUntilGolden$' -count=1 -v`: PASS, 0.002s, exact baseline output equality.
- Independent baseline `go test -run '^TestUntilEmptyStopsAtHorizon$' -count=1`: expected FAIL, 0.207s, exhaustion at MAXYEAR rather than effective default horizon.
- Lane `go test ./...`: PASS (cached).
- `git diff --check`: PASS; clean lane.
- `git log b9e8e109..4c9dba -- go.mod go.sum LICENSE README.md`: empty.

Inspected implementer evidence: full ordinary suite pass (~0.415s), full race pass 147.861s, vet/native build/GOARCH=386 build pass, gofmt output empty, red/green timing logs and final test refinement race log. Full race started before the final test-only refinement; my uncached targeted race run above closes that exact-test gap. Did not claim independent reruns of every build gate.

## Eight-item review checklist

1. Dead code: N/A — no new public function/module or production helper.
2. Synthetic-fixture test: PASS — tests invoke public constructors/query/iterator paths or actual production generate; baseline negative verified.
3. Silent empty return: PASS — the one new return documents exhaustion and sets existing len/finished state after proving no eligible candidate remains.
4. FR coverage: PASS — FR-001 public timing/positive controls, FR-002 baseline golden equality, FR-003 production exhaustion/setter controls.
5. Frozen surface: PASS — API, module files, license, README and prior safety changes preserved.
6. Locked decision: PASS — UTC pointer identity only; strict bound; unchanged non-UTC fallback; no prohibited machinery/caps.
7. Shared-file ownership: N/A — single WP owns the three changed files; later lane commits are lifecycle metadata.
8. Production fragility: PASS — no new panic/raise/error path.

## Limits and nonblocking observations

No cross-platform CI was executed by this reviewer. Existing CI uses Go 1.18/1.19 across Linux/macOS/Windows and a 3m race timeout. The measured full local race run leaves about 32 seconds of margin; the 18 subprocesses contribute roughly 18 seconds of race exit sleeps. The harness uses portable os/exec/context APIs, excludes process overhead from the 100ms operation budget, and gives a separate 3s watchdog. Slower or heavily loaded CI may need attention, but there is no observed failure warranting a speculative change. Root plans to inspect actual private PR matrix runs.

Non-UTC never-matching scans retain the documented MAXYEAR performance limitation. Fixed-zone/name lookalike tests establish compatibility, while the exact pointer gate itself establishes fallback selection. UTC sub-daily pruning permits conservative extra work through the final civil day. Default-horizon work has no universal 100ms guarantee. These are approved limits, accurately documented.
