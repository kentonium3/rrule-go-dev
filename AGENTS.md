# AGENTS.md: patched rrule-go

This is Kent Gale's development copy of [teambition/rrule-go](https://github.com/teambition/rrule-go) (MIT). Its first missions fixed bugs that let **accepted** RRULE input crash or hang the iterator. Those fixes are now published as the maintained fork `kentonium3/rrule-go` v1.8.3, and this repo is where further fixes are developed. Upstream looks dormant: the last change to Go code was 2023-01-13.

Read this file, then `docs/CONTEXT.md`, then the issue you are working on in `docs/issues/`.

## Status

- Upstream base: teambition `master` at `e74d163` (2023-04-01). `rrule.go` is byte-identical to tag `v1.8.2` (blob `4a49f3e8`).
- **Updated 2026-10-25.** The upstream author confirmed teambition/rrule-go is no longer maintained, and the embargo ended. Kent published the fixes from missions 1 to 3 as the **public** repo `kentonium3/rrule-go`: tag `v1.8.3` on branch `main` (named `release` until 2026-10-25; protected: PRs and green CI only, no bypass), with the same library code as this repo's `main` (then named `master`) at `f0711ce`. Vikunja PR #3071 depends on it. teambition/rrule-go#72 points users at it.
- **Updated 2026-11-17. This repo is public** (`kentonium3/rrule-go-dev`), republished as a single-commit snapshot. It carries `kitty-specs/`, `.kittify/`, `docs/` and this file, which the library repo does not; the library repo gets code-only patch series (see `docs/runbooks/teambition-backlog.md`). The history, issues and PRs from before 2026-11-17 are in the private archive `kentonium3/rrule-go-dev-archive`.
- Remotes: `origin` is `kentonium3/rrule-go-dev` (public). `upstream` is teambition, fetch-only; its push URL is `no_push`. There is no remote for the library repo here, on purpose.
- Tracking: **the library repo `kentonium3/rrule-go` is the only live issue tracker; issues are turned off here.** Numbers below that name this repo refer to the private archive `kentonium3/rrule-go-dev-archive`. Historically there were **two issue trackers, and the numbers collide.** This repo's #4, #5 and #6 mirror `docs/issues/01-03` (done). The public repo's #5 and #6 are current mission work (`docs/issues/05-06`); #7 is inherited contributor baseline76c646d. Always write the repo name with the number, and never put a closing keyword (`Fixes #5`) in a commit here.

## Hard rules

1. **Never push to `upstream` or the public `kentonium3/rrule-go`, open or comment on a PR or issue there or on teambition/rrule-go, or post anywhere public** without Kent's sign-off on the exact copy.
2. **This repo is public: keep private material out.** Never commit content from private systems or other private repositories: session or chat IDs, coordination-bus traffic, internal issue registers, other projects' private branch or issue references, or absolute home-directory paths. Issues are turned off here; send issues and contributions to `kentonium3/rrule-go`.
3. **Keep the exported API source-compatible:** existing types, functions, methods and their signatures. **RFC 5545 is normative in all cases** (Kent, 2026-10-25). Behaviour may change only where it breaks the RFC today, or where input panics or hangs, as an issue brief's acceptance defines. Any other rule must give byte-identical occurrences and `String()` output, proven by a golden comparison against v1.8.3 plus inherited contributor fix76c646d. Kent explicitly approved private ROption metadata and targetv1.9.0: existing signatures/keyed callers remain, but external positional ROption literals break and require keyed migration. Adding an exported identifier needs a plan justification and means a minor release.
4. **Keep `LICENSE`** (MIT, Teambition). Add a short NOTICE in the README saying this is a modified copy.
5. **Design guardrails (Kent's standing rule):**
   - Fix realistic failure modes. Do not build mechanisms for exotic cases.
   - Every finding states its chain of conditions and how likely it is.
   - Anything that can crash or hang a service on accepted input is in scope. Rare cosmetic oddities are known limits.
   - Prefer the smallest fix at the root cause over new machinery.
6. **No silent workarounds of spec-kitty faults.** Normally stop, capture the exact command/output, and tell Kent. For the autonomous iterator-panics run authorized 2026-10-25, D-3 applies: preserve evidence before mutation, diagnose, validate and record a safe workaround, then continue. See `docs/runbooks/autonomous-iterator-panics.md`. Never bypass substantive gates or publish upstream.
7. **Implementer and reviewer must be different models.** Record which model implemented and which reviewed each work package, with the date. A review by the implementing model is a degraded check; say so in the review record.

## Setup: kittify, then run missions

The machine's spec-kitty is `spec-kitty/spec-kitty@1a0c8b467` (uv tool install from git). Identify the build by SHA, not by version string. Prefix spec-kitty commands with `SPEC_KITTY_NO_UPGRADE_CHECK=1`.

```bash
cd ~/repos/rrule-go
git switch -c kittify
SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty init --ai codex --non-interactive
# Charter: run the charter interview and keep it small: a Go library, a test-first
# discipline, no new dependencies, API compatibility.
```

Commit the kittified scaffolding on `kittify`, merge it to `main` when Kent agrees, then create missions with `$spec-kitty-specify`. The arc is specify → plan → tasks → analyze → implement/review → accept → consolidate.

### Installed-build configuration notes

For build `1a0c8b467`, `.kittify/charter/activation.yaml` owns activation lists;
putting them in `charter.yaml` does not configure the active ledger. Explicitly
set `languages: [go]` in `.kittify/config.yaml`; the interview language recognizer
does not infer Go from prose. Project rules use `directives: {directives: [...]}`
in the schema-version `2.0.0` charter. Validate the charter model and runtime rule
loading as well as CLI generation: generation alone does not prove valid policy.
The acceptance-test-first tactic references Python/TypeScript mutation toolguides;
these are correctly scope-filtered for Go and do not require adding those tools.


### Suggested missions, in order

1. **Iterator panics** (`docs/issues/01-iterator-panics.md`, issue #4). Three panics on input that `StrToRRule`/`NewRRule` accept: a BYDAY ordinal outside the month, a BYEASTER offset outside the year mask, and INTERVAL overflow. Include a fuzz test over the BY* domains and INTERVAL.
2. **Unreachable sub-daily BY-list hang** (`docs/issues/02-unreachable-byset-hang.md`, issue #5). HOURLY, MINUTELY and SECONDLY rules whose BYHOUR/BYMINUTE/BYSECOND can't be reached at the INTERVAL loop forever, even with UNTIL set.
3. *(Optional, Kent decides)* **UNTIL does not stop a never-matching scan** (`docs/issues/03-until-scan.md`, issue #6). This is a performance limit, not a crash.
4. **Adoption in Vikunja** (`docs/issues/04-vikunja-adoption.md`). Done 2026-10-25: Vikunja PR #3071 uses v1.8.3 through a `replace`.
5. **Teambition backlog** (`docs/runbooks/teambition-backlog.md`). **Current.** One mission covering public issues kentonium3/rrule-go#5 (sub-daily DST, P1) and #6 (floating UNTIL, P2), one WP each. Contributor#7 fix76c646d is inherited baseline, not assessed or extended. Kent approved private ROption parse metadata with external positional-literal break and targetv1.9.0; keyed callers/signatures remain. Briefs: `docs/issues/05-07`. Teambition PRs #66 and #68 are prior art with probed defects; reuse ideas, don't trust them.

## Working conventions

- **Changelog.** User-facing changes include a consumer-focused entry in root
  `CHANGELOG.md` under Unreleased, or explain why none is needed. Use public issue
  links and migration instructions; keep private mission IDs/evidence out of the
  candidate public notes. Run `go run ./tools/changelog check`. The dev → public
  handoff and release reconciliation are in
  `docs/runbooks/changelog-flow.md`. Private implementation is not proof of public
  availability; public tags remain the released-version authority.
- **Tests.** Plain Go: `go test ./...`, `go test -race ./...`, `go vet ./...`, `gofmt -l .`. Fuzz with `go test -run=^$ -fuzz=FuzzX -fuzztime=60s`. A new test must fail on the unfixed code first (red, then green).
- **Reproductions.** `docs/repro/*/main.go` are `//go:build ignore` programs that import this module, so they run against the local code: `go run ./docs/repro/01-iterator-panics/main.go`. `out-v1.8.2.txt` beside each shows the unfixed output. `02-unreachable-byset-hang` times each case out at 10 s.
- **Go version.** `go.mod` says `go 1.16`. Kent authorizes a minimum-version increase if needed for fuzzing; choose and justify the version during mission planning. The machine has Go 1.27 through mise.
- **Commits.** Conventional Commits, scoped `git add`, never force-push, never `git stash`.


<!-- spec-kitty:orientation -->
**Spec Kitty v4.0.0rc9** - project: rrule-go-dev (healthy)

Two usage patterns:
- **Full mission** (spec -> plan -> tasks -> implement -> review -> merge):
  trigger: "spec out", "create a mission", "write a spec", "plan this"
  -> run `/spec-kitty.specify` (slash-command agents), `$spec-kitty-specify` (codex), `/skill:spec-kitty-specify` (pi) or `/spec-kitty-specify` (antigravity, letta, vibe)
- **Lightweight dispatch** (ad-hoc fix, question, or advice - no mission created):
  trigger: "hey spec kitty", "use spec kitty to", "spec kitty <anything>"
  -> **ALWAYS run `spec-kitty dispatch "<request verbatim>"` - do NOT answer directly.**
  If you know the right profile, pass it to skip routing:
  `spec-kitty dispatch "<request verbatim>" --profile <profile-id>`
  Reason: `spec-kitty dispatch` loads governance context, routes the request,
  and opens the Op. Skipping it produces ungoverned, untracked responses.
  After finishing the work, close the Op with the command printed in the capsule
  (`spec-kitty profile-invocation complete --invocation-id <id> --outcome <done|failed|abandoned>`).
<!-- /spec-kitty:orientation -->
