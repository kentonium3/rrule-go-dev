---
work_package_id: WP01
title: Safe iterator arithmetic and masks
dependencies: []
requirement_refs:
- FR-001
- FR-002
- FR-003
- FR-004
- FR-005
- FR-006
- NFR-002
- NFR-003
- C-001
- C-002
- C-003
- C-005
- C-006
- C-007
- SC-001
- SC-002
- SC-003
planning_base_branch: fix/iterator-panics
merge_target_branch: fix/iterator-panics
branch_strategy: Planning artifacts for this mission were generated on fix/iterator-panics. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into fix/iterator-panics unless the human explicitly redirects the landing branch.
base_branch: kitty/mission-iterator-panics-01M5RT9D
base_commit: 684315d89987d127f4ac9a1e88dac68e6a57a627
created_at: '2026-10-25T06:17:59.942372+00:00'
subtasks:
- T001
- T002
- T003
- T004
- T005
- T006
- T007
history: []
agent_profile: implementer-ivan
authoritative_surface: rrule.go
create_intent:
- rrule_safety_test.go
execution_mode: code_change
owned_files:
- rrule.go
- rrule_safety_test.go
- docs/issues/01-iterator-panics.md
role: implementer
tags: []
tracker_refs: []
---

## ⚡ Do This First: Load Agent Profile

Load implementer-ivan using the spk-charter-profile-load skill before coding.
Read AGENTS.md, docs/CONTEXT.md and the mission spec, plan and contracts.
This is an authorized autonomous run. Ordinary implementation choices are delegated.
D-3: capture faults before changing state, diagnose, prove and log safe recovery.
No master merge, public posts, upstream pushes or new dependencies.

## Branch Strategy

Planning/base and consolidation target: fix/iterator-panics. The final private
PR targets master. Use the workspace allocated by the runtime from lanes.json;
never invent a worktree/base or implement in the primary checkout. The orchestrator
owns workflow transitions. Implementation agent is codex/gpt-6.1-sol; review agent
is codex/gpt-6-astra. Record actual model and date in a durable handoff file.

## Shared verification contract

Use the real NewRRule and StrToRRule paths followed by Iterator, not helper-only
unit tests. Do not recover panics as passing outcomes. Invalid-input refusal or
empty matches must have working same-fixture positive controls. Never claim a
bounded output count guarantees the next iterator call returns.

Use GOCACHE=/tmp/rrule-go-cache if default cache writes are restricted. Bound
commands with timeout; record exact command and exit code. Poll a long gate until
completion; never rerun a mutation just because output contains prose before JSON.

Write handoff evidence to /tmp with the WP id, give its absolute path to the
orchestrator, and keep source/docs edits in your owned files. The orchestrator
persists evidence in the mission. Do not alter other WP files or workflow state.
Commit only owned files via safe-commit, confirming the actual commit contents.

The code is private. Keep all reproductions and logs local/private. Preserve MIT.

## Objective

Eliminate the three reported panic families and impossible interval results with
small root fixes. Preserve correct rules and all public signatures. This WP owns
all production edits to rrule.go; WP02 adds fuzzing after your review passes.

## T001 — Baseline and red tests

Run the existing suite first. Create rrule_safety_test.go with a constructor
adapter exercising equivalent ROption and RRULE text forms. Fix DTSTART for
reproductions and request enough occurrences to reach advancement, not just first
emission. Test each failure family independently before its code fix. Use finite
command timeouts; record logs and exit status. On unfixed ordinal cases Iterator()
itself may panic, so the test must include iterator construction.

Cases include MONTHLY -6MO/8MO/53MO, YEARLY BYMONTH=1 -6MO and BYMONTH=12 8MO;
YEARLY BYEASTER -100/300; MONTHLY huge INTERVAL. Also test the DAILY backwards and
HOURLY fabricated output controls. Do not let the first panic prevent collecting
red evidence for the other independent cases; run focused subtests separately.

## T002 — Ordinal indexing

Validate the preliminary candidate in [first,last] before either wdaymask read.
Retain the final range check after weekday adjustment. n is already bounded ±53.
Preserve empty impossible ordinals, mixed valid+impossible lists, YEARLY ±53
without BYMONTH, and WEEKLY's ignored ordinal semantics. A global ±5 rejection
is forbidden because it changes currently correct mixed/empty rules.

## T003 — Easter indexing

Compare offset with -eyday and len(eastermask)-eyday before adding it. A bounds
check after addition is too late for machine extremes. Keep the existing mask
tail and skip candidates beyond it. Test MinInt/MaxInt and mixed {0,MaxInt} cases.
Preserve BYEASTER257's current YEARLY spillover-drop as an explicit known limit,
with a case and nearby valid 257 result in issue 01 fix notes. Do not expand into
cross-year Easter semantic redesign. Ensure case years really exercise boundaries.

## T004 — Overflow-safe interval advancement

Avoid constructor interval caps. A huge interval may correctly produce DTSTART
then exhaust; on 32-bit MaxInt seconds is only about 68 years and can produce
several occurrences. Preserve that behavior.

YEARLY/MONTHLY: compare remaining calendar years/months before adding. MONTHLY
normalization must preserve month alignment. WEEKLY: compare before *7 and
retain WKST subtraction/alignment. DAILY: bound addition before normalizing days.
Subdaily: decompose carries or use bounded int64 arithmetic after comparing the
interval against the remaining civil calendar horizon. Merely casting MaxInt to
int64 is not enough on 64-bit. Do not use time.Duration for the multi-century
horizon; it saturates around 292 years. UTC civil Unix seconds can express it
without DST elapsed-time drift, but preserve the original wall-clock semantics.

Check horizon on EVERY iteration of the subdaily BY-filter loops, because day
carry accumulates before normalization. Avoid adding the separate modular
reachability solver. Preserve already-buffered valid occurrences on exhaustion.
A small private helper is acceptable if it reduces repeated safe arithmetic;
no exported API, generic framework or broad refactor.

## T005 — Compatibility and width controls

Golden checks: ±5 monthly ordinals; valid YEARLY ±53 without BYMONTH; WEEKLY
ordinals; BYEASTER=-80/257 and ordinary offsets; ordinary intervals and large
valid intervals; mixed lists retain their valid entries. Include starts near
9999, all seven frequencies, COUNT=1 and >1, interval zero normalization and
existing negative rejection. Include a real 32-bit-compatible 2147483647 SECONDLY
case with independently calculated expected dates, not the implementation reused
as oracle. Exercise subdaily filters with multiple carries across the horizon.

## T006 — Sensitivity and validation

For each independent production fix, demonstrate its corresponding regression
fails if that fix alone is absent (disposable scratch copy or equivalent precise
mutation). Do not mutate the active lane while a reviewer is inspecting it.
Run timeout 180 go test ./..., timeout 180 go test -race ./..., timeout 180 go vet
./..., gofmt -l ., and GOARCH=386 timeout 180 go test ./.... Formatting must have
no output. Record compiler/test identity and the baseline. Go test compiles typed
sources; also run go build ./... as the explicit compile gate.

## T007 — Handoff

Commit only the three owned paths. Report commit hash, change rationale,
red/green/sensitivity evidence, exact validation outputs, limitations and actual
model/date. Include any targeted out-of-map need before editing. No workflow
approval or master merge by implementer; parent claims review independently.

## Definition of Done

All safety regressions pass, correct-occurrence controls stay exact, all declared
gates pass, no extra dependency/API change, known limit documented, full review
evidence available. A no-op or half-fix must not pass the defect suite.

## Reviewer guidance

Inspect each arithmetic expression BEFORE normalization, especially weekly
multiplication, 64-bit extreme addition and 32-bit carries. Check UTC civil-time
reasoning and original DST behavior, buffered results at exhaustion, and actual
public entrypoint coverage. Confirm red evidence exists for each root fix and
large valid intervals were not rejected. Do not approve from aggregate pass alone.
