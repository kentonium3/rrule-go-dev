---
work_package_id: WP01
title: UTC period guard and compatibility evidence
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
- SC-001
- SC-002
- SC-003
planning_base_branch: fix/until-scan
merge_target_branch: fix/until-scan
branch_strategy: Planning artifacts for this mission were generated on fix/until-scan. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into fix/until-scan unless the human explicitly redirects the landing branch.
base_branch: kitty/mission-until-scan-01M5S0WJ
base_commit: 1368c3772bb46050c697299f22c146e536774ac8
created_at: '2026-10-25T08:15:56.856473+00:00'
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
- rrule_until_test.go
execution_mode: code_change
owned_files:
- rrule.go
- rrule_until_test.go
- docs/issues/03-until-scan.md
role: implementer
tags: []
tracker_refs: []
---

## Load context

Load implementer-ivan via spk-charter-profile-load. Read AGENTS.md then docs/CONTEXT.md then docs/issues/03-until-scan.md, and this mission's spec/plan/research/data-model/quickstart. The approved scope is UTC-only pruning with unchanged fallback for all other locations. Prefix CLI with SPEC_KITTY_NO_UPGRADE_CHECK=1; no upgrades. Stop/report actual Spec Kitty faults with exact output. Previous mission D-3 exception does not apply.

## Branch and ownership

Work only in the runtime-allocated lane. Consolidate into fix/until-scan, never master. Do not push or create a PR or change lifecycle; orchestrator owns those. Own only the three frontmatter files. Write evidence to /tmp for orchestrator persistence. Implementer gpt-6.1-sol and reviewer gpt-6-astra; record date/model and exact commits. No public posting.

## T001 — Baseline and red tests

Record toolchain and baseline revision. Add focused public-path tests exercising NewRRule and StrToRRule with DAILY DTSTART 20260105T100000Z UNTIL 20260201T000000Z and BYSETPOS=2, then 2..366. Test Iterator, All and After where practical. Each operation must finish within 100ms, excluding startup/compile/teardown. Run baseline before production edits and record genuine failures, not predicted ones. Use child-process isolation and a separate bounded parent watchdog; kill and collect every timeout. Avoid a goroutine timeout that leaves work running. Positive BYSETPOS=1 controls must use the same harness. Keep tests economical under race; do not spawn a child per ordinary golden row unnecessarily.

## T002 — Safe guard

Inside generate's outer empty-remain loop, prune only if r.dtstart.Location() == time.UTC (pointer identity). Lower bound is UTC midnight Jan1 for YEARLY; month day1 for MONTHLY; current iterator civil day for WEEKLY and finer. Stop strictly after effective until, preserving equality. Use existing finished/total/return termination pattern, no buffered-result loss. Retain existing per-candidate UNTIL checks. No cached state or default horizon changes. No timezone name/offset sampling, no transitions, no arbitrary cap, no version/API/dependency changes. Keep earlier overflow/reachability protections intact. Read research proof and independent review first.

## T003 — Compatibility and horizon controls

Capture exact baseline/fixed golden sequences for all seven frequencies, inclusive before/equal/after candidate boundaries, yearly earlier-month and monthly earlier-day after advancement, partial/cross-year weeks and alternate week start. Include COUNT, BYSETPOS and pending buffers, empty selection, differing UNTIL offset. Non-UTC controls include ordinary DST gap/fold, Sao_Paulo 2018-11-03 and Apia 2011-12-29 duplicate-normalization examples, FixedZone offsets and a location merely named UTC. Existing odd outputs are preserved; do not fix calendar semantics here. Test fresh iterators after DTSTART and UNTIL setter changes (shorten, extend, zero default, UTC/non-UTC location change). Demonstrate default horizon is honored without claiming every default scan is under100ms. No active-iterator mutation guarantee.

## T004 — Validation

Run targeted tests, go test ./..., go test -race ./..., go vet ./..., go build ./..., GOARCH=386 go build ./... and gofmt -l . Record commands/results/timings. Use honest defect-sensitive evidence: both measured red/green repros plus conservative-period controls/default tests. Production instrumentation is unnecessary. If a timing case is below100ms on baseline, report that instead of inventing failure; use the heavier fixture and deterministic or mutation controls to prove sensitivity. Independent review must check actual baseline golden equivalence and no accidental scope expansion.

## T005 — Issue notes and handoff

Update issue #6 local notes with actual baseline/fixed measurements, exact UTC pointer coverage, non-UTC fallback limitation, toolchain/platform and tests. Do not claim general timezone repair. Save handoff in /tmp/rrule-until-implementation.md with test evidence/log paths, actual model/date, source commit and any risks. Safe scoped Conventional Commit only the three owned files; do not stage generated mission state. Stop for a substantive new design issue; routine implementation choices are delegated.

## Review acceptance

All requirements mapped; both repro timings below100ms fixed; exact baseline controls pass; public APIs/deps/Go version unchanged; different model independently reviews. The private PR is not a public release and is not authorization to merge master.
