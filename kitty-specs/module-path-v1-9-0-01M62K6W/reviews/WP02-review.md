# WP02 advisory review — APPROVE

Date: 2026-10-28 (America/New_York).
Actual reviewer model: gpt-6-astra (Reviewer Renata, profile resolved earlier in this same review session).
Actual implementer model: gpt-6.1-sol, separately assigned by orchestrator.
Implementation commit: 9bf894f4b899d0c1cb814e2f30753eeec6e01dab.
Workspace: ~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b.
Scope: WP02's README.md, CHANGELOG.md, and new MIGRATION.md only. WP01's inherited source delta was already independently approved.

No significant findings. The dependency preflight precedes removal commands and expressly covers loss of fork panic/hang fixes for retained old-path consumers even without shared types. It distinguishes module requirements from package imports, notes why-chain and build-selection limits, requires coordination or staying on the pinned v1.8.3 arrangement, and does not treat passing tests as proof of preserved fixes. Named-type incompatibility is described separately. Both replacement forms have the correct mutually exclusive removal commands, with the actual left-side version identified for the version-specific form.

Candidate/public language is conditional throughout new consumer text. Local replacements are labeled disposable verification scaffolding, not production installation or proof of tag/download/checksum/indexing availability. Unchanged rrule qualification and the separate keyed-ROption migration are explicit. No official-succession claim is added. All released changelog sections from v1.8.3 onward remain byte-identical. The inherited recurrence notes and attribution are preserved.

README CI/CodeQL links correctly select the private development repository and existing workflow names; Go Reference selects the intended new identity without claiming verified indexing. Its existing private preamble/internal guide and unrelated stale prose remain outside selected public hunks. A later public port must reconcile public badge targets and exclude this private preamble/internal guide; this review is not approval to copy the README wholesale. New MIGRATION and changelog prose contain no private mission IDs or evidence paths.

## Independent checks

- GOWORK=off GOPROXY=off GOCACHE=/tmp/rrule-module-path-wp02-review-cache go run ./tools/changelog check: PASS, CHANGELOG.md: valid.
- Three disposable candidate fixtures under /tmp/rrule-wp02-review-cases-kmarr9jq: direct, all-versions replacement, and version-specific replacement. Each starts with an old requirement and its applicable original replacement directive, exercises the exact documented removal form where applicable, asserts removal, applies the documented candidate requirement/local replacement, runs tidy and go list -m -json all, and passes go test ./.... Test invokes production NewRRule with keyed options and checks three DAILY occurrences ending at the expected timestamp. Tidy leaves no old-path requirement in these fully migrated fixtures. This tests candidate migration command semantics, not historical public downloads or working old-state fixtures; WP03 owns complete integrated walkthroughs.
- Local links/section anchors and private CI/CodeQL workflow files exist.
- Programmatic comparison of released CHANGELOG sections: byte-identical.
- git diff --check: PASS. Only three owned files changed in the implementation commit; repository remains clean.

Environmental limitation: the separate documented go list -deps -test -f '{{.ImportPath}}: {{join .Imports " "}}' ./... command in an earlier disposable fixture failed with 'go: can’t load test package: error obtaining VCS status: exit status 128; Use -buildvcs=false to disable VCS stamping.' This was surfaced to the orchestrator; the original command is not claimed as passing. With explicit orchestrator authorization, a fixture-only rerun using process-scoped GOFLAGS=-buildvcs=false succeeded in /tmp/rrule-wp02-review-cases-kmarr9jq/direct. Its output shows the consumer test imports github.com/kentonium3/rrule-go and the selected library package imports only standard-library packages; no old-path package appeared. This isolates sandbox VCS stamping from module migration and does not change the public guide. The ordinary candidate tests and migration commands passed without special build flags. This is not a demonstrated defect in the changed documentation and does not block this scoped review; WP03 should retain the original limitation and scoped rerun distinction.

## Required anti-pattern checklist

1. Dead code: N/A — documentation only; no new public implementation.
2. Synthetic-fixture test: PASS — lightweight walkthrough invokes actual Go module editing and production recurrence code; no string-mirroring test added to repository.
3. Silent empty return: N/A — no new runtime paths.
4. FR coverage: PASS for allocated documentation portions — FR-003's three shapes and safety instructions inspected, commands exercised; FR-004's consumer changelog contribution checked. FR-007's final release/deprecation notification packet and final integrated walkthrough evidence belong to WP03, not a claimed WP02 delivery.
5. Frozen surface: PASS — only three owned files modified; historical released changelog sections verified identical.
6. Locked decision: PASS — no source/API/dependency changes, new notification mechanism, synthetic proxy, public availability assertion, or publication.
7. Shared-file ownership: PASS — implementation delta only touches WP02-owned files; selected public-copy handoff remains required.
8. Production fragility: N/A — no runtime code.

The explicit orchestrator instruction supersedes prompt directions to perform a verdict transition. This advisory approves WP02 only. No repository edits, workflow transitions, public actions, network requests, or contacts were performed. Final consumer/golden/race integration and release-packet checks remain WP03 responsibilities; public-tag retrieval remains a separately authorized release check.
