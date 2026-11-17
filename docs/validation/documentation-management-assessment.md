# Documentation management assessment — 2026-10-29

Scope: documentation organization under Kent's explicit no-mission instruction,
tracked in private issue 20. Read-only investigations by separate explorer agents
grounded this assessment in source and current library tests.

## Spec Kitty provenance and distinction

Inspected source line: `spec-kitty/spec-kitty`, commit
`d074349fdf91c91de20c13af48557698f5ee8205`, through the existing local checkout at
`~/repos/spec-kitty`. The checkout was clean before and after inspection;
GitHub commit reads confirmed this SHA in both upstream and the local fork's
repository. This identifies inspected source, not the installed CLI build. No
installation, upgrade, initialization or workflow command was run for this study.

There are two different things to assess: the documentation mission the product
provides to users, and the infrastructure Spec Kitty uses to publish its own
documentation.

| Evidence at the inspected commit | Finding | Decision for rrule-go |
|---|---|---|
| [Documentation design prompt](https://github.com/spec-kitty/spec-kitty/blob/d074349fdf91c91de20c13af48557698f5ee8205/packs/built-in/missions/mission-steps/documentation/design/prompt.md#L43) | Allocates content by learning/task/reference/explanation purpose, with justified deferrals | Adopt reader-purpose navigation; do not create four empty directory trees |
| [Audit prompt](https://github.com/spec-kitty/spec-kitty/blob/d074349fdf91c91de20c13af48557698f5ee8205/packs/built-in/missions/mission-steps/documentation/audit/prompt.md#L41) and [gap analysis](https://github.com/spec-kitty/spec-kitty/blob/d074349fdf91c91de20c13af48557698f5ee8205/src/specify_cli/doc_analysis/gap_analysis.py#L56) | Inventory/classification compares existing docs with implementation; automatic classification is heuristic | Apply an explicit inventory and source-backed gap check during this reorganization |
| [Validation prompt](https://github.com/spec-kitty/spec-kitty/blob/d074349fdf91c91de20c13af48557698f5ee8205/packs/built-in/missions/mission-steps/documentation/validate/prompt.md#L39) | Calls for executable examples, valid links and clean-environment reproducibility; the prompt is not itself a universal link checker | Compile/run examples and check links/anchors without installing a documentation platform |
| [Living-documentation directive](https://github.com/spec-kitty/spec-kitty/blob/d074349fdf91c91de20c13af48557698f5ee8205/packs/built-in/directives/037-living-documentation-sync.directive.yaml#L15) | Behavior changes should carry corresponding documentation/example updates | Adopt this as a contributor/review convention and keep canonical topic ownership |
| [Documentation mission](https://github.com/spec-kitty/spec-kitty/blob/d074349fdf91c91de20c13af48557698f5ee8205/packs/built-in/missions/documentation/mission.yaml#L7) and [setup wiring](https://github.com/spec-kitty/spec-kitty/blob/d074349fdf91c91de20c13af48557698f5ee8205/src/specify_cli/cli/commands/agent/mission_setup_plan.py#L749) | Discover/audit/design/generate/validate/publish uses mission artifacts, metadata and workflow commits | Skip for this approved documentation reorganization; it is not a lightweight standalone Markdown housekeeping operation |
| [Generator registry](https://github.com/spec-kitty/spec-kitty/blob/d074349fdf91c91de20c13af48557698f5ee8205/src/specify_cli/cli/commands/agent/mission_setup_plan.py#L817) | The inspected detection registry includes JSDoc, Sphinx and rustdoc, with no Go adapter there | Keep Go-native source comments, examples and pkg.go.dev as the API reference |
| [Spec Kitty's own navigation](https://github.com/spec-kitty/spec-kitty/blob/d074349fdf91c91de20c13af48557698f5ee8205/docs/toc.yml#L8) | Separates users and contributors | Separate consumer manual and private development entry points |
| [Relative-link fixer](https://github.com/spec-kitty/spec-kitty/blob/d074349fdf91c91de20c13af48557698f5ee8205/scripts/docs/relative_link_fixer.py#L2) | Repository-specific structural-move tooling; documented limitations include reference-style and HTML links | Borrow the move-map and before/after validation practice, not the tool wholesale |
| [Own publishing workflow](https://github.com/spec-kitty/spec-kitty/blob/d074349fdf91c91de20c13af48557698f5ee8205/.github/workflows/docs-pages.yml#L102) and [advisory lint](https://github.com/spec-kitty/spec-kitty/blob/d074349fdf91c91de20c13af48557698f5ee8205/.github/workflows/ci-router.yml#L224) | DocFX/.NET 8 plus Python generation, glossary, redirects and rendering are bespoke infrastructure; the cited lint command uses `|| true` | Do not import this stack or describe advisory lint as a required passing gate |

Conclusion: adopt the information architecture and verification practices. The
full workflow and publishing infrastructure are disproportionate to this small
Go library and this documentation-only task. Its publishing of mission records
would also be unsuitable for this private repository's public port.

## Concrete adoption here

- A five-page consumer manual with task-oriented index, usage, timezones, design
  and limitations. Keep API declarations in Go comments and executable examples.
- A separate private development index for issue briefs, runbooks, release
  research, reproductions and validation. Do not publish that index or evidence.
- Preserve root document URLs. Relocate long README demos into usage with working
  references. Check documented claims against current tests; mark old evidence
  as historical rather than silently treating it as current.
- Validate relative links and heading anchors plus runnable examples before
  review. No new publishing framework, CI language/runtime dependency or generator.
- Use an explicit public allowlist, as the existing changelog-flow runbook already
  requires. Documentation does not change the tested library candidate.

The independent review found a concrete reason to cross-check comments against
implementation and tests: the source comment at `str.go` describing RDATE/EXDATE
DATE support is stale, while `TestRDateValueDateStr` and the parser establish
acceptance. The guide was corrected to describe DATE inputs as midnight instants
without promising preserved DATE serialization. The Go comment is not changed
in this documentation-only port, preserving the active tested-source fingerprint.

## testdata/ decision

The public tree currently has no standalone fixture directory. Regression inputs
and expected values are inline Go tables; timezone tests synthesize TZif data
inline (`rrule_dst_test.go`), and fuzz seeds use `f.Add` (`rrule_fuzz_test.go`).
These are small, inspectable fixtures colocated with their assertions.

Private `docs/repro/05-07-backlog/golden/` contains JSON corpora, frozen historical
outputs and source-specific drivers. Their row counts, hashes and baseline
selection are evidence for completed compatibility proofs. They are not a folder
to move blindly into public packaging.

A future public `testdata/` would be justified by substantial reusable rule/input
and expected-output vectors, a stable round-trip text corpus, or nontrivial TZif
fixtures that are clearer as files. Each fixture needs a schema/format, named
cases, provenance/license, deterministic anchors/zones, reviewed expected results,
and a documented update procedure. Readers must fail on missing/malformed data;
golden regeneration must not redefine expected behavior to hide a regression.
No production/application data belongs there.

Moving current inline fixtures would modify Go tests. Extracting private baselines
would change evidence paths and possibly publishing scope. Both require a
separate scoped issue, reviewed equivalence and normal code-change governance.
Create no empty `testdata/` or `examples/` directory in this documentation task.
