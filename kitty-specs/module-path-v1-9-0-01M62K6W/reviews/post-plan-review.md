Two actionable findings remain.

1. **P1 — Removing the replacement can silently restore unfixed upstream code for transitive users.**

   **Affected:** [plan.md:76](~/repos/rrule-go-dev/kitty-specs/module-path-v1-9-0-01M62K6W/plan.md:76), [release-communication.md:13](~/repos/rrule-go-dev/kitty-specs/module-path-v1-9-0-01M62K6W/drafts/release-communication.md:13), [teambition-transition-pr.md:88](~/repos/rrule-go-dev/kitty-specs/module-path-v1-9-0-01M62K6W/drafts/teambition-transition-pr.md:88).

   **Conditions and likelihood:** An application adopted v1.8.3 through a replacement, migrates its own imports, but retains a dependency importing the old path. This is a realistic migration case already contemplated by the plan. Removing the replacement returns that dependency to the selected original upstream release. Tests can pass, especially where no rrule types cross package boundaries, while the dependency loses the fork’s panic/hang fixes.

   **Smallest correction:** Put a dependency-graph check before replacement removal. Explain that retained old-path dependencies also lose replacement-provided fixes. Require coordinated migration where those fixes matter; otherwise advise remaining on the existing v1.8.3 arrangement pending dependency updates. “Update when possible” understates this consequence. No shim or additional test framework is needed.

2. **P2 — Existing golden reproduction instructions become incompatible with the renamed driver.**

   **Affected:** [plan.md:48](~/repos/rrule-go-dev/kitty-specs/module-path-v1-9-0-01M62K6W/plan.md:48), [golden/README.md:1](~/repos/rrule-go-dev/docs/repro/05-07-backlog/golden/README.md:1), [golden/main.go:4](~/repos/rrule-go-dev/docs/repro/05-07-backlog/golden/main.go:4).

   **Conditions and likelihood:** The candidate driver gains the new import, while its header and adjacent README continue directing readers to build the same current driver against the old76c646d module. Following those instructions will no longer select the historical local library; offline it cannot resolve the new import. This is certain when those retained instructions are followed.

   **Smallest correction:** Explicitly include these executable instructions in the occurrence inventory and implementation scope. Distinguish the historical classified comparison from the migration comparison, and specify an old-import driver for historical source and the new-import driver for candidate source. Preserve corpus and baseline JSON unchanged.

The occurrence map’s explicit exceptions cover the known declaration, external-test and reproduction-import paths despite the classifier’s path-based limitations. I found no additional significant privacy, publication-authorization or conditional-release wording flaw.

Planning can proceed to task authoring after these corrections, review resolution and Kent’s occurrence-map approval. This does not approve implementation or authorize publication. No files were edited, workflow commands run, or Git/worktrees mutated.