# WP01 independent implementation review

**Verdict: APPROVE.** No blocking correctness, compatibility, scope, or verification findings.

- Reviewer: codex / **gpt-6-astra**, reviewer-renata, 2026-10-25.
- Implementer: codex / **gpt-6.1-sol**, 2026-10-25. This is a different-model review.
- Reviewed implementation commit: `6177350bd0be94a1750317574821e1e6b50025d9`.
- Workspace: `~/repos/rrule-go/.worktrees/iterator-panics-01M5RT9D-lane-a`.
- The later lane HEAD merge (`8808379`) has no changes to this commit's three owned files. The broader mission-base diff also contains inherited planning/runbook commits; those are not implementation scope expansion.
- Source/planning files and workflow state were not modified by this review. The orchestrator records the authoritative verdict transition and persists this artifact.

## Review basis

Read the full generated review prompt, AGENTS.md, docs/CONTEXT.md, issue 01, authoritative mission spec/plan, contracts/iterator-safety.md, WP01 handoff evidence, independent plan review, the full production/test/doc implementation diff, and surrounding iterator construction/generation/buffering code. Loaded spk-charter-profile-load and spk-run-review-wp. Resolved reviewer-renata via `SPEC_KITTY_NO_UPGRADE_CHECK=1 spec-kitty agent profile show reviewer-renata` and loaded action-scoped review context. The optional code-review-checklist query reported the section is not yet authored; existing project directives and explicit requirements governed the review.

WP01 has no dependencies. WP02 is the planned dependent in lane-b and owns the separate fuzz work; this approval does not approve WP02 or claim the mission fuzz gate is complete.

## Correctness findings

No actionable defect found.

1. Ordinals: both preliminary indices are bounded against the selected period before reading wdaymask. Existing post-adjustment bounds remain. With the existing ±53 input validation, preliminary integer arithmetic is itself bounded. The skip preserves impossible selections, mixed valid lists, YEARLY ±53 without BYMONTH, and WEEKLY ignored ordinal behavior.
2. Easter: offset comparisons occur before addition and cover machine MinInt/MaxInt without negating the supplied value. The mask length and weekly tail are unchanged. The 1818/1943 boundary fixtures actually exercise the early/late Easter extremes. The YEARLY 257 spill/drop remains explicitly documented and tested against a nonempty 1944 control; no correction is falsely claimed.
3. Calendar intervals: YEARLY and MONTHLY compare before adding; the zero-based monthly normalization preserves December alignment. WEEKLY derives the same nonnegative week-start offset as the former branch, then checks before multiplying by seven. DAILY checks before addition. Reaching the final supported date remains allowed; stepping beyond it exhausts.
4. Subdaily intervals: the helper compares before adding even on 64-bit, where a mere int64 cast would be insufficient. It converts day counts to int64 before unit multiplication. For the supported civil-year horizon, the bounded total is below approximately 3.2e11 seconds; day quotients fit 32-bit int. Temporary minute/second values after filtered-day shortcuts are included in the complete clock value. Every inner BY-filter iteration checks again, and daysRemaining includes carried days that have not yet been normalized. This prevents cumulative carry overflow. The 2147483647 SECONDLY and MINUTELY cases retain valid multi-occurrence behavior.
5. Civil time and buffers: UTC Unix-second differences measure calendar days without time.Duration saturation or local DST elapsed-time distortion. Local occurrence construction remains unchanged. Early exhaustion marks finished without clearing already-appended occurrences, and next continues draining the buffer.
6. Scope/API: production changes are local mask/advancement guards plus two private helpers. No constructor cap, panic recovery, public symbol/signature change, dependency, reachability solver, or general UNTIL scan optimization was introduced. LICENSE, README, go.mod and go.sum were not changed by the implementation commit.

## Requirement and test assessment

- FR-001: TestSafetyOrdinal exercises real construction and Iterator creation through options and RRULE text, covers every reported ordinal family, nearby ±5, mixed lists, valid yearly ±53, and weekly ordinals.
- FR-002: TestSafetyEaster covers unsafe and extreme offsets, actual boundary years, mixed ordinary/unsafe values, and exact valid controls.
- FR-003: TestSafetyInterval covers every frequency at native MaxInt and COUNT=1/3; original MONTHLY/DAILY/HOURLY cases force advancement. IntervalControls, InnerCarries, Buffered and SubdailyBuffered cover exact limits, accumulated carries, and pending emissions.
- FR-004: exact serialized expectations cover ordinary and large valid intervals, 32-bit-compatible values, WKST alignment, mixed masks, nearby valid ordinal/Easter values and DST wall-clock steps. Review independently compared 1,701 deterministic ordinary-rule cases against the pre-fix implementation, all equal.
- FR-005: safetyOccurrences invokes NewRRule or StrToRRule(option.String()), fails on constructor errors, constructs Iterator, compares every expected timestamp, and requests another occurrence to verify exhaustion. No recover swallows a panic. The inspected red logs show independent ordinal/Easter/monthly panics plus DAILY/HOURLY wrong output before their fixes. Precise scratch fix-removal logs reproduce failures through both constructors; the inner-horizon mutation times out and is correctly recorded as failure. Fuzz seeds/run remain WP02's assigned portion.
- FR-006: issue 01 and the explicit 1943/1944 tests honestly retain the known 257 limitation.
- NFR-002/NFR-003 and constraints: inspected baseline/toolchain, full test/race/vet/gofmt/build/386 evidence, bounded commands, distinct model records, safe commit ownership, and privacy/scope discipline. No unexplained gate failure is counted as passing.

## Independent verification

Toolchain independently confirmed: `go version go1.25.14 linux/amd64`.

In the lane workspace:

- `timeout 180 env GOCACHE=/tmp/rrule-go-cache go test ./...` — exit 0, cached pass.
- `timeout 180 env GOCACHE=/tmp/rrule-go-cache go test -count=1 ./...` — exit 0, `ok github.com/teambition/rrule-go 0.430s`.
- `git diff --exit-code 6177350..HEAD -- rrule.go rrule_safety_test.go docs/issues/01-iterator-panics.md` — exit 0, no output.
- `git status --short` — no output.

Additional disposable deterministic compatibility check in `/tmp/rrule-WP01-review-parity/{original,fixed}`:

- Original production source obtained from `git show 6177350^:rrule.go`; fixed production copied from the lane. No active checkout mutation.
- Identical public NewRRule/Iterator harness compares serialized occurrences for seven frequencies, nine ordinary intervals, three WKST values, three starts (spring/fall DST vicinity and final supported week), and UTC/America-New_York/Australia-Lord_Howe. Calendar time lists and weekly weekday selections exercise buffered output.
- In each scratch: `timeout 60 env GOCACHE=/tmp/rrule-go-cache go test -run '^TestReviewParity$' -count=1 -v .` — exit 0.
- Both produced `REVIEW_PARITY 1701 cf88f859b1e76297b2caeb39e5bdfbb3b03cffcdaf2657e12000d7ea8362ee58`.
- Harness and logs remain in that scratch directory. This is deterministic parity checking, not fuzzing.

Inspected actual `/tmp/rrule-WP01-red.log`, `/tmp/rrule-WP01-sensitivity.log`, `/tmp/rrule-WP01-gates.log` and durable implementation evidence. Full test, race, vet, clean gofmt, build, and 386 execution are reported passing. The original sandboxed 386 `bad system call` failure remains recorded separately from the successful same-command escalated run; it was not reclassified as a pass or bypassed through test changes. The Go 1.16 language-mode compile is accurately distinguished from testing with an actual Go 1.16 toolchain.

## Required anti-pattern checklist

1. **Dead code — PASS:** no new public function/module; daysRemaining is called by WEEKLY/DAILY/advanceClock, and advanceClock by all three subdaily branches (call sites inspected with rg).
2. **Synthetic-fixture test — PASS:** behavioral assertions exercise real constructors and iterator generation; expected timestamps are literal independent oracles, not mock implementation output.
3. **Silent empty return — PASS:** out-of-period mask skips and horizon exhaustion are the explicit behavior contract, and valid/mixed/buffered controls distinguish them from unconditional empty output.
4. **FR coverage — PASS:** FR-001 through FR-006 mapped above; WP02's fuzz responsibility remains explicit.
5. **Frozen surface — PASS:** implementation commit changes only its three owned paths; protected LICENSE/public API/dependency surfaces remain unchanged.
6. **Locked decision — PASS:** no arbitrary caps, recover, global ±5 rejection, unrelated recurrence redesign or unbounded new solver.
7. **Shared-file ownership — N/A:** WP01 exclusively owns rrule.go, rrule_safety_test.go and issue 01; dependent WP02 owns its separate fuzz file.
8. **Production fragility — PASS:** no new production panic/raise or fail-loud path; exhaustion is controlled and preserves the buffer.

## Remaining limits

Known YEARLY BYEASTER=257 spill/drop, general never-matching UNTIL scanning, and separate modular subdaily reachability remain outside this repair. Finite COUNT is not represented as bounding one iterator call. No fuzz run was performed or claimed by this reviewer; WP02 must still deliver its separately reviewed target and at least 60-second run. This is WP01 approval, not mission acceptance or authorization to merge master/publish.
