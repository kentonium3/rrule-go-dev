# PR 9 workflow configuration implementation

Implementer: codex / gpt-6.1-sol, 2026-10-25.
Independent reviewer requested: codex / gpt-6-astra.
Spec Kitty Op (opened and owned by parent): 01M5RYFFX07KTDKP0Z0E3AJZXK.
Workspace ~/repos/rrule-go, branch fix/iterator-panics.
Owned edits only: .github/workflows/codeql.yml and .github/workflows/ci.yml.
No commit, push, merge or Op closure performed by implementation agent.

Read AGENTS.md and docs/CONTEXT.md; reloaded implementer-ivan via resolver and implement action context. Applied spk-charter-profile-load skill, boring-code-review styleguide, efficient-local-tooling guidance. All Spec Kitty commands prefixed SPEC_KITTY_NO_UPGRADE_CHECK=1.

## Changes and rationale

CodeQL analyze has an explicit job-level `if: ${{ github.event.repository.private == false }}` with a comment explaining this private personal repository lacks code-scanning availability and private organization repositories require GitHub Code Security. For configured push, pull_request and schedule events the repository privacy flag is true here, so the job skips before matrix expansion or CodeQL action execution. This does not run a scan or represent security coverage; no continue-on-error or fabricated passing analysis was introduced. The existing triggers and action versions remain.

CI now runs for pull_request events targeting master, in addition to existing master pushes. The matrix is the same six combinations (Ubuntu/macOS/Windows, Go 1.18.x/1.19.x). Job names now include both Go version and OS so each check is distinguishable. Explicit workflow permissions contents:read grant checkout access while other token permissions default to none. Existing action versions, commands, Go versions, and library code remain unchanged. No Dependabot merge or dependency installation.

## Verification

`git diff --check`: exit 0, no output.

Parsed both workflow files with the existing Spec Kitty virtualenv Python and PyYAML 6.0.3, using yaml.BaseLoader to avoid YAML 1.1 interpreting the Actions `on` key as a boolean. Parsing exit 0. Read back events, permissions, job condition/name, matrix and action references from the parsed structures:

- CodeQL: push master, pull_request master, weekly schedule; analyze condition exactly `${{ github.event.repository.private == false }}`; language go; checkout v3 and CodeQL init/autobuild/analyze v2 unchanged.
- CI: push master and pull_request master; permissions contents:read; name `Test Go ${{ matrix.go-version }} on ${{ matrix.os }}`; unchanged three OS/two version matrix; setup-go v4 and checkout v3 unchanged.

System python lacks yaml and actionlint is not installed; used an already-installed parser without installing anything. This is local YAML parsing plus documented semantics review, not a claim of GitHub server validation or hosted run success. No implementation-mirroring tests were added, and Go tests were not repeated for workflow-only edits. Parent can verify hosted behavior after independent review/push.

Official docs consulted directly on 2026-10-25:

- https://docs.github.com/en/code-security/concepts/code-scanning/code-scanning : availability is public repositories or eligible organization repositories with Code Security. Parent's GitHub API evidence identifies this repository as private/personal and SARIF upload error as code scanning not enabled.
- https://docs.github.com/en/actions/reference/workflows-and-actions/workflow-syntax#jobsjob_idif : job conditions prevent execution and are evaluated before the matrix.
- https://docs.github.com/en/actions/reference/workflows-and-actions/workflow-syntax#onpull_requestpull_request_targetbranchesbranches-ignore : pull_request branch filters select the target branch.
- https://docs.github.com/en/actions/reference/workflows-and-actions/workflow-syntax#permissions : specifying a permission sets unspecified permissions to none.

`git status --short` after edits shows only the two owned tracked workflow changes and the parent's existing untracked kitty-ops directory. Ready for independent review; parent retains commit/push and Op close ownership.
