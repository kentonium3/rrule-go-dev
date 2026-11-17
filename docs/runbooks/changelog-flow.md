# Changelog: dev → public → release

`rrule-go-dev` is the authoring repository. `kentonium3/rrule-go` is the public
distribution repository. Both use a regular root `CHANGELOG.md` and the same
standard-library-only Go helper. Public tags establish what has shipped.

## Normal change flow

1. Implement and review a change in dev, with a consumer-focused Unreleased entry
   in the same dev PR. Explain old/new behavior and migration when relevant.
   Use public issue links and contribution credits. Mission records and evidence
   paths stay in this repository's development documentation, and private material
   (internal IDs, other private systems) is never committed here.
2. Merge the dev PR through the normal review process. Dev's Unreleased
   section may now be ahead of public: it describes implemented release
   candidates, not changes already available to public users.
3. Prepare the public port from current public main. Include the reviewed code,
   tests, and matching changelog entries in the same public PR, with any necessary
   public-facing documentation. Inspect an explicit file allowlist; do not merge
   the development repository wholesale or overwrite public README/CONTRIBUTING with
   their development counterparts. Preserve attribution and port provenance.
4. Review the exact public diff and pass public CI. Public posting still requires
   Kent's authorization under AGENTS.md. Once the public PR merges, those entries
   describe changes on public main awaiting a release tag. Deferred dev changes
   remain only in dev's Unreleased section.

Routine internal changes need no consumer note. Public roadmap items and
milestones describe intent; they are not implemented or released changelog entries.

## Release preparation and reconciliation

The public release PR selects the changes actually included, moves only their
entries into `## [X.Y.Z] - YYYY-MM-DD`, and retains `## [Unreleased]`. Use an actual
release date and the exact intended public tag. A dated heading alone is not proof
that the tag or release exists. `go 1.16` in go.mod is the compatibility floor,
not a release version.

Run from the public checkout before tagging, substituting the intended version:

```sh
go run ./tools/changelog check
go run ./tools/changelog check -version v1.9.0
go run ./tools/changelog extract -version v1.9.0
```

Use the extracted body as the GitHub release notes after the normal code gates and
release approval. The helper and CI validate notes; neither publishes a release.
Do not create release tags or GitHub releases in dev as part of this flow.

After the public release is verified, bring its finalized section and actual
version/date back to dev in a small reconciliation PR. Remove only the dev
Unreleased entries included in that public release; retain pending entries. Keep
released sections consistent between repositories, correcting factual errors
rather than restyling history. This narrow metadata return is the expected reverse
flow; normal implementation continues dev → public.

## Tooling and documentation changes

Future helper, test, and changelog-workflow changes start in dev too. Port those
files together and verify that the public copies match the reviewed dev versions.
The current bootstrap is the exception: public PRs 12/13 established the practice
first; this dev adoption brings it back into the normal authoring flow.

The workflow checks PRs/main and matches finalized notes on version-tag pushes.
Its tag trigger is a check, not authorization to create a dev or public tag.

## Entry format

Use only populated categories, in this order: Breaking, Upgrade Notes, Added,
Changed, Fixed, Deprecated, Removed, Security. Start entries with a short bold
headline; use public links and inline URLs. The helper checks structure, duplicate
versions, dates, populated sections, and conflict markers; human review checks
accuracy, migration, public suitability, and which changes the port includes.

Detailed release research can remain in `docs/releases/` (this repository is public,
so it must not contain private material), but it supplies
background rather than becoming a second independently maintained public release
history. The public changelog and its extracted notes are the consumer record.
