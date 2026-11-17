---
affected_files: []
cycle_number: 1
mission_slug: teambition-backlog-01M5TEBG
reproduction_command:
reviewed_at: '2026-10-25T22:06:25Z'
reviewer_agent: codex
wp_id: WP01
---

# WP01 independent implementation review

- Reviewer: **gpt-6-astra**, **2026-10-25**; implementation model **gpt-6.1-sol**, 2026-10-25.
- Reviewed code: **af862491333c8a06fdb30bb60a58619e65e9baa6**; lane-a was clean before and after review.
- Workspace: `~/repos/rrule-go/.worktrees/teambition-backlog-01M5TEBG-lane-a`.
- Runtime review invocation: `2735f21bbd04415ebd13b8cd24c62e30`; root already claimed WP01. This reviewer did not claim or transition it. The generated prompt retains implementer identity text; this record states the actual independent reviewer identity.
- Compared with the approved two-issue spec/plan/WP01 prompt and pinned inherited baseline `76c646dde8d33ff7c3325df0613b31b058243bea`. Contributor #7 was not re-reviewed or extended.

## Verdict: REJECT — changes required

One concrete correctness defect remains in the interaction between an explicit gap DTSTART and BYSETPOS. Existing regression, compatibility and minimum-Go gates pass, but they do not cover this combination. Root should record the canonical rejection and return WP01 for the scoped correction before WP02 proceeds.

## R1 [P2] Remove anchor collisions before computing BYSETPOS ranks

**Locations:** `civil_time.go:177–198` (`dayRanks`); `rrule.go:687–711` (selection and subsequent deduplication).

`dayRanks` deliberately retains the explicit gap DTSTART's nominal slot. If the timeset also contains its resolved local clock, both slots represent the same instant. They both contribute to `count` and positional rank. `timeContains(selected, res)` removes duplicate selected instants only after positions have already been chosen, so a valid later occurrence can disappear or an out-of-range position can incorrectly select an occurrence.

Production input:

```text
DTSTART;TZID=America/New_York:20070311T023000
RRULE:FREQ=DAILY;COUNT=4;BYHOUR=2,3,4;BYMINUTE=30;BYSETPOS=1,2
```

The explicit DTSTART resolves to 03:30 EDT / 07:30Z. The generated 03:30 candidate is that same instant; 04:30 EDT / 08:30Z is the second distinct candidate. Thus the first day's positions 1 and 2 select 07:30Z and 08:30Z. Position 1 includes the explicit start, so this case does not rely on an unsynchronized DTSTART exception.

Expected first four UTC instants:

```text
2007-03-11T07:30:00Z
2007-03-11T08:30:00Z
2007-03-12T06:30:00Z
2007-03-12T07:30:00Z
```

Actual:

```text
2007-03-11T07:30:00Z
2007-03-12T06:30:00Z
2007-03-12T07:30:00Z
2007-03-13T06:30:00Z
```

Changing BYSETPOS to `1,3` exposes the converse error: the gap day has only two distinct instants, yet the implementation treats 04:30 as a third and emits it. Expected then is March 11 07:30Z; March 12 06:30Z, 08:30Z; March 13 06:30Z. Both failures reproduce through All, Iterator and Set.All. Paired signed `1,-1` and negative `-2,-1` controls pass; this review does not claim those controls are defective.

**Condition and likelihood:** an explicit parsed DTSTART falls in a DST gap; a same-period BY time expansion includes both its nominal clock and its resolved clock; positional selection crosses those duplicate slots. This is less common than an ordinary daily DST schedule but is a direct combination of WP01's expressly required explicit-anchor collision and positional-set semantics. It needs no malformed input, unusual timezone or new mechanism.

**Smallest correction:** make the valid positional set count a resolved anchor collision once before rank calculation, while retaining the explicit initial instant and nominal cadence. Adjust sparse holes/rank mapping rather than expanding the annual timestamp product. Add the positive-collision and positive-out-of-range regressions, keep the signed/negative controls, and test COUNT and inclusive UNTIL/query consistency for the corrected set. Retain direct UTC selection and resource bounds. This is a WP01 correction, not new duplicate-selector work from inherited issue #7.

The applicable semantics are the explicit gap interpretation and recurrence-set selection rules in [RFC 5545 §§3.3.5 and 3.3.10](https://www.rfc-editor.org/rfc/rfc5545.html), together with FR-001/FR-002 and WP01 T004's explicit-anchor and BYSETPOS requirements.

## Independent reproduction evidence

- Source: `/tmp/rrule-wp01-rank-collision-review.go`; durable exact copy: `evidence/WP01-review-rank-collision.go.txt`.
- Output: `/tmp/rrule-wp01-rank-collision-review.txt`; durable exact copy: `evidence/WP01-review-rank-collision.txt`.
- Command in the reviewed lane: `GOCACHE=/tmp/rrule-go-cache go run /tmp/rrule-wp01-rank-collision-review.go`.
- **Exit 1; six failed assertions**, corresponding to two failing fixtures × three public paths; all six paired-control assertions pass.
- Expected-array SHA256 for positive collision: `ab928aa3e9865b2591680436071a5e0724dd922003d69cccb53a6c7703746ec0`.
- Expected-array SHA256 for positive out-of-range: `cab647ab13c8d9a3dbfb5242e0eac3ffbe493a4b9827baaa2f81a9aa5651d6ee`.
- Hashes cover each compact JSON expected string array, as emitted by the standalone probe. The log also preserves actual outputs, control hashes, exact commit, model/date and exit status.

No production file or permanent test was edited. The new probe lives in /tmp; only this review and its evidence copies were authored in the primary checkout.

## Checks performed and useful passing evidence

The full code diff, new resolver, parser/set adapters, DST tests, golden driver and handoff were reviewed. The six legacy ignore tags and three inherited UNTIL golden-row updates match the root-authorized, disclosed scope. The private ROption field and external positional-literal break are explicitly user-approved for v1.9.0 and disclosed in release notes. The standalone/shared parser adaptations are coordinated with sequential WP02. The inherited selector-normalization logic was not changed by this diff.

Commands actually rerun in lane-a:

| Command | Result |
|---|---|
| `GOCACHE=/tmp/rrule-go-cache go test ./...` | Pass; cached full-suite result |
| `GOCACHE=/tmp/rrule-go116-cache /tmp/rrule-go116-toolchain/go/bin/go test -count=1 ./...` | Pass; fresh full suite, Go 1.16.15 |
| `GOCACHE=/tmp/rrule-go-cache go test -race -count=1 -run 'TestDST\|TestUntil' ./...` | Pass; fresh focused race run |
| `GOCACHE=/tmp/rrule-go-cache go vet ./...` | Pass; no diagnostics |
| `gofmt -l .` | Pass; no paths |
| `GOCACHE=/tmp/rrule-go-cache go run /tmp/rrule-wp01-review-probes.go` | Prepared independent signed-gap, half-hour, later-fold, gap-collision, dense-UTC and horizon examples match expected outputs |
| Golden driver build, current evaluation and baseline comparison | Pass; **799 rows, 102 changed, 0 unexplained, 0 current timeouts** |
| New anchor/rank collision probe above | **Fail; R1** |

The independent golden run wrote `/tmp/rrule-review-current-wp01.json` and `/tmp/rrule-review-golden-wp01.txt` using `/tmp/rrule-review-golden-wp01`. The comparator verified the committed baseline hash and per-row current hashes. Changed-row rationale groups remain gap/fold, positional-gap and explicit/shared-anchor classes; the existing corpus misses R1. The driver is sensitive to output changes, but its passing hashes cannot establish correctness for an unrepresented interaction.

Prepared dense UTC selection remains arithmetic by inspection; the independent 31,536,000-nominal-candidate `BYSETPOS=-1` probe completed selection in approximately 81 microseconds in one sample. This is a smoke observation, not a benchmark claim. The handoff's repeated same-source baseline/current benchmark evidence and red-before-green logs were inspected; the documented fixture-authoring corrections are not represented as library defects. Full 60-second fuzzing remains the planned integrated WP02 gate, not claimed here.

The review skill was applied. The read-only charter checklist command reported that no Code Review Checklist section is authored. The pre-review hook was reported skipped because of undeclared configuration; no hook-based test success is claimed.

## Runtime anti-pattern checklist

| Item | Result | Evidence/qualification |
|---|---|---|
| Dead code | PASS | New civil module is called by buildRRule, generation and DTSTART parsing; metadata hooks have live production calls and the expressly planned WP02 foundation role. |
| Synthetic-fixture tests | PASS | Regression assertions call constructors/parsers/iteration/query APIs. |
| Silent empty return | PASS | Missing civil candidates and out-of-range positions are deliberate exclusions; metadata absence denotes ordinary/stale values. |
| FR coverage | PASS | FR-001/002/007/008 have behavior assertions or release/golden evidence; R1 shows their implementation is incomplete despite existing coverage. |
| Frozen surface | PASS | No contributor-fix rework, dependency change or unapproved public API expansion observed. |
| Locked decision | **FAIL** | R1 counts the same resolved instance in two positional slots, violating required recurrence-set ranking. |
| Shared-file ownership | PASS | Narrow str.go/shared-start adapters and exact ancillary edits are explicitly authorized and disclosed; WP02 remains sequential. |
| Production fragility | PASS | No new production panic/raise path; bounded local timezone operations and iterator state. |

## Re-review requirement

Return a scoped fix commit with the new regression red on this reviewed commit and green after correction, preserve the paired negative/signed controls, rerun affected DST/query/metadata tests and golden classification, and retain dense selection/minimum-Go gates. Root owns the canonical verdict transition. This review does not approve WP01 or permit dependent WP02 implementation.
