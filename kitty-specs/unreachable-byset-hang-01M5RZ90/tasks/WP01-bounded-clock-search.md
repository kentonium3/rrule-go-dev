---
work_package_id: WP01
title: Bounded clock search and regression evidence
dependencies: []
requirement_refs:
- FR-001
- FR-002
- FR-003
- NFR-001
- NFR-002
- C-001
- C-002
- C-003
- C-004
- SC-001
- SC-002
- SC-003
planning_base_branch: fix/unreachable-byset-hang
merge_target_branch: fix/unreachable-byset-hang
branch_strategy: Planning artifacts for this mission were generated on fix/unreachable-byset-hang. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into fix/unreachable-byset-hang unless the human explicitly redirects the landing branch.
base_branch: kitty/mission-unreachable-byset-hang-01M5RZ90
base_commit: 8cefe32384904930db48deaa1443592a4ae60bd4
created_at: '2026-10-25T07:39:08.481471+00:00'
subtasks:
- T001
- T002
- T003
- T004
- T005
history: []
agent_profile: implementer-ivan
authoritative_surface: rrule.go
create_intent:
- rrule_reachability_test.go
execution_mode: code_change
owned_files:
- rrule.go
- rrule_reachability_test.go
- docs/issues/02-unreachable-byset-hang.md
role: implementer
tags: []
tracker_refs: []
---

## ⚡ Do This First: Load Agent Profile

Load implementer-ivan through spk-charter-profile-load, then read AGENTS.md, docs/CONTEXT.md, docs/issues/02-unreachable-byset-hang.md and this mission's spec/plan/research/data-model/quickstart. Prefix CLI commands SPEC_KITTY_NO_UPGRADE_CHECK=1. Do not upgrade Spec Kitty. Stop and report tooling faults with exact output; do not silently extend the previous mission's D-3 exception.

## Objective

Deliver FR-001/FR-002 promptly without violating FR-003. Follow the confirmed plan: local attempt bounds of 24/1440/86400 in the existing three inner searches, unchanged constructor acceptance and setters. No public API, dependency, Go version, unrelated calendar semantics or general UNTIL scan changes.

## Branch Strategy

Use only the runtime-allocated lane worktree from lanes.json. Consolidation target fix/unreachable-byset-hang; later private PR targets master. Do not merge master, push, create a PR or change mission state; orchestrator owns those actions. Implementation model gpt-6.1-sol, independent reviewer gpt-6-astra; record actual identities and date in handoff.

## T001 — Baseline and red tests

Record go version and baseline tests. Write public-path fixtures using NewRRule and StrToRRule and After/All (Iterator where relevant). Original reported cases: hourly interval2 BYHOUR1, minutely interval2 BYMINUTE1, secondly interval2 BYSECOND1 at 2026-01-05T10:00:00Z, plus explicit UNTIL hourly. Isolate stuck calls in the test binary as a child process using os.Executable/exec.CommandContext. Kill and Wait on timeout; no leaked goroutines or children. Measure the actual operation inside the child against one second and use a separate bounded parent watchdog to accommodate startup and race teardown. Record each frequency red before production edits.

## T002 — Bounded search

Read existing generate/advanceClock/remaining-buffer code first. M consecutive failed candidates suffice because clock residues repeat within M steps. Preserve advanceClock before assignment and predicate; allow a match on the final candidate. Retain filtered-day shortcuts and their first-step normalization. On exhausting unsuccessful attempts, set len to total and finished true, return without clearing buffered results. Preserve huge interval/MAXYEAR protections.

## T003 — Compatibility and setters

Use same-fixture positive controls for each negative case. Exact hourly control begins 2026-01-06T02:00:00Z for interval2 BYHOUR2. Add minute/second controls, joint unreachability, mixed lists, missing filters, midnight/year carry, filtered-day shortcuts, non-UTC location, final-candidate cycle match and buffered occurrence preservation. Assert order and repeated exhaustion. Test DTSTART reachable-to-unreachable-to-reachable and UNTIL rebuilding on fresh iterators, both constructors. No concurrent active-iterator mutation guarantee.

## T004 — Verification

Run targeted green, full go test ./..., go test -race ./..., go vet ./..., go build ./..., GOARCH=386 go test ./..., and gofmt -l on changed Go files. GOCACHE may be /tmp/rrule-go-cache. Record commands/exits. Independently disable each frequency bound in disposable scratch copies and demonstrate its regression fails. Do not alter tracked checkout for mutation runs. Timing evidence must report measured operation time, not claim parent watchdog equals one second. Existing fuzz coverage remains unchanged; no new fuzz campaign required here.

## T005 — Handoff

Update issue 02 notes with outcome, reproduction conditions, likelihood, measured validation and separate general-calendar-scan limitation. Write detailed evidence and log paths under /tmp for orchestrator persistence; code_change WP cannot own kitty-specs files. Use spec-kitty safe-commit for the three owned paths only. Verify actual commit diff. Do not change other files or lifecycle state. Return commit SHA, evidence path, test results and any concerns.

## Definition of Done and reviewer guidance

All FR/NFR/C/SC refs satisfied; clean formatting and compiler checks; red/green evidence for each frequency and exact controls; no API/dependency drift. Reviewer independently inspects modular proof, last-attempt off-by-one, filtered-day carry, constructor/setter wiring and child cleanup. Different-model review is required before acceptance, and no self-approval is allowed.
