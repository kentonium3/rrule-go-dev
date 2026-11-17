# Mission brief: teambition backlog fixes (#5, #6)

> **Historical (completed; released in v1.9.0).** Written while this repository was private.
> Where it says "private", read it as the state at the time. Current rules: [AGENTS.md](../../AGENTS.md).

Set up 2026-10-25 at Kent's request. One spec-kitty mission covering the two public issues on
kentonium3/rrule-go that came from teambition's open backlog:

| Public issue | Brief | Seam | Priority |
|---|---|---|---|
| #5 sub-daily rules across DST | `docs/issues/05-dst-sub-daily.md` | `rrule.go` `generate()` | P1 |
| #6 floating UNTIL read as UTC | `docs/issues/06-floating-until.md` | `str.go` parsing and `String()` | P2 |

**Scope update (Kent,2026-10-25):** Public issue7 is removed after contributor publicPR9 landed on private main at76c646dde8d33ff7c3325df0613b31b058243bea. It is inherited baseline, not assessed or extended here. Golden baseline is v1.8.3 plus76c646d. Kent also explicitly approved private ROption metadata, the external positional-literal source break, and targetv1.9.0; migrate positional callers to keyed literals.

**Normative rule (Kent, 2026-10-25): adhere to RFC 5545 in all cases.** Where python-dateutil, a
teambition PR, today's behaviour or one of these briefs disagrees with the RFC, the RFC wins. If you
find a brief that contradicts the RFC, follow the RFC and record the correction in the spec.

The remaining seams have shared parser/civil-anchor helpers; use two serial WPs to avoid competing designs. One work package per issue is the expected shape, and they can run in
parallel lanes. Reproductions: `docs/repro/05-07-backlog/main.go`, with v1.8.3 output beside it.

## Authorization

Kent sets the autonomy level when he starts the session. If he says to run it autonomously, use
the posture of `docs/runbooks/autonomous-iterator-panics.md`: the full arc through a **private** PR
to `main` on `kentonium3/rrule-go-dev`, the D-3 fault protocol, and a stop before merging to
`main`. No public action of any kind is authorized by this brief. That includes comments on the
public issues, pushes to the public repo, tags and releases.

## Step 0: bring the charter up to date (before `$spec-kitty-specify`)

The charter in `.kittify/charter/` was written for missions 1 to 3 and now contradicts this mission.
Its directives have `severity: error`, so they would block or misdirect the work:

- **RRULE-001** allows behaviour changes only for panics, hangs and impossible results. The new rule
  is RFC 5545 in all cases, as AGENTS.md rule 3 now says.
- **RRULE-003** says the repo stays private until 2026-11-13. That embargo is over and the fixes are
  public. Still valid: this dev repo stays private, nothing is pushed upstream, and public copy
  needs Kent's sign-off.
- **RRULE-004** sets the scope to panics, then hangs. The scope is now the public #5 and #6; contributor#7 is inherited baseline only.
- **`branch_strategy.main_branch: master`.** The branch was renamed to `main` on 2026-10-25.

**Corrected 2026-10-25, after the first attempt stopped (see Progress).** An earlier version of this
step said to make the change only through `charter interview` and `generate`. That was wrong, and
Codex correctly stopped on it. On build `1a0c8b467`, `governance` and `directives` in `charter.yaml`
are **hand-authored** sections. The state contract says so (`specify_cli/state/contract.py`:
"governance/directives are hand-authored sections"), and `charter generate` refreshes only the
derived `catalog` and `metadata`, by design. `charter.md` is likewise a curated companion that
generate never writes. Editing those sections is the intended authoring path, not a workaround.

So:

1. On a branch (`chore/charter-backlog`), hand-edit `.kittify/charter/charter.yaml`. Change only the
   `governance` and `directives` sections:
   - RRULE-001, RRULE-003 and RRULE-004 as listed above;
   - `branch_strategy.main_branch: main`, and the branch rules that name `master`.

   Keep the nested `directives: {directives: [...]}` shape. Mirror the same facts in `charter.md`,
   which says "merge to master" at line 26.
2. Leave `interview/answers.yaml` alone. It records the original interview. If its stale "master"
   answer causes trouble later, report that rather than rewriting the record.
3. Run `spec-kitty charter generate`, which refreshes the derived sections and must leave the
   authored ones byte-for-byte, then `spec-kitty charter validate`, `spec-kitty charter bundle
   validate` and `spec-kitty charter status`. Also confirm that runtime rule loading picks up the
   new directive text (AGENTS.md: generation alone does not prove valid policy).
4. PR to `main`, merge it, then specify the mission.

Stop and report only if generate alters the authored sections, or validation rejects a well-formed
edit. Either of those would be a spec-kitty fault.

## Teambition PRs: use, don't trust

Kent's direction (2026-10-25): reuse work from teambition PRs where it helps, but verify it
independently. Both relevant PRs were probed, and both have defects that rule out adopting them as
is (details in the issue briefs):

- **PR #66 (DST)** steps sub-daily rules in elapsed time. That isn't RFC 5545, and it shifts BYHOUR
  matches by an hour on the day clocks go back.
- **PR #68 (UNTIL)** rejects the RFC-required `TZID` + `UNTIL=…Z` form. Its `TZID=…:…Z` check is
  RFC-correct and reusable.

Treat their tests as candidate cases, re-derive each expected value from RFC 5545 rather than
copying it, and credit any idea or test you keep in the commit message ("Based on an idea from
teambition/rrule-go#66").

## Compatibility gate (both WPs)

AGENTS.md rule 3 is refined for this mission. Public signatures/keyed callers stay compatible. Kent explicitly approved private ROption metadata and the resulting external positional-literal source break; targetv1.9.0 and document keyed-literal migration. Adding an
exported identifier is allowed only if a fix needs it, with a justification in the plan, and it
makes the release v1.9.0 instead of v1.8.4. Behaviour may change **only** where RFC 5545 requires
it, for the cases each issue brief's acceptance names. That includes rejecting input the RFC forbids
(#6) and dropping instances at nonexistent local times for every frequency (#5). Every such change,
and every newly rejected input form, goes in the release notes.

Prove that with a **golden comparison**: build a corpus of rules and record each one's first N
occurrences and `String()` output on v1.8.3 plus contributor fix76c646d (private baseline76c646dde8d33ff7c3325df0613b31b058243bea; pristine local snapshot `/tmp/rrule-backlog-v183-plus-duplicates`). Then diff against the fixed code. Every
difference must be one an issue's acceptance predicts. For the corpus, start from:

- the existing test suite's rules;
- `docs/repro/matrix/` and `docs/repro/edge-cases/`;
- every frequency × {UTC, America/New_York, Australia/Sydney, Australia/Lord_Howe}, with DTSTARTs in
  the week of each DST change.

Commit the harness as a test or a `//go:build ignore` program, so the release port can re-run it.

## Local gate

- `go vet ./...`, `gofmt -l .` (no output), `go test ./...` and `go test -race ./...`.
- At least 60 s of `FuzzIterate`.
- The golden comparison.
- Each defect test is shown red on unfixed code first.
- Go 1.16 stays the minimum. The public repo's CI tests 1.16 and the latest Go.

## Models

The implementer and reviewer are different models, ideally the most capable available from each
vendor. Record which model did what, with the date, per WP. The last run used gpt-6.1-sol to
implement and gpt-6-astra to review. A same-model review is a degraded check and must be recorded
as one.

## After the mission: reaching the public repo (Kent-gated, not part of the mission)

This repo carries `kitty-specs/`, `.kittify/`, `docs/` and `AGENTS.md`, which the library repo does not.
The public repo (`kentonium3/rrule-go`, branch `main`) gets only the library changes:

1. After the private PR merges to `main`, export the code-only commits (`*.go`, `go.mod`,
   `go.sum`), one or two per issue, as a patch series.
2. Apply them on a branch of the public repo with author and committer **Kent Gale <kentgale@gmail.com>**. Use
   Conventional Commits with `Fixes #5` / `Fixes #6` / `Fixes #7`. Those numbers refer to the
   public repo, so **don't put closing keywords in commits here**: in this repo #5–#7 are different,
   older issues.
3. Update the "What this fork fixes" section of the public README.
4. Kent signs off on the exact copy. The branch goes up as a PR to the public `main`, which only
   accepts PRs whose four CI jobs pass; there is no bypass. After the merge, tag `v1.8.4` (or
   `v1.9.0`; see above) on the merge commit.
5. Vikunja adoption is a separate cross-repo step (`docs/issues/04-vikunja-adoption.md`). Vikunja
   evaluates rules with a floating DTSTART in `time.UTC`, so #5 and #6 shouldn't change its
   results, but its recurrence suites are the check.

## Progress

- 2026-10-25: Step 0 blocked by the explicitly requested charter-command capability
  hard stop. Read AGENTS.md, CONTEXT, this runbook, and briefs 05–07 in order;
  then loaded the autonomous D-3 posture and charter-governance skill. Initial
  checkout was clean on `main`; origin is the private `kentonium3/rrule-go-dev`
  URL, upstream push URL is `no_push`.
- Pinned build verified from installed `direct_url.json`:
  `1a0c8b467af253eb0185fa19d0d50880b2b6475e`. Prefixed interview/generate help
  and status commands succeeded. However, interview offers prose answers and
  doctrine ID selections, no authored project-directive or branch-strategy
  amendment. Its writer preserves existing unmodeled YAML keys. The generator's
  `write_compiled_charter` explicitly updates only derived `catalog` and
  `metadata`, preserving authored `governance` and `directives`, even with
  `--force`. Thus this command sequence cannot express the RRULE-001/-003/-004
  and `main_branch` changes. Source inspection established the limitation before
  mutation; no interview write or generate mutation was attempted.
- Exact capability commands, outputs, initial HEAD, build identity, and installed
  source excerpts: `evidence/teambition-backlog-step0-charter.txt`. Existing status
  reports fresh/synced, which does not establish that policy matches this brief.
  No direct charter YAML edits, installation changes, mission creation, pushes,
  PRs, merges, or public actions. Stopped as requested; no workaround applied.
  Resumption needs a command surface capable of these authored amendments, or
  Kent's explicit revision of the command-only update constraint. No mission
  record exists because specification is gated on the charter PR being merged.

- 2026-10-25 resumed at corrected step 0 after pulling main (533ab66).
  Authored only the requested governance/directive amendments and mirrored the
  companion; original interview unchanged. Generate passed and authored sections
  remained byte-identical. All eight doctrine artifacts validate, CharterYaml
  model validation passed, bundle validation passed, status fresh/synced, runtime
  context loads the revised RRULE-001/-003/-004 summaries and main_branch=main.
  Evidence: evidence/teambition-backlog-step0-validation.txt.
  Initial pull/generate hit the known read-only Git metadata sandbox boundary;
  retries with authorized metadata access succeeded. CLI invocation corrections:
  validate has no --json and requires a path; charter.yaml is not a doctrine
  artifact filename, so validate the directory and separately validate the
  CharterYaml model. These were invocation errors, not rejected policy edits.

- 2026-10-25: Charter PR12 merged privately at 6ac755d. Created mission
  `teambition-backlog-01M5TEBG` on `fix/teambition-backlog`; canonical planning
  and consolidation target is that private fix branch, final delivery remains
  an unmerged PR to main. Discovery/specification completed (250222c), with
  official RFC-derived corrections for explicit gap DTSTART, DATE serialization,
  legacy EXRULE/multiple-RRULE behavior and fragment semantics. Plan authored;
  tasks intentionally await independent gpt-6-astra post-plan review. Baseline
  `go test ./...` passed (Go1.25.14, root report). Planned implementation model
  gpt-6.1-sol; review model gpt-6-astra. No public action.
- Planning operational record: `kitty-specs/teambition-backlog-01M5TEBG/evidence/planning-faults.md`.
  An empty CLI decision-index lock was accidentally staged with the decisions
  directory, then preserved outside the checkout and removed in corrective
  commit62face6 after zero-byte/no-holder inspection. No decision data lost;
  future staging enumerates individual decision files. No substantive gate bypass.
- Independent post-plan review (gpt-6-astra, 2026-10-25) requires changes before
  tasks. Full record: mission `reviews/post-plan-review.md`. BYSETPOS design was
  corrected to preserve arithmetic selection on gap-free periods and rank around
  transition holes, with effective-UTC UNTIL bounds retained; re-review pending.
  The proposed private timezone-marker design cannot preserve an arbitrary caller
  *time.Location by reloading its name. A fixed-zone/IANA-name collision is
  reproduced in mission `evidence/location-name-counterexample*`. A clean fix
  uses private ROption metadata but changes its field layout, breaking external
  positional literals despite being permitted by Go's general compatibility
  policy. Kent's explicit compatibility decision is pending; no tasks,
  implementation or gate waiver. Private v1.9.0 notes are proposed if authorized.

- 2026-10-25: Kent explicitly approved private ROption parse metadata and
  targetv1.9.0 (decision01M5TFFVFY5YXYPVD4JN9SS90M). Existing signatures
  and keyed callers remain; external positional literals break and need keyed
  migration. Kent then removed issue7 from scope after contributor fix76c646d
  landed; main was merged into fix/teambition-backlog. Scope is now only issues5/6,
  two WPs. Golden baseline=v1.8.3 plus76c646d, pristine snapshot
  `/tmp/rrule-backlog-v183-plus-duplicates`. Root baseline tests passed modern
  Go (.383s) and Go1.16.15 (.345s). Charter scope/compatibility text and all active
  planning artifacts revised; independent gpt-6-astra re-review precedes tasks.
- Canonical metadata decision wrote an empty decision-index lock. No holder;
  preserving by Path.rename to `/tmp` failed with Errno18 cross-device link and
  made no mutation. Root then copied with shutil.copy2 and unlinked the empty
  source, preserving `/tmp/rrule-backlog-private-metadata-decision-index.lock`.
  Original decision/index/events retained; no gate bypass.
- Independent post-plan re-review APPROVED by gpt-6-astra on2026-10-25,
  revision3b776c5, durable `reviews/post-plan-rereview.md`. Two WPs authored,
  seven subtasks each, all seven active FRs and all NFR/C/SC mapped; retired
  FR006/SC003 not reused. Finalize validate-only passed with no ownership or
  requirement warnings; canonical finalize committed7c5973e. WP02 depends on
  WP01 despite two computed lanes; narrow sequential parser/foundation adapters
  are explicitly recorded. Runtime now at analyze; no code implementation yet.

- 2026-10-25: Independent post-plan re-review approved the revised two-issue,
  private-metadata design (gpt-6-astra); full review in mission reviews/.
  Tasks finalized at7c5973e with exactlyWP01(DST) andWP02(time forms), serial
  dependency; requirements mapped. Analysis atc374554 passed ready with zero
  findings. WP01 claimed through canonical action at56c4b63 in lane-a and
  dispatched to gpt-6.1-sol, with gpt-6-astra review planned. Runtime's empty
  resolved_model projection is not model evidence; actual native dispatch and
  per-WP handoff/review records identify models/date. WP02 remains planned.

- 2026-10-25: Read-only accept preflight completed (evidence/
  accept-diagnose-before-implementation.json). No missing artifacts or path
  violations. Expected unfinished WP/matrix blockers remain. Claim records agent
  but no assignee; supply canonical assignee on the eventual for_review transition
  and verify metadata then. Planned WP02 is intentionally unclaimed. Restricted
  sandbox DNS caused ORIGIN_UNREACHABLE warning; rerun actual acceptance with
  authorized network/Git access. Diagnose did not mark acceptance or bypass gates.

- 2026-10-25 WP01 implementation checkpoint (gpt-6.1-sol): red-first outputs
  saved under /tmp/rrule-WP01-logs before source edits, including all-frequency
  gaps/folds, parsed anchors, public query paths and signed BYSETPOS holes.
  Expanded seven-frequency × three-zone × gap/fold regressions now pass.
  Three inherited TestUntilGolden gap-normalization expectations require RFC
  corrections (NY, SaoPaulo, Apia); original full-suite failures preserved.
  One new fixture's expected Sydney date was corrected because the originally
  expected following-year date was also a gap; this is a test-authoring correction,
  not independent defect evidence. Reusable golden corpus contains796rows:
  304static option fixtures,92FREQ literals,392transition matrix and8targeted;
  39dynamic options are individually listed as extraction skips and remain
  full-suite tested. Golden classification/resource/local gates and review pending.

- 2026-10-25 WP01 implemented by gpt-6.1-sol at
  af862491333c8a06fdb30bb60a58619e65e9baa6; lane clean. Full handoff and48logs
  persisted in mission evidence/WP01-implementation.md andWP01-logs/.
  Final799-row golden:102classified changes,0unexplained,0timeouts; unchanged
  baseline comparison and deliberately corrupted UTC row both fail as intended.
  Modern/Go1.16.15 full tests, vet, zero-path gofmt and focused race passed.
  Dense365×86400 BYSETPOS benchmark stayed near baseline1.15ms without full
  product scanning. Six old repros gained equivalent legacy ignore tags after
  whole-checkout Go1.16 exposed missing tags (the earlier baseline snapshot was
  code-only). Shared parsed Set gap red evidence and narrow descriptor transfer
  fix included. Out-of-map edits and initial fixture/benchmark corrections are
  disclosed in handoff. Independent gpt-6-astra review pending.
- Qualified public issue matrix explicitly authored through canonical
  `write_issue_matrix`/ProtectionPolicy with fail-closed write-surface resolution,
  mission lock and reread/splice. Private commit2494441 contains only matrix;
  readback verifies repo=`kentonium3/rrule-go`, #5→WP01 and #6→WP02, both
  `in-mission`, siblings preserved. Neither is claimed fixed; issue7 not added.
  Exact script/output: mission `evidence/qualified-issue-matrix-authoring*`.
  Command: `SPEC_KITTY_NO_UPGRADE_CHECK=1 ~/.local/share/uv/tools/spec-kitty-cli/bin/python kitty-specs/teambition-backlog-01M5TEBG/evidence/qualified-issue-matrix-authoring.py.txt`.
  No production edit, WP transition or public action.

- 2026-10-25 WP01 review cycle1: gpt-6-astra independently reran probes,
  modern/Go1.16 tests, relevant race, vet/format and799-row golden successfully,
  then rejectedaf86249 for a compound explicit-gap-anchor/BYSETPOS collision.
  Two nominal slots resolve to one instant before rank selection, missing a
  valid position2 or inventing position3. Only proven1,2 and1,3 cases are
  findings; signed variants that pass are controls. Durable review/probe evidence
  is in reviews/WP01-review.md and evidence/WP01-review-rank-collision*.
  Canonical rejection durably recorded; all7WP01subtasks reset by workflow,
  implementation re-claimed with focused cycle1prompt and dispatched to the
  original gpt-6.1-sol implementer. WP02 remains gated. No gate bypass.

- 2026-10-25 WP01 correction committedef373a5872dba30063a6a706275a28e46508af77
  by gpt-6.1-sol; lane clean. New same-day/coarse/half-hour and same-period
  cross-day Apia collision regressions were shown red onaf86249 before correction.
  Filtered-period membership plus a sorted-clock lookup keeps one instant per
  positional rank; daily/filtered-boundary controls protect ordinary period
  semantics. Independent rejection probe now passes. Modern/Go1.16 tests,
  relevant race, vet/format and dense selection checks passed. Golden811rows,
  114classified,0unexplained; original799rows unchanged. Rework evidence is
  persisted under evidence/WP01-rework1*. Independent re-review pending.

- 2026-10-25 WP01 independently approved by gpt-6-astra at exact
  ef373a5872dba30063a6a706275a28e46508af77 (reviews/WP01-rereview.md),
  original12collision assertions now pass; refreshed gates and811-row golden
  verified. Canonical review verdict durably approved and private issue-matrix
  row for public kentonium3/rrule-go#5 marked fixed with branch/review evidence
  (no public tracker action). WP02 claimed in lane-b and dispatched to
  gpt-6.1-sol using the runtime prompt; integrated gates/floating forms next.

- 2026-10-25 WP02 completed by gpt-6.1-sol: implementation62b2527,
  final clean lane head0f7154e7bdb98b7e2de3e22ed2db78b8bfe0ce53. Red matrix
  precedes parser/set changes; valid UTC/TZID, fragment and constructor controls
  remain. Modern/Go1.16 tests, full race(148.174s), vet/zero-path format,
  resource controls and FuzzIterate60.077s/~2,980,037executions passed.
  Golden DST811/116changed/0unexplained plus forms158/141changed/0unexplained,
  no timeouts; exact per-row deltas and repeated baseline regeneration proof
  persisted in evidence/WP02-implementation.md andWP02-logs/.
  Two authoring faults preserved and corrected: a combined command used the
  pristine baseline cwd, whose replacement wrote identical bytes and whose
  no-test green output is explicitlyINVALID; all15baseline source/module/license
  files verified byte-identical to76c646d, and actual lane tests rerun. One late
  shared-start baseline control accidentally retained COUNT3 and used old-parser
  implicit now; its nondeterminism was caught by repeat cmp, preserved, replaced
  with explicit2020UNTIL and documented2026-10-25 context, then repeated outputs
  matched. Neither invalid result counted as a gate. Independent review pending.
- Acceptance definitions authored canonically at eb9f7b9: seven concrete FR
  descriptions/notes replace scaffold TODO text; all judgments remain pending,
  proof fields and siblings preserved. Script/output in mission
  `evidence/acceptance-definitions-authoring*`; durable evidence proposal in
  `evidence/acceptance-evidence-proposal.md` explicitly awaits WP02 independent
  approval. Spec Open tables are intent snapshots, canonical reviews/status
  record actual completion. No accept, consolidate or WP transition performed.
- Strict closure2026-10-25: both independent WP reviews/canonical approvals
  complete; all seven live FR verdicts pass with actual reviewed evidence.
  Diagnose ok=true, strict local accept97b3bc6/residual27032ae. Preview inspected
  then merge consolidation succeeded124c372 onto `fix/teambition-backlog`,
  retained branches/worktrees, no push. Both WPs done, mission_number4;
  auto-retrospective HEAD9c517fd. Go/module/LICENSE parity with reviewed0f7154e
  verified; main remains76c646d. Exact commands/results and remaining work:
  mission `evidence/accept-consolidate-closeout.md` and linked outputs. Root
  still owes actual target gates, integrated independent final review and the
  private PR URL for SC005; final PR must remain unmerged. No public action.

- 2026-10-25 consolidated-target gates passed: modern and directGo1.16 full tests,
  fullrace148.336s, vet/no diagnostics, format/no paths, diffcheck, and final
  FuzzIterate60.078s/3,009,261executions. Fresh golden811DST+158forms rows
  have0unexplained differences and no timeouts. Sourcehashes for all12actual
  target command records matchd482828; evidence/final-target-verification.md
  and final-target-gates/ retain actual commands/logs/outputs. Integrated
  independent final review closing verdict and privatePRdelivery next.

- 2026-10-25 integrated final review PASS by gpt-6-astra atd482828,
  durable reviews/mission-review.md. All7liveFRs traced;34approved source,
  module/license/README/golden/releasefiles byte-identical afterconsolidation;
  all12actualtarget command records sourcehashverified. No unresolved blocker.
  Privateorigin visibilityverifiedprivate; remote main still76c646d. Final
  privatePRmain delivery pending actualURL; no merge/publicaction authorized.

- 2026-10-25 delivery complete: privatePR
  https://github.com/kentonium3/rrule-go-dev-archive/pull/13, base main,
  headfix/teambition-backlog, open/unmerged/mergeable verified. SC005actual
  delivery is recorded in evidence/delivery.md; bothWPsdone, allrequiredlocal
  gates and independent integrated reviewPASS. InitialremoteCIqueued,
  CodeQLskipped byprivatepolicy; no falsegreenclaim. Main remains76c646d.
  No publicaction/tag/release/upstreampush orfinalPRmerge. Stopping atPR
  asauthorized. Read-only connector argumentvalidation errors corrected
  without mutation and disclosed in delivery evidence.
