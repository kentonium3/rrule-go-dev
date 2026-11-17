---
work_package_id: WP01
title: Adopt the fork module identity in active source
dependencies: []
requirement_refs:
- FR-001
- FR-002
- NFR-002
- C-001
- C-002
- C-003
- C-006
planning_base_branch: rrule-go-dev-xfer
merge_target_branch: rrule-go-dev-xfer
branch_strategy: Planning artifacts for this mission were generated on rrule-go-dev-xfer. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into rrule-go-dev-xfer unless the human explicitly redirects the landing branch.
base_branch: kitty/mission-module-path-v1-9-0-01M62K6W
base_commit: b430bfe9437f3fe683914556ac918ad3cf5a5c5e
created_at: '2026-10-29T03:16:42.944423+00:00'
subtasks:
- T001
- T002
- T003
- T004
- T005
history: []
agent_profile: implementer-ivan
authoritative_surface: go.mod
create_intent: []
execution_mode: code_change
owned_files:
- go.mod
- example_test.go
- str_timeforms_api_test.go
- docs/repro/01-iterator-panics/main.go
- docs/repro/02-unreachable-byset-hang/main.go
- docs/repro/03-until-scan/main.go
- docs/repro/05-07-backlog/main.go
- docs/repro/05-07-backlog/golden/main.go
- docs/repro/edge-cases/main.go
- docs/repro/matrix/main.go
- docs/repro/05-07-backlog/golden/README.md
role: implementer
tags: []
task_type: implementation
tracker_refs: []
---

# WP01 — Adopt the fork module identity in active source

## ⚡ Do This First: Load Agent Profile

Use the `/spk-charter-profile-load` skill to load the agent profile specified in the frontmatter, and behave according to its guidance before parsing the rest of this prompt.

- **Profile**: `implementer-ivan`
- **Role**: `implementer`
- **Agent/tool**: `codex`

If no profile is specified, run `SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty agent profile list` and select the best match for this work package's task type and authoritative surface.

---

## Objective

Declare the library as `github.com/kentonium3/rrule-go` and make its active external tests and runnable reproduction drivers import that identity.
Keep package name `rrule`, exported declarations, recurrence behavior and Go 1.16 compatibility unchanged.
Correct the golden-driver instructions so historical and candidate source builds remain reproducible after imports diverge.

## Context

The fork's earlier v1.8.3 release retained the original declaration as a transition.
Kent has selected the fork-owned path for candidate v1.9.0.
Changing a dependency identity is established Go behavior; this package needs a small explicit declaration check and ordinary builds, not another publishing system.
A local replacement alone can hide an incorrect declaration, so preserve a before/after declaration assertion.
Actual public-tag retrieval remains a later release action and is not authorized here.

Read these mission artifacts from the runtime-selected mission location:

- `spec.md`: FR-001, FR-002 and the compatibility/privacy constraints.
- `plan.md`: source ownership, simple verification, zero-change golden strategy and review fixes.
- `occurrence_map.yaml`: Kent-approved classification; inspect actual matches rather than assuming approval proves correctness.
- `quickstart.md`: the intended declaration check and later consumer/golden steps.
- `reviews/post-plan-review.md` and `reviews/post-plan-disposition.md`: why the golden instructions must change.

WP02 owns README/changelog/consumer migration prose.
WP03 owns final local-consumer, golden and integration evidence plus private communication packets.
This package must provide a reproducible code candidate and a useful command transcript for those successors.
Do not edit their owned artifacts to make this package look complete.

## Branch Strategy

Planning/base and mission consolidation target: `rrule-go-dev-xfer`.
Eventual delivery target: a private PR to `main`, not a direct merge or public port.
Use the lane/workspace returned by the workflow after finalization and analysis.
Do not construct a worktree path, choose a base branch, or move workflow state manually.
Dependencies are empty, but this does not bypass the required current analysis report.
The orchestrator owns claiming, lane transitions, review/verdict recording and consolidation.

Claim entry point, for the authorized orchestrator:

```sh
SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty agent action implement WP01 --agent codex --mission module-path-v1-9-0-01M62K6W
```

Do not claim or execute this package during task generation.
Record the actual implementing model and the actual reviewer's different model with the date.
Using another agent name or role on the same model does not satisfy independent review.
A same-model review is degraded and cannot silently satisfy the requirement.

## Ownership and Safeguards

Owned files are exactly the frontmatter list.
All already exist; this package creates no repository files.
The module declaration, two external imports, seven reproduction imports and golden instructions are the entire intended change surface.
Do not modify runtime files such as `rrule.go`, `civil_time.go` or parsing implementations.
Do not modify assertions to conceal failures, serialized output, public declarations or package names.
Do not add aliases, a compatibility shim, another Go module, a synthetic module proxy or a verification framework.
Keep the `go 1.16` directive and all dependencies unchanged.
Preserve LICENSE, Teambition attribution, historical issue briefs and prior mission records.
Do not update inherited corpus JSON or recorded historical outputs.
Do not add candidate-source replacement directives to the library's go.mod.
No public pushes, tags, releases, PRs, comments, indexing requests or messages are authorized.
Do not contact upstream or other Vikunja sessions.
You are not alone in the repository: preserve others' work and stop on unexpected concurrent ownership changes.

### Subtask T001: Capture pre-change identity and recheck active occurrences

**Purpose**: Establish identity-specific red evidence before changing the declaration and verify that approved policy still fits the source.

**Steps**:

1. From the allocated candidate workspace, record the source commit, workspace path and `go.mod` contents through the ordinary command transcript.
2. Confirm the source still declares `github.com/teambition/rrule-go` and `go 1.16`.
3. Record `GOWORK=off GOPROXY=off go list -m` without altering global Go settings.
4. Compare that value with `github.com/kentonium3/rrule-go` and record the expected nonzero result.
5. The failure must be the old-versus-new declaration, not a missing Go executable, unrelated syntax error or network failure.
6. Inspect the exact old-path matches in every owned source file against the occurrence map.
7. Check that the golden README/header contain active instructions dependent on the old import identity, even though the README need not contain the exact token.
8. Capture unexpected new matches or contradictory policy as findings for the orchestrator; do not silently change the occurrence map.

Example declaration assertion, before and after the edit:

```sh
GOWORK=off GOPROXY=off go list -m
test "$(GOWORK=off GOPROXY=off go list -m)" = github.com/kentonium3/rrule-go
```

**Files**: Read owned files; no mutation in this subtask.

**Validation**: The transcript identifies the precise old declaration and expected failing assertion.

**Handoff**: Report command/output and source identity to the orchestrator so WP03 can retain final evidence in its owned artifact.
Do not create a `kitty-specs/` evidence file from this code-change package.

### Subtask T002: Change the module and root external imports

**Purpose**: Make the declaration and external API tests agree on the fork-owned identity.

**Steps**:

1. Set only the go.mod module directive to `github.com/kentonium3/rrule-go`.
2. Keep the Go version line unchanged and introduce no requires, replaces or excludes.
3. Change the import in `example_test.go` to the new path.
4. Change the aliased import in `str_timeforms_api_test.go` to the new path while retaining the `rrule` qualifier.
5. Preserve package declarations, test bodies, expectations and keyed constructor examples.
6. Repeat the exact declaration assertion from T001 and require success.
7. Inspect the targeted diff to ensure only the declaration/import values changed.

**Files**: `go.mod`, `example_test.go`, `str_timeforms_api_test.go`.

**Validation**: New declaration passes, the old active external-test imports are absent, and no compatibility/dependency metadata changed.
Do not infer public availability from this local source check.
Keep errors visible rather than adding a local old-module replacement to repair a missed import.

### Subtask T003: Update the seven active reproduction imports

**Purpose**: Keep private runnable reproductions bound to the local module after its identity changes.

**Steps**:

1. Update the import in `docs/repro/01-iterator-panics/main.go`.
2. Update the import in `docs/repro/02-unreachable-byset-hang/main.go`.
3. Update the import in `docs/repro/03-until-scan/main.go`.
4. Update the import in `docs/repro/05-07-backlog/main.go`.
5. Update the aliased import in `docs/repro/05-07-backlog/golden/main.go`.
6. Update the imports in `docs/repro/edge-cases/main.go` and `docs/repro/matrix/main.go`.
7. Preserve build-ignore directives, options, fixtures and program behavior.
8. Compile individual drivers to temporary binaries if needed to prove import resolution; do not run programs whose purpose includes panic/hang cases without their established bounds.
9. Review the diff for unrelated formatting or algorithm changes and remove accidental edits within this package's scope.

**Files**: The seven `main.go` paths listed in frontmatter.

**Validation**: Active imports select the new module, and the driver source remains otherwise unchanged except for the golden instruction header covered by T004.

**Boundary**: WP03 performs bounded golden output comparison; do not claim it has happened merely because these imports compile.
No stored corpus or historical output is a rename target.

### Subtask T004: Correct the golden header and README build instructions

**Purpose**: Preserve executable verification instructions when the candidate and historical snapshots have different import identities.

**Steps**:

1. Read the golden driver header and adjacent README as commands, not historical prose alone.
2. Remove the implication that the current new-import driver can be built unchanged in an old-declaration source checkout.
3. Explain that candidate builds use the candidate's new-import driver from candidate cwd.
4. Explain that old-source builds use a frozen old-import driver from a preserved pre-migration snapshot, with the old source cwd selecting its library.
5. Use `737ae01be40048d1d2932dbafb623230efa8d361` as this migration's pre-change source baseline.
6. Where explaining the older classified comparison against `76c646d`, ensure an old-import driver actually exists: a frozen driver exported from the pre-migration snapshot may be built from that historical source cwd.
7. Do not invent an unavailable driver path in the older snapshot or change the driver's program logic to force compatibility.
8. Distinguish the historical comparator from this migration's zero-change comparison.
9. State that the historical `-baseline`/`-current` comparator deliberately rejects zero changed rows.
10. Explain that WP03 instead collects fresh before/after `-out` files for the same corpus and compares their bytes.
11. Retain the historical corpus/baseline facts, hashes, classified behavior explanations and attribution.
12. Do not alter corpus JSON to make an old instruction appear valid.

**Files**: Header comments in `docs/repro/05-07-backlog/golden/main.go`; `docs/repro/05-07-backlog/golden/README.md`.

**Validation**: Each command names a driver whose import matches its intended source module and identifies the correct working directory.
The README keeps history truthful and explains the zero-change proof without requiring a synthetic proxy or new framework.

### Subtask T005: Validate the focused candidate and hand off exact results

**Purpose**: Deliver a locally usable identity change with failures exposed before consumer/documentation finalization.

**Steps**:

1. Run `go test ./...` in the candidate source and retain the command result.
2. Run `go vet ./...` and require success.
3. Run `gofmt -l .` and require zero reported files relevant to the final candidate.
4. Recheck `go list -m` and inspect go.mod after testing for accidental generated replaces or dependency changes.
5. Search active owned Go sources for the old import and reconcile every remaining match with policy.
6. Compare exported declaration and runtime-source diffs to confirm this package contains only the allowed identity/instruction edits.
7. Report exact commands, results, source commit/workspace, changed files and any limits to the orchestrator.
8. Include the T001 identity-specific red result and T002 green result in the handoff.
9. Hand off the corrected driver instructions and any temporary build locations to WP03 without writing its owned evidence files.
10. Race testing, final combined gate coverage, consumer walkthroughs and bounded golden equality remain explicit WP03 responsibilities; do not claim they passed here unless actually run and recorded.

**Files**: Read/validate all owned paths; do not add test infrastructure or new repository evidence files.

**Validation**: Focused checks pass, formatting is clean, and the diff contains no runtime/API/minimum-version/dependency expansion.
If failures need a substantive out-of-scope change, stop and report rather than weakening checks or changing another package's files.

## Definition of Done

- T001 captures the expected declaration mismatch before edits.
- The same declaration check succeeds after edits.
- The declaration, two root external imports and seven driver imports agree on the new identity.
- The golden header/README distinguish correct old-source and candidate-driver builds.
- Runtime code, exported declarations, package name and Go minimum remain unchanged.
- Historical corpus/output files remain byte-identical and no new dependency or module replacement is added.
- Focused test/vet/format results and exact source identity are handed off for WP03 evidence.
- Actual implementing and reviewing model identities and date are recorded through the authorized review process; models differ.
- Subtask completion is event-sourced, not marked by editing task checkboxes.

The orchestrator records completion through the workflow, for example:

```sh
SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty agent tasks mark-status T001 T002 T003 T004 T005 --status done --mission module-path-v1-9-0-01M62K6W
```

Do not move lanes, fabricate review history or write event logs manually.

## Risks

A broad replacement can alter historical attribution; constrain edits to reviewed active occurrences.
An old import left behind can resolve unintended source; do not add a replacement to conceal it.
A candidate driver used from historical cwd can test the wrong library; identify each source/driver explicitly.
An intentional hang/panic repro can stall or kill a process; compile it or use existing bounded verification later.
A same-model reviewer is not independent merely because it has a different role or agent name.

## Reviewer Guidance

Inspect every changed hunk against the approved map; Kent's approval is not a substitute for technical checking.
Confirm the red failure was identity-specific and the green assertion used the identical check.
Confirm every owned import changed, but APIs, algorithm bodies, Go minimum and dependencies did not.
Read the golden commands literally and verify their driver paths/cwd/import identities are coherent.
Check that no inherited JSON, old mission record, historical release, public remote or notification surface changed.
Require exact test output and a clear handoff; distinguish pending WP03 verification from completed work.
If a condition needs a correction, state the condition, realistic likelihood and smallest fix.
Reject any hidden publication, fabricated review result, manual workflow-state change or workaround of an unexpected workflow fault.
