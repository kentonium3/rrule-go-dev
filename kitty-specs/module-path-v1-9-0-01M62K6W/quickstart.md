# Implementation verification guide (planned, not executed)

Run these checks during the later implementation phase. This plan does not change library code or claim the candidate is available publicly. Use disposable paths and record actual values privately; do not create unmanaged Git worktrees or publish source.

## Declaration and source

Before edits, from the unchanged candidate cwd, record `go list -m` and the intentional mismatch:

```sh
GOWORK=off GOPROXY=off go list -m
test "$(GOWORK=off GOPROXY=off go list -m)" = github.com/kentonium3/rrule-go
```

The test must fail specifically because the declaration is still github.com/teambition/rrule-go. After edits the identical assertion passes. Inspect go.mod to verify the declaration and unchanged Go minimum. Review active import search results against occurrence_map.yaml; old historical references are expected. Do not mistake a local replacement for proof of correct identity.

## External consumer

Create a temporary consumer directory outside the repository. Set `candidate_dir` to the actual absolute implementation source directory. Use a private temporary GOCACHE/GOMODCACHE if isolation is needed; keep them process-scoped. The example below assumes a current working Go toolchain and no third-party imports.

```sh
GOWORK=off GOPROXY=off go mod init example.invalid/rrule-migration-check
GOWORK=off GOPROXY=off go mod edit -require=github.com/kentonium3/rrule-go@v1.9.0
GOWORK=off GOPROXY=off go mod edit -replace=github.com/kentonium3/rrule-go="$candidate_dir"
GOWORK=off GOPROXY=off go mod tidy
GOWORK=off GOPROXY=off go list -m -json all
GOWORK=off GOPROXY=off go test ./...
```

Write a small consumer test importing the new path, parsing a fixed UTC DAILY/COUNT rule and checking expected occurrence timestamps and String output; include a keyed ROption constructor if desired. Record the replacement directory to show which code ran. The temporary replacement is fixture scaffolding only. Remove it from any proposed production instructions. No synthetic proxy or publication is needed.

## Migration walkthrough matrix

Before removing a protective old-path replacement, inspect module requirements and package imports for remaining old-path consumers. Removal can silently restore original upstream code and lose the fork's panic/hang fixes even when no rrule types cross interfaces. Coordinate those consumers' migration first, or remain on the existing pinned v1.8.3 replacement arrangement. Passing application tests alone do not establish preservation of those fixes.

Demonstrate a retained-old-import dependency case and the documented stop/coordinate/stay-pinned outcome.

For each case, preserve initial/final go.mod/import text and actual command output:

| Case | Initial state | Required action |
|---|---|---|
| Direct | Old import and old module requirement | Change import; require new identity; tidy. |
| All-version replacement | Old requirement; replace old => fork v1.8.3 | Remove old replacement with bare old path; migrate import/require; tidy. |
| Version-specific replacement | Replace old v1.8.2 => fork v1.8.3 | Remove `old@v1.8.2` (actual left-side version); migrate import/require; tidy. |

Use pre-migration source for the working old-state fixture and candidate source for the new-state fixture, each through local fixture-only replacements with GOPROXY=off. Keep an untouched copy of the documented initial directives; do not confuse source-location normalization for the intended production migration. Show that the exact removal command eliminates its old directive. Assert no obsolete old replacement remains and new imports select candidate source. Old direct requirements should disappear when unused; retained old-path consumers require the preflight above before protective replacement removal.

The public guide's normal `go get github.com/kentonium3/rrule-go@v1.9.0` is a later release-verification command, not evidence produced by local replacements.

## Behavior equivalence

Update the golden driver header and docs/repro/05-07-backlog/golden/README.md active build instructions: use a historical snapshot's old-import driver for old source and the candidate's new-import driver for candidate source. Never tell readers to build a new-import driver against an old-declaration snapshot. Preserve historical facts, corpora and baseline JSON.

Export pre-migration source at `737ae01be40048d1d2932dbafb623230efa8d361` into a temporary directory with `git archive` extraction. Build that snapshot's golden/main.go from its own cwd, using its old imports. Build candidate golden/main.go from candidate cwd, using new imports. Set GOWORK=off GOPROXY=off for builds. Record source identities and binary hashes.

For both binaries, run with the same absolute `corpus-dst.json` path and distinct `-out` files; repeat with `corpus-forms.json`. Verify 811 and 158 rows respectively, matching IDs/order and no unexpected timeout/error rows. Compare baseline/candidate outputs with `cmp` and require zero differences. Corrupt only a temporary output copy and confirm comparison fails. Never rewrite inherited corpora or baseline JSON.

The historical runner's `-baseline/-current` comparator deliberately requires differences, so do not use it as the zero-change migration assertion. Do not run unbounded panic/hang repros; compile changed drivers or use the bounded golden harness.

## Required repository checks

From candidate source:

```sh
go test ./...
go test -race ./...
go vet ./...
gofmt -l .
go run ./tools/changelog check
```

Require zero formatting paths. Preserve existing minimum-version CI. No new dependency or minimum-version increase is planned.

## Later checks

Before code changes: obtain Kent's occurrence-map approval. Before public actions: separate exact-copy/public-action authorization. After approved public tagging: verify declaration and clean `go get` against actual public source with ordinary proxy/checksum settings, then documentation and release metadata. Do not create v1.9.0 if it already exists.

When ready, prepare a private Vikunja handoff naming candidate commit, required imports/replacement changes and desired recurrence integration checks. Contact other sessions only after Kent authorizes that handoff. No Vikunja edits, messages or public distribution are part of this planning command.
