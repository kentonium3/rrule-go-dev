# Implementation Plan: Bound unreachable sub-daily recurrence iteration

**Branch**: `fix/unreachable-byset-hang` | **Date**: 2026-10-25 | **Spec**: [spec.md](spec.md)

## Summary

Bound each existing hourly, minutely and secondly inner clock-filter search by the number of distinct clock states in a day. If no applicable clock filter matches within that bound, finish the iterator normally. Keep constructors, setters, clock advancement and calendar filtering unchanged. Tests demonstrate prompt termination and preserve exact reachable sequences.

## Engineering Alignment

Kent confirmed this design in decision `01M5RZHPBXMKMK5YBHQB6R8BTR`. Current checkout, planning/base branch and mission consolidation target are `fix/unreachable-byset-hang`; they match. The later private PR targets `master`; this plan does not authorize merging it. No further design questions remain.

## Technical Context

**Language/Version**: Go; module minimum remains Go 1.16. Observed development toolchain: Go 1.25.14 linux/amd64.
**Primary Dependencies**: Standard library only; no dependency changes.
**Storage**: Existing in-memory rule and iterator state; no persistent state.
**Testing**: Public-constructor/query regression tests, isolated child-process timeouts, exact reachable sequences, race/vet/gofmt and linux/386 checks.
**Target Platform**: Existing Go targets, including Linux, macOS and Windows CI.
**Project Type**: Single Go library module.
**Performance Goals**: Each reported unreachable fixture finishes evaluation within one second; each inner clock search tries at most 24, 1440 or 86400 candidates respectively.
**Constraints**: Preserve all exported signatures, correct sequences and prior overflow guards. Keep development private. Different models implement and review each work package.
**Scale/Scope**: Three sub-daily loops in rrule.go, focused regression tests and issue 02 documentation. No general recurrence solver or UNTIL scan optimization.

## Charter Check

Pre-design and post-design: PASS. Small root-cause fix, no new dependency or public API, no arbitrary interval caps. Go 1.16-compatible loop counters and subprocess test facilities suffice; no minimum-version increase needed. Public-path red/green evidence and independent review are required during implementation, not claimed complete by this plan. Preserve MIT and README notice. Stop and report Spec Kitty faults under current project policy; the previous mission's unattended-run exception is not silently extended.

## Project Structure

- `rrule.go`: bound the existing three inner searches, retaining advanceClock and existing calendar handling.
- `rrule_reachability_test.go`: public-path fixtures, subprocess harness, setter transitions and exact compatibility controls.
- `docs/issues/02-unreachable-byset-hang.md`: document outcome, conditions, verification and exclusions.
- This mission directory: plan, research, data-model, quickstart, reasons canvas and later review/evidence records.

## Design

### Complete clock-cycle bound

Let M be 24 for hours, 1440 for minutes of day, or 86400 for seconds of day. Existing advancement follows x -> (x + interval) mod M. For a positive accepted interval, the sequence repeats after M/gcd(M, interval) steps, which is at most M. Therefore M failed clock-filter candidates prove that this inner search will never succeed. Use a local attempt counter bounded by M; do not add a gcd helper or precomputed reachability state.

Retain each filtered-day shortcut before entering the bounded search. It advances by a multiple of the interval and may leave a clock component beyond one day's range. The existing advanceClock normalizes that value on the first step, so the same at-most-M argument applies. Do not compare raw unnormalized clock fields as a cycle sentinel.

For each attempt, first call advanceClock and respect its horizon/overflow failure; then assign normalized clock fields and test the existing combined BY predicate. A match on the final allowed attempt must still succeed. Only after every candidate fails should the iterator record r.len = iterator.total, mark finished and return. This mirrors existing exhaustion behavior and preserves buffered occurrences. Do not clear remain, bypass a successful match or alter count handling.

This bound concerns clock predicates only. If a clock match exists but calendar filters never match, the existing outer scan remains; that is explicitly outside scope. No wall-clock deadline or arbitrary maximum result count is added to production code.

### Constructor and setter behavior

Keep constructor acceptance unchanged. No new errors, fields or caches. Each fresh iterator executes the bounded search using current rule state. DTStart and Until retain their current rebuild behavior and signatures. Tests cover reachable -> unreachable -> reachable transitions, Until rebuilding after explicit DTSTART, and both constructors. Mutation concurrent with an active iterator gains no new guarantee.

### Regression harness

Use a named test-binary child selected by environment variable and `-test.run`, invoked via os.Executable and exec.CommandContext. Each fixture exercises a real public constructor plus After, All or Iterator as appropriate. The parent must kill and Wait for a timed-out child and retain diagnostic output, avoiding leaked spinning goroutines. Keep race-test process teardown delay out of the one-second operation measurement: assert elapsed operation time inside the child and use a separate bounded parent watchdog for startup/teardown overhead. A pre-fix stuck operation must trip that watchdog and fail, not be accepted as an expected passing result.

Record failures independently for HOURLY, MINUTELY and SECONDLY on the merged-master baseline. Pair negative cases with reachable controls using the same fixture/harness; test missing and mixed filters, jointly impossible combinations, explicit UNTIL, day-filter shortcuts, midnight/year carry and a non-UTC location. Include a reachable case whose match occurs on the final candidate in a complete cycle to catch an off-by-one bound. Assert exact timestamps and repeated exhaustion, not merely completion.

Retain prior huge-interval regressions and add targeted overlap checks where clock-cycle termination meets MAXYEAR exhaustion. Use small finite COUNT on reachable cases and explicit bounded compatibility datasets; COUNT alone does not limit a stuck Next call.

### Alternatives considered

- Constructor rejection: rejected because DTSTART setters can change reachability and cannot return a new error without breaking APIs; duplicating checks adds state and paths.
- Exact gcd-based cycle length: mathematically valid but unnecessary machinery for maxima of 24/1440/86400. Reconsider only if measured tests miss the one-second bound.
- UNTIL checks inside the search: insufficient when UNTIL is absent/distant, and expands calendar semantics.
- Arbitrary timeout or iteration cap: rejected; use the complete finite clock-state bound so reachable schedules are never truncated.

## Validation and Risks

1. Record toolchain and baseline tests; add public-path tests before production changes and capture all three families red.
2. Apply bounds and prove targeted green plus exact reachable controls. Remove each frequency's bound in disposable scratch copies and confirm its regression fails independently.
3. Run go test ./..., go test -race ./..., go vet ./..., gofmt -l on changed Go files, go build ./..., and GOARCH=386 go test ./....
4. Independently review the modular-cycle proof, final-attempt match, filtered-day shortcuts, combined predicates and buffered-result preservation. Record implementing/reviewing models and date.
5. Record measured operation times and clean child termination. Full fuzz expansion is not required by this mission; existing fuzz tests remain intact.

Risk: terminating before the last candidate could drop reachable dates; final-attempt controls address it. Risk: modifying normalization changes DST/carry semantics; retain advanceClock unchanged and compare exact sequences. Risk: timeouts leave CPU-consuming children or race overhead causes false failures; use cleanup and distinguish operation time from watchdog time. These are realistic integration risks in the three touched loops, not hypothetical new subsystems.

## Implementation Concern Map

### IC-01 — Bounded clock search and public-path regressions

- Purpose: prompt normal exhaustion while preserving correct results and setter behavior.
- Requirements: FR-001, FR-002, FR-003, NFR-001, NFR-002, C-001, C-002, C-003.
- Surfaces: rrule.go and rrule_reachability_test.go.
- Dependencies: none.
- Risk: shared production file; keep one implementation owner and verify each frequency independently.

### IC-02 — Verification and private handoff evidence

- Purpose: make the patch reviewable with reproducible red/green and compatibility evidence.
- Requirements: FR-003, NFR-001, NFR-002, C-003, C-004.
- Surfaces: issue 02 notes and mission evidence/reviews.
- Dependencies: evidence reflects IC-01's final patch.
- Risk: distinguish measured results from planned checks and keep unrelated scan limits explicit.

## Complexity Tracking

No charter deviations. A local finite loop bound adds no exported API, dependency, persistent state or solver. Concern labels are planning boundaries; executable work packages are authored in the tasks phase.
