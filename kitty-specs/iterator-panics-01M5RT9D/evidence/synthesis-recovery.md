# Charter synthesis recovery — 2026-10-25

Installed build: spec-kitty/spec-kitty 1a0c8b467af253eb0185fa19d0d50880b2b6475e.

Parent-observed failure: implementation claim refused `synthesized_drg missing`; `SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty charter synthesize --dry-run --json` failed with GeneratedArtifactMissingError for `.kittify/charter/generated/directives/001-project-intent-directive.directive.yaml`, expected ID PROJECT_001.

Diagnosis: charter generation produced structured governance but no agent-authored synthesis inputs or promoted manifest. The installed generated adapter requires those YAML bodies. This is resolved via the ordinary documented harness-authoring path, not fixture generation, hash restamping, installation changes, or gate bypass.

Read spk-charter-governance skill and synthesis workflow, current charter/interview, AGENTS.md, CONTEXT.md, installed target mapping, generated adapter and all three artifact schemas. The installed `_build_synthesis_request` + normalize/resolve/build/order APIs identified exactly eight targets. Authoring produced two directives, two tactics and four styleguides derived from existing project policy; fault guidance refers to the current AGENTS.md/run adapter so the operator-authorized D-3 exception is retained. No new substantive policy chosen.

Commands (repo cwd ~/repos/rrule-go):
- `SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty charter synthesize --dry-run --json` → exit 0, artifact_count 8, validated true, conflicts [], warnings []. Full output /tmp/rrule-synthesis-dry-run.log.
- `SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty charter synthesize --json` → exit 0, result success, generated adapter 1.0.0, eight artifacts. Full output /tmp/rrule-synthesis-promote.log.
- `SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty charter status --json` → exit 0, generation_state promoted, manifest valid, 8 live artifacts and 8 provenance records, no missing provenance, charter_source/synced_bundle/synthesized_drg all fresh. Full output /tmp/rrule-synthesis-status.log.
- `SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty charter validate .kittify/charter-packs` → exit 0. Full output /tmp/rrule-synthesis-validate.log.

Protected-byte comparison confirms charter.yaml, interview/answers.yaml and config.yaml unchanged. Promotion itself appended review-policy-tactic and testing-philosophy-tactic to activation.yaml (existing acceptance-test-first retained; empty directive/paradigm selections unchanged). This canonical CLI side effect was reported to parent; it was not manually reverted. No source/planning/runbook files edited by this agent. No commits.

Minor diagnostic-only failures: initial target introspection attempted `model_dump()` on a dataclass and returned AttributeError; corrected to print target dataclass. An attempted source read used an obsolete charter_runtime/freshness path and returned not found; canonical `charter status` supplied the actual freshness proof. Neither changed repository state.

Changed paths follow; all listed inputs/live artifacts/provenance/graph/manifests must be included when root commits. Activation is the sole modified existing file attributable to synthesis; parent's concurrent docs/runbooks change is not ours.
.kittify/charter-packs/directive/001-project-intent-directive.directive.yaml
.kittify/charter-packs/directive/002-risk-appetite-directive.directive.yaml
.kittify/charter-packs/graph.yaml
.kittify/charter-packs/styleguide/documentation-policy-style-guide.styleguide.yaml
.kittify/charter-packs/styleguide/go-style-guide.styleguide.yaml
.kittify/charter-packs/styleguide/quality-gates-style-guide.styleguide.yaml
.kittify/charter-packs/styleguide/testing-style-guide.styleguide.yaml
.kittify/charter-packs/tactic/review-policy-tactic.tactic.yaml
.kittify/charter-packs/tactic/testing-philosophy-tactic.tactic.yaml
.kittify/charter/generated/directives/001-project-intent-directive.directive.yaml
.kittify/charter/generated/directives/002-risk-appetite-directive.directive.yaml
.kittify/charter/generated/styleguides/documentation-policy-style-guide.styleguide.yaml
.kittify/charter/generated/styleguides/go-style-guide.styleguide.yaml
.kittify/charter/generated/styleguides/quality-gates-style-guide.styleguide.yaml
.kittify/charter/generated/styleguides/testing-style-guide.styleguide.yaml
.kittify/charter/generated/tactics/review-policy-tactic.tactic.yaml
.kittify/charter/generated/tactics/testing-philosophy-tactic.tactic.yaml
.kittify/charter/provenance/directive-project-intent-directive.yaml
.kittify/charter/provenance/directive-risk-appetite-directive.yaml
.kittify/charter/provenance/styleguide-documentation-policy-style-guide.yaml
.kittify/charter/provenance/styleguide-go-style-guide.yaml
.kittify/charter/provenance/styleguide-quality-gates-style-guide.yaml
.kittify/charter/provenance/styleguide-testing-style-guide.yaml
.kittify/charter/provenance/tactic-review-policy-tactic.yaml
.kittify/charter/provenance/tactic-testing-philosophy-tactic.yaml
.kittify/charter/synthesis-manifest.yaml
.kittify/charter/synthesized-artifacts.yaml
