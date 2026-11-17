# Implementation Plan: rrule-go v1.9.0 module path migration

**Branch**: `rrule-go-dev-xfer` | **Date**: 2026-10-28 (America/New_York) | **Spec**: [spec.md](spec.md)
**Input**: Confirmed scope and engineering decision `01M62N0BZ44HJDCVQPNTC7D0F8`.

## Summary

Change the maintained fork's module declaration and active imports to `github.com/kentonium3/rrule-go`; keep package name `rrule`, exported declarations and recurrence code unchanged. Classify occurrences instead of replacing every old URL. Document direct adoption and migration from direct upstream use and both replacement forms. Keep historical releases, prior evidence, MIT attribution and the original upstream's identity intact.

Kent approved simpler verification: an explicit declaration/import check, a small external consumer using a local candidate replacement, walkthroughs of the migration commands, ordinary repository gates, and a zero-difference golden comparison. Do not build a module proxy, custom registry, identity-testing framework, compatibility shim, or notification service. A clean public-tag download is an explicitly deferred release check. Later Vikunja core-change sessions can exercise adoption when Kent authorizes that coordination; do not contact or message them now.

## Technical Context

**Language/Version**: Go library; retain `go 1.16` in go.mod. Use the existing installed Go toolchain for verification; the inherited golden driver uses newer test-tool APIs and is not part of the library's Go 1.16 runtime surface.
**Primary Dependencies**: Standard library only; no new or upgraded third-party dependencies.
**Storage**: Source and Markdown/YAML planning files; ephemeral consumer directories and evidence outputs. No application database or persistent runtime state.
**Testing**: Explicit declaration check; external consumer and three migration walkthroughs; existing unit/race/vet/format/changelog gates; 811 DST plus 158 forms golden rows against this mission's pre-migration source.
**Target Platform**: Existing supported Go environments; no new platform-specific binary distribution.
**Project Type**: Single Go module/library; private authoring and separate public distribution repositories.
**Performance Goals**: No recurrence algorithm or runtime performance change; zero golden-output differences.
**Constraints**: No publication, public indexing requests, messages to other sessions/people, tag changes, API additions, new dependencies, compatibility shim or synthetic proxy. Preserve license and attribution.
**Scale/Scope**: One declaration, two root external-test imports, seven runnable reproduction-driver imports, selected README/Unreleased/migration documentation; historical occurrences preserved. No filesystem/module-directory moves.

## Engineering Alignment and Branch Contract

The confirmed approach is a direct identity rename with local candidate checks. A local replacement does not prove public downloadability or declaration agreement; the explicit declaration check covers identity, the consumer build covers API usability, and an eventual public-tag retrieval covers distribution. These are separate claims.

Current/planning/consolidation branch: `rrule-go-dev-xfer`, matching the runtime target. Pre-mission source base: private `main` at `737ae01be40048d1d2932dbafb623230efa8d361`. Eventual private PR base: `main`; no PR submission or merge is performed in this planning step. Public code porting follows `docs/runbooks/changelog-flow.md` under later exact-copy approval.

## Charter Check

| Rule | Plan disposition |
|---|---|
| Compatibility/minimal change | Rename identity only. Preserve algorithms and declarations; compare current-baseline and candidate outputs. Existing DST/time-form fixes remain baseline, not new work. |
| Test-first evidence | Before editing, record the expected new-declaration check failing on the old declaration. After editing, repeat the same check successfully and run consumer/golden checks. Failure must be identity-specific, not missing tooling. |
| Go/API/license | Preserve Go 1.16 floor, exported declarations, package name and MIT attribution. No new dependency. |
| Independent review | Read-only plain Codex post-plan review on spec/plan/research/data-model/map/contracts; repair findings before tasks. Different-model WP implementation/review and integrated final review remain later gates. |
| Scope authority | Kent's confirmed migration scope supersedes the earlier DST/UNTIL-only charter scope. Its substantive compatibility and privacy rules still apply. |
| Privacy | No public writes or contacts. Separate candidate verification from later public availability verification. |
| Occurrence classification | All eight categories and file exceptions appear in occurrence_map.yaml. Schema/admissibility validation and Kent's map approval precede implementation. |

Re-evaluate after review. No dependency-install decision requires a third-party supply-chain addition: no package is added, upgraded or removed in the library. Consumer changes select the same reviewed library under its new identity. Verify local source provenance, leave process-global proxy/checksum settings untouched, and use GOPROXY=off for private checks. No security-impacting third-party dependency decision is introduced; no additional supply-chain squad is needed for this plan.

## Project Structure

- `go.mod`: new module identity; minimum Go line unchanged.
- `example_test.go`, `str_timeforms_api_test.go`: new external imports; no assertion/algorithm rewrites.
- `docs/repro/{01-iterator-panics,02-unreachable-byset-hang,03-until-scan,05-07-backlog,edge-cases,matrix}/main.go` and `docs/repro/05-07-backlog/golden/main.go`: imports follow local module. Existing historical output/corpus files stay unchanged.
- `README.md`: active example/install/reference links; retain source attribution. Private README text is not copied wholesale into public.
- `CHANGELOG.md`: consumer Breaking/Upgrade Notes entries under Unreleased; preserve prior sections and existing entries.
- `MIGRATION.md` (new consumer guide): verified migration commands and dependency-safety preflight. Private verification details stay in this mission. Public excerpts contain no internal IDs/paths.
- `kitty-specs/module-path-v1-9-0-01M62K6W/`: plan, research, data-model, contracts, quickstart, occurrence map, drafts, reviews and evidence. No task/WP generation in this phase.

No production code restructuring. README's unrelated stale wording/badges are classified explicitly; fix only identity-related active pointers here, preserving upstream provenance and leaving unrelated documentation cleanup out of scope.

## Design and Verification

### 1. Occurrence policy

Use `occurrence_map.yaml` and the accompanying inventory in research.md. Inventory the exact module-path token with tracked-file searches, then review broader repository URLs separately. Never change `teambition` to `kentonium3` globally. The whole-file manual-review exceptions for README/changelog allow mixed old/new meanings, not wholesale replacement. Historical source links, inherited release records, issue briefs, old evidence and unrelated missions stay unchanged.

Kent approved the map before task generation; `status.approved: true` records that explicit approval. Schema validity and a favorable review alone are not approval, and technical inventory rechecks remain mandatory. Refresh inventory at implementation start and resolve new occurrences against policy.

### 2. Simplest adequate identity/consumer check

In the unchanged source, capture `go list -m` and compare its output with the intended new path; record the specific expected mismatch. After changing declaration/imports, the identical check must pass. Inspect candidate go.mod explicitly and scan remaining old-path imports. No permanent test framework or generated proxy is needed.

Create an ephemeral consumer module outside the repository. Require `github.com/kentonium3/rrule-go v1.9.0` and replace that new path with the absolute candidate directory **for local testing only**. Import the new path and call parsing plus a small deterministic recurrence/serialization example using keyed ROption literals. Set `GOWORK=off GOPROXY=off` for those commands. Record `go list -m -json` including the replacement directory. Remove the temporary directory afterward or retain only sanitized evidence privately.

This local replacement is a known limitation, not a distribution proof. Explicit identity inspection is mandatory because replace can hide a declaration mismatch. Do not add production replace directives or recommend this development-only replacement to consumers.

### 3. Migration walkthroughs

Before removing a protective old-path replacement, inspect module requirements and package imports for remaining old-path consumers. Removal can silently restore original upstream code and lose the fork's panic/hang fixes even when no rrule types cross interfaces. Coordinate those consumers' migration first, or remain on the existing pinned v1.8.3 replacement arrangement. Passing application tests alone do not establish preservation of those fixes.

Add a retained-old-import dependency walkthrough showing that the guide gates removal and calls for coordination or remaining pinned; no new test framework is needed.

Prepare three disposable consumer go.mod files: direct original dependency; unconditional old-path replacement to fork v1.8.3; version-specific old-path replacement to fork v1.8.3. Record their initial text and expected dependency/import state. Verify each begins from a working old-source fixture using the preserved pre-migration source and controlled local replacements; distinguish that private scaffolding from the actual user-facing commands.

Apply the documented import change; remove the original replacement using the exact left-side form (`go mod edit -dropreplace=github.com/teambition/rrule-go` or `...@v1.8.2`). For private execution replace the new path with candidate source and require its candidate version locally; then run tidy/build/tests without network. Verify no obsolete replacement remains, the selected new module resolves to candidate, and direct old requirements disappear when unused. Do not treat retained old-path consumers as harmless: apply the dependency-safety gate above before removal. Public guide uses normal `go get newpath@v1.9.0`; clearly label that actual command's network retrieval as a later release check, not executed successfully here.

Document distinct old/new named types with a short example and coordinated-update instructions. Do not build a separate compiler-behavior test suite. Migration command success and a representative API call are sufficient here.

### 4. Zero-change golden proof

Update the golden driver header and docs/repro/05-07-backlog/golden/README.md active build instructions: use a historical snapshot's old-import driver for old source and the candidate's new-import driver for candidate source. Never tell readers to build a new-import driver against an old-declaration snapshot. Preserve historical facts, corpora and baseline JSON.

Keep the inherited corpora and old76c646d baseline JSON immutable. The inherited comparator intentionally fails on zero changes and therefore is not the migration comparator.

Read a pristine source snapshot of `737ae01` using `git archive` extraction into a temporary directory (read-only source export, not an unmanaged worktree). Build its own old-import golden driver from that baseline cwd. Build the candidate's new-import driver from candidate cwd. Use the same unchanged absolute corpus files for both, emit separate outputs with `-corpus`/`-out`, assert 811 DST and 158 forms rows with matching IDs/order, and compare each before/after output byte-for-byte with `cmp`. Record zero differences. Check for timeout rows or unexplained errors rather than treating two equal failures as success; expected parser-rejection rows remain corpus behavior. Corrupt a temporary copy of a known output to demonstrate the comparison detects a difference. Do not modify pristine evidence.

A current-baseline comparison preserves already approved DST/time-form behavior. Optional historical classified comparison can remain as supporting evidence but must not rewrite historical baselines or demand zero differences from the pre-DST release. Repro programs that intentionally panic or time out are not all executed without bounds; compile their import changes, and use the existing bounded drivers/tests for behavior checks.

### 5. Documentation and relationship boundary

Prepare complete verified migration prose, the Unreleased entry, and release excerpt together. Update drafts with exact commands and distinguish prepared/implemented/publicly released states. Retain positional ROption migration as separate inherited v1.9.0 information. Keep previous release sections truthful.

The upstream packet proposes a separately maintained fork pointer, offers README-only communication and an advisory deprecation alternative, preserves the original module declaration, and asks original maintainers to decide about a final old-module release. No official succession is asserted. Existing maintenance audit is dated evidence, not permission to contact anyone or proof of current access.

### 6. Later integration and release checks (not executed now)

The proposed later public-port allowlist is go.mod, example_test.go, str_timeforms_api_test.go, MIGRATION.md, selected README migration/install/reference hunks, and the matching Unreleased CHANGELOG.md entry. Any additional file needs explicit review. Private docs/repro, mission records, agent context and the private README preamble are excluded; reconcile selected hunks with current public README rather than overwrite it.

When code is ready, prepare a short private Vikunja handoff with candidate commit, import/replacement changes, test commands and requested recurrence scenarios. Kent can then authorize contacting the relevant core-change sessions. No session messaging, Vikunja source edits, coordination automation or task creation is authorized here. Vikunja integration is complementary and does not block this planning phase.

After separately approved public code port and tag: verify the actual tag's declaration, a clean new consumer's `go get github.com/kentonium3/rrule-go@v1.9.0`, checksum/proxy retrieval and pkg.go.dev documentation. Check that v1.9.0 has not already shipped before publication; stop for version direction if it has. Publish a GitHub Release only with separate approval; tag creation alone does not deliver GitHub Release notifications. Upstream deprecation is a separate maintainer action and cannot promise warnings through unconditional replacements.

## Risks and Conditions

| Condition and likelihood | Effect | Response |
|---|---|---|
| Mixed history and active references; certain | Blind replacement falsifies provenance | Explicit per-file policy and manual review. |
| Local replacement masks declaration mismatch; realistic | Consumer build can give false confidence | Explicit declaration check plus recorded source directory; actual download later. |
| Old-path consumer survives replacement removal; realistic | Original upstream code may regain execution and lose fork fixes | Dependency/import preflight; coordinate adoption or stay pinned. |
| Old-path dependency exposes public rrule types; plausible | Old/new values cannot be exchanged directly | Explain coordinated adoption; no shim. |
| Reuse of inherited comparator for zero-change test; likely without care | Correct migration wrongly fails | Fresh before/after outputs and cmp, preserving old corpora. |
| Wrong driver/cwd loads unintended source; plausible | Golden evidence tests wrong library | Use each snapshot's own driver and record module/source paths. |
| Private README/mission artifacts copied to public; realistic | Privacy leak | Explicit public-file/content allowlist and exact-copy review. |
| Contributor history mistaken for exclusive authority; plausible | Misleading takeover claim | Neutral fork wording and upstream decision boundary. |

## Implementation Concern Map

Concerns inform later decomposition; they are not executable work packages.

### IC-01 — Identity and classified source references
- **Purpose**: Change declaration and active imports with preserved history.
- **Relevant requirements**: FR-001, FR-002, C-002, C-003.
- **Affected surfaces**: go.mod, two external test imports, seven reproduction imports and golden README/header instructions, README active pointers, occurrence map.
- **Dependencies**: Approved occurrence map before edits.
- **Risks**: Mixed semantics and accidental old-module resolution.

### IC-02 — Lightweight consumer and equivalence evidence
- **Purpose**: Prove local usability, command correctness and unchanged recurrence outputs without new infrastructure.
- **Relevant requirements**: FR-003, FR-008, NFR-001, NFR-002, SC-001, SC-002.
- **Affected surfaces**: Disposable consumer directories, baseline source export, golden driver/corpora, private evidence and quickstart.
- **Dependencies**: IC-01 candidate for green checks; capture pre-change declaration failure before IC-01.
- **Risks**: Local-replace masking, wrong cwd and equal failing outputs.

### IC-03 — Consumer/release communication and upstream proposal
- **Purpose**: Deliver reviewable migration and relationship copy.
- **Relevant requirements**: FR-003 through FR-007, NFR-003, C-001, C-004, C-005, SC-003, SC-004.
- **Affected surfaces**: README, Unreleased changelog, migration guide, draft packet, maintenance audit, future allowlist/Vikunja handoff.
- **Dependencies**: Verified commands/evidence from IC-02 before finalizing claims.
- **Risks**: Premature availability claims or implied upstream endorsement.

## Review and Exit

Run plain `codex exec -s read-only` with prompt on stdin, no write-enabled review profile, on spec/plan/research/data-model/contracts/occurrence-map/quickstart. Require concrete conditions and likelihood for findings; challenge excessive machinery as well as missing coverage. Save prompt, output and dispositions privately. Apply significant findings and re-review affected sections. This checkpoint is separate from later WP and post-consolidation reviews.

Finish planning only after the plan is substantive, map validation passes, decisions verify clean and the independent review is resolved. Stop before tasks or implementation. Kent's occurrence-map approval is now recorded; refresh inventory before code edits.

### Recorded post-plan outcome

Independent plain Codex read-only review found two actionable issues: retained old-path dependencies can lose protective replacement fixes, and historical golden commands need old/new driver separation. Both were corrected in plan/spec/contracts/map/quickstart/drafts as applicable. Focused independent rereview confirms both resolved, with no significant planning blockers. Review prompts, findings and dispositions are in reviews/. The occurrence map passes schema/runtime/admissibility validation; Kent approval is recorded in evidence/occurrence-map-approval.md; technical inventory rechecks remain required before implementation. No task generation, library changes, integration contact or public verification occurred.

### Occurrence-map approval received

Kent approved the map on 2026-10-28 and authorized task generation, while explicitly delegating confidence in individual entries to the agents. Technical responsibility remains with implementers/reviewers; refresh inventory before edits. See evidence/occurrence-map-approval.md. Earlier review-time pending-approval statements are superseded by this approval record; no policy entry changed. No implementation or public action is authorized by this planning/task-generation approval.
