# WP02 implementation and validation handoff

Date 2026-10-25. Implementer: codex/gpt-6.1-sol (wp02_implement); authored
rrule_fuzz_test.go in runtime lane-b. Its turn was twice interrupted by an
automated content-risk flag after the code and passing seed run were saved.
The root orchestrator inspected that work and completed local verification and
commit without editing its source. No disabled safeguard, public target, or
external test was involved. Independent reviewer: gpt-6-astra, pending.

Commit 3dfc313, branch kitty/mission-iterator-panics-01M5RT9D-lane-b, contains
only rrule_fuzz_test.go (162 lines); actual git show and clean status verified.
Dependency WP01 approved at 6177350. Production code unchanged by WP02.

## Domain and compatibility

Both go1.18 build constraints keep go.mod 1.16 intact. Current verification
compiler is Go 1.25.14; actual historical Go 1.16 was not run.

FuzzIterate decodes frequency, interval, and at most 30 data bytes into small
option lists, including invalid scalar bounds. Calendar frequencies start
January 9999, bounding calendar iteration by the last supported year. Each BY
list has <=2 elements; calendar time cross products <=8. Subdaily starts in
the final minute of 9999, removes calendar and higher/equal-order time filters,
and restricts BYSETPOS to reachable +/-1. Lower-order expansions remain covered.
These restrictions intentionally exclude known separate unreachable-BY-list
hangs. COUNT is an output limit, not asserted as a bound on Next.

Every accepted option is evaluated through NewRRule and StrToRRule, followed
by Iterator, with no panic recovery. Assertions check bounded output count,
nondecreasing dates, DTSTART/UNTIL and MAXYEAR. Same-decoder positive controls
require a real first DTSTART for every frequency and both constructors.
178 committed f.Add calls (169 unique coverage seeds) include independent
ordinal, Easter and monthly interval cases, every explicit scalar boundary,
and ordinary nonempty controls. The original early-year regressions remain in
WP01, so the near-horizon fuzz normalization does not replace them.

## Original-version seed sensitivity

Scratch /tmp/rrule-WP02-original copied module sources, this exact fuzz test,
and original rrule.go from git show 6177350^:rrule.go. Active lane unchanged.
For each seed#0,seed#1,seed#2 and each options,text constructor:
`timeout 30 go test -run '^FuzzIterate/seed#N/PATH$' -count=1 .`
with GOCACHE=/tmp/rrule-go-cache exited 1 with an actual panic. All six runs
completed without timeout. Full output WP02-red.log. Seed 0 is ordinal -6;
seed 1 Easter 300; seed 2 native MaxInt monthly interval. These are committed
regression inputs, not only transient random mutations.

## Passing gates

All commands executed in lane-b, GOCACHE=/tmp/rrule-go-cache:
- `timeout 180 go test -count=1 ./...`: exit 0, 0.346s.
- `timeout 180 go test -race ./...`: exit 0, 4.666s.
- `timeout 180 go vet ./...`: exit 0, no output.
- `timeout 180 go build ./...`: exit 0, no output.
- `timeout 180 gofmt -l .`: exit 0, no output.
- `timeout 180 env GOARCH=386 go test ./...`: exit 0, executed with the already
  diagnosed required host permission for 32-bit binaries; output WP02-386.log.
- `timeout 150 go test -run='^$' -fuzz='^FuzzIterate$' -fuzztime=60s -parallel=2 .`:
  exit 0, 60.084s, 3,437,785 executions, 188 new interesting inputs, 357 total.
  No observed panic or property failure. Actual output WP02-fuzz.log.
- Agent's initial deterministic seed run also passed; output WP02-seeds.log.

The bounded command timeout was not reached. No timeout or absent hook is
counted as passing. The separately known Easter 257 spill/drop, modular filter
hang, and general UNTIL-scan limits remain. No master merge/public posting.
