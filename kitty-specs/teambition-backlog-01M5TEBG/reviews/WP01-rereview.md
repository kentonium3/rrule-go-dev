# WP01 independent re-review — cycle 2

- Reviewer: **gpt-6-astra**, **2026-10-25**. Implementation: **gpt-6.1-sol**, 2026-10-25.
- Exact reviewed code: **ef373a5872dba30063a6a706275a28e46508af77**; parent **af862491333c8a06fdb30bb60a58619e65e9baa6**.
- Workspace: `~/repos/rrule-go/.worktrees/teambition-backlog-01M5TEBG-lane-a`; clean before and after review.
- Runtime review invocation: `bd99dfe54f714d9eb1985ccf56556b0e`; root claimed the review and owns the canonical verdict transition. This reviewer made no status transition, production edit, Git mutation or public action.

## Verdict: APPROVE

**R1 is resolved. No remaining blocking finding was identified within WP01's approved scope.** This approval supersedes the cycle 1 rejection for the exact code above. Root may record the canonical approval and advance to sequential WP02. This is not final mission acceptance or an approval of unimplemented WP02 behavior.

## Scope and correction

Reviewed the complete six-file rework delta, full regenerated review prompt, rework handoff and relevant root evidence, retaining the original full WP01 review as context. The correction removes the redundant nominal gap-anchor slot before positional counting when a normal candidate in the filtered current period represents the same instant. Its clock lookup is binary, and its selected-day lookup is bounded by the current period; it does not enumerate the day-timeset product. Ordinary UTC starts return immediately from the extra check. Sparse gap ranks and the existing effective-UTC termination bound remain intact.

The selected-day membership check matters: it covers Apia's skipped date resolving onto another selected date in a coarse period, while preserving the explicit anchor if the resolved date was filtered out. Separate DAILY periods keep their own positional sets. New tests cover those controls, YEARLY/MONTHLY/WEEKLY collisions, Lord Howe's half-hour collision, both original positive-position failures, and the original signed/negative controls.

The rework has no parser/UNTIL algorithm, metadata layout, dependency or public-signature change. The earlier authorized parser/set adapter wiring, three RFC expectation corrections and six legacy Go 1.16 ignore tags are unchanged. The private ROption metadata and positional-literal break remain explicitly user-approved for v1.9.0. Inherited contributor issue #7 at baseline `76c646dde8d33ff7c3325df0613b31b058243bea` was not re-reviewed or extended.

## R1 verification and independent probes

The unchanged independent cycle 1 probe now passes **all 12 assertions**, including six formerly failing assertions. Positive `BYSETPOS=1,2` returns March 11 07:30Z and 08:30Z before March 12 results; positive `1,3` no longer invents a third distinct position on the gap day. All, Iterator and Set.All agree. Signed `1,-1` and negative `-2,-1` remain passing controls, not newly fixed defects. The original expected-array hashes remain unchanged.

New production-path regression tests also verify Between inclusive/exclusive, After, Before, and COUNT-free inclusive UNTIL boundaries one second before, at and after the selected occurrence. Reviewed the red-before-green evidence with its stated provenance: original positive/coarse cases were red before production edits; supplemental cross-day cases were red on the intermediate same-day correction before the final period-local change. The latter is not presented as an initial baseline run.

The prepared independent probes also still match their expected instants: Sydney signed gap ranks, a Lord Howe non-hour gap, parsed New York minutely/hourly gap collisions, parsed first-fold selection, and authoritative later-fold constructor and setter behavior. The later-fold cases start at `2026-11-01T06:30:00Z`, while the parsed fold starts at `05:30Z`; both then progress to `07:30Z` and `08:30Z`. Dense annual UTC selection returns `2026-12-31T23:59:59Z`, and short UTC-Z/default-floating UNTIL probes terminate immediately without scanning the default horizon. Timing observations are smoke checks, not benchmark claims.

Durable rerun outputs: `evidence/WP01-rereview-rank-collision.txt` and `evidence/WP01-rereview-probes.txt`. The original assertion source remains in `evidence/WP01-review-rank-collision.go.txt`.

## Commands actually rerun against the reviewed commit

All commands used the lane above as their working directory.

| Command | Result |
|---|---|
| `GOCACHE=/tmp/rrule-go-cache go test -count=1 ./...` | PASS, fresh full suite, 0.475s |
| `GOCACHE=/tmp/rrule-go116-cache /tmp/rrule-go116-toolchain/go/bin/go test -count=1 ./...` | PASS, fresh minimum-Go full suite, 0.334s |
| `GOCACHE=/tmp/rrule-go-cache go test -race -count=1 -run 'TestDSTGapAnchor\|TestDSTMetadata\|TestDSTCustom\|TestDSTAuthoritative\|TestDSTUTCShortUntil' ./...` | PASS, fresh affected-path race run, 1.036s |
| `GOCACHE=/tmp/rrule-go-cache go vet ./...` | PASS, zero diagnostics |
| `gofmt -l .` | PASS, zero paths |
| `GOCACHE=/tmp/rrule-go-cache go run /tmp/rrule-wp01-rank-collision-review.go` | PASS, 12 assertions, failures=0 |
| `GOCACHE=/tmp/rrule-go-cache go run /tmp/rrule-wp01-review-probes.go` | Expected outputs confirmed by inspection |
| Fresh golden executable build, current evaluation and exact baseline comparison below | PASS, 811 rows / 114 changed / 0 unexplained |

The fresh full-suite runs include the new coarse/half-hour/cross-day/filter tests and dense arithmetic test. The implementer's full focused DST/Until race log and resource benchmark also pass; this review distinguishes those supplied results from its own reruns. Full integrated fuzzing remains the planned WP02/mission gate. The pre-review hook was skipped for undeclared configuration; no hook-based test success is claimed.

Golden commands:

```sh
GOCACHE=/tmp/rrule-go-cache go build -o /tmp/rrule-review2-golden-wp01 docs/repro/05-07-backlog/golden/main.go
/tmp/rrule-review2-golden-wp01 -out /tmp/rrule-review2-current-wp01.json
/tmp/rrule-review2-golden-wp01 -baseline docs/repro/05-07-backlog/golden/baseline-dst.json -current /tmp/rrule-review2-current-wp01.json
```

The comparator validates baseline and classified-row hashes. Independent JSON comparisons additionally confirm that the first **799 current rows equal the cycle 1 independent output**, and the first **799 baseline rows equal the committed af86249 artifact**. Only the 12 appended cases extend the comparison; zero current timeouts were observed. Baseline SHA256: `68fbd11d9cf3ca187dc473a543876196c23552ea5c2bb1b45f1ca561f75bd226`. Full independent comparison output is `/tmp/rrule-review2-golden-wp01.txt`; root's durable rework evidence contains the corresponding complete classified output and negative sensitivity controls.

## Runtime review checklist

| Item | Verdict | Basis |
|---|---|---|
| Dead code | PASS | Collision helper is called before day-rank construction; existing metadata/resolver paths remain live. |
| Synthetic-fixture tests | PASS | New assertions use public parsing, constructors, iteration and queries. |
| Silent empty return | PASS | Absent collisions and excluded positional slots are intentional; no new swallowed error. |
| FR coverage | PASS | R1 closes the explicit-anchor/positional interaction in WP01 FR-001/002; earlier documentation and compatibility coverage remains. |
| Frozen surface | PASS | Six-file scoped delta; inherited contributor fix untouched; approved API exception unchanged. |
| Locked decision | PASS | Unique anchor cardinality precedes BYSETPOS, with bounded rank computation and period-local membership. |
| Shared-file ownership | PASS | No new out-of-map/shared parser edit in rework; WP02 remains sequential. |
| Production fragility | PASS | No new panic, registry or unbounded timestamp expansion; matching exact instants guards the collision decision. |

No additional correction is required before root records this WP01 verdict. Existing scope limits and the planned WP02/integrated gates remain in force.
