# Retrospective summary selects retained coordination checkout

Date: 2026-10-25. Installed build: 1a0c8b467af253eb0185fa19d0d50880b2b6475e.

After successful consolidation with --keep-branch --keep-worktree, automatic
commit155321d writes kitty-specs/iterator-panics-01M5RT9D/retrospective.yaml on
fix/iterator-panics. Its findings_status is has_findings, with three helped rows.
No mutation was made while diagnosing this fault.

Command: `SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty retrospect summary`
Exit0. Counts show mission_count1/completed1, but Record State Summary shows
missing1/has_findings0. `agent retrospect synthesize --mission
iterator-panics-01M5RT9D` dry-run exits0, no planned applications.

Diagnosis: installed cli/commands/retrospect.py lines1315-1329 discovers primary
mission directories, then replaces each with candidate_feature_dir_for_mission
when that directory exists. For this retained coordination topology, the latter
points to the coordination checkout, where the post-consolidation retrospective
has not been written. The summary therefore reads a different surface.

Direct read-only verification with installed Python and the exact functions:
- classify_mission_record(primary kitty-specs directory) => has_findings;
  retrospective.yaml exists.
- candidate_feature_dir_for_mission(root,slug) =>
  .worktrees/iterator-panics-01M5RT9D-coord/kitty-specs/iterator-panics-01M5RT9D;
  record does not exist; classify_mission_record => missing.

Safe continuation: inspect and validate the actually committed primary record;
retain the summary discrepancy as a tooling finding. Do not fabricate a second
record, modify the installation, or treat the missing summary as evidence the
retrospective was not authored. Existing record contains no staged proposals,
so the dry-run's no-application result is consistent. No upstream report sent;
this private evidence is ready for later dedupe/copy approval.
