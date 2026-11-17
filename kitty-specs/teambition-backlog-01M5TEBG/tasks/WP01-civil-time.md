---
work_package_id: WP01
title: RFC civil time and explicit DTSTART
dependencies: []
requirement_refs:
- FR-001
- FR-002
- FR-007
- FR-008
- NFR-001
- NFR-002
- NFR-003
- C-001
- C-002
- C-003
- SC-001
planning_base_branch: fix/teambition-backlog
merge_target_branch: fix/teambition-backlog
branch_strategy: Planning artifacts for this mission were generated on fix/teambition-backlog. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into fix/teambition-backlog unless the human explicitly redirects the landing branch.
base_branch: kitty/mission-teambition-backlog-01M5TEBG
base_commit: a0ea3af6f1e7896334e50189610c393c2a68b0b5
created_at: '2026-10-25T21:40:11.489213+00:00'
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
- civil_time.go
- rrule_dst_test.go
- docs/repro/05-07-backlog/golden/main.go
- docs/repro/05-07-backlog/golden/corpus-dst.json
- docs/repro/05-07-backlog/golden/baseline-dst.json
- docs/releases/v1.9.0-dst.md
execution_mode: code_change
model: gpt-6.1-sol
owned_files:
- rrule.go
- civil_time.go
- rrule_dst_test.go
- docs/repro/05-07-backlog/golden/main.go
- docs/repro/05-07-backlog/golden/corpus-dst.json
- docs/repro/05-07-backlog/golden/baseline-dst.json
- docs/releases/v1.9.0-dst.md
role: implementer
tags: []
tracker_refs: []
---

## ⚡ Do This First: Load Agent Profile

Load implementer-ivan using the spk-charter-profile-load skill before code work.
Read AGENTS.md, docs/CONTEXT.md, the current runbook, spec.md, plan.md and data-model.md.
The current scope is only public kentonium3/rrule-go issues5 and6.
Contributor issue7 fix76c646d is inherited baseline; do not assess or extend it.
Kent approved private ROption metadata and targetv1.9.0, including the external positional-literal break.
Existing signatures/keyed callers remain; no additional API or dependency expansion.

## Branch Strategy

Planning/base and consolidation target: fix/teambition-backlog.
Final delivery: private PR to main, unmerged. No public action, tags, releases or upstream pushes.
Use the runtime-allocated lane workspace from lanes.json; do not invent a worktree/base.
The root orchestrator owns claims, review and workflow transitions.
Implementation model: gpt-6.1-sol; reviewer: gpt-6-astra. Record actual model/date in handoff.
All Spec Kitty commands start with SPEC_KITTY_NO_UPGRADE_CHECK=1.
Claim command (orchestrator only):
`SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty agent action implement WP01 --agent codex --mission teambition-backlog-01M5TEBG`

## Shared evidence and safety contract

Write code only in your allocated workspace. No concurrent production editing.
Use production entry points in regression tests; helper tests alone do not prove wiring.
A negative/rejection case has a same-fixture conformant positive control.
Each independently defective behavior must be shown red before its implementation change.
Do not represent already-correct baseline controls as failing defects.
Save exact commands, output/exit, baseline commit, code commit and model/date in /tmp handoff artifacts.
The root orchestrator persists mission evidence; code WPs do not own kitty-specs/ paths.
D-3: preserve fault command/output/state before diagnosing/recovering; no blind mutation retries.
A substantive gate must not be bypassed. Genuine new scope/API decisions go to root.
Never alter Spec Kitty installation or host/security settings. Never stash or force-push.
Use scoped Conventional Commits without issue-closing keywords (private issue numbers collide).
Safe cache: GOCACHE=/tmp/rrule-go-cache. Bound long commands and poll to completion.
The Go1.16 toolchain is /tmp/rrule-go116-toolchain/go/bin/go, cache /tmp/rrule-go116-cache.
Baseline:76c646dde8d33ff7c3325df0613b31b058243bea, snapshot /tmp/rrule-backlog-v183-plus-duplicates.
Baseline modern and Go1.16.15 tests passed before this mission; do not alter baseline.

## Objective and boundary

Own public kentonium3/rrule-go#5: RFC civil-time generation and explicit DTSTART interpretation.
Provide the immutable metadata foundation needed for gap nominal fields and later issue6 forms.
Do not implement issue6 lexical validation/serialization beyond narrow gap/fold DTSTART wiring.
Keep the original exact *time.Location rather than a marker or reloaded name.

## T001 — Red civil-time and DTSTART regressions

Create rrule_dst_test.go with named independently failing production-path tests.
Use fixed wall starts and RFC-derived UTC expected instants, not dateutil or the resolver under test.
Sydney spring HOURLY COUNT6:00,01AEST then03,04,05,06AEDT.
Sydney fold HOURLY COUNT6:00,01,02AEDT then03,04,05AEST.
MINUTELY BYMINUTE0 spring:00,01,03,04; DAILY02:30 skips October2.
DAILY BYHOUR2,3 October1 compatibility fixture plus aligned DTSTART02 normative fixture.
Dailyfold02:30 must use firstAEDT; named BYHOUR9 control remains09 after fold.
Add NY, LordHowe half-hour cases and every frequency YEARLY through SECONDLY.
Test gap holes before positive/negative BYSETPOS positions and COUNT replenishment.
Exercise Iterator,All,Between,After,Before and Set; UNTIL inclusive before/at/after boundary.
Show explicit parsed gap DTSTART pre-gap offset, nominal next-day cadence and first-count behavior.
Include a subdaily gap anchor whose resolved first instant overlaps later generated civil candidates.
No duplicate/backward results permitted in this overlap case.
Constructor/public DTStart given the later fold instant must emit that authoritative first instant once.
Parsed ambiguous civil DTSTART chooses first offset instead; keep these tests distinct.
Run red cases separately so one failure does not conceal an independent missing behavior.
Record each red output against baseline before any corresponding production edit.

## T002 — Small bounded civil resolver

Create civil_time.go (unexported helpers only) for requested civil fields + actual location.
Determine valid candidate instants by offsets and exact round-trip field equality.
No exact match means generated gap; multiple matches resolve earliest.
Explicit gap DTSTART resolves using offset before gap, not generated-candidate skip.
Keep a fast normal/UTC path and bounded lookups; no hard-coded one-hour arithmetic.
State the bounded-search invariant and realistic supported-transition assumptions in handoff.
Prove southern/northern and30-minute shifts; do not claim arbitrary historical discontinuity support.
Never step subdaily BY selectors in elapsed duration from midnight (PR66 defect).
Retain actual *Location, including custom FixedZone/TZif behavior.
All production helpers must compile on Go1.16: no ZoneBounds, generics or slices package.
Avoid new caches/global registries or dependency downloads.
Review helper use at each candidate construction seam, not only one branch.

## T003 — Immutable parse metadata and nominal anchor foundation

Add the user-approved private metadata field to ROption, with immutable descriptors in civil_time.go.
Descriptors carry lexical kind as needed, original nominal fields, exact *Location and parsed time snapshots.
RRule may retain private nominal anchor state distinct from resolved lower-bound instant.
ROption value copies share immutable metadata; never mutate it in place.
Validate snapshot by current time value INCLUDING location identity, not Equal-instant alone.
Replacing Dtstart invalidates only its descriptor; replacing Until invalidates only its descriptor.
Expose only private hooks needed by WP02 forms and internal shared-start transfer.
Do not hide metadata in Wkst/subseconds, use sentinel locations, unsafe, reflection or name reload.
Direct ordinary NewRRule options have no metadata and preserve programmatic semantics.
A caller-supplied Go time is authoritative; do not infer nonexistent original wall fields.
Public time-only setters replace relevant semantics; laterfold setter remains authoritative.
Narrow anticipated out-of-map edit: str.go StrToDtStart/private DTSTART helper wiring for RFC gap/fold.
Record that exact rationale in handoff; WP02 owns str.go and runs only after WP01 approval.
Do not move general UNTIL/type parser work into this WP.
Preserve ordinary public parsed time.Location observability.

## T004 — Candidate validity and bounded BYSETPOS ranks

Route both generation output branches through valid civil resolution.
Keep wall-clock advancement and BY values matching civil fields.
Skip generated gaps before COUNT and before BYSETPOS positions.
Resolve repeated generated local times once at the first instant.
Treat explicit DTSTART separately; do not drop its authorized gap/laterfold first instance.
Enforce unique ascending output when a resolved gap anchor and later candidate overlap.
Preserve inclusive UNTIL and exhaustion bookkeeping.
Keep existing arithmetic direct-index BYSETPOS selection for UTC/fixed/gap-free periods.
For gap periods subtract nonexistent local ranges from per-day candidate counts.
Use sorted-timeset binary searches and rank/select across holes, not annual Cartesian enumeration.
Handle positive/negative positions and selected-instant dedup across overlapping positions.
No full product timestamp materialization or unconditional product scan.
Dense annual UTC rule can represent31,536,000 nominal candidates but must still select arithmetically.
Retain effective-UTC short-UNTIL termination, including parsed default-floating numerical evaluation.
Do not regress never-matching short-horizon rules to MAXYEAR scanning.
Record any necessary private metadata evaluation hooks for WP02.

## T005 — Reusable golden harness and baseline corpus

Create docs/repro/05-07-backlog/golden/main.go as ignored standalone program or equivalent driver.
Provide deterministic corpus-dst.json and baseline-dst.json (or small documented structured equivalents).
Same driver runs baseline76c646d snapshot and current lane; no public fetch/tag/replace committed.
Record occurrence instants, String and parse error status for each case; fixed DTSTARTs only.
Take first16 values or exhaustion, with explicit near-boundary subdaily starts/intervals.
Use aligned day/month selectors for coarser frequencies so samples actually reach transition days.
Cover every frequency ×UTC/NY/Sydney/LordHowe ×week around spring/fall dates.
Extract bounded fixtures from existing tests and repro matrix/edge-cases; report provenance/included/skipped counts.
Pathological cases need bounded calls/timeouts with a reason, not silent omission.
Compare exact inherited baseline behavior outside issue5 corrections; contributor issue7 is not new work.
Classify every changed row: generated gap/fold plus downstream COUNT/BYSETPOS, explicit anchor correction.
No blanket DST-zone allowlist. Name the RFC/test rationale per difference.
Include changed rows as sensitivity controls against a broken/empty comparison.
Write golden usage beside harness; WP02 adds separate forms corpus without parallel edits to your files.

## T006 — Green and resource verification

Run focused DST regression tests and full modern go test ./....
Run direct Go1.16.15 go test ./... with its separate cache.
Check go vet ./... and gofmt -l . (zero paths) on final code.
Benchmark dense annual positive/negative BYSETPOS against pristine baseline.
Prove work remains arithmetic on regular/UTC periods; memory-bounded enumeration alone is insufficient.
Run transition-rank fixtures proving gaps removed before signed positional indexing.
Verify no repeated selector implementation change beyond inherited baseline.
Run golden baseline/current and save every classified diff plus zero-unexplained result.
Include authoritative laterfold constructor/setter and custom location controls where foundation touches them.
Verify compound fix parts each have independent failing production-path tests.
Do not mutate source solely to manufacture test failures after the baseline evidence exists.
Full race/fuzz integrated gates are repeated by WP02 closeout; WP01 review may run relevant race tests.
Record actual toolchain and gate exit status, not merely command strings.

## T007 — Release note, credit and independent handoff

Create docs/releases/v1.9.0-dst.md describing all-frequency gap/fold result changes.
Explain parsed gap DTSTART nominal cadence versus resolved first instant and programmatic-time authority.
Document private ROption field and external positional-literal source break; keyed migration; targetv1.9.0.
Credit retained Sydney transition fixtures/ideas from teambitionPR66 with RFC-rederived expectations.
Do not credit/copy its elapsed-time algorithm or invalid double-fold expected results.
Preserve MIT LICENSE/README notice. No public issue comments or closing keywords.
Commit only owned files plus explicitly justified narrow parser adapter.
Provide root an absolute /tmp handoff path with changed files/code commit/red-green/golden/benchmark evidence.
Record implementation model gpt-6.1-sol and date2026-10-25; reviewer records actual gpt-6-astra/date.
No same-model review represented as independent approval.
Do not mark sibling WP status or begin issue6 implementation.

## Reviewer guidance and definition of done

Check FR001/002 behavior on all public paths, not helper-only correctness.
Check earliest-fold generation versus authoritative laterfold DTSTART exception.
Check gap-start collision, signed BYSETPOS ranking, COUNT and UNTIL boundaries.
Check helper boundedness/performance and minimumGo compatibility.
Check immutable metadata copies, exactLocation retention and per-part invalidation contracts.
Check golden comparison is sensitive and baseline76c646d includes contributor change untouched.
Require durable red-before/green-after and exact code identity; independently run selected controls.
Approve only with gpt-6-astra review evidence; leave implementation lane for orchestrator transition.
