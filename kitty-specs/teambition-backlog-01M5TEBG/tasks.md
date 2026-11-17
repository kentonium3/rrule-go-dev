# Work packages: teambition backlog

Planning/consolidation target: fix/teambition-backlog. Final private PR targets main, unmerged.
Exactly two serial code WPs, one per scoped public issue5/6. Contributor issue7 is baseline76c646d, not assessed or extended.
Post-plan re-review APPROVED gpt-6-astra2026-10-25 at3b776c5; private metadata/v1.9.0 explicit user decision.

## WP01 — RFC civil time and explicit DTSTART

Priority P1; dependencies:none. Prompt:tasks/WP01-civil-time.md.
Own rrule.go/civil_time.go/DST tests/golden driver+DST corpus/release note5.
Independent test: generated gaps/folds and explicit gap/laterfold anchors match RFC/authoritative instants on all public evaluation paths, preserving bounded BYSETPOS.

T001 Establish red civil-time and DTSTART production-path regressions (WP01)
T002 Implement bounded civil resolution with exact actual location (WP01)
T003 Add immutable metadata/nominal anchor foundation and narrow DTSTART parser wiring (WP01)
T004 Integrate candidate validity and efficient BYSETPOS ranking (WP01)
T005 Commit reusable golden harness, baseline and provenance corpus (WP01)
T006 Verify green, minimum-Go, compatibility and dense positional resource controls (WP01)
T007 Document result changes/source break/prior art and independent handoff (WP01)

## WP02 — RFC DTSTART UNTIL forms and integrated compatibility

Priority P2; dependencies:WP01. Prompt:tasks/WP02-time-forms.md.
Own str.go/rruleset.go/form tests/forms corpus/golden usage/release note6.
Independent test: every forbidden form rejects with paired valid forms accepted; floating/DATE roundtrip and exact supplied custom location survive copy/Set/setters.

T008 Establish red form-validation and semantic round-trip regressions (WP02)
T009 Collect lexical descriptors and validate explicit DTSTART UNTIL forms (WP02)
T010 Retain descriptors across shared Set DTSTART transfers and validate every rule line (WP02)
T011 Serialize forms and enforce immutable per-part copy/replacement semantics (WP02)
T012 Extend golden corpus and classify integrated compatibility differences (WP02)
T013 Run full modern/race/vet/format/fuzz/minimum-Go and golden gates (WP02)
T014 Document all rejection/serialization/source changes and independent handoff (WP02)

## Ownership and sequencing safeguard

Declared owned_files are disjoint. No concurrent production edits; WP02 claim waits for WP01 approved/done.
WP01 has one planned narrow out-of-map str.go adapter for explicit gap/fold DTSTART parsing; record exact rationale. WP02 owns the file and starts later.
WP02 may need a minimal metadata adapter in WP01 foundation or golden driver; likewise record exact rationale rather than broad shared ownership. Do not move whole parser toWP01.
Code WPs write handoff evidence to/tmp; root persists kitty-specs evidence/reviews and manages workflow. No fourth/third integrationWP.

## Requirement ownership and delivery

WP01 owns FR001/002 and shared FR007/008; WP02 owns FR003/004/005 and shared FR007/008. Both map NFR001/002/003 and C001/002/003; WP01 SC001, WP02 SC002/004/005.
Retired FR006/SC003 are not mapped or reused. Contributor duplicate behavior must match baseline without new mission work.
Tasks -> analyze -> WP01 implement/review -> WP02 implement/review -> strictaccept -> consolidate fixbranch -> integrated final independent review -> privatePRmain; no finalPRmerge/publicaction.
Subtasks are event-sourced with mark-status; reference rows above are not checkboxes.
