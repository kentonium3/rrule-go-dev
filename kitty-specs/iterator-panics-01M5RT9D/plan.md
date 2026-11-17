# Implementation Plan: Iterator panic safety

**Branch**: `fix/iterator-panics` | **Date**: 2026-10-25 | **Spec**: [spec.md](spec.md)

## Summary

Repair the existing iterator at its arithmetic and indexing boundaries. Keep
constructor acceptance where safe iteration is possible, preserving mixed valid
and impossible selections. Avoid arbitrary interval caps. Use overflow-safe
advancement that stops beyond MAXYEAR, and retain the established calendar and
wall-clock semantics for valid rules. Add defect-sensitive regressions first,
then a separately reviewed bounded fuzz target.

## Technical Context

**Language/Version**: Go; go.mod remains 1.16, development toolchain Go 1.25.14.
**Primary Dependencies**: Standard library only; no new dependencies.
**Storage**: In-memory masks and iterator state; no persistence changes.
**Testing**: Go table-driven regressions, golden compatibility, race/vet/gofmt,
32-bit tests, and a Go 1.18+ build-constrained fuzz target with committed seeds.
**Target Platform**: Existing Go targets; verify linux/amd64 and GOARCH=386.
**Project Type**: Single Go library module.
**Performance Goals**: No arithmetic overflow or out-of-range access in the three
failure families; bounded verification, no new unbounded search mechanism.
**Constraints**: Exported API and all currently correct occurrences unchanged;
no master merge or public posting; distinct implementation/review models.
**Scale/Scope**: rrule.go arithmetic/masks, focused new tests and fuzz file, issue
notes and validation evidence. No new generalized recurrence engine.

## Engineering alignment

Kent confirmed the spec and delegated ordinary design decisions for the unattended
run. Current/base/consolidation branch is fix/iterator-panics; it matches the
workflow target. Master is the later private PR target, not the consolidation
target. Design decisions below are within that authorization.

## Charter Check

Pass: root fixes only, no dependency or exported-signature changes, existing
correct sequences protected, regression evidence before fixes, independent
model review, private delivery. D-3 is authorized and documented in the run
adapter. No minimum-version increase is necessary merely to add fuzz tests:
use both `//go:build go1.18` and legacy `// +build go1.18` on the fuzz file.
If testing demonstrates a genuine incompatibility, record the reason before
changing go.mod; permission is already granted for a necessary fuzzing increase.

## Project Structure

- rrule.go: ordinal/Easter mask generation and frequency advancement.
- rrule_safety_test.go: red-green reproductions and nearby exact controls.
- rrule_fuzz_test.go: bounded Go 1.18+ fuzz target and deterministic seeds.
- docs/issues/01-iterator-panics.md: fix disposition and known limits.
- docs/repro/01-iterator-panics/: existing reproduction stays available.
- Mission evidence/reviews: planning-authority artifacts; not code-WP ownership.

## Design

### Ordinal weekday indexing

Compute the candidate index, prove it lies within the selected period before
reading wdaymask, then perform weekday adjustment and the existing final range
check. Keep ±53 validation and current YEARLY/WEEKLY semantics. Impossible
ordinals match nothing; mixed valid/impossible lists keep valid occurrences.
Do not globally reject ±6/±7 merely to avoid the separate scan limitation.

### Easter offsets

Check offset bounds before adding the Easter day index, using comparisons that
cannot overflow for MinInt/MaxInt. Skip candidates outside the existing mask.
Preserve in-mask offsets and existing weekly spill behavior. Retain and document
the known BYEASTER=257 silent drop when the indexed day is outside the year's
emitted dayset; correcting cross-year Easter semantics is a broader behavior
change than the panic repair. Test the known limit and valid 257 years explicitly.

### Interval advancement

Use per-frequency bounds and quotient/remainder arithmetic so addition and
multiplication never overflow before the bound check. YEARLY/MONTHLY compare
remaining supported years/months before addition. WEEKLY checks before *7,
retaining week-start alignment. DAILY and subdaily carries must be bounded by
remaining supported calendar days, without a time.Duration conversion that would
truncate the horizon to ~292 years. Preserve wall-clock calendar behavior.

For HOURLY/MINUTELY/SECONDLY, split large intervals into bounded remainders and
carried days rather than adding MaxInt to clock components. Check carried days
on every inner BY-filter iteration before accumulating/advancing. Do not add a
reachability solver for the separate hang mission. A zero-day unreachable loop
may remain a known limit, so fuzzing cannot blindly enter it. If the next step
is beyond MAXYEAR, finish normally after any currently buffered occurrences.

Exercise the actual constructor/iterator path, COUNT=1 and COUNT>1, boundary
starts, huge increments, and 32-bit MaxInt seconds/minutes that are still valid.
Do not reject values solely on size or change correct result ordering.

### Fuzz verification

Run no fewer than 60 seconds on a recent toolchain. Cover every BY field via
small lists and both valid/invalid scalar boundaries, including explicit extreme
BYEASTER/INTERVAL values. Bound the calendar horizon near MAXYEAR for calendar
cases, requested count, Cartesian list sizes, and subdaily search combinations.
Subdaily filters must be constructed to be reachable or omitted; do not hide
invalid input crashes behind blanket recover. A finite COUNT does not bound a
single Next call; use a whole-command timeout and document domain restrictions.
Keep deterministic unbounded-domain regressions separately so near-MAXYEAR fuzz
normalization cannot erase early-year defects. Public constructor paths must be
exercised, including text serialization where representable.

## Verification and evidence

1. Record baseline full test pass and toolchain.
2. Add tests and show each failure family independently red on unfixed code.
3. Fix one root cause at a time; show targeted green and valid golden controls.
4. Verify each independent change matters (revert/mutate in disposable scratch
   copies, never disturb the production checkout); preserve evidence summaries.
5. Run full tests, race, vet, formatting, 32-bit tests, and at least 60s fuzzing.
6. Independent reviewer checks arithmetic, public-path coverage, byte-exact
   sequences, known-limit honesty, and evidence. Keep model/date records outside
   the known lossy approval-feedback path.

## Implementation Concern Map

### IC-01 — Iterator arithmetic and masks
- Purpose: eliminate three crash families and overflow-generated dates.
- Requirements: FR-001, FR-002, FR-003, FR-004, FR-005, FR-006.
- Surfaces: rrule.go, rrule_safety_test.go, issue 01 notes.
- Dependencies: none.
- Risk: shared production file; keep one owner and independently verify each fix.

### IC-02 — Bounded fuzz evidence
- Purpose: exercise accepted option domains without hanging on a separate bug.
- Requirements: FR-005, NFR-001, NFR-003; reinforces FR-001 through FR-004.
- Surfaces: rrule_fuzz_test.go only.
- Dependencies: IC-01, so known panic seeds become passing regression seeds.
- Risk: COUNT/UNTIL alone do not bound Next; domain design must prove tractability.

## Complexity Tracking

No charter deviations. Two sequential WPs avoid overlapping ownership of rrule.go.
Fuzzing is a separate reviewable concern; no generic arithmetic framework or new
library dependency is justified.
