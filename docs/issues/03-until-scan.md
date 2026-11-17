# 03: UNTIL does not stop a scan that never matches (optional)

GitHub issue: #6 (this repo). An optional mission; Kent decides whether to do it.

## Problem

UNTIL is only compared with an occurrence when one is emitted (`rrule.go` around L646 and L675 at v1.8.2). A rule that never matches therefore scans every period up to MAXYEAR (9999), even when UNTIL is weeks after DTSTART. `docs/repro/03-until-scan/main.go` shows two cases on DAILY with BYSETPOS and UNTIL a month after DTSTART: `BYSETPOS=2` takes 128 ms, and `BYSETPOS=2..366` takes 13.2 s. Both return 0 occurrences. The default UNTIL of DTSTART + about 292 years (upstream #24, PR #27) does not bound this either.

It terminates, so it is a design limit, not a crash. It is a CPU-cost problem for services that evaluate user rules. Vikunja bounds the cost on its side with BYSETPOS, occurrence and horizon limits.

## Proposed change

After advancing to the next period, finish iteration once that period's start is later than UNTIL. No later occurrence could be emitted, so the results don't change.

## Acceptance (if taken)

- The two repro cases finish in milliseconds.
- Occurrence output is unchanged across the existing test suite, plus a golden set of UNTIL rules: DST, sub-daily, and an UNTIL on an occurrence boundary.

## Private repair (2026-10-25)

Iteration now stops when the earliest possible UTC date in the current period is
strictly later than effective UNTIL. This applies only when DTSTART's location is
exactly `time.UTC` (pointer identity). YEARLY uses January 1, MONTHLY uses day 1,
and WEEKLY and finer frequencies use the cursor's civil day. Existing inclusive
candidate checks and buffered results are preserved. The default horizon and
fresh-iterator setter behavior are unchanged.

Every other location follows the original iterator path, including fixed zones
and locations merely named UTC. Never-matching non-UTC rules can still scan to
MAXYEAR. This restriction preserves existing timezone normalization behavior,
including duplicate dates around historical Sao_Paulo and Apia midnight gaps.

Measured on Go 1.25.14 linux/amd64, baseline source `b9e8e1097ebd9ccb2dab50cee18a230bfdd86aa0`:
public NewRRule/StrToRRule and Iterator/All/After paths for BYSETPOS=2 took
approximately 140–151ms; all six BYSETPOS=2..366 executions exceeded the separate
3s parent watchdog and were killed and collected. Fixed operation times were
below 1ms for both fixtures (process startup/teardown excluded). BYSETPOS=1
positive controls used the same harness. These measurements are local evidence,
not a universal latency guarantee. A default-horizon DAILY BYSETPOS=2 scan took
approximately 9ms; the 100ms acceptance limit applies to the short explicit cases.

`rrule_until_test.go` embeds exact pre-fix golden output for all seven frequencies,
inclusive boundaries, expanded yearly/monthly periods, cross-year weeks, COUNT,
BYSETPOS/buffers, differing UNTIL offsets, DST gaps/folds, midnight normalization
and fixed offsets. Fresh-iterator setter/default controls and isolated timing
regressions accompany the existing suite. Validation includes full tests, race,
vet, native/386 builds and gofmt. Implementation: gpt-6.1-sol, 2026-10-25;
independent gpt-6-astra review is required before acceptance and will be recorded
in the mission evidence.
