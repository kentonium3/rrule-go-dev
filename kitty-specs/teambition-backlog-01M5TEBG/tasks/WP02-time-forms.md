---
work_package_id: WP02
title: RFC DTSTART UNTIL forms and integrated compatibility
dependencies:
- WP01
requirement_refs:
- FR-003
- FR-004
- FR-005
- FR-007
- FR-008
- NFR-001
- NFR-002
- NFR-003
- C-001
- C-002
- C-003
- SC-002
- SC-004
- SC-005
planning_base_branch: fix/teambition-backlog
merge_target_branch: fix/teambition-backlog
branch_strategy: Planning artifacts for this mission were generated on fix/teambition-backlog. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into fix/teambition-backlog unless the human explicitly redirects the landing branch.
base_branch: kitty/mission-teambition-backlog-01M5TEBG
base_commit: a0ea3af6f1e7896334e50189610c393c2a68b0b5
created_at: '2026-10-25T22:20:17.423665+00:00'
subtasks:
- T008
- T009
- T010
- T011
- T012
- T013
- T014
history: []
agent_profile: implementer-ivan
authoritative_surface: str.go
create_intent:
- str_timeforms_test.go
- docs/repro/05-07-backlog/golden/corpus-forms.json
- docs/repro/05-07-backlog/golden/baseline-forms.json
- docs/repro/05-07-backlog/golden/README.md
- docs/releases/v1.9.0-timeforms.md
execution_mode: code_change
model: gpt-6.1-sol
owned_files:
- str.go
- rruleset.go
- str_timeforms_test.go
- docs/repro/05-07-backlog/golden/corpus-forms.json
- docs/repro/05-07-backlog/golden/baseline-forms.json
- docs/repro/05-07-backlog/golden/README.md
- docs/releases/v1.9.0-timeforms.md
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
`SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty agent action implement WP02 --agent codex --mission teambition-backlog-01M5TEBG`

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

Own public kentonium3/rrule-go#6: lexical forms, pair validation and correct serialization.
Depend on WP01 approved civil/anchor/immutable-metadata foundation.
Own integrated golden/gates closeout without adding an integration WP.
No contributor issue7 completeness assessment or extension.

## T008 — Red form and round-trip regressions

Create str_timeforms_test.go and external-package keyed-API compatibility controls.
Use real StrToRRule,StrToROption/InLocation and Set/slice/InLoc parser entry points.
Floating DTSTART+floating UNTIL must round-trip noZ/noTZID with expected daily occurrences.
Canonical DTSTART;VALUE=DATE plus DATE UNTIL must preserve DATE serialization.
Reject TZID/UTC DTSTART plus floating UNTIL, floating DTSTART plus UTC UNTIL.
Reject both DATE/DATE-TIME mismatch directions and TZID applied to DATE or Z.
Every rejection names the relevant part and has a conformant same-fixture positive control.
Pin legal TZID DTSTART+UTC-Z UNTIL (PR68's rejected RFC-required form).
Pin UTC+UTC, floating+floating, DATE+DATE, canonical VALUE=DATE-TIME and legacy accepted bare DATE input.
Record each defective family red on preceding-WP/baseline code before parser changes.
Do not represent already-correct legal controls as newly fixed defects.
Test inline DTSTART before/after UNTIL, separate lines in either order and shared set DTSTART placement.
Check all RRULE/legacy EXRULE lines, even overwritten ones.
Split old mixed set fixture with TZID+DATE EXRULE UNTIL into explicit rejection and validUTC control.

## T009 — Collect lexical descriptors then validate

Use private helpers to collect DTSTART/UNTIL lexical forms before pair checking.
Consume WP01 immutable metadata hooks; retain exact original *Location pointers.
Preserve API signatures and existing keyed callers; no new exported identifier/dependency.
Add narrow canonical DATE/DATE-TIME parameter parsing; no full component framework.
Validate TZID-on-Z/DATE independently from DTSTART/UNTIL kind matching.
Distinguish lexicalUTC from numericalUTC floating storage.
A supplied InLocation does not convert floating syntax to boundUTC output.
Bare RRULE without DTSTART preserves today's loc parsing and legacyUTC-Z UNTIL String.
Do not attach explicit-start serialization semantics to no-DTSTART fragments.
Inline DTSTART remains a library extension; preserve acceptance with order-independent checks.
Errors should identify DTSTART,UNTIL or TZID as relevant.
Prior-art PR68 suffix equality is wrong; do not port it.
Use exact expected forms/instants derived from RFC5545.
No generalized validation expansion unrelated to this acceptance.

## T010 — Shared Set DTSTART and internal transfers

Collect shared DTSTART across the supplied set lines before processing RRULE/EXRULE.
Validate every rule line against shared DTSTART before historical last-wins/ignore behavior.
An inline DTSTART may not bypass shared-start compatibility or silently replace its anchor.
Use private metadata-aware transfer helpers so nominal gap/floating/date form survives Set.RRule handoff.
Public time-only setters remain authoritative and replace the relevant descriptor.
Set may hold private shared-anchor descriptors; no new public Set API.
Legacy EXRULE undergoes form validation then remains ignored; do not implement exclusion evaluation.
Multiple RRULE lines retain last-wins after validation because RFC semantics are undefined.
Do not change general unknown-property handling.
Preserve RDATE/EXDATE behavior unless exact DTSTART form preservation requires a narrow adapter.
Test before/after shared placement, copied options and serialization after Set construction.
Document any small out-of-map rrule.go/civil_time.go adapter with exact rationale; WP01 is already approved.
No concurrent edits and no broad ownership claims for those foundation files.

## T011 — Serialization, immutable copies and public replacement

ROption.String,RRuleString,RRule.String and Set.Recurrence must preserve explicit floating/DATE forms.
TZID DTSTART keeps its TZID and UNTIL remains UTCZ.
Programmatic NewRRule with Until in any ordinary location retains UTCZ output.
No-DTSTART fragments keep legacyUTC-Z UNTIL output, including DATE and floating input.
Compare current public Dtstart/Until to per-part snapshots including Location identity before metadata use.
Replacing one part invalidates only its own descriptor; never mutate shared immutable metadata.
Copies followed by NewRRule and independent public-field assignments do not corrupt the source option.
Include same-instant/different-location replacements to catch Equal-only stale checks.
Public DTStart/Until setters deliberately replace supplied-part semantics; preserve other descriptors.
Private Set transfer preserves explicit-start metadata intentionally.
Include ordinary FixedZone("America/New_York",-5h) and IANA NY paired controls crossing March8.
Use both with/without UNTIL and through copy/Set/setter paths.
Add custom TZif same-name control to prove actual pointer data survives, not a name reload.
No registry/marker/unsafe/reflection or hidden metadata in Wkst/subseconds.
Standalone StrToDtStart returns correct resolved instant; detached time.Time has no lexical metadata promise.

## T012 — Extend golden corpus and integrated compatibility

Add separate corpus-forms.json/baseline-forms.json and classifications/usage files under golden directory.
Consume WP01 driver; a minimal needed driver adapter is a justified sequential out-of-map edit.
Baseline exact commit76c646d, snapshot /tmp/rrule-backlog-v183-plus-duplicates.
Compare firstN occurrences,String/error status on the same deterministic corpus.
Cover form pairs, ordering, sharedsets, barefragments, copied options and custom evaluation locations.
Add DST+BYSETPOS+parsed form integrated cases at actual boundaries.
Inherited duplicate-selector cases are compatibility controls only; do not assess/extend contributor fix.
Report provenance and included/skipped source fixture counts.
Classify new differences only floating/DATE serialization, enumerated forbidden-form errors and issue5 RFC corrections.
No broad ignore lists, no zone-wide allowance and no unexplained String changes.
Keep changed-row sensitivity controls so an empty/disabled comparison fails.
Save exact baseline/current outputs, comparison results and rationale per changed row.
Do not commit a module replace or use public network actions for baseline.

## T013 — Full integrated local gates

Run go vet ./..., gofmt -l . (zero output), go test ./... and go test -race ./....
Run modern go test -run='^$' -fuzz=FuzzIterate -fuzztime=60s, polling to completion.
At least60seconds actual fuzz duration; do not count build/setup time as fuzz time.
Run direct /tmp/rrule-go116-toolchain/go/bin/go test ./... with GOCACHE=/tmp/rrule-go116-cache.
Preserve existing minimum-Go fuzz build constraints; go.mod stays go1.16.
Run complete golden comparison against76c646d and prove zero unexplained differences.
Repeat dense BYSETPOS resource controls and effectiveUTC short-UNTIL cases on integrated parsing paths.
Run all WP01 civil/anchor/public setter tests together with form tests.
Record actual exit statuses, duration/toolchain, corpus counts and classified differences.
Avoid rerunning already-passed broad gates without new code/failure reason.
If a gate fails, preserve evidence, diagnose root cause and repeat affected gates after fix.
No acceptance waiver, fake empty directories or silent lenient mode.

## T014 — Release note and independent handoff

Create docs/releases/v1.9.0-timeforms.md listing every newly rejected form and serialization change.
State floating/DATE exceptions, validTZID+UTC retention, barefragment conventions and legacyEXRULE validation.
Explain approved private ROption field, external positional-literal break and migration to keyed literals.
Targetv1.9.0; no tag/release/public copy is authorized.
Credit any retained TZID-on-Z refusal idea from teambitionPR68; reject its suffix/localUNTIL defects.
Preserve MIT LICENSE and README notice; Conventional Commits without closingkeywords.
Provide root absolute /tmp handoff with code commit/model/date/red-green/golden/fullgates/release notes.
Implementation gpt-6.1-sol and independent review gpt-6-astra must be recorded with actual date.
Root persists mission evidence and owns strictaccept/consolidation/final review/privatePR.
Do not push, create public PR/comment or merge final privatePR.

## Reviewer guidance and definition of done

Review each invalid-form branch and corresponding conformant control through production parsers.
Confirm every shared RRULE/EXRULE is validated before ignored/overwritten semantics.
Check form retention across copies/setters and exact customLocation data; no name token reconstruction.
Check direct ordinary NewRRule and no-DTSTART fragment controls remain unchanged.
Check inherited contributor code is not expanded under this issue.
Require each independent defect red/green evidence, all integrated gates and classified golden diffs.
Check Go1.16 direct test and private-field positional-break release notes.
Approval is implementation evidence only; final integrated review occurs after consolidation.
