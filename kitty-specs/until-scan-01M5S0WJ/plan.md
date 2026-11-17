# Implementation Plan: UTC UNTIL scan bound

**Branch**: `fix/until-scan` | **Date**: 2026-10-25 | **Spec**: [spec.md](spec.md)

## Summary

Stop empty scans once the current UTC period's earliest possible candidate exceeds effective UNTIL. Kent approved UTC-only pruning with unchanged fallback on 2026-10-25. Planning and consolidation target fix/until-scan; later private PR targets master, for Kent to merge.

## Technical Context

**Language/Version**: Go; module minimum 1.16 unchanged; installed toolchain recorded during tests.
**Primary Dependencies**: Standard library only; no dependency changes.
**Storage**: None.
**Testing**: Plain Go public-path regressions, isolated timing harness, exact baseline golden sequences; full tests, race, vet, build, 386 build and gofmt.
**Target Platform**: Existing supported Go platforms.
**Project Type**: Single Go library.
**Performance Goals**: Both UTC reported short-horizon cases below 100ms operation time, excluding startup/teardown.
**Constraints**: Unchanged exported API, constructor/setter acceptance, default UNTIL and prior safety fixes. No timezone-name heuristics, transition machinery or arbitrary cap.
**Scale/Scope**: Small guard in rrule.go; focused new test file and issue documentation.

## Charter Check

RRULE-001: smallest proved guard, unchanged public API/dependencies; approved scope narrowing avoids exotic timezone machinery. RRULE-002: red-first public paths and independent different-model review. RRULE-003: private branch/PR only, no upstream or master merge. RRULE-004: optional mission and narrowed scope explicitly approved; record and stop on Spec Kitty faults. Checks pass before/after design.

## Algorithm and proof

At the top of generate's outer `for remain.Len() == 0`, enable pruning only when `r.dtstart.Location() == time.UTC`. Construct midnight UTC of Jan 1 for YEARLY, current month day 1 for MONTHLY, current civil day for WEEKLY and finer frequencies. These are lower bounds on every possible current-period candidate and subsequent period; UTC has no timezone normalization. Strictly `periodStart.After(r.until)` means no eligible candidate can remain. Finish using existing total/finished/return semantics. Existing candidate comparisons remain authoritative; equality remains eligible. Check inside empty-buffer loop so buffered results are not discarded. Subdaily scans may continue through the final civil day; this is deliberate safe slack.

Use current iterator state and current rebuilt rule; no cached bound. Setters retain their existing semantics for fresh iterators. The existing default UNTIL stays unchanged. Comparing time instants supports UNTIL with a different offset. Non-UTC locations follow the existing path, including timezone normalization oddities. Research/period-bound-design.md and period-bound-review.md record the proof and counterexamples.

## Validation strategy

Record current-baseline red failures for both BYSETPOS fixtures using operation timing inside an isolated child. Parent deadline kills and collects the child; timeout must not leak work. Same-fixture BYSETPOS=1 controls exercise the same public harness. Include NewRRule/StrToRRule and Iterator/All/After as appropriate. Avoid timing-only claims: assert computed bounds for representative periods or use a far-future nonmatch watchdog/negative control to demonstrate horizon sensitivity without production instrumentation.

Capture exact baseline golden sequences for all seven frequencies: inclusive endpoints, YEARLY earlier-month and MONTHLY earlier-day candidates after advancement, weekly cross-year/nondefault week start, count/buffer, differing UNTIL offset. Preserve midnight DST/date-line normalization, normal gaps/folds, fixed offsets and a noncanonical location named UTC. Default/setter cases must use UTC and demonstrate shortened/extended/zero UNTIL and DTSTART location changes on fresh iterators. Non-UTC performance is not an acceptance requirement.

## Project Structure

rrule.go owns the guard; rrule_until_test.go owns focused tests; docs/issues/03-until-scan.md states exact coverage and measured results. Mission artifacts hold evidence. No new API contract: meta.json marks contracts none.

## Implementation Concern Map

### IC-01 — Safe UTC period exhaustion
- Purpose: bound empty scans without suppressing eligible candidates.
- Requirements: FR-001, FR-003, C-001.
- Surfaces: rrule.go.
- Dependencies: none.
- Risk: using raw cursor for yearly/monthly or terminating buffered output.

### IC-02 — Defect-sensitive evidence and compatibility
- Purpose: prove both speed repair and unchanged occurrence sequences.
- Requirements: FR-002, NFR-001, NFR-002, C-002, C-003 and all success criteria.
- Surfaces: rrule_until_test.go, docs/issues/03-until-scan.md and mission evidence.
- Dependencies: baseline evidence before IC-01 implementation.
- Risk: flaky wall time, leaked child processes or silently claiming non-UTC pruning.

## Engineering alignment

Confirmed conservative bound design, then explicitly selected UTC-only fallback option. No open questions or unjustified charter violations. One implementation package with gpt-6.1-sol implementation and gpt-6-astra review is sufficient.
