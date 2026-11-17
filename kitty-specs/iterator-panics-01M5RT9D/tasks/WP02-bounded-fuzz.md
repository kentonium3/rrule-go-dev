---
work_package_id: WP02
title: Bounded fuzz coverage
dependencies:
- WP01
requirement_refs:
- FR-005
- NFR-001
- NFR-002
- NFR-003
- C-001
- C-002
- C-003
- C-004
- C-005
- C-006
- C-007
- SC-004
- SC-005
planning_base_branch: fix/iterator-panics
merge_target_branch: fix/iterator-panics
branch_strategy: Planning artifacts for this mission were generated on fix/iterator-panics. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into fix/iterator-panics unless the human explicitly redirects the landing branch.
base_branch: kitty/mission-iterator-panics-01M5RT9D
base_commit: f9b0f1ac5393aee0405674e145f5b53367272a99
created_at: '2026-10-25T06:28:32.634821+00:00'
subtasks:
- T008
- T009
- T010
- T011
history: []
agent_profile: implementer-ivan
authoritative_surface: rrule_fuzz_test.go
create_intent:
- rrule_fuzz_test.go
execution_mode: code_change
owned_files:
- rrule_fuzz_test.go
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

Add FuzzIterate, committed deterministic seeds, and honest bounded-domain panic
coverage, after WP01 is approved. Production changes belong to WP01 and require
coordinated rework, not an unreviewed edit in this WP.

## T008 — Fuzz domain design

Cover every BY* field using bounded list lengths and positive/negative/invalid
boundaries, plus ordinals, Easter offsets and interval extremes. Include seeds
that exercise each fixed defect. Bound list Cartesian expansion. Use starts near
MAXYEAR for impossible calendar scans, with separate deterministic regressions
from WP01 retaining original-year cases.

A single Next call can hang even with COUNT and UNTIL. For HOURLY omit BYHOUR;
for MINUTELY omit BYHOUR/BYMINUTE; for SECONDLY omit all three time filters.
Lower-order expanding lists may be tested in bounded form. Restrict subdaily
BYSETPOS to empty or a reachable position (±1) so it cannot force billions of
second-by-second scans. Calendar filters combined with a small subdaily interval
also need a proven finite cheap bound; separate calendar and subdaily modes or
choose sufficiently late starts. Explain those choices in comments and evidence.
Do not silently drop a discovered panic or claim the excluded hang family fixed.

## T009 — Public-path fuzz target

Use //go:build go1.18 and // +build go1.18 with required blank line. Keep module
minimum 1.16 unless evidence requires an increase. F.Add seeds in the committed
file satisfy committed corpus; generated crash inputs, if any, must be preserved
and turned into regressions rather than discarded.

Both NewRRule and StrToRRule should receive equivalent inputs where representable.
Rejected invalid options are normal; accepted inputs must reach Iterator and
bounded Next calls. Do not blanket recover panics or test helper-only paths.
Use portable int conversions so 386 builds work. Do not use newest-library APIs
that prevent compilation with Go1.18+ merely for convenience.

## T010 — Evidence campaign

Show deterministic seeds fail against an unfixed disposable copy, at least for
the covered root defects; verify normal input seeds produce expected nonempty
results so skipped execution cannot pass. Record exactly what domain is covered.
Run timeout 120 go test -run='^$' -fuzz=FuzzIterate -fuzztime=60s -parallel=2.
Poll to actual exit. A timed-out command is not a successful 60-second campaign.
If a fuzz input hits known out-of-scope hang, preserve evidence and design a clear
reachability restriction, then rerun the complete campaign; never relabel failure.

## T011 — Integrated validation and handoff

Run full tests, race, vet, build, gofmt and 386 tests with finite command limits.
Record duration/executions from fuzz output, seed checks, exact commands and exit
codes. Commit only rrule_fuzz_test.go. Hand off /tmp evidence with actual model
and date to parent; parent saves it under the mission and requests independent
review. Do not change workflow state or open a public issue.

## Definition of Done

Fuzz seeds are committed and meaningful, both construction paths exercised,
60-second bounded run completes without panic, ordinary suite and width checks
pass, restrictions are explicitly stated, and API/minimum/dependencies remain
consistent with the plan.

## Reviewer guidance

Check domain restrictions mathematically, not just a lucky completed campaign.
Confirm list-size and near-horizon normalization cannot erase the seeded defects.
Check single-call nontermination exposure from BY filters, BYSETPOS, and calendar
scan combinations. Verify seed sensitivity and no blanket panic suppression.
