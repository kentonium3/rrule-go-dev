# 04: Adoption in Vikunja (cross-repo; not a mission here)

This work happens in the Vikunja repository, not here (Vikunja PR #3071 adopted v1.8.3 and later moved to v1.9.0). It's listed so this repo's release shape suits it.

- **Module path.** Vikunja depends on `github.com/teambition/rrule-go`. It can adopt the fix through a `replace` directive to this repo's module at a tag, or by vendoring a patch. Vikunja is public, so a `replace` to a **private** repo cannot ship. Making this repo public (after 2026-11-13, with Kent's say-so) or vendoring is a plan-phase decision there.
- **Tags.** Tag releases here as `v1.8.3-kg.N`, or similar, so the `replace` pins an exact version.
- **What stays in Vikunja.** Vikunja keeps its own guards (`docs/CONTEXT.md`) as defence in depth. The test of adoption is that Vikunja's recurrence suites pass on the patched module.

**Done 2026-10-25.** Vikunja PR #3071 uses `replace github.com/teambition/rrule-go => github.com/kentonium3/rrule-go v1.8.3`. Releases are plain semver tags on the public repo (no `-kg.N` suffix, which would sort below the release).
