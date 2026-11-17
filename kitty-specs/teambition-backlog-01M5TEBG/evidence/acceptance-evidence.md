# Acceptance evidence for strict pre-consolidation gate

Finalized2026-10-25 by gpt-6.1-sol after independent gpt-6-astra approval of both WPs and root canonical transitions. WP01 reviewed codeef373a5872dba30063a6a706275a28e46508af77; integrated WP02 reviewed head0f7154e7bdb98b7e2de3e22ed2db78b8bfe0ce53. Both canonical WPs are approved, both qualified public issue rows fixed; no final mission/PR merge inferred. Independent WP02 approval is reviews/WP02-review.md, with296 prepared checks,19 supplemental checks,12 inherited WP01 assertions, fresh modern/minimumGo/focusedrace/vet/format and independently rerun golden corpora.

## Reviewed evidence map

All relative paths below are under kitty-specs/teambition-backlog-01M5TEBG/.

| Live criterion | Evidence to cite | Actual current support / remaining condition |
|---|---|---|
| FR-001 generated civil times | reviews/WP01-rereview.md; evidence/WP01-rework1.md; evidence/WP02-logs/green-focused.txt | gpt-6-astra approved WP01; affected DST/Until full integrated tests supplied on final source; WP02 independently approved |
| FR-002 explicit DTSTART | reviews/WP01-rereview.md; evidence/WP01-rereview-rank-collision.txt; evidence/WP01-rereview-probes.txt | independent12-assertion collision probe and authoritative laterfold/gap/query/minimumGo proof; no outstanding WP01 blocker |
| FR-003 type/form validation | evidence/WP02-implementation.md; evidence/WP02-logs/red-timeforms.txt; evidence/WP02-logs/red-additional.txt; evidence/WP02-logs/red-separate-inline.txt; evidence/WP02-logs/green-focused.txt | red chronology explicit, six parser adapters/order/shared every line; independent WP02 review approved |
| FR-004 floating/DATE preservation | evidence/WP02-implementation.md; evidence/WP02-logs/green-focused.txt; evidence/WP02-logs/golden-forms-final.txt | immutable descriptors/exact customLocation/value copies/Set/setters and canonical forms tested; independent review approved |
| FR-005 ordinary constructor/fragment controls | evidence/WP02-implementation.md; evidence/WP02-logs/golden-forms-final.txt; reviews/WP02-review.md | no-DTSTART legacyUTC-Z, programmaticUntilUTC-Z, per-part replacement/copy controls; independent review approved |
| FR-007 golden/unaffected compatibility | evidence/WP02-logs/golden-dst-final.txt; evidence/WP02-logs/golden-forms-final.txt; evidence/WP02-logs/golden-dst-delta-proof.json; evidence/WP02-logs/final-source-and-gate-proof.json; evidence/WP02-logs/golden-forms-baseline-cmp.txt |811DST rows/116changed/0unexplained and158forms/141changed/17controls/0unexplained, no timeouts; exact inherited76c646d baseline; negative sensitivity exit1; independently approved in reviews/WP02-review.md |
| FR-008 release/credit | docs/releases/v1.9.0-dst.md and docs/releases/v1.9.0-timeforms.md at exact lane-b head; evidence/WP02-implementation.md; reviews/WP01-rereview.md and reviews/WP02-review.md | both notes list changes/priorart, keyed migration/sourcebreak/v1.9.0; confirmed in independent WP02 review |

FR008 should use manual_qa with actual review actor/evidence; FR001–005/007 automated_test references may cite a composite durable acceptance-evidence file linking these logs. Independent approval evidence is the reviewer-authored records and canonical verdicts, distinct from implementation handoff.

## Supplied final-head local gate evidence

WP02-logs/final-source-and-gate-proof.json reports finalhead0f7154e, clean=true and unchanged rootGo hashes for all recorded gates. Supplied actual results: modern full test0.505s; minimumGo1.16.15 full test0.409s; vet0/no diagnostics; gofmt0/no paths; fullrace148.174s package/149.316runner; fuzz60.077s actual/2,980,037executions; resource benchmark1,034,487ns/op plus six-adapter dense/horizon controls. Actual companion command.json pin cwd/argv/head/sourceSHA. These logs exist durably; root will run final-target gates after consolidation.

The initial wrong-cwd focused run is marked INVALID and excluded; baseline integrity15files proven. The first one-row forms baseline determinism cmp failed, was preserved, fixture corrected without production change, then exact regeneration/cmp passed. The README states the2026-10-25 old-fallback convention. Do not omit these qualifications when explaining evidence provenance.

## Pre-consolidation status and remaining delivery

Both WPs approved by independent gpt-6-astra at exact code heads; both public issue rows fixed canonically. This evidence supports the seven live FR verdicts. Strict accept and consolidation have not yet run at this evidence-authoring commit. Root will rerun actual target gates and obtain integrated independent final review after consolidation. SC005 remains owed until the actual final privatePRURL exists; no future PR completion is represented here. Private PR base remains main and final PR must remain unmerged.

## Gate command contract

Per live FR001,002,003,004,005,007,008 with real criterion-specific reference:
SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty agent mission acceptance-verdict --mission teambition-backlog-01M5TEBG --criterion FR-001 --result pass --verification-method automated_test --actor codex --evidence ACTUAL_DURABLE_EVIDENCE_PATH --json

FR008 uses --verification-method manual_qa. Canonical seam stamps actual actor/time. Concrete descriptions/notes were authored through the protected canonical locked writer at eb9f7b9, preserving all seven pending judgments/proof fields and siblings. Exact script/output are evidence/acceptance-definitions-authoring*. The spec tables remain the specification intent snapshot; canonical WP approvals and verdict evidence record actual completion. Verdict CLI changes proof/result/evidence/identity, not authored descriptions/notes. No raw matrix/state edit is permitted.

SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty accept --mission teambition-backlog-01M5TEBG --diagnose --json
SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty accept --mission teambition-backlog-01M5TEBG --mode local --actor codex --json
SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty consolidate --mission teambition-backlog-01M5TEBG --target fix/teambition-backlog --strategy merge --keep-branch --keep-worktree --dry-run --json
SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty consolidate --mission teambition-backlog-01M5TEBG --target fix/teambition-backlog --strategy merge --keep-branch --keep-worktree

Origin enforce remains default. No --lenient/allow-fail/skip-review/attestation; --test merely records command strings and is no substitute for actual gate execution. Nonzero/exit75 requires D3 state inspection rather than blind retry.
