# Planning operational evidence

2026-10-25, build1a0c8b467; gpt-6.1-sol.

- Attempt to read `.codex/prompts/spec-kitty.specify.md` returned file not found. Read CLI help and actual runtime-composed prompts instead; no mutation or gate bypass.
- `git add .../decisions` included the CLI-created `decisions/index.json.lock` in scoped commit e82c746. Inspection: zero bytes, permissions0600; lsof returned no holder (only unrelated filesystem stat warnings). This is the same known analysis-lock residue documented in the autonomous adapter. Preserve a copy at `/tmp/teambition-backlog-planning-index.json.lock`, remove the tracked empty lock, and commit correction. No decision/index/event data removed. Future staging must enumerate decision artifact/index files individually, never the directory. This was orchestration staging error, not an unsafe runtime state or substantive gate exception.

- Metadata-decision lock preservation (root report): empty index.json.lock had no holder. Path.rename to `/tmp/rrule-backlog-private-metadata-decision-index.lock` failed Errno18 Invalid cross-device link and made no mutation; shutil.copy2 plus unlink then preserved the empty file. No decision content removed. Scope-decision open/resolve similarly leaves an empty lock; no holder observed via lsof (unrelated filesystem warnings only); preserve at `/tmp/rrule-backlog-scope-decision-index.lock` with copy2/unlink rather than cross-device rename before scoped staging.
