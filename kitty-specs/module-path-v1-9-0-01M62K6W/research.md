# Planning research and decisions

Date: 2026-10-28 (America/New_York). Planning only; no library edits or public actions.

## Approved verification proportionality

**Decision:** Use an explicit declaration/import check, disposable external consumers, migration command walkthroughs and existing golden outputs. Do not build a synthetic module proxy or permanent identity-validation framework.

**Rationale:** Go module naming/replacement behavior is established. Verify this repository's edits and instructions instead of testing the ecosystem. A local replace can mask a declaration mismatch, so inspect go.mod separately and record the selected source. Actual public downloadability is not claimed until a later authorized tagged release is checked.

**Alternatives:** Temporary file:// proxy would exercise the real download validator privately, but Kent selected the simpler approach. Old-path shims/type-alias bridges and duplicate modules add maintenance/API machinery without a requirement; reject them. Do not insert runtime update warnings or network calls.

**Authority:** Planning decision `01M62N0BZ44HJDCVQPNTC7D0F8`. This resolves the earlier adjustment request. Future Vikunja integration is a useful complementary check, with contact/coordination deferred until ready and authorized.

## Occurrence inventory

Tracked exact-token scan at planning start found 122 files: 102 mission records, one agent-instruction file and 19 other active/history files. Counts are a starting snapshot, not an implementation acceptance threshold. Refresh before edits.

| Surface | Classification |
|---|---|
| go.mod | Rename module declaration only; retain go 1.16. |
| example_test.go; str_timeforms_api_test.go | Rename external import only. |
| Seven docs/repro/*/main.go drivers, including nested golden/main.go | Rename active local-module import only. Retain logic and recorded corpus/output data. |
| README.md | Change active demo/install/reference pointers. Preserve Teambition source/license attribution. CI/CodeQL links identify the actual repository on each side of the private/public handoff; no blind prefix replacement. Existing unrelated Codecov/stale prose cleanup is outside this migration. |
| CHANGELOG.md | Add an Unreleased migration entry; preserve v1.8.3's truthful old declaration and historical upstream links. |
| docs/issues/01,02,04,05,06 | Historical reports, examples and adoption records: preserve. |
| docs/releases/v1.9.0-dst.md | Prior-work provenance links: preserve. |
| docs/runbooks/evidence/teambition-backlog-step0-charter.txt | Recorded original remote evidence: preserve. |
| Prior kitty-specs records and AGENTS.md | Preserve as historical/authority context. This plan does not edit agent context files. |
| Current mission and new migration guide | Deliberate old/new language; manual review, never blanket rename. |

The golden README has no exact-token match but its commands depend on identity, so it is also in scope. Update the golden driver header and docs/repro/05-07-backlog/golden/README.md active build instructions: use a historical snapshot's old-import driver for old source and the candidate's new-import driver for candidate source. Never tell readers to build a new-import driver against an old-declaration snapshot. Preserve historical facts, corpora and baseline JSON.

The executable policy is occurrence_map.yaml. All eight categories are explicit. File-level manual_review means inspect each hunk, not that every token in that file should change. Schema/path heuristics do not understand semantic meaning and cannot replace the human review. No structural moves or structural-exemption bypasses are needed. No code edits until Kent approves the map.

## Golden harness finding

**Decision:** Collect fresh before/after outputs from the two corpora and compare bytes, preserving the inherited historical baselines.

**Evidence:** `docs/repro/05-07-backlog/golden/main.go` accepts `-corpus`, `-out`, `-baseline`, `-current`; its historical comparator exits nonzero when `changes == 0`. It expects approved differences against old76c646d, so it cannot directly prove this migration changes nothing.

**Rationale:** Baseline private main already includes accepted DST/time-form corrections. Export source at `737ae01be40048d1d2932dbafb623230efa8d361`; run that snapshot's old-import driver and candidate's new-import driver against identical corpus files. Assert 811 DST and 158 forms rows and zero new differences. Check subprocess timeouts/errors; pair equality with working deterministic consumer behavior and a temporary corrupted-output sensitivity control.

**Alternatives:** Rewriting old corpus baselines would erase provenance; using the old-import driver in the candidate cwd can silently load unintended code. Reject both. Historical classified comparison remains optional supporting evidence, not the zero-change check.

## Dependency and interface decisions

No new/updated third-party dependency, registry, server, schema, HTTP API or exported Go identifier is introduced. Standard Go tooling and existing module source suffice. Local candidate replaces are scoped to disposable consumer fixtures, not library go.mod. Set GOWORK=off and GOPROXY=off for private candidate commands; leave ordinary checksum settings intact. Avoid downloads and executing unreviewed dependencies.

Existing exported API declarations stay the same; import identity intentionally changes. Named types from old/new import paths are distinct. Explain coordinated migration if dependencies expose those types. Do not build another test suite for Go's type system.

## Publication and upstream relationship

**Decision:** Describe a separately maintained fork and propose upstream changes for maintainer consideration. Keep upstream module declaration unchanged in the deprecation draft. No official succession claim, contact or publication.

**Evidence:** `evidence/upstream-maintenance-audit.md` records the latest upstream Go-source change in January 2023 and master commit in April 2023, historical merger roles for zensh and damoye, recent unanswered community activity and no public handoff approval in the inspected record. Public history cannot enumerate present permissions or establish universal abandonment. The private author conversation was not independently inspected.

**Release prerequisites:** v1.9.0 remains a candidate target. Recheck public tags before an eventual public port/release; if already used, stop for version direction. Confirm exact public README/copy rather than overwrite it with private content. Go deprecation metadata requires latest-version lookup and may be suppressed by unconditional replacement. A final upstream tag requires upstream authority. GitHub Release notifications are opt-in and are separate from tags.

## Planning review disposition

Independent plain Codex read-only post-plan review is required after artifacts are complete. Record its prompt/output and each significant finding with conditions, likelihood and disposition under reviews/. This section will be updated with the actual outcome; no review success is presumed.

## Independent post-plan findings and dispositions

- P1, realistic retained-old-path dependency: accepted. Before removing a protective old-path replacement, inspect module requirements and package imports for remaining old-path consumers. Removal can silently restore original upstream code and lose the fork's panic/hang fixes even when no rrule types cross interfaces. Coordinate those consumers' migration first, or remain on the existing pinned v1.8.3 replacement arrangement. Passing application tests alone do not establish preservation of those fixes. Plan, specification, contract, quickstart and drafts now include this gate and a lightweight walkthrough.
- P2, certain if following old golden commands: accepted. Update the golden driver header and docs/repro/05-07-backlog/golden/README.md active build instructions: use a historical snapshot's old-import driver for old source and the candidate's new-import driver for candidate source. Never tell readers to build a new-import driver against an old-declaration snapshot. Preserve historical facts, corpora and baseline JSON.
- Planner clarification: root MIGRATION.md is the consumer guide; record a concrete later public-port file/content allowlist and exclude private artifacts.

See reviews/post-plan-review.md and reviews/post-plan-disposition.md. Focused plain Codex read-only rereview completed: both findings resolved and no significant planning blockers. Kent subsequently approved the map before task generation; agents retain responsibility for technical rechecks.
