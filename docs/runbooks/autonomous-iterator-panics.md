# Autonomous run: iterator panics

Authorized by Kent on 2026-10-25, after scope confirmation. Queue: private issue
kentonium3/rrule-go#4 only, mission iterator-panics-01M5RT9D. Finish specification,
plan, tasks, analysis, implementation, independent review, verification, and a
private PR to master. Stop before merging to master, public disclosure, upstream
posting, release publication, or Vikunja changes. No additional queue item is
implicitly authorized. Preserve work in scoped, resumable commits.

## Repo adapter

- Full Spec Kitty mission; pinned build 1a0c8b467af253eb0185fa19d0d50880b2b6475e.
  Prefix all commands with SPEC_KITTY_NO_UPGRADE_CHECK=1.
- D-3 supersedes AGENTS.md's blanket stop-on-fault for this run. Capture exact
  commands, output, build and perishable state before modifications; diagnose,
  prove a safe workaround, record it in the private fault log, then continue.
  Stop for unresolved unsafe failures, genuine scope decisions, usage limits,
  or host/security changes. Never bypass a substantive acceptance/review gate.
- Only private origin kentonium3/rrule-go is a publication destination. Verify
  visibility before push/PR. Upstream is fetch-only. Public copy waits for Kent.
- Local gate: go test ./..., go test -race ./..., go vet ./..., gofmt -l .
  (no output), targeted regression and at least 60s bounded fuzzing. Establish
  the unfixed test baseline and prove each defect red before its fix.
- The Go test CI triggers master pushes; the separate CodeQL workflow triggers
  PRs to master. Report their actual status separately from local validation. No merge
  is authorized even with green checks. Do not expand into unrelated CI upgrades.
- Implementation model gpt-6.1-sol; review model gpt-6-astra, with explicit model
  records/date per WP. Native subagents; at most two children active at once.
  Post-plan and integrated-final review checkpoints use independent review.
- No deployment target, release publishing, live service verification, host
  changes, audited-surface rebaseline, or service architecture modifications.
- Ordinary implementation choices are autonomous, with rationale in plan/log.
  Keep exported API and correct occurrences, no new dependencies. A minimum Go
  version increase is authorized if needed for fuzzing; justify it in planning.
- Maintain progress and fault evidence with the mission. Own the rrule-go changes;
  Spec Kitty code faults are findings, not permission to edit its installation.

## Progress

- Setup merged to master at 3fc4eb5; all 14 command skills healthy.
- Mission scaffold on fix/iterator-panics; confirmed intent decision
  01M5RTA7A5FAHWV04DNQBJ9V6B resolved.
- Origin visibility verified private through GitHub connector.
- Specification, plan, tasks, analysis, both WPs and independent reviews completed.
- Strict acceptance passed; consolidation ac594a5 landed on fix/iterator-panics.
  Both WPs are done. Integrated gates passed; final review PASS WITH NOTES.
  Private PR https://github.com/kentonium3/rrule-go-dev-archive/pull/9 delivered; master is unchanged.

## Fault log

No new mission workflow faults yet. Bootstrap history: Git staging needed sandbox
escalation; stale skill instructions put activation keys in charter.yaml rather
than activation.yaml; Go required explicit config languages; authored directives
needed the nested directives mapping. All were diagnosed and corrected before
this run, with source/model validation and healthy charter status recorded in chat.

### Preflight evidence

- Baseline: `GOCACHE=/tmp/rrule-go-cache timeout 180 go test ./...` passed on
  Go 1.25.14 linux/amd64 (2026-10-25).
- Bounded read-only sweep of known Spec Kitty tooling issues before the run.
- Relevant known tooling issues (context-only; no work owed by this mission):
  see spec-kitty/spec-kitty#6652 (approval evidence lost), #6630 (model projection),
  #6653 (single-digit issue matrix terminality), #6651 (mixed prose/JSON),
  #6650 (shared-lane context warning), #6628 (analysis lock residue), #3132
  (no PR landing guard), #6429 (merge may mutate before post-validation failure).
  Keep full reviews as separate tracked artifacts, record identities in notes,
  inspect every matrix row and actual commits; never blindly retry a mutation.
- Sleep inhibitor attempt `systemd-inhibit --what=idle:sleep --mode=block
  --why='Private rrule-go autonomous mission' tail -f /dev/null` was rejected by
  automatic approval review as a host power-management action. No process started;
  no workaround attempted. Run has no active host-sleep protection.
- Task requirement mapping first attempt failed before mutation: `map-requirements
  --batch ... --mission iterator-panics-01M5RT9D --json` returned Errno 30 for
  `.git/spec-kitty-locks/iterator-panics-01M5RT9D.status.lock`. This is the known
  sandbox Git metadata boundary; validate-only passed all ownership and requirement
  coverage checks. Retry the same mapping with approved Git metadata access.
- Corrected an orchestration command typo (`agent mission map-requirements` has
  no such command) using CLI help: `agent tasks map-requirements` succeeded with
  all six FRs mapped. finalize-tasks committed a7629bc; two WPs, no warnings.
- Independent plan review approved (gpt-6-astra, 2026-10-25); durable review saved.
- Analysis preflight found the known #6628 empty decisions/index.json.lock
  residue (0 bytes; lsof found no holder). Preserve it outside the checkout while
  recording analysis; do not commit lock files. No decision data is removed.

### Implementation claim preflight (2026-10-25)

Pinned build 1a0c8b467: `spec-kitty agent action implement WP01 --mission
iterator-panics-01M5RT9D --agent codex --profile implementer-ivan` exited 1:
`Error: synthesized_drg missing; run spec-kitty charter synthesize`.
Git status was clean, both WPs remained planned, and no code workspace was claimed.
CLI help documents a fresh-project seed path for absent generated artifacts;
inspect its dry-run and freshness before promoting. Do not bypass the guard.

Recovery: dry-run first reported GeneratedArtifactMissingError for PROJECT_001,
proving this configured project needs agent-authored synthesis rather than the
help text's fresh-project fallback. Authored eight targets from existing answers
and charter: two directives, two tactics, four styleguides. Generated-adapter dry
run passed without warnings/conflicts, then canonical promotion succeeded; source,
bundle, and synthesized DRG all report fresh. CLI appended the two synthesized
tactics to activation.yaml. Existing policy and charter.yaml were unchanged.

Canonical runbook reference: the operator's autonomous-run protocol (kept outside this repository).
Analysis passed ready, all requirements mapped, private issue #4 explicitly
registered in-mission in the issue matrix (manual single-digit coverage).

### Root-layout acceptance preflight

`accept --mission iterator-panics-01M5RT9D --diagnose --json` reported missing
src/ and tests/ on pinned build, matching upstream #6633. Both production .go
and *_test.go files actually live at repository root. Installed supported
project.path_conventions overrides preserve strict enforcement. Set workspace
and tests to explicit `.`; do not create empty directories or use --lenient.
Original diagnostic saved as evidence/accept-preflight.json. Other pending
checks are expected before implementation/review; sandbox DNS prevents the
read-only origin check, which will be rerun with network permission at acceptance.
Validation: repeated the same read-only diagnose after the override; path_violations
is empty, while unfinished work and acceptance evidence still correctly block.

### Verification progress

WP01: committed 6177350; gpt-6.1-sol implementation, gpt-6-astra approval on
2026-10-25. Red evidence for all three families and wrong interval output;
full tests/race/vet/format/build/386 pass. Independent 1,701-case original/fixed
parity matched. Runtime pre-review hook is undeclared and reports skipped;
explicit actual validation is preserved, never represented as hook execution.

WP02: gpt-6.1-sol authored the fuzz file and deterministic seeds, then its agent
turn was interrupted twice by an automated content-risk flag. Root preserved
work, ran direct local validation without source edits, and committed 3dfc313.
Six original-version seed runs fail as expected; 60.084s fuzz campaign passed
3,437,785 executions. Full gates and386 pass. Independent review pending.

The full tooling-issue sweep later completed; its record is not part of the public
snapshot.

## Closeout evidence

Final independent review: reviews/mission-review.md (gpt-6-astra, 2026-10-25).
All implementation files are byte-identical to independently approved commits.
Retrospective captured in kitty-specs/iterator-panics-01M5RT9D/retrospective.yaml;
summary routing discrepancy diagnosed separately, with no installation changes.
Follow-up: https://github.com/kentonium3/rrule-go-dev-archive/issues/8 tracks tooling findings
privately. It owes no recurrence implementation work and needs approval before
any upstream copy is posted. Fault ledger reconciled; no substantive gate bypass.

Delivery completed: private PR https://github.com/kentonium3/rrule-go-dev-archive/pull/9,
head fix/iterator-panics, base master. No merge performed. All required source
validation and independent reviews passed. SC-005 delivery evidence is now the
PR and its linked review/test records. Tooling follow-up remains private #8.

CI correction at delivery: codeql.yml runs Analyze (go) on PRs to master;
its check is pending. The earlier blanket CI note was incomplete; the PR body
and this adapter now distinguish the Go test workflow from CodeQL.
