# Documentation organization validation — 2026-10-29

Tracking: private issue20. Vehicle: Kent explicitly waived a full mission for this
documentation-only reorganization. Author model gpt-6.1-sol; independent gpt-6-astra review closed APPROVE after correcting a DATE-support
statement that had followed a stale Go comment. No runtime behavior, Go source/test, module declaration,
checksum file, LICENSE, historical baseline or release tag was changed.

## Authored scope and canonical topics

- Five public-safe docs pages: README/index, usage with relocated demonstrations,
  timezones, design/source map, limitations. Root README is shortened to entry
  points; old Demo section anchors remain as compatibility links.
- Root MIGRATION, CHANGELOG and other established root documents stay put.
  Changelog adds one Unreleased documentation entry. API comments and
  example_test.go remain unchanged.
- Private-only docs/development.md indexes existing private material.
  The source-backed documentation-management assessment stays in docs/validation/.
- Private README corrects stale embargo/candidate/completeness wording and removes
  the inherited upstream coverage badge; the private/public identities remain
  distinct. Public README preserves its maintainer, roadmap, install and badge
  content and changes only documentation organization plus the incorrect Gear
  license label (LICENSE bytes/attribution are unchanged).
- Relocated sample comments now describe seven DAILY occurrences without claiming
  weekend exclusion, and show Between output in the helper's actual line format.

## Executed checks

```json
{
  "links": [
    {
      "tree": "~/repos/rrule-go-dev",
      "relative_links_checked": 53
    },
    {
      "tree": "/tmp/rrule-docs-organization/public-tree",
      "relative_links_checked": 45
    }
  ],
  "usage_go_blocks_compiled": 4,
  "usage_functions_run": 4,
  "timezone_output": "2023-03-12T07:30:00Z\n2023-03-13T06:30:00Z\n",
  "source_fingerprints_unchanged": 21,
  "public_privacy_allowlist_passed": true,
  "Go1.16_example_runs": [
    {
      "example": "usage-check",
      "exit": 0
    },
    {
      "example": "timezone-check",
      "exit": 0
    }
  ]
}
```

The temporary validation export at /tmp/rrule-docs-organization/public-tree was
built from public5bad5ace with an explicit seven-file allowlist. Its relative
links and anchors were checked separately from the private tree. Public marker
checks exclude private repository identifiers, mission/config paths and local
filesystem references. These checks do not assert remote URL availability.

All four relocated Go blocks were combined with their shared imports/helper,
compiled and executed (including every demonstration function). The standalone
timezone program's exact two UTC outputs were asserted. Both programs were also
run with the existing Go1.16.15 toolchain. Temporary consumers used a local
replacement only for documentation validation; they were not committed.

The 21 Go/module/LICENSE SHA256 values in both the private checkout and proposed
public export match the released/tested1f4 candidate manifest. No fixture directory
or fixture extraction is part of this change. Historical evidence stays fixed.

Changelog validation passed. Independent review and actual PR/CI/merge references
are recorded in issue20; this file does not claim gates that are still pending.

## Public export allowlist

README.md, CHANGELOG.md, docs/README.md, docs/usage.md, docs/timezones.md,
docs/design.md, docs/limitations.md. Consumer guide bytes match dev; root README
and CHANGELOG use selected public counterparts. Exclude docs/development.md,
all docs/validation records, private docs/issues/runbooks/releases/repro,
kitty-specs, kitty-ops and private configuration/history.

## Authoring check correction

The first complete staged whitespace check reported `docs/usage.md:197: new blank
line at EOF` and stopped before commit. The outer tool sequence incorrectly
continued: it pushed the unchanged base branch and PR creation returned422
`No commits between main and codex/docs-organization`. No PR or documentation
content was published by that sequence. The surplus EOF blank line was removed
in dev and public preview; the checked commit/push sequence was rerun with failure
propagation. This changed no example semantics or library source.
