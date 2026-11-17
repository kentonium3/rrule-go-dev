# WP01 cycle1 rework handoff — R1 anchor collisions before BYSETPOS

**Completed code commit:** `ef373a5872dba30063a6a706275a28e46508af77` (`fix(WP01): deduplicate gap anchors before positional ranking`). Parent reviewed implementation: `af862491333c8a06fdb30bb60a58619e65e9baa6`. Implementation model **gpt-6.1-sol**, actual work date **2026-10-25**. Author/committer Kent Gale, 18:14:19 -0400. Required independent re-review remains root-owned, by **gpt-6-astra**; this implementer claims no approval.

Allocated lane `~/repos/rrule-go/.worktrees/teambition-backlog-01M5TEBG-lane-a`, branch `kitty/mission-teambition-backlog-01M5TEBG-lane-a`. Safe-commit succeeded once for six explicit files; git log verified the commit and git status was clean. No mission status/evidence edits, sibling implementation, push, PR, merge, public action, installation/security change or new agent. Root owns all transitions and durable mission evidence.

New exact command/output/exit logs: **`/tmp/rrule-WP01-rework1-logs`**. Original implementation context and full initial evidence remain at `/tmp/rrule-backlog-WP01-handoff.md` and `/tmp/rrule-WP01-logs`.

## Review finding and scoped correction

Read regenerated cycle1 fix prompt, full `reviews/WP01-review.md`, and both durable `evidence/WP01-review-rank-collision*` files before editing. R1 is valid: an explicit gap anchor and its normal resolved candidate both consumed positional ranks; deduplication happened only after BYSETPOS selection. Positive `1,2` lost the second distinct gap-day occurrence, while positive `1,3` invented a third position. The signed `1,-1` and negative `-2,-1` controls already worked in af86249 and remain controls, not newly claimed defects.

The final correction keeps existing sparse rank holes. Before building positional counts, `rIterator.anchorCollision` checks whether the original nominal tuple differs from the resolved start; ordinary/UTC starts return immediately. For a gap it binary-searches the already-sorted timeset for the resolved clock, checks that the resolved civil date is in the **filtered current period dayset**, and verifies that normal candidate's exact instant. If a normal slot already represents the anchor, `dayRanks` leaves its nominal gap slot removed and uses that normal slot once, at its correct physical position. If absent or filtered out, the explicit gap anchor retains its nominal slot as before.

The check is period-local: later DAILY periods keep their ordinary BYSETPOS ranks, while the existing global emission boundary suppresses repeated output instants across periods. It also handles a verified skipped date resolving onto another selected day in a coarse period, without enumerating timestamps. No UNTIL algorithm, parser forms, metadata layout, public signature, contributor selector normalization or dependency changed.

Complexity remains one clock binary search plus one selected-day scan, and at most one normal candidate resolution for gap anchors. There is no annual day×timeset product scan or materialization. The UTC/gap-free arithmetic selection branch and short-UNTIL horizon remain unchanged. Specific existing Apia skipped-date coverage is strengthened; arbitrary historical transitions and rapidly repeated transitions are still not claimed supported. No broader timezone resolver redesign was introduced.

## Changed files (all6 scoped)

- `civil_time.go`: accept period collision information in dayRanks and exclude the redundant nominal anchor before counting positions.
- `rrule.go`: add bounded period-local anchorCollision lookup and pass its result into positional day ranking.
- `rrule_dst_test.go`: production-path R1 regressions, paired signed/negative controls, half-hour/coarse variants, cross-day coarse collision tests and daily/filter controls.
- `docs/releases/v1.9.0-dst.md`: disclose unique-anchor cardinality before BYSETPOS.
- `docs/repro/05-07-backlog/golden/corpus-dst.json`: append12 individually classified sensitivity/control rows; update provenance and exact regenerated-baseline SHA256.
- `docs/repro/05-07-backlog/golden/baseline-dst.json`: appended old-baseline outputs generated only by executing the unchanged old-baseline driver. Original799 baseline rows are identical.

There are no new out-of-map file edits in this rework. The original approved narrow str.go, UNTIL golden expectations and six Go1.16 ignore-tag adaptations remain untouched.

## Red-before-green evidence

Before any production edit, `HEAD=af86249`; only new public regression tests were added:

1. `red-TestDSTGapAnchorBysetposCollision.txt` — command `GOCACHE=/tmp/rrule-go-cache go test -run '^TestDSTGapAnchorBysetposCollision$' -v -count=1`, **exit1**. Positive1,2 and1,3 fail independently through All, Iterator, Set.All, Between inclusive/exclusive, After and COUNT-free UNTIL before/at/after boundaries. The relevant Before query also runs; its already-correct positive-out-of-range case is not called defective. Every signed/negative control path passes.
2. `red-TestDSTGapAnchorRankPeriods.txt` — command `GOCACHE=/tmp/rrule-go-cache go test -run '^TestDSTGapAnchorRankPeriods$' -v -count=1`, **exit1**. YEARLY/MONTHLY/WEEKLY and Lord Howe half-hour positive collision variants fail.

A first narrow same-day correction made the review cases green. Root then pointed out that existing Apia skipped-date coverage makes cross-day coarse cardinality a useful bounded edge. Before the corresponding period-level correction, `red-cross-day.txt` records `GOCACHE=/tmp/rrule-go-cache go test -run '^TestDSTGapAnchorCrossDayRanks$' -v -count=1`, **exit1** for all6 YEARLY/MONTHLY/WEEKLY×positive1,2/1,3 variants. This is explicitly supplemental red evidence on the same-day intermediate correction, not falsely claimed to precede all production work. No production code was changed merely to manufacture failures.

The final period-local correction makes these all green in `green-rank-all.txt` and the final `gate-focused.txt`. Added paired `TestDSTGapAnchorPeriodControls` verifies a later DAILY period retains its normal rank set and that a filtered-out resolved date cannot suppress the explicit anchor. Both pass. Existing helper/data/API behavior is covered by the unchanged DST suite.

The exact reviewer standalone probe was rerun: `GOCACHE=/tmp/rrule-go-cache go run /tmp/rrule-wp01-rank-collision-review.go`, **exit0**, **failures=0** in `gate-independent-probe.txt`. All six former failing assertions and all six original controls now pass.

## Final gates (all exit0, fresh tests)

| Command | Result | Log |
|---|---|---|
| `GOCACHE=/tmp/rrule-go-cache go test -run TestDST -v -count=1` | Pass,0.007s | `gate-focused.txt` |
| `GOCACHE=/tmp/rrule-go-cache go test -count=1 ./...` | Pass,0.464s | `gate-modern.txt` |
| `GOCACHE=/tmp/rrule-go116-cache /tmp/rrule-go116-toolchain/go/bin/go test -count=1 ./...` | Pass,0.336s; actual Go1.16.15 toolchain | `gate-go116.txt` |
| `GOCACHE=/tmp/rrule-go-cache go vet ./...` | Pass, zero diagnostics | `gate-vet.txt` |
| `gofmt -l .` | Pass, zero paths | `gate-gofmt.txt` |
| `GOCACHE=/tmp/rrule-go-cache go test -race -count=1 -run 'TestDST|TestUntil' ./...` | Pass,19.504s | `gate-race.txt` |
| `git diff --check` | Pass, zero diagnostics before commit | Native tool record; no extra mutation |

Modern toolchain remains Go1.25.14; minimum actual toolchain remains Go1.16.15, as verified in initial handoff and direct executable path. Separate caches were retained. Full integrated race/fuzz closeout remains the planned root/WP02 responsibility; no claim of new60s fuzz here.

Resource control command: `GOCACHE=/tmp/rrule-go-cache go test -run '^TestDSTDenseAnnualArithmetic$' -bench BenchmarkDSTAnnualBysetpos -benchmem -benchtime=1s -count=1`, **exit0** (`gate-resource.txt`). The read-only dense arithmetic assertion still selects ranks1,366,-366,-1 from365×86,400=31,536,000 nominal candidates, proving Jan1 and Dec31 endpoints. Benchmark sample: **1,055,523ns/op,2,226,026B/op,47allocs/op**. This is consistent with prior bounded resource evidence, not a new speedup claim; baseline median/resource provenance remains in the original handoff. No annual product enumeration was added.

## Golden comparison and sensitivity

Final **811 rows,114 changed,0 unexplained,0 baseline/current timeouts**. Original corpus had799 rows/102 corrections. Appended12 rows cover review positive collisions through rules/sets, signed/negative controls, half-hour/coarse variants, Apia same-period cross-day collisions and period/filter controls. Every new changed row carries an individual RFC/test rationale and exact current-row SHA256, with controls distinguished from defects.

- All original799 baseline rows exactly match the committed af86249 artifact.
- All original799 current rows exactly match `/tmp/rrule-current-dst.json` from the initial implementation. Thus rework introduced **zero changes to existing golden rows**.
- The old baseline artifact was regenerated only by executing `/tmp/rrule-dst-golden-baseline`, the unchanged driver compiled against pinned76c646dde8d33ff7c3325df0613b31b058243bea; neither the pristine snapshot nor any old-baseline source was edited.
- New baseline artifact SHA256: **68fbd11d9cf3ca187dc473a543876196c23552ea5c2bb1b45f1ca561f75bd226**.
- Exact full baseline/current changed rows and rationale: `golden-classified.txt`; prefix equality/hash proof: `golden-original-rows-unchanged.txt`.
- Signed/negative controls' occurrence outputs were already correct on af86249. Old76c646d may differ only in parsed DTSTART physical interpretation/String for a control; that is explicitly classified, not overclaimed as a rework defect.
- Comparing baseline with itself still fails exit1, changed0 (`negative-empty-comparison.txt`). Corrupting an unchanged UTC row still fails exit1, changed115/unexplained1 (`negative-unclassified-UTC.txt`). No broad DST allowlist or empty comparison pass.

Actual successful commands (exit0; golden runs reached completion):

```
# Lane cwd, build against fixed current module
GOCACHE=/tmp/rrule-go-cache go build -o /tmp/rrule-rework1-golden-current docs/repro/05-07-backlog/golden/main.go
/tmp/rrule-rework1-golden-current -out /tmp/rrule-rework1-current-dst.json
# Unchanged old-baseline executable; pristine snapshot cwd
/tmp/rrule-dst-golden-baseline -corpus ~/repos/rrule-go/.worktrees/teambition-backlog-01M5TEBG-lane-a/docs/repro/05-07-backlog/golden/corpus-dst.json -out ~/repos/rrule-go/.worktrees/teambition-backlog-01M5TEBG-lane-a/docs/repro/05-07-backlog/golden/baseline-dst.json
# Lane cwd
/tmp/rrule-rework1-golden-current -baseline docs/repro/05-07-backlog/golden/baseline-dst.json -current /tmp/rrule-rework1-current-dst.json
```

No operational fault, gate waiver or scope/API decision was required during this rework. A first same-day draft was generalized before commit to respect root's directly relevant cross-day observation; only the final period-local implementation is committed. All Spec Kitty commands used `SPEC_KITTY_NO_UPGRADE_CHECK=1`; safe-commit used the authorized Git metadata escalation and explicit destination branch. Review and status transitions remain for root.
