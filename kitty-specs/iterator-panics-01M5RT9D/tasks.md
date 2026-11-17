# Work packages: iterator panics

Planning/consolidation branch: fix/iterator-panics. Private PR target: master.
Two sequential code WPs with disjoint ownership. No parallel production edits.

## WP01 — Safe iterator arithmetic and masks

Priority P1. Dependencies: none. Prompt: tasks/WP01-safe-iterator.md.
Owns rrule.go, rrule_safety_test.go and docs/issues/01-iterator-panics.md.
Independent test: each original panic family fails before its fix and then passes
through both public constructors, while exact valid occurrences stay unchanged.

T001 Establish baseline and add independently failing public-path regressions (WP01)
T002 Guard ordinal weekday mask reads without narrowing accepted rules (WP01)
T003 Guard Easter offset addition/indexing and document spillover limit (WP01)
T004 Make interval advancement overflow-safe across all frequencies (WP01)
T005 Add exact compatibility, boundary and 32-bit controls (WP01)
T006 Verify independent fix sensitivity and run full validation (WP01)
T007 Commit focused code/docs and provide review evidence (WP01)

Risks: buffer preservation at exhaustion, weekly WKST alignment, subdaily carry
accumulation, masking a failing family with another. No blanket interval cap.

## WP02 — Bounded fuzz coverage

Priority P1. Dependencies: WP01. Prompt: tasks/WP02-bounded-fuzz.md.
Owns rrule_fuzz_test.go. Independent test: deterministic seeds and a 60-second
fuzz run exercise both constructors and the iterator without panic.

T008 Design tractable fuzz domains and committed defect seeds (WP02)
T009 Implement build-constrained FuzzIterate through public entry points (WP02)
T010 Prove seed sensitivity against unfixed code and run bounded fuzz campaign (WP02)
T011 Run integrated checks and record scope limits/model attribution (WP02)

Risks: COUNT/UNTIL do not bound Next; filter reachability, BYSETPOS and sparse
calendar selections can cause excessive scans. Domain restrictions must be explicit.

## Delivery sequence

Post-plan independent review -> tasks finalization -> analysis gate -> WP01
implementation/review -> WP02 implementation/review -> acceptance -> consolidation
onto fix/iterator-panics -> integrated independent review -> private PR. No master
merge. A task is complete only through event-sourced mark-status, not this file.
