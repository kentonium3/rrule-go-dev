# 07: Repeated BY values produce duplicate occurrences and shorten COUNT

GitHub issue: kentonium3/rrule-go#7 (public). From teambition/rrule-go#70. Priority P2. No teambition PR exists.

## Problem

A value repeated inside a BY list (`BYSECOND=0,30,0`, `BYHOUR=9,9,10`) is used twice:

- `RRule` returns the same instant twice, and each copy uses up a COUNT position. `FREQ=DAILY;COUNT=6;BYHOUR=9,9,10` covers two days instead of three.
- `RRuleSet` de-duplicates its output, so the same rule there loses occurrences: `FREQ=MINUTELY;COUNT=12;BYSECOND=0,30,0` gives 8 instead of 12.

Measured on v1.8.3: `docs/repro/05-07-backlog/out-v1.8.3.txt`.

## RFC 5545

§3.3.10 defines BY-list processing as operating on a recurrence *set*. A value repeated in a list
names the same selector, so it can't add an instance or use up a COUNT position. The RFC doesn't
forbid repeats, so they stay valid input; don't reject them. (Kent, 2026-10-25: adhere to RFC 5545
in all cases.)

## Acceptance

- A repeated value in any BY list counts once. That covers BYSECOND, BYMINUTE, BYHOUR, BYDAY (the same weekday and ordinal), BYMONTHDAY, BYYEARDAY, BYWEEKNO, BYMONTH, BYSETPOS and BYEASTER. Both reproductions give the same results as their de-duplicated forms.
- Rules without repeats give byte-identical occurrences to v1.8.3.
- Plan decision: whether `String()` keeps the lists as the caller gave them (today's behaviour; recommended, since `OrigOptions` is exported) or writes them de-duplicated.
- The smallest fix is to de-duplicate when the rule is built (`buildRRule`). Keep in mind that BYDAY values with and without an ordinal (`MO` and `1MO`) are different values.
