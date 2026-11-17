# PR 9 CI configuration review — APPROVE

Date: 2026-10-25.
Reviewer: codex / gpt-6-astra (reviewer-renata).
Implementer: codex / gpt-6.1-sol. Independent-model requirement satisfied.
Parent-owned Spec Kitty Op: 01M5RYFFX07KTDKP0Z0E3AJZXK.
Workspace: ~/repos/rrule-go, branch fix/iterator-panics.
Review base HEAD: a836d3a805baeb3fd7336f156e2cfdf54cbced01.
Scope: uncommitted .github/workflows/ci.yml and .github/workflows/codeql.yml only.

APPROVE. No blocking findings. This approves the configuration change for commit/push and hosted verification; it does not claim that hosted CI has run or passed, or authorize a merge.

## Checks and reasoning

- The new `pull_request` trigger targets master and retains the existing master push trigger. It uses the ordinary PR event, not pull_request_target. The job checks out the standard PR merge ref; it does not interpolate PR text into shell commands, use a privileged custom checkout, reference secrets, or acquire a write token.
- CI now explicitly grants only contents:read. This supports checkout and leaves unspecified token permissions disabled. Existing setup/test steps need no additional repository write permission. GitHub documents the target-branch filter and permission semantics in [workflow syntax](https://docs.github.com/en/actions/reference/workflows-and-actions/workflow-syntax).
- All six existing Go/OS combinations are preserved. Rendered job names include both dimensions and are unique: Go 1.18.x and 1.19.x, each on ubuntu-latest, macos-latest and windows-latest. Existing race/test commands and action versions are unchanged. The separate tag-only coverage workflow is unchanged and its job name does not collide with these names.
- CodeQL has a job-level boolean condition `github.event.repository.private == false`. For this private repository's configured push/PR/schedule events, it is false and prevents the analysis job from starting. Job conditions precede matrix expansion. This is an explicit skip with an explanatory comment, with no continue-on-error, fake successful scan, or failed-analysis suppression. Existing narrowly scoped CodeQL permissions are unchanged and the skipped job executes no actions. See [job condition semantics](https://docs.github.com/en/actions/reference/workflows-and-actions/workflow-syntax#jobsjob_idif).
- The stated private-personal-repository limitation matches current [GitHub code-scanning availability](https://docs.github.com/en/code-security/concepts/code-scanning/code-scanning): public repositories and eligible organization repositories with GitHub Code Security. The private/personal account classification and observed SARIF error are parent-supplied context; I did not independently query account settings. The condition intentionally excludes all private repositories. If this repository is later transferred to an eligible organization and scanning is enabled, the condition should be revisited. That future possibility does not block this current, explicitly authorized private-personal-repository fix.
- The diff is only ten added/two removed lines across these two workflows. No library code, API, dependencies, Go-version matrix, action-version updates, disclosure settings, or coverage publishing behavior changed. The existing untracked kitty-ops directory belongs to the parent and was not modified.

## Verification performed

Read AGENTS.md, docs/CONTEXT.md, /tmp/rrule-ci-implementation.md, both full workflows, and the complete diff. Reused the already resolved reviewer-renata profile and refreshed review-scoped governance via `SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty charter context --action review --json` (exit 0).

`git diff --check` passed with no output. An independent bounded parse (`timeout 20 ~/.local/share/uv/tools/spec-kitty-cli/bin/python`, existing PyYAML BaseLoader) exited 0. Parsed current/base structures confirm CI steps and strategy are identical; all other CodeQL data are identical after removing the new job condition. Expanded CI names are six distinct strings. No dependencies were installed. No production tests were rerun for this workflow-only diff.

Reviewed content hashes:
- .github/workflows/ci.yml: SHA-256 0b4d840c029687c1f620876a4c1e014e6e4d4dfc7dea2d94f984cdcacf3c7fe2
- .github/workflows/codeql.yml: SHA-256 bb5867e5f8df57efcea0a2ffba0de5e829193f882d2e8290ed68dcc14cbb5ceb

Hosted acceptance remains for the parent after push: observe all six matrix jobs on the new PR head and the explicit skipped CodeQL job. Existing older Go/action versions and current runner images require that real execution evidence; local YAML parsing cannot establish hosted success. I made no source, workflow, repository-state, or Op-state changes, and performed no commit, push, merge, or public posting.
