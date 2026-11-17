---
work_package_id: WP02
title: Document safe consumer migration and release notes
dependencies:
- WP01
requirement_refs:
- FR-003
- FR-004
- FR-007
- NFR-003
- C-001
- C-002
- C-004
- C-005
planning_base_branch: rrule-go-dev-xfer
merge_target_branch: rrule-go-dev-xfer
branch_strategy: Planning artifacts for this mission were generated on rrule-go-dev-xfer. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into rrule-go-dev-xfer unless the human explicitly redirects the landing branch.
base_branch: kitty/mission-module-path-v1-9-0-01M62K6W
base_commit: 0db3ab758e7b8ad80648780df884bd301bd9cdca
created_at: '2026-10-29T03:23:35.008582+00:00'
subtasks:
- T006
- T007
- T008
- T009
- T010
history: []
agent_profile: comms-cleo
authoritative_surface: README.md
create_intent:
- MIGRATION.md
execution_mode: code_change
owned_files:
- README.md
- CHANGELOG.md
- MIGRATION.md
role: communicator
tags: []
task_type: documentation
tracker_refs: []
---

# WP02 — Document safe consumer migration and release notes

## ⚡ Do This First: Load Agent Profile

Use the `/spk-charter-profile-load` skill to load the agent profile specified in the frontmatter, and behave according to its guidance before parsing the rest of this prompt.

- **Profile**: `comms-cleo`
- **Role**: `communicator`
- **Agent/tool**: `codex`

If no profile is specified, run `SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty agent profile list` and select the best match for this work package's `task_type` and `authoritative_surface`.

---

## Objective

Prepare accurate consumer documentation for the new module identity and safe migration from upstream or the transitional fork replacement.
Keep the change privately reviewable and distinguish local candidate checks from an actual public release.
Do not publish, announce an official succession, or imply that v1.9.0 can already be downloaded.

## Context

- This WP depends on WP01, which changes the module declaration and active source imports.
- The new identity is `github.com/kentonium3/rrule-go`; the package qualifier remains `rrule`.
- This mission does not change recurrence algorithms or exported API declarations.
- The existing positional `ROption` literal migration is separate inherited v1.9.0 context.
- Kent approved simple verification rather than a synthetic module proxy or identity-test framework.
- Local candidate replacements support private walkthroughs; they do not prove public distribution.
- WP03 performs final integrated consumer, golden-output and release-packet verification.
- The post-plan review identified a critical migration hazard: removing an old replacement can restore unfixed upstream code for dependencies still importing the old path.

Read the mission's `spec.md`, `plan.md`, `occurrence_map.yaml`, `quickstart.md`, and `contracts/module-migration.md` in the mission directory.
Consult `docs/runbooks/changelog-flow.md` for the dev-to-public release boundary.
The existing draft packets provide source material; WP03 owns their edits.

## Branch Strategy and Workflow

Planning/base and consolidation target: `rrule-go-dev-xfer`.
The eventual private PR targets `main`; this WP does not submit or merge it.
Use the exact runtime-allocated lane workspace from the finalized lane metadata.
Do not invent a branch, create an unmanaged worktree, or choose a dependency base manually.
The orchestrator owns claims, lane transitions, review verdict capture and consolidation.

Implementation claim, when the analyze gate is current and WP01 is approved/done:

```sh
SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty agent action implement WP02 --agent codex --mission module-path-v1-9-0-01M62K6W
```

Subtask completion is event-sourced, never a manually edited checkbox:

```sh
SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty agent tasks mark-status T006 --status done --mission module-path-v1-9-0-01M62K6W
```

Use the actual supported completion command through the orchestrator; do not fabricate events or status fields.
A command failure, unexpected branch condition or workflow authority mismatch must be preserved and reported under the repository rules.

## Ownership and Safeguards

Modify only `README.md`, `CHANGELOG.md`, and the new `MIGRATION.md`.
The runtime classifies this WP as `code_change` because these are root delivery files; that does not authorize production algorithm edits.
You are not alone in the repository: do not revert another agent's work or overwrite unrelated documentation changes.
Do not edit `go.mod`, Go source/tests, reproduction drivers, golden files, mission drafts, evidence records, or workflow metadata.
If documentation reveals a code defect or incorrect source identity, report it for the owning WP through review rather than editing unowned code.
Do not globally replace the old repository URL.
Retain historical release facts, original-source links, MIT attribution and the modified-copy notice.
No new dependencies, runtime notifications, compatibility shim, synthetic proxy or registry account are in scope.
No public pushes, tags, releases, PRs, comments, indexing requests, other-session messages or maintainer contact are authorized.

### Subtask T006: Write the consumer migration guide

**Purpose**: Give a new adopter and each existing-consumer shape explicit, usable instructions.

**Steps**:
1. Create root `MIGRATION.md` as the consumer guide, keeping private verification details outside it.
2. Explain the old and new module paths, unchanged `rrule` qualifier, and candidate v1.9.0 intent.
3. Start with the dependency-safety preflight in T007 before giving replacement-removal commands.
4. Show the import change using both full paths; do not rename package-qualified calls.
5. Cover direct upstream dependencies that have no replacement.
6. Cover the documented all-version old-path replacement to the fork's v1.8.3.
7. Cover a version-specific old-path replacement and explain that the actual left-side version controls removal.
8. Show the exact removal syntax for both replacement forms.
9. Separate normal post-release installation commands from temporary local candidate verification scaffolding.
10. State that existing immutable releases remain available to pinned consumers.

Required command forms for the eventual released module:

```sh
go mod edit -dropreplace=github.com/teambition/rrule-go
# Instead, for a version-specific left side such as old-module v1.8.2:
go mod edit -dropreplace=github.com/teambition/rrule-go@v1.8.2
go get github.com/kentonium3/rrule-go@v1.9.0
go mod tidy
go test ./...
```

Do not present both removal commands as a sequence everyone must run.
Explain that applications without a replacement omit the removal step.
Do not claim the real public `go get` has succeeded based on a local replacement.

**Files**: `MIGRATION.md`, new consumer-facing document of the smallest sufficient length.
**Validation**: Each of the three migration shapes has an unambiguous path through the instructions.
**Handoff**: Tell WP03 which exact command sequences and document sections require final walkthrough verification.

### Subtask T007: Make retained old-path consumers a preflight gate

**Purpose**: Prevent migration instructions from silently removing the fork's panic/hang fixes from other dependencies.

**Steps**:
1. Before any replacement removal, ask consumers to inspect their module graph and package imports.
2. Explain that module requirements alone do not establish which packages still import the original path.
3. Include practical inspection guidance using ordinary Go tooling rather than a new analysis framework.
4. Distinguish the application's own imports from imports in dependencies and relevant test/build configurations.
5. State the concrete hazard: an old-path consumer can revert to selected original upstream code once its replacement is removed.
6. Explain that this can lose the fork's panic/hang fixes even when no rrule types cross an API boundary.
7. Require coordinated dependency migration first, or recommend keeping the existing pinned v1.8.3 arrangement until ready.
8. Do not describe a passing application test as proof that those fixes remain active.
9. Separately explain that old/new named types differ, requiring coordinated migration where APIs exchange them.
10. Do not require deletion of legitimate old transitive dependencies merely to make the graph look clean.

Ordinary inspection commands can include `go mod graph`, `go mod why -m github.com/teambition/rrule-go`, and a package-import listing for the relevant application packages.
Describe their limits: a single `why` chain is not a complete inventory, and build tags/platforms can affect package selection.
Keep this as focused consumer guidance, not a request to prove every possible dependency configuration.
A retained-old-import walkthrough in WP03 must demonstrate the stop/coordinate/stay-pinned outcome.

**Files**: `MIGRATION.md`; concise safety link or summary in `README.md` and relevant Upgrade Notes.
**Validation**: The warning precedes removal commands and names the loss of protective fixes, not only type incompatibility.
**Safeguard**: Do not invent a dual-module alias bridge or tell users to discard their protective replacement indiscriminately.

### Subtask T008: Update active README adoption pointers

**Purpose**: Make current examples and installation/reference pointers agree with the candidate identity without falsifying provenance.

**Steps**:
1. Recheck each README occurrence against the approved manual-review policy.
2. Update the active demo import to `github.com/kentonium3/rrule-go`.
3. Add or revise a concise installation/usage section linking to `MIGRATION.md`.
4. Keep candidate/release language accurate: prepared private code is not public availability.
5. Point the Go Reference badge/link at the intended new module identity, without claiming that indexing was verified.
6. Review CI and CodeQL badges for the repository they actually describe.
7. In this private development README, use the actual development repository for applicable active workflow links.
8. Explain in handoff that later public README changes must target the public repository and be reconciled as selected hunks.
9. Preserve Teambition source attribution and the MIT/source license relationship.
10. Leave unrelated stale prose or Codecov cleanup outside the identity migration unless an owned, justified correction is essential.

The private README currently contains a private preamble and internal links.
Do not copy that document wholesale into the public repository.
The candidate public-copy selection must exclude private mission IDs, internal evidence paths and the private preamble.
A link to the upstream project can remain correct as attribution even when active install pointers use the fork.

**Files**: `README.md`, selected identity-related hunks only.
**Validation**: Active imports/install/reference links are coherent; historical/source links retain their intended meaning.
**Safeguard**: Do not edit license contents or agent-context files to make a bulk search return zero matches.

### Subtask T009: Add consumer-focused Unreleased notes

**Purpose**: Record the identity change and migration consequence in the established release record.

**Steps**:
1. Add a short consumer-focused entry under the existing `## [Unreleased]` heading.
2. Use the existing populated category ordering; do not add empty sections.
3. Explain before/after module identity and point to the migration guide.
4. Make the source import-path change explicit rather than describing it as a transparent repository rename.
5. Include the retained-old-dependency preflight in Upgrade Notes or link to its prominent guide section.
6. Keep the inherited positional `ROption` migration and other accepted v1.9.0 notes distinct and intact.
7. Do not create a dated v1.9.0 release section or imply that the public tag exists.
8. Preserve every released section, particularly v1.8.3's truthful statement that it retained the original module declaration.
9. Retain upstream PR/source links and contribution credits that still describe historical work.
10. Keep private mission IDs and evidence paths out of the proposed consumer entry.

**Files**: `CHANGELOG.md`, Unreleased additions only.
**Validation**: The entry explains what consumers must do and does not overstate public availability.
**Safeguard**: Do not use issue-closing keywords, retroactively rewrite tags or retract working releases for publicity.

### Subtask T010: Verify documentation commands and hand off exact copy

**Purpose**: Give WP03 a coherent consumer packet and catch documentation defects before integrated verification.

**Steps**:
1. Read the three owned files as a new consumer and check every path through the guide.
2. Walk the documentation command forms against disposable local candidate fixtures with `GOWORK=off GOPROXY=off`.
3. Clearly label any temporary new-path-to-candidate replacement as private verification scaffolding, never a production recommendation.
4. Record which public-tag retrieval command remains deferred; do not execute it as part of this private WP.
5. Run the established changelog format check from the candidate source.
6. Check local Markdown targets such as `MIGRATION.md` and section anchors without requesting public indexing.
7. Inspect the final diff for accidental historical edits or private details in new consumer prose.
8. Report source identity, commands, outcomes and remaining deferred checks through the normal WP handoff.
9. Give WP03 exact-copy source sections for the private release/upstream draft packets.
10. If a command or source mismatch requires changes outside ownership, return it through the workflow to the owning WP.

Required repository documentation check:

```sh
go run ./tools/changelog check
```

WP03 owns durable verification evidence and edits to mission draft packets.
Do not write those artifacts from this code-change WP or claim their integration checks have already passed.
Do not add tests that merely mirror strings; use the agreed lightweight walkthroughs and existing helper.

**Files**: Final corrections within `README.md`, `CHANGELOG.md`, `MIGRATION.md` only.
**Validation**: Changelog check passes; walkthrough results and deferred public checks are accurately handed off.

## Definition of Done

- All five subtasks have clear completion evidence in the event-sourced workflow.
- Root migration guidance covers direct adoption, both replacement forms and dependency-safety preflight.
- README active examples/pointers agree with the new identity, with preserved source attribution and private/public distinctions.
- Unreleased notes are consumer-focused and historical released sections are unchanged.
- The changelog check passes and the verified/deferred command boundary is explicit.
- No production source, API, algorithm, workflow state or unowned evidence/draft file was edited.
- No public post, push, tag, release, indexing request, session message or upstream contact occurred.
- Actual implementing model, reviewing model and review date are recorded in the handoff/review record.

## Risks

- **Realistic**: Remaining old-path imports lose fork fixes after indiscriminate removal; require the preflight and coordination/stay-pinned instruction.
- **Certain if blindly replaced**: Historical release statements or attribution become false; review each mixed-purpose documentation occurrence.
- **Realistic**: Local fixtures are mistaken for a verified public release; keep claims and command labels separate.
- **Realistic**: Private README material leaks into a later public port; preserve the explicit selected-hunk allowlist boundary.

## Reviewer Guidance

Review the consumer journey before polishing prose.
Check that the safety gate comes before commands and covers loss of fixes without shared-type exchange.
Check exact version-specific removal syntax and unchanged package qualifier.
Confirm that public availability and upstream endorsement are not asserted prematurely.
Confirm historical v1.8.3 facts, attribution and inherited release notes remain intact.
The reviewer must be a different model from the implementing model; a different profile or role alone does not satisfy this rule.
Record actual model identities and date, and report same-model review as degraded rather than silently treating it as approval.
