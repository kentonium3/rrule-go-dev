# v1.9.0 migration communication — private draft

Prepared 2026-10-28 for Kent. Not published. These are proposed migration excerpts, not a complete v1.9.0 release changelog. Reconcile with all changes actually selected for the public release. Keep this file and all mission records private.

## Release gate

Use the announcement below only after the public v1.9.0 tag has the new module declaration and a clean consumer can retrieve and use it. Before that point, describe this as a planned change. Obtain Kent's sign-off on all exact public copy before posting. No release date has been selected here.

## Proposed release announcement excerpt — after verification

rrule-go v1.9.0 uses `github.com/kentonium3/rrule-go` as its module path. This is a separately maintained fork of teambition/rrule-go, retaining its MIT license and attribution.

Before removing a protective old-path replacement, inspect module requirements and package imports for remaining old-path consumers. Removal can silently restore original upstream code and lose the fork's panic/hang fixes even when no rrule types cross interfaces. Coordinate those consumers' migration first, or remain on the existing pinned v1.8.3 replacement arrangement. Passing application tests alone do not establish preservation of those fixes.

Update your imports from `github.com/teambition/rrule-go` to `github.com/kentonium3/rrule-go`. The package name remains `rrule`. If your application uses the transitional `replace` directive, remove it before adding the new module:

```sh
go mod edit -dropreplace=github.com/teambition/rrule-go
go get github.com/kentonium3/rrule-go@v1.9.0
go mod tidy
go test ./...
```

The replacement removal above applies to the documented all-version replacement. If the left side names a specific version, remove that exact directive instead; for example, `go mod edit -dropreplace=github.com/teambition/rrule-go@v1.8.2` removes a directive whose left side names v1.8.2. Substitute its actual left-side version. Applications without a replacement can omit that step.

Changing your own imports does not migrate dependencies that still use the original package. The old and new packages define distinct types; update dependent libraries together where their APIs exchange rrule values. Review the full changelog for all v1.9.0 changes, including the previously announced migration from positional `ROption` literals to keyed literals.

Existing release tags remain unchanged. Pinned consumers can keep their existing dependency. Go will not automatically rewrite imports or migrate a version-pinned replacement to the new module identity.

## Communication delivery notes — private

- Put installation and migration instructions in the public README and consumer-facing Unreleased changelog entry during implementation. Preserve historical release facts.
- Publish a GitHub Release only under separate authorization; users who watch Releases can receive GitHub notifications. A Git tag alone is not that notification.
- The teambition PR is a proposal to upstream maintainers. No official succession or endorsement is asserted. A README-only pointer remains an alternative if they do not want deprecation.
- An upstream `Deprecated:` comment requires a new old-module release selected as latest before it can supply tagged-release deprecation metadata. It is not a guaranteed warning for consumers using unconditional replacements.
- Do not retract a working historical release merely to advertise this move. Do not add runtime logging or network-based notifications to library consumers.
- Before publication, reconcile selected changes through `docs/runbooks/changelog-flow.md`, verify installation/proxy/documentation using the actual public tag, and recheck all links and copy.

## Final private reconciliation — 2026-10-28

Local candidate verification is recorded in [migration verification](../evidence/migration-verification.md); the unsent [Vikunja handoff](../evidence/vikunja-handoff.md) lists later integration and public-port boundaries. The candidate guide and Unreleased entry agree on dependency preflight, both alternative removal forms, and separate inherited keyed-ROption migration. The excerpt above is conditional after-release copy, not a claim that v1.9.0 is currently retrievable.

Before posting, confirm the proposed version is unused; if v1.9.0 has already shipped, stop for version direction. Verify the actual public tag declaration, ordinary clean consumer retrieval with normal checksum/proxy settings, package documentation, links, and separately approved release metadata. Reconcile the full selected recurrence/input-validation notes with this migration excerpt; it is not the complete release notes. GitHub Release notifications are opt-in and a tag alone does not deliver them. The 2026-10-28 maintenance audit supports neutral fork language only; current permissions and private author conversations were not independently established.
