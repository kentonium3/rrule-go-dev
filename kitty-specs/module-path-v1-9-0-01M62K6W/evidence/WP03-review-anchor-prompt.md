---
work_package_id: WP03
created_at: '2026-10-29T03:34:58.247452Z'
base_ref: unknown
invocation_id: e3dc8695519b4fa9b2a95b8eceaf6f03
lane_branch: unknown
lane_worktree: ~/repos/rrule-go-dev
mission_branch: kitty/mission-module-path-v1-9-0-01M62K6W
mission_id: 01M62K6WADRS86HE5GKZ1KKHEY
mission_slug: module-path-v1-9-0-01M62K6W
prompt_path: /tmp/spec-kitty-review-prompts/rrule-go-dev-df27c292f109/module-path-v1-9-0-01M62K6W/WP03/e3dc8695519b4fa9b2a95b8eceaf6f03.md
repo_root: ~/repos/rrule-go-dev
---

================================================================================
REVIEW: WP03
================================================================================

Source: ~/repos/rrule-go-dev/kitty-specs/module-path-v1-9-0-01M62K6W/tasks/WP03-verification-handoff.md

Workspace: ~/repos/rrule-go-dev
Workspace contract: lane lane-planning shared by WP03

Resolved agent identity:
  tool       : unknown
  model      : unknown-model
  profile_id : implementer-ivan
  role       : implementer

Governance:
  - Template set: software-dev-default
  - Paradigms: (none)
  - Tools: git, spec-kitty
Directive IDs:
  - RRULE-001
  - RRULE-002
  - RRULE-003
  - RRULE-004
Tactic IDs:
  - acceptance-test-first
  - review-policy-tactic
  - testing-philosophy-tactic
Styleguide IDs:
  - boring-code-review
Procedure IDs:
  - adversarial-squad-deployment
  - development-assist-test-cleanup
  - disciplined-defect-diagnosis
  - legacy-codebase-triage
  - mission-tracer-files
  - mission-wrap-up-sequence
  - post-merge-arch-gate-adjudication
  - red-main-release-discipline
  - refactoring
  - situational-assessment
  - test-first-bug-fixing
Section Anchors:
  - rrule-go project charter
  - Diagnostics: No available_tools selection provided; using runtime tool registry fallback. | Merged 4 project-local directive(s) onto the activation set: RRULE-001, RRULE-002, RRULE-003, RRULE-004. | Template set not selected in charter; fallback 'software-dev-default' applied.
  - Languages: go
  - Advisory: No specialist guidance exists for this project's language; add tech-specific guidelines to your local charter (see 'Extend your charter for an unsupported language').
  - Project layer root: ~/repos/rrule-go-dev/.kittify/charter-packs

Project authority paths:
  - docs/issues/    (consult when you change content under this directory)

Required Governance Reading:
  - AGENTS.md
  - docs/CONTEXT.md

Action-Critical Charter Sections (review):

### Terminology Canon
  Run: spec-kitty charter context --include section:terminology-canon
  When you rename or introduce a term in the diff, run this command and apply the returned rule.

### Code Review Checklist
  Run: spec-kitty charter context --include section:code-review-checklist
  When you prepare a WP for review, run this command and apply the returned rule.

### Regression Vigilance
  Run: spec-kitty charter context --include section:regression-vigilance
  When you perform a terminology cutover, run this command and apply the returned rule.

╔==============================================================================╗
║  🚨 CRITICAL: WORK PACKAGE ISOLATION RULES                              ║
╠==============================================================================╣
║  YOU ARE REVIEWING: WP03                                                     ║
║                                                                          ║
║  ✅ DO:                                                                  ║
║     • Only modify status of WP03                                            ║
║     • Ignore git commits and status changes from other agents           ║
║                                                                          ║
║  ❌ DO NOT:                                                              ║
║     • Change status of any WP other than WP03                               ║
║     • React to or investigate other WPs' status changes                 ║
║     • Review or approve any WP other than WP03                             ║
║                                                                          ║
║  WHY: Multiple agents work in parallel. Each owns exactly ONE WP.       ║
║       Git commits from other WPs are other agents - ignore them.        ║
╚==============================================================================╝

<!-- WORKTREE_TOPOLOGY -->
{
  "mission_slug": "module-path-v1-9-0-01M62K6W",
  "mission_number": "",
  "mission_type": "software-dev",
  "target_branch": "rrule-go-dev-xfer",
  "mission_branch": "kitty/mission-module-path-v1-9-0-01M62K6W",
  "current_wp": "WP03",
  "diff_command": "unavailable: no deterministic implementation claim commit found",
  "shared_lane": false,
  "note": "WP03 runs in the repository root planning workspace.",
  "entries": [
    {
      "wp": "WP01",
      "status": "approved",
      "execution_mode": "code_change",
      "workspace_kind": "lane_workspace",
      "lane_id": "lane-a",
      "lane_wp_ids": [
        "WP01"
      ],
      "branch": "kitty/mission-module-path-v1-9-0-01M62K6W-lane-a",
      "base": "kitty/mission-module-path-v1-9-0-01M62K6W",
      "commits_ahead": 11
    },
    {
      "wp": "WP02",
      "status": "approved",
      "execution_mode": "code_change",
      "workspace_kind": "lane_workspace",
      "lane_id": "lane-b",
      "lane_wp_ids": [
        "WP02"
      ],
      "branch": "kitty/mission-module-path-v1-9-0-01M62K6W-lane-b",
      "base": "kitty/mission-module-path-v1-9-0-01M62K6W",
      "commits_ahead": 16,
      "dependencies": [
        "WP01"
      ]
    },
    {
      "wp": "WP03",
      "status": "in_review",
      "execution_mode": "planning_artifact",
      "workspace_kind": "repo_root",
      "lane_id": "lane-planning",
      "lane_wp_ids": [
        "WP03"
      ],
      "branch": "rrule-go-dev-xfer",
      "base": null,
      "commits_ahead": 0,
      "dependencies": [
        "WP02"
      ]
    }
  ]
}
<!-- /WORKTREE_TOPOLOGY -->

─── GIT REVIEW CONTEXT ─────────────────────────────────────────────────────────
Review commands unavailable: no deterministic implementation claim commit found for this WP.
Re-run review after the WP has a committed implementation claim on this mission.
────────────────────────────────────────────────────────────────────────────────

================================================================================
WHEN YOU'RE DONE:
================================================================================
✓ Review passed, no issues:
  spec-kitty agent tasks move-task WP03 --to approved --mission module-path-v1-9-0-01M62K6W --agent codex --note "Review passed"

⚠️  Changes requested:
  1. Write feedback to (in-repo, committed with the project):
     ~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-coord/kitty-specs/module-path-v1-9-0-01M62K6W/tasks/WP03-verification-handoff/review-feedback-1.md
  2. spec-kitty agent tasks move-task WP03 --to planned --review-feedback-file ~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-coord/kitty-specs/module-path-v1-9-0-01M62K6W/tasks/WP03-verification-handoff/review-feedback-1.md --mission module-path-v1-9-0-01M62K6W --agent codex
  3. move-task stores feedback reference in the event log and WP frontmatter
================================================================================

📍 WORKING DIRECTORY:
   cd ~/repos/rrule-go-dev
   # Review the implementation in this workspace
   # Read code, run tests, check against requirements
   # When done, return to repo root: cd ~/repos/rrule-go-dev

📚 SHARED MISSION ARTIFACTS:
   Spec, plan, and tasks are visible from the primary checkout: ~/repos/rrule-go-dev/kitty-specs/module-path-v1-9-0-01M62K6W/
   Status authority resolves through the coordination worktree for modern missions.
   Use this lane workspace for code/tests; do not expect shared mission artifacts here

📋 STATUS TRACKING:
   kitty-specs/ status is tracked in rrule-go-dev-xfer branch (visible to all agents)
   Status changes auto-commit to the coordination branch (visible to all agents)
   ⚠️  You will see commits from other agents - IGNORE THEM
================================================================================

Review the implementation against the requirements below.
Check code quality, tests, documentation, and adherence to spec.

## Anti-pattern checklist (WP-level cheap version of mission-review)

For each item below, state PASS / FAIL / N/A in your verdict. A FAIL on any
item blocks approval.

1. **Dead code**: every new public function, class, or module created by this
   WP has at least one live caller from production code, excluding tests. For
   each new module, run targeted import/call-site greps, for example
   `grep -r -e "from <new_module> import" -e "import <new_module>" src/ --include="*.<ext>"`.
   Zero production hits means dead code.

2. **Synthetic-fixture test**: every test marked as covering an FR listed in
   this WP's frontmatter actually invokes the production code path that would
   produce the asserted shape. A test that constructs a literal dict matching
   the assertion is a synthetic fixture. Ask: if I delete the implementation
   code, does this test still pass? If yes, the FR is untested.

3. **Silent empty return**: search every new code path for `except ...:
   return ""`, `return None`, `return []`, `return {}`, or `pass`. Each hit
   must have a documented reason; absent that, it is a silent failure
   candidate.

4. **FR coverage**: every functional requirement in `requirement_refs`
   (`FR-###`, including a letter-suffixed `FR-###a`) has at least one test
   assertion that references the behavior it names, not just a comment or
   frontmatter entry. `SC-###` refs are tracked, not gated, and a
   `<mission-slug>#<ID>` ref cites another mission, so neither needs a test
   here.

5. **Frozen surface**: no commit in this WP modifies a file the spec,
   contract, or WP prompt marks as frozen or untouchable. For each frozen file,
   `git log --oneline <base>..HEAD -- <frozen-file>` must be empty.

6. **Locked decision**: no new code path contradicts a `MUST NOT` clause in
   `spec.md`, `plan.md`, or `contracts/`. Grep the diff for forbidden patterns
   named by those clauses.

7. **Shared-file ownership**: any file modified by this WP that is also
   modified by another WP, visible in `lanes.json`, shared lane metadata, or
   the same mission merge, has an explicit coordination note in the move-task
   reason or review feedback.

8. **Production fragility**: any new `raise` in a production code path has a
   documented fail-loud rationale. A bare `raise` in a request handler, worker,
   or CLI path that can fire on a transient race is a fragility risk.


╔==============================================================================╗
║  WORK PACKAGE PROMPT BEGINS                                            ║
╚==============================================================================╝

---
work_package_id: WP03
title: Verify the migration and complete private handoff packets
dependencies:
- WP02
requirement_refs:
- FR-003
- FR-004
- FR-005
- FR-006
- FR-007
- FR-008
- NFR-001
- NFR-002
- NFR-003
- C-001
- C-002
- C-003
- C-004
- C-005
- C-006
- SC-001
- SC-002
- SC-003
- SC-004
planning_base_branch: rrule-go-dev-xfer
merge_target_branch: rrule-go-dev-xfer
branch_strategy: Planning artifacts for this mission were generated on rrule-go-dev-xfer. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into rrule-go-dev-xfer unless the human explicitly redirects the landing branch.
subtasks:
- T011
- T012
- T013
- T014
- T015
- T016
history: []
agent_profile: implementer-ivan
authoritative_surface: kitty-specs/module-path-v1-9-0-01M62K6W/
create_intent:
- kitty-specs/module-path-v1-9-0-01M62K6W/evidence/migration-verification.md
- kitty-specs/module-path-v1-9-0-01M62K6W/evidence/vikunja-handoff.md
execution_mode: planning_artifact
owned_files:
- kitty-specs/module-path-v1-9-0-01M62K6W/drafts/release-communication.md
- kitty-specs/module-path-v1-9-0-01M62K6W/drafts/teambition-transition-pr.md
- kitty-specs/module-path-v1-9-0-01M62K6W/evidence/migration-verification.md
- kitty-specs/module-path-v1-9-0-01M62K6W/evidence/vikunja-handoff.md
role: implementer
tags: []
task_type: verification
tracker_refs: []
---

# WP03 — Verify the migration and complete private handoff packets

## ⚡ Do This First: Load Agent Profile

Use the `/spk-charter-profile-load` skill to load the agent profile specified in the frontmatter, and behave according to its guidance before parsing the rest of this prompt.

- **Profile**: `implementer-ivan`
- **Role**: `implementer`
- **Agent/tool**: `codex`

If the assigned profile cannot be loaded, report the failure to the orchestrator rather than silently substituting a role.

---

## Objective

Verify the completed local module-path migration using the approved lightweight approach, preserve durable private evidence, and finish the reviewable release/upstream packets.
Prepare a later Vikunja handoff without contacting any session or modifying Vikunja.
This WP produces implementation evidence and artifacts; it is not its own independent review.

## Context and prerequisites

- Mission: `module-path-v1-9-0-01M62K6W`.
- WP02 must be approved/done before this WP can be claimed; its dependency on WP01 supplies the candidate code and consumer guide.
- Read `spec.md`, `plan.md`, `quickstart.md`, `contracts/module-migration.md`, the approved occurrence map, and both post-plan review records.
- Read the candidate `MIGRATION.md`, README migration sections and Unreleased changelog before validating their commands.
- The unchanged source baseline is commit `737ae01be40048d1d2932dbafb623230efa8d361`.
- That baseline already includes approved DST/time-form work. Do not compare it as though it were the old pre-DST release.
- Kent approved simpler verification: explicit declaration inspection, disposable consumers, documented command walkthroughs and existing golden outputs.
- No synthetic proxy, registry, compatibility shim, extra dependency or permanent identity-test framework is needed.
- Actual public-tag retrieval remains deferred until a separately authorized release.

## Branch strategy and ownership

Planning/consolidation target is `rrule-go-dev-xfer`; eventual private PR target is `main`.
Use the lane/workspace allocated by Spec Kitty from `lanes.json`; do not create a worktree or select a base manually.
All owned deliverables are confined to the four private mission paths in frontmatter.
Do not edit `go.mod`, root consumer documentation, test imports, reproduction drivers, golden README, historical JSON, the specification, plan or workflow state.
WP01 owns source/import/instruction corrections; WP02 owns root consumer documentation.
If verification exposes a defect there, send findings to the orchestrator for the owning WP's workflow rework/review; do not patch an unowned surface to make this WP pass.
Do not revert other agents' work. Re-read candidate state before interpreting earlier logs.

The orchestrator owns claims, lane transitions, review verdict capture and consolidation.
The claim form, when authorized, is:

```sh
SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty agent action implement WP03 --agent codex --mission module-path-v1-9-0-01M62K6W
```

Record completed subtasks through the event-sourced task surface, not Markdown checkboxes:

```sh
SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty agent tasks mark-status T011 --status done --mission module-path-v1-9-0-01M62K6W
```

Have the orchestrator execute workflow transitions as required by the active handoff contract.
Prefix every Spec Kitty command with `SPEC_KITTY_NO_UPGRADE_CHECK=1`.
On an unexpected workflow/source/permission/state failure, stop, preserve the exact command/output and report it; do not invent a bypass.

### Subtask T011: Verify declaration identity and a local external consumer

**Purpose**: Prove this source declares the intended identity and is usable through the new import, without claiming distribution has been tested.

**Steps**:
1. Record candidate commit, source directory, toolchain and relevant source fingerprints in `evidence/migration-verification.md`.
2. Inspect candidate `go.mod`: module must be `github.com/kentonium3/rrule-go`; `go 1.16` remains unchanged.
3. Capture the explicit declaration assertion failing against old source and passing against candidate source:

   ```sh
   GOWORK=off GOPROXY=off go list -m
   test "$(GOWORK=off GOPROXY=off go list -m)" = github.com/kentonium3/rrule-go
   ```
4. Prefer WP01's actual pre-edit transcript. If it is unavailable or ambiguous, reconstruct the red case from a pristine `git archive` export of the pinned baseline into a temporary directory.
5. A read-only source export is not a new workflow worktree. Never reset the implementation workspace to recreate red.
6. Record that failure specifically reports/matches the old declaration, not missing tools, an invalid directory or a dependency-download failure.
7. Create a small disposable module outside the repository, importing `github.com/kentonium3/rrule-go`.
8. Require the candidate version and add a fixture-only replacement from the **new** path to the absolute candidate directory.
9. Use `GOWORK=off GOPROXY=off` for private fixture commands; leave global proxy/checksum configuration alone.
10. Run tidy, inspect `go list -m -json all` and run the consumer's tests.
11. Exercise a fixed UTC DAILY/COUNT parse and expected timestamps/serialization; use keyed `ROption` literals if a constructor is included.
12. Record the selected replacement directory and fixture source so a reader can identify what was actually tested.

**Files**: Write this WP's owned `evidence/migration-verification.md`; temporary consumer files are disposable, never production replacements.

**Validation**: Same declaration assertion red/green is explained, candidate consumer succeeds and its selected source is explicit.
A local replacement can mask a declaration mismatch, so consumer success alone is insufficient.
Do not execute or report a public `go get` as part of this private check.

### Subtask T012: Walk all three consumer migrations and the dependency-safety stop

**Purpose**: Verify the actual guide, including the review finding that replacement removal can restore unfixed upstream code.

**Steps**:
1. Prepare three small disposable consumer states: direct original dependency, all-version old-path replacement to fork v1.8.3, and version-specific old-path replacement to fork v1.8.3.
2. Preserve the original `go.mod`/import text in the evidence before normalizing local source locations.
3. Use baseline source for an old-state working fixture and candidate source for its migrated state, through clearly labeled fixture-only local replacements.
4. Keep local location scaffolding separate from the production migration instructions; do not imply the published `go get` was executed.
5. Follow the guide's import/dependency changes, checking its actual command arguments rather than silently correcting them during the walkthrough.
6. For the all-version form, check removal with `go mod edit -dropreplace=github.com/teambition/rrule-go`.
7. For the version-specific form, check removal with `go mod edit -dropreplace=github.com/teambition/rrule-go@v1.8.2`, substituting the actual left-side version if different.
8. Establish that the intended obsolete directive is gone and the new import selects candidate source; then tidy, build/test and inspect the resulting graph.
9. Old direct requirements should disappear when unused. A remaining old-path package consumer is not harmless merely because tests pass.
10. Walk a retained-old-import dependency case before protective replacement removal. Inspect both module requirements and package-import evidence.
11. Confirm the guide requires coordinating that consumer's migration first or remaining on the pinned v1.8.3 arrangement; do not demonstrate blind removal as a success.
12. Explain that losing replacement-provided panic/hang fixes can occur even without old/new rrule types crossing interfaces.
13. Confirm distinct named-type guidance and the separate keyed-`ROption` migration are present, without constructing another Go type-system test suite.

**Files**: Record each initial/final state, command result and safety-stop outcome in `evidence/migration-verification.md`.

**Validation**: All three intended migrations succeed under private scaffolding; the retained-old-consumer case reaches the documented stop/coordinate/stay-pinned outcome.
If commands are wrong, report the smallest correction to the orchestrator/WP02 and rerun after authorized rework.
Do not edit the root migration guide from this WP.

### Subtask T013: Prove zero new recurrence differences

**Purpose**: Preserve the approved current behavior rather than reclassify earlier DST/time-form changes.

**Steps**:
1. Export pristine source at `737ae01be40048d1d2932dbafb623230efa8d361` to a temporary directory using `git archive` extraction.
2. Build that snapshot's **own old-import** golden driver from its own cwd.
3. Build the candidate's **new-import** driver from candidate cwd, with `GOWORK=off GOPROXY=off` for both builds.
4. Record source paths/commits, module identities and binary hashes before running fixtures.
5. Use the same unchanged absolute `corpus-dst.json` path for both binaries, distinct `-out` files, and repeat for `corpus-forms.json`.
6. Require 811 DST rows and 158 forms rows, matching IDs/order and nonempty complete outputs.
7. Use `cmp` on each baseline/candidate output pair and require byte equality.
8. Inspect unexpected timeouts/errors: equal broken or timed-out outputs are not proof of a healthy migration. Existing expected parser-rejection rows remain valid corpus behavior.
9. Corrupt only a temporary output copy and confirm comparison fails, demonstrating that the comparison can detect a difference.
10. Do not use the historical `-baseline`/`-current` classified comparator for this assertion: it intentionally fails when `changes == 0`.
11. Do not alter old76c646d baseline JSON, corpus JSON, expectations or historical result files.
12. Check WP01's golden header/README correction now directs readers to properly imported source-specific drivers.
13. Do not run every standalone panic/hang reproduction without bounds; use the existing bounded harness or compile changed drivers only.

**Files**: Record commands, source/corpus/binary/output hashes, counts, zero-difference results and sensitivity-control outcome in `evidence/migration-verification.md`.

**Validation**: Both corpora have zero migration differences and a working positive consumer control from T011.
Temporary result paths are not durable evidence after cleanup; preserve sufficient hashes, commands and summaries in the owned report.
A source/output mismatch or unexplained error is a failure to report, not permission to adjust baselines.

### Subtask T014: Account for every required repository check

**Purpose**: Establish the final candidate's required checks without unnecessary repeated runs.

**Steps**:
1. Inventory WP01/WP02 evidence and candidate source fingerprints before deciding what must run again.
2. Account for every required command:

   ```sh
   go test ./...
   go test -race ./...
   go vet ./...
   gofmt -l .
   go run ./tools/changelog check
   ```
3. A previous result is reusable only when its exact command, toolchain, outcome and relevant source fingerprint are available and unchanged.
4. Run any missing check, including race tests, and repeat a check when new changes or unresolved failures invalidate earlier evidence.
5. Require zero filenames from formatting; a zero exit alone does not mean formatting is clean.
6. Record exit/results with relevant source identity. Do not claim checks executed here when they were reused from WP01/WP02.
7. Inspect the candidate diff for unchanged exported declarations, package name, Go minimum, dependencies, MIT license and historical sections.
8. Recheck the approved map against actual changed files and intentional remaining old references; approval does not replace this technical verification.
9. Return any source/documentation failure to the orchestrator and owning WP rather than editing unowned files or waiving a gate.

**Files**: Consolidate the verification ledger and actual implementation model/date in `evidence/migration-verification.md`.

**Validation**: Every required check has applicable passing evidence and all intended preservation assertions are supported.
No runtime performance/new fuzz infrastructure is added solely for this identity migration.
This WP's evidence author is not its independent reviewer.

### Subtask T015: Complete the private release and upstream draft packets

**Purpose**: Make Kent's later exact-copy review concrete, accurate and bounded by what has been verified.

**Steps**:
1. Reconcile `drafts/release-communication.md` with the final consumer guide and selected Unreleased entries.
2. Preserve the distinction between locally prepared/implemented candidate work and an actually available public tag.
3. Keep after-release copy clearly gated on real tag declaration, clean consumer retrieval and separate public-action approval.
4. Retain both replacement-removal forms and the retained-old-consumer safety check before removal commands.
5. Keep keyed-`ROption` migration as separately inherited v1.9.0 information; do not claim this WP implemented that behavior.
6. Reconcile `drafts/teambition-transition-pr.md` title/body, proposed README notice and exact advisory deprecation diff.
7. Preserve the upstream declaration `github.com/teambition/rrule-go`; original maintainers decide about README/deprecation and any final old-module release.
8. Use “separately maintained fork,” retain the README-only alternative and do not assert official succession, endorsement or transfer.
9. Cite the dated existing maintenance audit and explicitly preserve the limits on knowledge of current permissions/private conversations.
10. Explain opt-in GitHub Release notifications, latest-version deprecation lookup and unconditional-replacement caveats; promise no universal warning or automatic migration.
11. Scan proposed public excerpts for private mission IDs, internal paths and workflow records; private wrapper/evidence sections must remain clearly separate.
12. Keep exact-copy approval and future public availability checks unresolved where genuinely not performed, rather than turning prerequisites into claimed successes.

**Files**: Modify only the two owned `drafts/*.md` files; document reconciliation results in the verification report.

**Validation**: Packets contain concrete proposed copy and prerequisites while making no unsupported availability or relationship claim.
Do not post a draft, create a PR/tag/release, request indexing or contact an upstream maintainer.

### Subtask T016: Prepare the private public-port and Vikunja handoff record

**Purpose**: Make later adoption/integration reviewable while preserving authorization boundaries.

**Steps**:
1. Create `evidence/vikunja-handoff.md` as a private handoff draft, not a sent message or task.
2. Name the verified candidate commit, module identity, import changes and applicable replacement-removal/safety guidance.
3. List the relevant Vikunja recurrence/application checks to request later; retain service-level bounds/guards as defense in depth.
4. State that integration sessions may exercise adoption only when ready and Kent authorizes coordination; do not choose a recipient or send anything now.
5. Record the proposed later public-port allowlist: `go.mod`, `example_test.go`, `str_timeforms_api_test.go`, `MIGRATION.md`, selected README migration/install/reference hunks, and the matching Unreleased changelog entry.
6. Explicitly exclude private mission records, `docs/repro`, authority files and private README preamble from that port.
7. Require reconciliation with the current public README rather than wholesale replacement of the private README.
8. List deferred release checks: confirm candidate version is unused, verify actual public tag declaration, clean normal `go get`, checksum/proxy retrieval, package documentation and separately approved release metadata.
9. If v1.9.0 has already shipped, stop for version direction; do not rewrite it.
10. Summarize remaining permissions and readiness conditions without scheduling, messaging, editing Vikunja or claiming integration was executed.

**Files**: Write `evidence/vikunja-handoff.md`; cross-reference it from the owned verification report.

**Validation**: Kent can later review a concrete handoff and allowlist without any external action having occurred.
Future integration is complementary verification, not a substitute for this WP's local evidence.

## Definition of Done

- T011–T016 have concrete evidence and are recorded complete through the event log by the authorized workflow actor.
- The verification report distinguishes observed red/green checks, reused results, limitations and deferred checks.
- All three migration walkthroughs pass; the retained-old-consumer example enforces coordination/staying pinned.
- Both golden corpora show complete, correctly sourced, byte-identical outputs with a comparison sensitivity control.
- All required repository checks have applicable passing evidence; approved map has been technically rechecked.
- Both draft packets are complete, conditional and relationship-accurate; the private allowlist/Vikunja handoff is ready for Kent's later review.
- No unowned source, root documentation, corpus, historical release or workflow state was manually altered.
- No public or cross-session action was taken.

## Risks and safeguards

Local `replace` may hide identity errors: pair the consumer build with explicit declaration evidence.
Retained old imports may regain unfixed upstream: enforce preflight before removal, even when tests pass.
Wrong snapshot/driver/cwd may produce irrelevant golden evidence: record provenance and source-specific drivers.
Equal failed outputs may pass a blind comparison: inspect timeouts/errors and pair with the consumer positive case.
Private/public copy can leak internal artifacts or imply a handoff: review exact excerpts and preserve authority limits.
Do not compensate for missing prerequisites by adding a proxy, shim, new test framework or another dependency.

## Reviewer guidance

A different actual model must independently review this WP; record implementer model, reviewer model and review date.
A different role/profile running the same model is a degraded check and does not silently satisfy that rule.
Verify the report's commands/source fingerprints and limits, not merely the presence of the report files.
Challenge unsupported public-download claims and evidence that depends only on a local replacement.
Check the dependency-safety preflight, both replacement forms, golden comparator choice, row counts and immutable historical JSON.
Confirm draft text does not assert official succession or a publicly available candidate without release verification.
Check the public-port allowlist and the unsent Vikunja handoff for privacy and authorization boundaries.
Reviewers report findings through the prescribed workflow; they do not fabricate event-log, approval or lane state.


╔==============================================================================╗
║  WORK PACKAGE PROMPT ENDS                                              ║
╚==============================================================================╝

================================================================================
🎯 REVIEW COMPLETE? RUN ONE OF THESE COMMANDS:
================================================================================

✅ APPROVE (no issues found):
   spec-kitty agent tasks move-task WP03 --to approved --mission module-path-v1-9-0-01M62K6W --agent codex --note "Review passed: <summary>"

❌ REQUEST CHANGES (issues found):
   1. Write feedback to the in-repo path (committed with the project):
      cat > ~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-coord/kitty-specs/module-path-v1-9-0-01M62K6W/tasks/WP03-verification-handoff/review-feedback-1.md <<'EOF'
**Issue 1**: <description and how to fix>
**Issue 2**: <description and how to fix>
EOF

   2. Move to planned with feedback:
      spec-kitty agent tasks move-task WP03 --to planned --review-feedback-file ~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-coord/kitty-specs/module-path-v1-9-0-01M62K6W/tasks/WP03-verification-handoff/review-feedback-1.md --mission module-path-v1-9-0-01M62K6W --agent codex

⚠️  NOTE: You MUST run one of these commands to complete the review!
     The Python script handles all file updates automatically.
================================================================================