# WP01 independent implementation review

Verdict: APPROVE. No blocking or actionable correctness findings.

Date: 2026-10-25.
Reviewer model: gpt-6-astra.
Implementation model: gpt-6.1-sol.
Reviewed implementation commit: 78f4ed596d28a3090819029e3550f96237a06deb.
Workspace: ~/repos/rrule-go/.worktrees/unreachable-byset-hang-01M5RZ90-lane-a.
Scope: rrule.go, rrule_reachability_test.go and docs/issues/02-unreachable-byset-hang.md; mission spec/plan/tasks, generated WP review prompt, implementation handoff and evidence logs. This report is a review verdict only; the orchestrator owns lifecycle transitions.

## Correctness and compatibility

The implementation adds only three local counters and exhaustion branches. At attempts 0 through M-1 it retains the original advanceClock call, assignments and existing predicate. Only the following iteration at attempts == M exhausts. Thus a match on candidate M succeeds. Each normalized clock coordinate advances by interval modulo M; its period divides M. No reachable predicate can be missed by the bound. Applicable combined predicates, including interval90 joint-impossibility fixtures, are unchanged.

Filtered-day shortcuts retain their existing multiple-of-interval phase and advanceClock normalization; no raw clock-field cycle sentinel is used. Existing horizon arithmetic, calendar carry, DST semantics and overflow checks are untouched. Constructor and setter wiring is unchanged and no cached reachability state is introduced.

The exhaustion branch assigns r.len = iterator.total and finished=true without clearing remain. next still drains the existing buffer after generate returns. COUNT and UNTIL paths are unchanged. A phase that just emitted a matching clock state remains reachable on its orbit, so horizon tests are the practical coverage of buffer retention on termination; prior TestSafety buffer/huge-interval regressions remain and pass.

I independently tested every new exact reachable control against pre-fix rrule.go in a disposable copy. All passed with the identical expected timestamps. This includes full-cycle final candidates at all three frequencies, filtered days, mixed lists, missing filters, non-UTC fixed offset, midnight/year carry and buffered UNTIL results. Existing safety tests additionally exercise actual DST and large/32-bit intervals. Scope and exported API are preserved; no dependency, minimum Go version, license or notice changes.

## Independent checks

- Verified implementation commit has exactly the three owned files (231 additions, 3 deletions).
- Workspace HEAD had advanced to lifecycle commit 86f9eead2c9c84e867de444eb9031a9d191d1a29. Diff of all three implementation paths against the requested fixed commit was empty; working tree was clean.
- GOCACHE=/tmp/rrule-go-cache go test ./... passed (cached).
- GOCACHE=/tmp/rrule-go-cache go test -count=1 -run 'TestClockReachability|TestSafety' passed, 0.622s; /tmp/rrule-hang-review-targeted.log.
- Disposable /tmp/rrule-hang-review-baseline was exported from the reviewed commit with only rrule.go replaced from its parent. GOCACHE=/tmp/rrule-go-cache go test -count=1 -run 'TestClockReachability/(hour-control|minute-control|second-control|mixed|local|carry|filtered-hour|filtered-minute|filtered-second|last-hour|last-minute|last-second|buffered|horizon)/' passed, 0.186s; /tmp/rrule-hang-review-compatibility.log. This checks exact old/new compatibility without modifying a tracked checkout.
- Inspected implementation evidence for full tests, full race (132.887s), targeted final joint-fixture race (13.288s), vet, build, formatting and linux/386. No reason arose to repeat the full race suite.

## Regression evidence and limitation

The three frequency tests exercise real public constructors and All/After/Iterator. Setters test reachable-to-unreachable-to-reachable behavior and Until rebuilds through both constructor routes. Repeated exhaustion is asserted for Iterator. The option route derives options using the parser but invokes NewRRule itself before evaluation.

The child process is directly invoked with exec.CommandContext and CombinedOutput. On watchdog expiry CommandContext kills the child and CombinedOutput waits/reaps it. No spinning goroutine is left in the parent. Operation timing is measured inside the child against one second; the separate four-second parent watchdog accommodates startup and race teardown.

The patched pre-fix hourly reproduction passes the one-second threshold on this machine (475.485776ms), and its ordinary guard-removal mutation also passes (496.943351ms). This is an explicitly documented limitation of the ordinary timing test's sensitivity, not a claimed red result. Minutely and secondly are red on patched baseline and with their guards individually removed. Historical upstream hourly is red at the watchdog, and hourly guard removal under the required race configuration fails the unchanged one-second operation threshold at 3.708002073s. I inspected these logs. This supplies credible frequency-specific evidence without tightening the agreed threshold or aggregating repeated operations to manufacture failure. No deterministic production instrumentation is necessary for this small, directly provable bound.

## Requirements and anti-pattern checklist

FR-001 PASS: each clock family, joint predicates, explicit UNTIL, both constructors and public queries covered.
FR-002 PASS: both-constructor setter transitions and exact restored occurrences covered.
FR-003 PASS: unchanged advancement/predicates plus independent pre-fix exact-control comparison and prior safety tests.
NFR-001 PASS: reported fixtures measured inside child against one second; fixed operation results are well below threshold.
NFR-002 PASS: isolated kill-and-wait harness and retained failure logs; hourly evidence qualification stated above.
C-001 through C-004 PASS within WP scope: API and prior guards intact, no broad calendar optimization, independent models and checks, no publication/master merge performed by reviewer.

Required eight-item review checklist:
1. Dead code: N/A — no new public function/class/module; production additions are live iterator branches.
2. Synthetic-fixture test: PASS — public recurrence paths execute; mutations demonstrate sensitivity, with hourly timing limitation disclosed.
3. Silent empty return: PASS — new exhaustion has the explicit mathematical unreachability reason, records length and preserves buffers.
4. FR coverage: PASS — behavioral assertions for FR-001/002/003 described above.
5. Frozen surface: PASS — only the three owned files changed; exported surfaces untouched.
6. Locked decision: PASS — local complete-cycle counters implement confirmed plan; constructor acceptance, setters and unrelated calendar behavior retained.
7. Shared-file ownership: N/A — sole WP/production owner; no competing source edits.
8. Production fragility: PASS — no new panic/raise/error path; normal iterator exhaustion only.

## Governance notes

The generated review prompt embeds an implementer profile and unknown-model identity. Actual review identity is explicitly recorded above, as directed by the orchestrator. The logged charter checklist query returns a notice that its section is not authored, rather than a command failure; this was already reported in the handoff. I made no runtime/lifecycle calls, source edits, commits, pushes or public posts, and encountered no new Spec Kitty fault.
