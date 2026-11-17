# WP02 independent review — APPROVE

Date: 2026-10-25.
Reviewer: codex/gpt-6-astra, resolved reviewer-renata profile.
Implementer: codex/gpt-6.1-sol. Distinct-model review requirement satisfied.
Reviewed commit: 3dfc31396a49e42a4be9aecfc2fe63dfcbddf0cf.
Workspace: ~/repos/rrule-go/.worktrees/iterator-panics-01M5RT9D-lane-b.
Scope: rrule_fuzz_test.go only (162 added lines). No source, planning, or workflow changes made by this reviewer. Parent owns transitions.

## Verdict and findings

APPROVE. No blocking findings or requested changes.

I applied Reviewer Renata's code-review mode: correctness, local scope, observable test behavior, and standards compliance; implementation and workflow ownership remain with their assigned agents. Read AGENTS.md, docs/CONTEXT.md, issue 01, runtime review prompt, mission spec/plan/contract and WP02 instructions. Resolved the profile using `SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty agent profile show reviewer-renata` (also JSON), and loaded `spec-kitty charter context --action review --json`. The requested code-review-checklist section is not authored in this charter; the runtime checklist, mission requirements, and four project directives were applied.

## Design checks

- Both real public entry paths run: NewRRule(option) and StrToRRule(option.String()), followed by Iterator and up to COUNT+1 Next calls for each accepted input. Serialization preserves all supplied fields; interval zero is omitted and normalizes to the same default. There is no panic recovery or helper-only stand-in. Errors are permitted for invalid boundary selections, with the reason documented at the return.
- Every BY field has explicit valid/invalid boundary seeds. Signed fields include both signs and zero; BYDAY includes monthly-impossible and outer validation boundaries. Easter and intervals include portable native MinInt/MaxInt values, and interval cases include zero, negative, ordinary, MaxInt-1, MaxInt and 2147483647. All seven frequencies are exercised. Lists are at most two elements, and at most 30 input bytes affect decoding.
- The near-horizon normalization preserves the three committed defect seeds. Independent original-code replays below demonstrated ordinal -6, Easter 300, and monthly native MaxInt failures through each constructor.
- Mathematical search bound: accepted calendar intervals normalize to a positive integer and start January 1, 9999. The remaining supported horizon has at most 365 daily periods, 53 weekly periods, 12 monthly periods or one yearly period. Empty calendar matches therefore exhaust within a small finite scan. The time expansion has at most eight elements; BYSETPOS has at most two selections. The default-filled lists cannot enlarge these limits.
- Subdaily starts at 9999-12-31 23:59:00 UTC. All calendar filters are cleared; HOURLY clears BYHOUR, MINUTELY also clears BYMINUTE, SECONDLY also clears BYSECOND. Constructor defaults do not reintroduce higher/equal frequency filters. Each inner advancement loop therefore breaks on its first accepted advance or exhausts beyond MAXYEAR. Only lower-order expansion lists remain, and any supplied BYSETPOS becomes reachable +1 or -1. Final-day horizon permits at most 60 second periods, one minute period, or one hour period. An expansion selected before DTSTART may yield nothing, but cannot extend that search. COUNT/UNTIL are correctly not claimed as termination guarantees.
- TestFuzzIterateControls uses the same decoder, horizon, construction and iteration harness and requires the exact first DTSTART for every frequency and both constructors. This detects blanket rejection/empty iteration. Existing approved WP01 tests supply the narrower same-fixture ordinal/Easter/interval controls and original-year regressions; this WP does not replace or weaken them.
- Property assertions cover count, nondecreasing order, DTSTART lower bound (initial previous), UNTIL and MAXYEAR. Nondecreasing rather than strict ordering is appropriate for the existing duplicate-list behavior. This remains bounded panic/property coverage, not a proof of arbitrary recurrence semantics or a fix for excluded hangs.
- Both go1.18 build constraints and the required blank line are present. go.mod remains 1.16; no newer convenience API is used. Native-width integer calculation is portable; the supplied 386 execution evidence passed. Actual historical Go 1.16/1.18 compilers were not run, and no such claim is made.

## Independent checks

1. `timeout 90 env GOCACHE=/tmp/rrule-go-cache go test -count=1 ./...` in lane-b: exit 0, 0.339s.
2. `timeout 60 env GOCACHE=/tmp/rrule-go-cache go test -run='^TestFuzzIterateControls$' -count=1 -v .` in lane-b: exit 0, 0.002s; all seven frequencies x two public paths passed.
3. Verified byte equality of /tmp/rrule-WP02-original/rrule.go against `git show 6177350^:rrule.go`, and of its fuzz file against the reviewed lane file. Independently ran all six commands `timeout 30 env GOCACHE=/tmp/rrule-go-cache go test -run '^FuzzIterate/seed#N/PATH$' -count=1 .` in that disposable original-code directory, with N=0,1,2 and PATH=options,text. Every run exited 1 with its expected actual panic, never timeout: ordinal index -5; Easter index 386 into length 372; interval index 9223372036854775807 into length 13. Full outputs: /tmp/rrule-WP02-review-red.log.
4. Inspected actual commit file list and clean lane status; inspected production construction/normalization/serialization and iterator loops to establish bounds, without editing them.

Inspected the saved WP02-gates.log, WP02-386.log, WP02-seeds.log handoff and WP02-fuzz.log. The completed campaign reports PASS, 60.084s, 3,437,785 executions, 188 new interesting inputs and 357 total. It reached a normal successful exit, not the timeout. Recorded ordinary/race/vet/build/format and 386 gates passed. A redundant 60-second campaign was unnecessary after independent seed reproduction and bound review.

## Required anti-pattern checklist

1. Dead code: N/A — no new production/public API; test helpers have live test callers, and testing discovers the fuzz/control entry points.
2. Synthetic-fixture test: PASS — all safety/property assertions invoke real public constructors and Iterator.
3. Silent empty return: PASS — constructor-error returns are explicitly justified as invalid boundary refusals; iterator exhaustion is expected and positive controls enforce nonempty execution.
4. FR coverage: PASS — FR-005 has committed defect-sensitive seeds through both public paths and independently verified original-code failures/current-code success. NFR-001/002/003 evidence also checked; SC refs are tracked rather than incorrectly treated as separate assertions.
5. Frozen surface: PASS — WP02 commit changes only its owned rrule_fuzz_test.go; no production, exported API, dependency, module minimum or license changes.
6. Locked decision: PASS — no blanket recover, interval cap, hang-family repair, UNTIL optimization, public posting or incompatible API change.
7. Shared-file ownership: PASS — rrule_fuzz_test.go is the sole WP02-owned addition; WP01 source/tests remain unmodified by this commit.
8. Production fragility: N/A — no production paths or new production panic/raise behavior introduced.

Remaining limits are explicit: excluded unreachable subdaily filter combinations, broader UNTIL scan behavior, retained Easter-257 spill/drop behavior, and no historical compiler execution. None is hidden or newly introduced by this test-only WP.
