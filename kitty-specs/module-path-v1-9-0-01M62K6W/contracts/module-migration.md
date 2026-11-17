# Consumer and release contract

This is a Go module/source distribution contract, not an HTTP API. No OpenAPI schema is applicable.

## Candidate identity

- Declared module: github.com/kentonium3/rrule-go.
- Package qualifier: rrule. Existing exported declarations/signatures and recurrence outputs unchanged by this mission.
- Candidate version intent: v1.9.0; not proof that the public tag exists or is installable.
- Local verification: explicit declaration/import checks plus disposable new-path-to-candidate replacement. That replacement is development-only and does not prove public availability.

## Consumer migration

Before removing a protective old-path replacement, inspect module requirements and package imports for remaining old-path consumers. Removal can silently restore original upstream code and lose the fork's panic/hang fixes even when no rrule types cross interfaces. Coordinate those consumers' migration first, or remain on the existing pinned v1.8.3 replacement arrangement. Passing application tests alone do not establish preservation of those fixes.

1. Change active imports from github.com/teambition/rrule-go to github.com/kentonium3/rrule-go.
2. If present, remove the old module replacement using its exact left-hand side:
   - All versions: `go mod edit -dropreplace=github.com/teambition/rrule-go`.
   - Version-specific example: `go mod edit -dropreplace=github.com/teambition/rrule-go@v1.8.2` (substitute actual left-side version).
3. After actual release, `go get github.com/kentonium3/rrule-go@v1.9.0`, `go mod tidy`, and application tests verify real adoption. Private walkthrough substitutes a clearly labeled local candidate requirement/replacement and cannot claim the public get was run.
4. Coordinate dependencies exposing old-path named types; retain legitimately needed old transitive requirements. Review the separate keyed-ROption migration in the selected release notes.

## Upstream and communication

Upstream proposal retains module github.com/teambition/rrule-go and asks maintainers about a README pointer, optionally deprecation/final release. No official transfer or universal warning is promised. Historical tags remain immutable. Public excerpts contain no private mission IDs/evidence paths. All public actions, messages and session coordination remain excluded until separately authorized.
