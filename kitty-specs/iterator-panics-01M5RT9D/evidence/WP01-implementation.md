# WP01 implementation and independent review handoff

- Implementer: codex / gpt-6.1-sol (required separate implementation model).
- Profile: implementer-ivan, loaded through `SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty agent profile show implementer-ivan`; action context loaded through `spec-kitty charter context --action implement --json`.
- Date: 2026-10-25; verification identity go1.25.14 linux/amd64.
- Workspace: ~/repos/rrule-go/.worktrees/iterator-panics-01M5RT9D-lane-a.
- Branch: kitty/mission-iterator-panics-01M5RT9D-lane-a.
- Commit: 6177350bd0be94a1750317574821e1e6b50025d9 (`fix(WP01): guard iterator masks and calendar advancement`).
- Safe-commit exit 0, JSON `committed: true`; actual `git show`, `git diff-tree`, committed test content and clean `git status --short` verified. Exactly rrule.go, rrule_safety_test.go and docs/issues/01-iterator-panics.md committed. No workflow transitions, planning edits, dependency/API changes, master merge or public/upstream actions.
- Independent reviewer required: gpt-6-astra. No implementing-model review is represented as approval.

## Reading / governance

Read full runtime prompt, AGENTS.md, docs/CONTEXT.md, issue 01, mission spec/plan, behavior contract and independent plan review. Loaded ~/.agents/skills/spk-charter-profile-load/SKILL.md and resolver profile/action context. Loaded boring-code-review and efficient-local-tooling guidance before production changes. Queried review-checklist and terminology-canon sections; CLI reports those sections are not yet authored. Existing project directives, spec and plan applied; no charter edits made.

## Baseline

Before adding tests/production changes:

`go version` => go version go1.25.14 linux/amd64 (exit 0).

`timeout 180 env GOCACHE=/tmp/rrule-go-cache go test ./...` => exit 0, `ok github.com/teambition/rrule-go 0.044s`.

Original rrule.go preserved at /tmp/rrule-WP01-baseline.go.

## Red before green

Full actual outputs and exact commands: /tmp/rrule-WP01-red.log. Each command was run on unfixed production before its fix. No panic recovery in regression code. Every command below exited 1; none timed out:

- `timeout 30 env GOCACHE=/tmp/rrule-go-cache go test -run ^TestSafetyOrdinal/monthly_negative6/options$ -count=1 .` => index -5 panic in rebuild during Iterator construction.
- `timeout 30 env GOCACHE=/tmp/rrule-go-cache go test -run ^TestSafetyEaster/-100/options$ -count=1 .` => index -6 panic in rebuild during Iterator construction.
- `timeout 30 env GOCACHE=/tmp/rrule-go-cache go test -run ^TestSafetyInterval/original_MONTHLY/options$ -count=1 .` => index 9223372036854775807 panic during advancement.
- `timeout 30 env GOCACHE=/tmp/rrule-go-cache go test -run ^TestSafetyInterval/original_DAILY/options$ -count=1 .` => unexpected 2026-12-31 occurrence after DTSTART (overflow artifact; later dates in original reproduction move backwards).
- `timeout 30 env GOCACHE=/tmp/rrule-go-cache go test -run ^TestSafetyInterval/original_HOURLY/options$ -count=1 .` => unexpected 2026-01-26T01:00:00Z after DTSTART.

Ordinal/Easter guards were applied first. `timeout 60 env GOCACHE=/tmp/rrule-go-cache go test -run 'TestSafety(Ordinal|Easter)$' -count=1 .` exited 0 (0.248s) before interval production changes. Test expectation development exposed two fixture assumptions: -100 can correctly match later years, so the empty fixture now has UNTIL 2027; 1944 Easter+257 is December 22, not December 23. These were corrected, with mixed -100 retaining ordinary Easter dates. No production workaround was added for test expectations.

After interval fix and completed controls: `timeout 90 env GOCACHE=/tmp/rrule-go-cache go test -run TestSafety -count=1 .` => exit 0, 0.227s. Final full gate includes the later buffer/alignment additions.

## Production changes / arithmetic reasoning

- Ordinal preliminary candidate must lie in [first,last] before either weekday-mask read; retain final check after weekday adjustment. Preserve impossible empty and mixed rules, YEARLY +/-53 without BYMONTH, WEEKLY ordinal semantics.
- Easter compares offset against -eyday and len(mask)-eyday before addition. This is safe for machine MinInt/MaxInt and preserves existing mask tail. Skip only the current year's out-of-mask candidate.
- YEARLY compares interval against MAXYEAR-currentYear before adding.
- MONTHLY compares against remaining calendar months before adding; zero-based month normalization preserves alignment.
- WEEKLY compares interval against (remainingDays+WKST-offset)/7 before multiplying; preserves alignment and current buffer.
- DAILY bounds interval by remaining days before adding.
- daysRemaining uses UTC civil-month-start/Jan-10000 Unix seconds divided by 86400, then subtracts current day, including any accumulated unnormalized carry. No local DST elapsed-time arithmetic or time.Duration horizon.
- advanceClock checks interval against remaining civil units BEFORE addition on EACH inner loop pass. It includes temporarily nonnormalized minute/second from filtered-day shortcuts. Once compared, total is bounded by <~3.2e11 seconds; quotient/day casts are <=~3.7e6, safe in 32-bit int. Direct int64 cast of MaxInt is never treated as sufficient protection.
- Exhaustion leaves the already-buffered current occurrences intact.

The helper consolidates the same arithmetic boundary across three existing failure paths; no generic framework or reachability solver was added.

## Coverage / compatibility

Both NewRRule(ROption) and equivalent StrToRRule(option.String()) enter real Iterator and advance to exact expected occurrences/exhaustion, including a final Next beyond the expected list. Constructor failures are not accepted as success. Tests include:

- MONTHLY -6MO/8MO/53MO; YEARLY Jan -6MO and Dec 8MO; monthly +/-5; mixed valid/impossible; YEARLY +/-53 without BYMONTH; WEEKLY ignored ordinal.
- Easter -100/300, machine MinInt/MaxInt, mixed {0,extreme}; -81/-80 with early Easter 1818; 258/257 with late Easter 1943; valid 257 in leap year 1944; ordinary offset -1 and -80 in 2026.
- All seven frequencies with native MaxInt and COUNT=1/3 near year 9999; original 64-bit MONTHLY/DAILY/HOURLY reproductions at 2026 start.
- Exact ordinary and large valid intervals; YEARLY 7973 reaches 9999 and MONTHLY 95687 reaches December 9999.
- Valid 2147483647 SECONDLY three dates and MINUTELY two dates followed by horizon exhaustion under COUNT=3, on amd64 and 386. Expected dates independently verified using Python datetime/timedelta, recorded in /tmp/rrule-WP01-expected-dates.log: SECONDLY 2094-01-23T13:14:07Z, 2162-02-11T16:28:14Z; MINUTELY 6109-01-28T12:07:00Z.
- All frequencies near final supported instant; zero interval normalization; existing negative errors and 32-bit text integer overflow rejection.
- Buffered calendar and subdaily occurrences survive exhaustion; explicit nondefault WKST alignment.
- Multiple subdaily BY-filter carries and loops crossing horizon; filtered-day shortcuts; New York DST wall-clock steps for DAILY/HOURLY/MINUTELY/SECONDLY.

Independent baseline compatibility scratch run: /tmp/rrule-WP01-compatibility.log. Original production with current tests passed the ordinary/large interval controls, filtered-day controls and DST exact sequences. Command `timeout 60 env GOCACHE=/tmp/rrule-go-cache go test -run 'TestSafetyWallClock|TestSafetyIntervalControls/(YEARLY|MONTHLY|WEEKLY|DAILY|HOURLY|MINUTELY|SECONDLY|default)|TestSafetyInnerCarries/filtered' -count=1 .` => exit 0 (0.008s).

## Independent fix sensitivity

Full commands, scratch locations and actual outputs: /tmp/rrule-WP01-sensitivity.log. Active lane was never reverted. Each scratch copied module sources/current tests and removed only the specified production fix:

- Ordinal guards alone removed: focused options and text regressions both exit 1 with panic.
- Easter guard alone removed: focused options and text regressions both exit 1 with panic.
- Interval advancement alone restored from original: focused MONTHLY options and text regressions both exit 1 with panic; ordinal/Easter fixes remain.
- Only advanceClock horizon comparison disabled: `timeout 5 env GOCACHE=/tmp/rrule-go-cache go test -run ^TestSafetyInnerCarries/horizon_HOURLY/options$ -count=1 .` exits 124. This is a deliberately failing mutation, not a passing verification or claim of separate hang repair. It proves the inner-loop comparison matters before accumulated day carry can be normalized.

## Final gates

Actual outputs /tmp/rrule-WP01-gates.log. All final gates pass:

- `timeout 180 env GOCACHE=/tmp/rrule-go-cache go test ./...` => exit 0, ok 0.359s.
- `timeout 180 env GOCACHE=/tmp/rrule-go-cache go test -race ./...` => exit 0, ok 4.373s.
- `timeout 180 env GOCACHE=/tmp/rrule-go-cache go vet ./...` => exit 0, no output.
- `gofmt -l .` => exit 0, no output.
- `timeout 180 env GOCACHE=/tmp/rrule-go-cache GOARCH=386 go test ./...` => final exit 0, ok 0.415s, executed with require_escalated.
- `timeout 180 env GOCACHE=/tmp/rrule-go-cache go build ./...` => exit 0, no output.
- `timeout 180 env GOCACHE=/tmp/rrule-go-cache go build -gcflags=-lang=go1.16 ./...` => exit 0, no output. This is a current-compiler language-mode compatibility check, not a claim that an actual Go 1.16 toolchain was run. Production adds no APIs newer than Go 1.16.
- `git diff --check` => exit 0, no output.

### Diagnosed environmental failure (D-3)

Initial GOARCH=386 gate inside sandbox exited 1 with `signal: bad system call (core dumped)` after compilation. Preserved exact output in gates log, reported parent promptly before mutation. Same bounded command requested through normal require_escalated tool path passed (0.415s), establishing sandbox syscall filtering as the failure cause. No tests, sandbox settings, Spec Kitty installation or workflow state were altered to bypass it. Both outcomes retained; timeout/failure never counted as pass.

## Explicit limits and review focus

- YEARLY BYEASTER=257 spill/drop retained: Easter April 25 1943 +257 is Jan 7 1944 in mask tail but outside the 1943 YEARLY dayset. A 1943-until fixture is empty; 1944 April 9 +257 correctly yields Dec 22 1944. Documented in issue 01; no cross-year Easter redesign claimed.
- Never-matching UNTIL scanning and separate modular BY-list reachability are outside scope. No reachability solver introduced. Finite output count alone is not claimed to bound Next; actual shell commands use timeouts.
- Fuzz target/run belongs to WP02 and is not claimed complete in WP01.
- Reviewer should inspect pre-add/multiply bounds, int width, temporary minute/second normalization, accumulated inner day carries, weekly alignment/buffering, exact valid 32-bit advancement and actual public entrypoint coverage.
