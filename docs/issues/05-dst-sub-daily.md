# 05: Instances at nonexistent or repeated local times (DST)

GitHub issue: kentonium3/rrule-go#5 (public). From teambition/rrule-go#63. Priority P1.

**Normative rule (Kent, 2026-10-25): adhere to RFC 5545 in all cases.** Where python-dateutil, a
teambition PR or an earlier brief disagrees with the RFC, the RFC wins.

## Problem

`generate()` builds every candidate with `time.Date(y, m, d, hour, minute, second, ns, loc)`. Across
a DST change Go normalises that wall-clock time in two ways, and both break RFC 5545:

- **Nonexistent local time** (Sydney 2022-10-02, 02:00 → 03:00). Go moves 02:00 forward to 03:00
  AEDT. `FREQ=HOURLY` yields 03:00 AEDT twice, and so does `MINUTELY;BYMINUTE=0`.
  `FREQ=DAILY;BYHOUR=2,3` yields 03:00 AEDT twice on that day. A daily 02:30 rule yields 03:30.
- **Repeated local time** (Sydney 2023-04-02, 03:00 AEDT → 02:00 AEST). Go picks the later offset,
  so 02:00 resolves to AEST and 02:00 AEDT is never produced.

The v1.8.3 output for each case is in `docs/repro/05-07-backlog/out-v1.8.3.txt`.

## RFC 5545 requirements (all frequencies, YEARLY to SECONDLY)

- **§3.3.10:** a recurrence instance with a nonexistent local time "MUST be ignored and MUST NOT be
  counted as part of the recurrence set". It is skipped, not moved forward.
- **§3.3.5:** a local time that occurs twice refers to its first occurrence, the pre-change offset.
- DTSTART itself is covered by §3.3.5, which interprets a nonexistent DTSTART using the offset
  before the gap. The plan states how DTSTART interacts with the rule and cites the RFC text.
- Floating and UTC DTSTARTs have no DST, so nothing changes for them. Vikunja evaluates rules this
  way.

## Acceptance

Expected results with Sydney DTSTARTs:

| Rule | Expected |
|---|---|
| `HOURLY;COUNT=6` from 2022-10-02 00:00 | 00:00, 01:00 AEST; 03:00, 04:00, 05:00, 06:00 AEDT |
| `HOURLY;COUNT=6` from 2023-04-02 00:00 | 00:00, 01:00, 02:00 AEDT; 03:00, 04:00, 05:00 AEST |
| `MINUTELY;BYMINUTE=0;COUNT=4` from 2022-10-02 00:00 | 00:00, 01:00 AEST; 03:00, 04:00 AEDT |
| `DAILY;BYHOUR=2,3;COUNT=4` from 2022-10-01 00:00 | 10-01 02:00, 03:00 AEST; 10-02 03:00 AEDT; 10-03 02:00 AEDT |
| `DAILY;COUNT=3` at 02:30 from 2022-10-01 | 10-01 02:30 AEST; 10-03 02:30 AEDT; 10-04 02:30 AEDT (10-02 skipped) |
| `DAILY;COUNT=3` at 02:30 from 2023-04-01 | 04-01 02:30 AEDT; 04-02 02:30 **AEDT**; 04-03 02:30 AEST |
| control: `HOURLY;BYHOUR=9;COUNT=2` from 2023-04-01 | 04-01 09:00 AEDT; 04-02 09:00 AEST (unchanged) |
| control: `DAILY;COUNT=3` at 09:00 from 2023-04-01 | unchanged from v1.8.3 |

- No result ever contains the same instant twice, and results stay in ascending order.
- BY values keep matching wall-clock values (see PR #66 below). Make the `BYHOUR=9` control a named
  regression test.
- Zones: Australia/Sydney, America/New_York, and Australia/Lord_Howe, whose DST shift is 30 minutes.
  `BYMINUTE` matters there.
- COUNT, UNTIL, BYSETPOS, `Between`, `After`, `Before`, the iterator and `RRuleSet` all respect the
  skip. A skipped instance takes no BYSETPOS position.
- **Golden comparison:** a difference from v1.8.3 is allowed only for an instance whose local time
  is nonexistent or repeated, and the COUNT/BYSETPOS shifts that follow from it. Everything else is
  byte-identical.
- This changes results for existing rules in DST zones on transition days, including daily ones.
  That is intended. The release notes must say so plainly.

## Prior art: teambition PR #66 (do not trust it)

https://github.com/teambition/rrule-go/pull/66, open and unreviewed since 2024-05-06. For sub-daily
frequencies it builds each candidate as `midnight.Add(hour*h + minute*m + second*s)`, which steps in
elapsed time. Probed on 2026-10-25 by applying its head `e7bed74dc0` to v1.8.2:

- It removes the clocks-forward duplicate for sub-daily rules only.
- **It is not RFC 5545.** When clocks go back it yields both 02:00 AEDT and 02:00 AEST. The RFC
  requires only the first. Do not adopt the elapsed-time model.
- **It regresses a working rule.** On the day clocks go back, `FREQ=HOURLY;BYHOUR=9` returns
  **08:00 AEST** instead of 09:00, because midnight plus nine elapsed hours is 08:00 after the
  change. Its tests use COUNT=3 and never get past the transition.
- It leaves DAILY and coarser untouched.

Its test cases are useful once their expected values are re-derived from the RFC. Its code is not.
