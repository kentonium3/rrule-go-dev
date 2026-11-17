# WP03 advisory review — APPROVE

Date: 2026-10-29, America/New_York.
Actual reviewer: gpt-6-astra, Reviewer Renata (profile resolved earlier in this same review session).
Actual implementer: gpt-6.1-sol, independently assigned.

## Authorized scope exception

Reviewed exactly commit 7dd0546b17b058bbd77742283525489829319278 against its parent, restricted to the four owned files: drafts/release-communication.md, drafts/teambition-transition-pr.md, evidence/migration-verification.md, and evidence/vikunja-handoff.md under kitty-specs/module-path-v1-9-0-01M62K6W/. Confirmed these are the entire commit's file set and the inspected artifact bytes match that commit.

This explicitly uses Kent's authorized scope workaround after the preserved planning-artifact review-anchor fault. The generated prompt has unknown base/branch and no deterministic claim anchor. That inferred runtime diff was NOT used or repaired. The original fault is preserved in /tmp/rrule-WP03-review-anchor-fault.md; its historical statement that no workaround had yet been approved describes capture time, superseded for this review by the user authorization delivered by the orchestrator. No claim, verdict, transition, acceptance, merge or workflow metadata command was executed by this reviewer.

Verified source remains approved candidate 9bf894f4b899d0c1cb814e2f30753eeec6e01dab in lane-b. Execution HEAD 9a4f64dc2169d96969f03c12ee5a9fa24bfe1154 differs only in status.events.jsonl/status.json, independently confirmed. Root's old module declaration is intentionally not the candidate. Baseline is 737ae01be40048d1d2932dbafb623230efa8d361.

## Findings

No significant substantive findings. The report contains actual commands, outputs, working directories, fixture source, source/binary/corpus/output hashes, and the disposable runner. It separates local verification from public availability, acknowledges cache/VCS limitations, and records the broad historical whitespace failure separately from the scoped candidate-source success. No evidence of source substitution, synthetic shape-only tests, changed golden expectations, or claimed publication/integration was found.

Nonblocking evidence-format observation: exact WP03 git diff --check exits 2 because preserved command output has trailing spaces after empty import lists and the embedded historical whitespace diagnostics preserve the offending spaces. These are inside verbatim evidence fences, not runtime code or proposed public copy. Likelihood: certain when running a whitespace check over this artifact commit. Effect: cosmetic diagnostics only. Retaining exact output is defensible; do not represent this exact artifact check as passing. The separately recorded candidate-source check remains valid.

## Independent verification

Go checks use process-scoped GOWORK=off GOPROXY=off GOCACHE=/tmp/rrule-module-path-wp03-review-cache.

- All report candidate source/document hashes match lane-b bytes. Existing pristine baseline export's root Go runtime/test files, go.mod and golden driver match the pinned baseline Git blobs.
- Explicit declaration assertion independently reproduced: baseline prints github.com/teambition/rrule-go and exits 1; candidate prints github.com/kentonium3/rrule-go and exits 0.
- Built fresh baseline and candidate golden binaries from the correctly imported driver/source cwd. Re-executed both unchanged corpora. DST: 811 rows, SHA256 1fca954470ccc972e6eadd40faaf2c9c4ebe99b9f779897c92d6e6f0a5c01a4b. Forms: 158 rows, SHA256 c632fe0db2b700ff2acdca385d2e74f9a3392a432982020898e85c0b5503a236. Each before/after pair is byte-identical and exactly matches the report hash. IDs/order also match the actual corpus, not only one another.
- Rejection inventory exactly reproduces all 39 DST and 60 forms expected parser/validation rejections. No panic or timeout. Both temporary corrupted-copy cmp controls independently exit 1. Reviewer outputs: /tmp/wp03-review-s7d4gblz/.
- Reran actual preserved direct/all/version migrated consumer tests with -count=1: all PASS. Inspected source assertions for the fixed UTC parsing, all three timestamps, exact serialization and keyed constructor. Earlier report transcript records working old states and exact left-side removal syntax before new-path replacement/tidy.
- Retained-consumer package query independently reveals example.invalid/dependency importing the old path while its public Count function returns int. Protective replacement remains present. Query used explicitly authorized fixture-only GOFLAGS=-buildvcs=false; this is not a claim that the original unmodified WP02 query succeeded.
- Independently reran go test -race ./... against lane-b with the review cache: PASS, library 148.174s and tools/changelog 1.014s (uncached output).
- Inspected all required candidate unit/race/vet/format/changelog results and seven standalone driver compilation results in the report. Formatting output is empty. Unit/race transcript honestly shows Go's cached successes; it does not claim uncached execution. Prior default-cache failures are explicitly preserved in WP01-implementation.md, and subsequent commands use temporary process-scoped caches. Broad inherited Markdown diff --check exit 2 is preserved and not converted to a passing claim.
- Public excerpts scanned: no private mission IDs, local paths or workflow records. Root workspace remains clean.

## Draft and handoff assessment

Both complete packets retain after-release gates, exact-copy approval, neutral separately-maintained-fork language, README-only alternative, unchanged upstream module identity/Go floor, and maintainer control of any deprecation/final release. The migration guide URL replaces the uncertain README anchor. The dated audit is used with explicit limits on present permissions/private exchanges; no current public-state verification is claimed. Opt-in Release notification and latest-version/unconditional-replacement warning limits are explicit. Version reuse/public retrieval/indexing and actual Vikunja execution are correctly deferred.

The unsent Vikunja record names exact candidate source, desired bounded application checks and existing service guards. Its public-port allowlist excludes private mission/authority/reproduction artifacts and requires selected README hunks/public badge reconciliation. It selects no recipient and claims no sent message or completed integration.

## Required anti-pattern checklist

1. Dead code: N/A — private evidence and draft authoring only; no production module added.
2. Synthetic fixture test: PASS — consumers and golden runners call actual library paths; declaration check is independent of local replacement masking.
3. Silent empty return: N/A — no new production paths.
4. FR coverage: PASS — FR-003/008 have actual migrations, safety-stop and declaration evidence; FR-004/005/007 have reviewed concrete conditional copy and notification limits; FR-006 references the dated primary-source maintenance audit and preserves authority uncertainty. SC criteria tracked; no actual public distribution/integration required here.
5. Frozen surface: PASS — exact commit only four owned artifacts; source/historical corpus/release preservation supported by reviewed fingerprints, existing WP reviews and evidence.
6. Locked decision: PASS — no public operation, compatibility shim, synthetic proxy, source/API/version/dependency change or misleading availability assertion.
7. Shared-file ownership: PASS — four exclusively WP03-owned paths; other work preserved.
8. Production fragility: N/A — no new production code.

Approval is advisory for WP03's exact artifact commit only. The orchestrator records the normal CLI verdict. This review does not approve publication, external contact, main merge, or a repair of the runtime fault.
