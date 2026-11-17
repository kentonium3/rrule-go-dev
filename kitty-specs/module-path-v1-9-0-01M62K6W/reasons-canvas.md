# REASONS Canvas: rrule-go v1.9.0 module path migration

## Requirements

Consumers need the maintained fork to declare its own identity and provide verified migration from original imports and transitional replacements. Kent needs accurate, reviewable release and upstream-proposal copy. The authoritative requirements and acceptance scenarios are in [spec.md](spec.md); Kent confirmed the intent in decision `01M62K7ADP565BH1469Q0VEJ12`.

Definition of done for implementation: new-identity adoption and three migration fixtures pass, recurrence golden outputs remain unchanged, required checks/reviews pass, and communication packets meet publication and upstream-authority boundaries. This specification step does not implement or release those outcomes. All public actions and upstream contact remain excluded.

## Entities

- Original module: `github.com/teambition/rrule-go`, whose maintainers retain authority over its public metadata.
- Maintained fork/new module: `github.com/kentonium3/rrule-go`, separately maintained with candidate target v1.9.0.
- Private development repository: `kentonium3/rrule-go-dev`; mission planning branch `rrule-go-dev-xfer`, eventual private PR target `main`.
- Consumer: fresh adopter or existing application using direct upstream dependency, all-version replacement or version-specific replacement.
- Release packet: proposed changelog/announcement/migration copy, gated on selected release content and verified availability.
- Upstream proposal: unposted request for maintainer agreement, not an official succession or transfer.

## Approach

Directly rename module declaration and active imports; preserve historical references and recurrence algorithms. Kent approved explicit identity inspection, a small local external consumer, migration command walkthroughs and zero-difference golden output comparison. Reject synthetic proxies, old-path compatibility shims and runtime notification machinery. See [plan.md](plan.md), [research.md](research.md) and [quickstart.md](quickstart.md). Later public-tag retrieval and authorized Vikunja integration are complementary deferred checks, not current actions.

## Structure

The library remains one Go module with unchanged exported declarations. The [occurrence map](occurrence_map.yaml) governs active imports and mixed-purpose documentation while preserving history. [data-model.md](data-model.md) and [consumer contract](contracts/module-migration.md) describe identities, local-fixture boundaries and release authority. Planning/consolidation targets rrule-go-dev-xfer; eventual private PR targets main. No tasks or implementation generated during planning.

## Operations

Task generation defines a sequential chain: WP01 owns the declaration/import migration and active golden-driver instructions; WP02 owns root consumer migration/README/Unreleased copy; WP03 owns private acceptance evidence and final draft/handoff packets. The authoritative dependency/requirement index is wps.yaml and its CLI-generated tasks.md. Before implementation, run the required analyze gate after task finalization. Implementation and public actions are not authorized by task generation.

## WP boundaries and safeguards

WP01 owns only go.mod, the two external test import files, seven reproduction driver files and golden README. WP02 owns README.md, CHANGELOG.md and planned-new MIGRATION.md. WP03 owns only the two draft files plus planned-new private migration-verification and Vikunja handoff records. No ownership overlap, no code_change WP owns kitty-specs. WP03 reports source/doc defects through the orchestrator for the owning package; it does not edit those surfaces itself. Runtime commands own worktrees, lanes, transitions and verdicts. Actual implementation/review model identities and date must be recorded and differ; profile names alone are not proof.
