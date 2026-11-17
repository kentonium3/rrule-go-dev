# Teambition transition PR — private draft for Kent

Prepared 2026-10-28. Nothing in this file has been posted or pushed. This is proposed public copy for Kent's review, followed by private release notes. No library source files have been changed.

## Relationship and maintainer-consent gate

This is a separately maintained fork, not an asserted official successor or takeover. Public history supports low upstream activity but does not prove that every maintainer has stepped away. Ask whether the upstream maintainers welcome a README pointer and module deprecation; offer a README-only alternative if deprecation is too strong. Any eventual PR is a proposal for their agreement, not an announcement of their decision. Kent must approve exact public copy before any contact or posting.

## Publication gate

The main PR copy and patches below are the **publish-after-release** variant. Do not submit them until the public fork has released and verified `github.com/kentonium3/rrule-go@v1.9.0` with that exact module declaration. The existing v1.8.3 fork retains the upstream module declaration and uses a replacement; it is not evidence that direct installation at the new path works.

If the first own-path release uses another version, revise every v1.9.0 reference and obtain Kent's sign-off on the resulting exact copy. Do not rewrite the v1.8.3 tag.

## Proposed PR title — after release

docs: propose a pointer to a separately maintained rrule-go fork

## Proposed PR body — after release

Would you welcome a README pointer to the separately maintained fork at [kentonium3/rrule-go](https://github.com/kentonium3/rrule-go), together with the proposed advisory module-deprecation message? If deprecation is not appropriate, the README pointer can be considered on its own. This proposal does not imply an official transfer or endorsement.

The fork's v1.9.0 release uses `github.com/kentonium3/rrule-go` as its module path. Users migrate by updating imports and dependencies; the README describes both ordinary upstream dependencies and applications using the fork through a `replace` directive. The old module declaration stays `github.com/teambition/rrule-go`, so existing users can continue building with it.

The fork incorporates fixes for the DST, DTSTART/UNTIL parsing, and repeated-BY-value problems reported in [#63](https://github.com/teambition/rrule-go/issues/63), [#67](https://github.com/teambition/rrule-go/issues/67), and [#70](https://github.com/teambition/rrule-go/issues/70). Related proposals in [#66](https://github.com/teambition/rrule-go/pull/66) and [#68](https://github.com/teambition/rrule-go/pull/68) were evaluated, with behavior independently checked against RFC 5545. See the [release notes](https://github.com/kentonium3/rrule-go/blob/main/CHANGELOG.md) for compatibility changes and migration instructions.

The `Deprecated:` comment is advisory metadata. Go 1.17 and later can surface it during relevant dependency queries; it does not emit a warning on every build or migrate a consumer automatically. Making it visible to users of tagged upstream releases requires a separate upstream release containing this comment, selected as the module's latest version. This PR does not create that release.

Validation: the proposed `go.mod` change adds only a comment, retaining the module path and Go 1.16 minimum. The migration commands must be verified against the fork's public v1.9.0 release before this PR is submitted.

## Exact proposed go.mod diff — after release

Baseline: `teambition/rrule-go` master at `e74d163475cf1ca1fd019752c5c41ea1f472d4c5`.

```diff
diff --git a/go.mod b/go.mod
--- a/go.mod
+++ b/go.mod
@@ -1,3 +1,4 @@
+// Deprecated: use github.com/kentonium3/rrule-go instead; see https://github.com/kentonium3/rrule-go/blob/main/MIGRATION.md for migration instructions.
 module github.com/teambition/rrule-go

 go 1.16
```

## Proposed README notice — after release

Insert the following immediately after `# rrule-go`, before the existing introductory sentence. The rest of the upstream README remains historical documentation of the old module.

> **Separately maintained fork available.** For additional fixes and releases, see [kentonium3/rrule-go](https://github.com/kentonium3/rrule-go).
>
> Starting with v1.9.0, the maintained fork uses the module path `github.com/kentonium3/rrule-go`. Existing releases of this original module remain available.

### Migrating to the maintained fork

Before removing a protective old-path replacement, inspect module requirements and package imports for remaining old-path consumers. Removal can silently restore original upstream code and lose the fork's panic/hang fixes even when no rrule types cross interfaces. Coordinate those consumers' migration first, or remain on the existing pinned v1.8.3 replacement arrangement. Passing application tests alone do not establish preservation of those fixes.

Change imports in your Go source from:

```go
"github.com/teambition/rrule-go"
```

to:

```go
"github.com/kentonium3/rrule-go"
```

If your application's `go.mod` currently replaces `github.com/teambition/rrule-go` with the fork, remove that replacement:

```bash
go mod edit -dropreplace=github.com/teambition/rrule-go
```

For a replacement that explicitly names an old version on the left-hand side, remove that exact version-specific directive instead. For example, if its left side is `github.com/teambition/rrule-go v1.8.2`, use:

```bash
go mod edit -dropreplace=github.com/teambition/rrule-go@v1.8.2
```

Substitute the actual left-side version from your application's directive. Then add the maintained module and tidy dependencies:

```bash
go get github.com/kentonium3/rrule-go@v1.9.0
go mod tidy
go test ./...
```

Review the fork's [release notes](https://github.com/kentonium3/rrule-go/blob/main/CHANGELOG.md) for behavior changes and migration requirements. In v1.9.0, external positional `rrule.ROption` literals need to be converted to keyed literals.

If another dependency still imports the original path, it may remain in your module graph. Coordinate that dependency migration before removing a replacement that supplied the fork fixes, or remain on the existing arrangement. Values from packages under the old and new import paths have different Go type identities; changing only your own imports does not migrate a dependency's public API.

## Preparatory wording — safe before the own-path release

This is an alternative README-only notice if Kent wants an earlier public pointer. It does **not** advertise v1.9.0 installation and does **not** include the go.mod deprecation patch. It still requires exact-copy approval before posting.

Proposed title:

`docs: propose a pointer to a separately maintained fork`

Proposed body:

`Would you welcome a README notice pointing users to the separately maintained fork and its current installation instructions? This does not propose an official transfer or imply endorsement. The fork is preparing a module-path migration; this change makes no claim that the planned own-path release is already available. It does not change this module's identity or release tags.`

Proposed README insertion after `# rrule-go`:

> **Separately maintained fork available.** Additional fixes are developed at [kentonium3/rrule-go](https://github.com/kentonium3/rrule-go). Follow that repository's [current installation instructions](https://github.com/kentonium3/rrule-go#using-it); its v1.8.3 release uses a `replace` directive while retaining the original import path. A move to the fork's own module path is planned. Check its released versions and migration instructions before changing imports.

## Private release coordination notes — not part of the PR body

1. Release and verify the fork's own-path version first. Verify the public tag's `go.mod`, an isolated consumer install, actual imports, tests, and public module/documentation availability. These checks have not been executed as part of this copy preparation.
2. Recheck upstream master and tags immediately before preparing a branch. The baseline fetched for this draft is `e74d163475cf1ca1fd019752c5c41ea1f472d4c5`; go.mod blob `c5cf4e2afc4b0195f82cbd4d10ddae146f836b83`, README blob `3b75cbdbe88a3ab47faa5eeefcbbfd2aeeb8c713`.
3. Kent approves the exact public title, body, README notice, and deprecation text. Only then may an authorized later action create/push a branch and submit the PR. This artifact does neither.
4. The upstream maintainer must merge the PR and choose/authorize a fresh upstream release tag. Do not assign v1.8.3 or any other tag in this draft: fork and upstream version histories are distinct, and current tag availability must be checked. The release containing deprecation metadata must be the old module's `@latest` selection; merge alone is insufficient for users following tagged releases.
5. Keep the old module declaration intact. Do not redirect it by changing the upstream repository's `module` line to the fork path, delete historical tags, or alter release contents.
6. Verify the old-module deprecation through Go queries after the upstream release. Native metadata is an additional signal, not a guaranteed notification channel.

### Go warning limitations

- The module-deprecation mechanism is available since Go 1.17. It is distinct from a minimum-version increase: adding a comment leaves the upstream `go 1.16` directive unchanged.
- Go gets deprecation metadata from the version selected by `@latest`, ignoring retractions/exclusions for that lookup. A comment only on a branch or an older tag will not reliably supply that metadata.
- Relevant `go get` operations and `go list -m -u` can expose deprecation. Ordinary build/test commands are not a dependable notification mechanism, and offline or cached queries may not immediately reflect new metadata.
- `CheckDeprecation` skips an original module if all its versions have been replaced. Thus a new upstream deprecation does not guarantee warnings to applications using `replace github.com/teambition/rrule-go => github.com/kentonium3/rrule-go@v1.8.3`.
- `go list -m -u` checks replacement metadata separately. `go get` does not promise the same replacement warning. Putting a deprecation comment into a transitional fork tag is not reliable either: a later healthy own-path release becomes fork `@latest` and supersedes that comment.
- Deprecation comments do not send email, identify all users, upgrade dependencies, or change import/type identities. README/release notes and explicit communications approved by Kent remain needed.

## Primary sources

- [Verified upstream baseline](https://github.com/teambition/rrule-go/commit/e74d163475cf1ca1fd019752c5c41ea1f472d4c5).
- [Upstream go.mod at baseline](https://github.com/teambition/rrule-go/blob/e74d163475cf1ca1fd019752c5c41ea1f472d4c5/go.mod).
- [Upstream README at baseline](https://github.com/teambition/rrule-go/blob/e74d163475cf1ca1fd019752c5c41ea1f472d4c5/README.md).
- [Fork installation instructions](https://github.com/kentonium3/rrule-go#using-it) — v1.8.3 uses replacement at preparation time.
- [Go module-deprecation reference](https://go.dev/ref/mod#go-mod-file-module-deprecation).
- [Go CheckDeprecation implementation](https://go.dev/src/cmd/go/internal/modload/modfile.go#L303).
- [Go get deprecation checks](https://go.dev/src/cmd/go/internal/modget/get.go#L1628).
- [Official ListModules source showing separate replacement processing](https://go.googlesource.com/go/+/67f7e16bcce0a3e68bf92e233ea16c8a9d2ac07a/src/cmd/go/internal/modload/list.go).

The historical Go source link documents replacement traversal; recheck the supported consumer toolchain when validating the eventual transition.

## Final private reconciliation — 2026-10-28

The approved local candidate declares the new identity and has a complete MIGRATION.md. The proposed deprecation points directly to that guide after public port. This is local implementation readiness only; no public tag or retrieval was verified here.

Relationship wording follows the dated 2026-10-28 [maintenance audit](../evidence/upstream-maintenance-audit.md). It found no public transfer or endorsement and could not establish current repository permissions or independently inspect Kent's private author exchange. Recheck those limits before contact; original maintainers decide README/deprecation acceptance and any final old-module tag. The audit is evidence, not contact authorization.

The preparatory notice's current-installation link must be rechecked against the actual public README before posting. Private development badges and README preamble are excluded from any proposed upstream copy. All public sections above require Kent's exact-copy approval and the applicable availability gate.
