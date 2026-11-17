# UNTIL bound research — 2026-10-25

Read-only investigation of AGENTS.md, docs/CONTEXT.md, docs/issues/03-until-scan.md, mission spec/research, generate/calcDaySet/fillTimeSet/rebuild and Go time implementation. No production changes.

## Source facts and ordinary safe shape

- YEARLY calcDaySet starts at Jan 1, MONTHLY at month day 1. The raw iterator day/month/hour can be later than eligible candidates; it is not a bound.
- WEEKLY starts at cursor day (including first partial week); subsequent cursor is week start. Other frequencies use cursor civil day.
- Candidate civil dates actually come from `ii.firstyday.AddDate(0,0,i).Date()`, not UTC ordinal conversion. firstyday is local Jan 1 midnight. A midnight transition can normalize this date, and AddDate can normalize again.
- Candidate time components come from timeset Clock(), then `time.Date(candidateDate, components, timeset.Location())`. Those components are bounded accepted clock fields, but Date may normalize a gap/fold. Existing per-occurrence UNTIL checks must remain unchanged, including BYSETPOS sorting and buffered remain semantics.
- A guard belongs at the top of the outer generate loop, inside `for remain.Len()==0`, before dayset filtering/BYSETPOS work. On termination copy existing `r.len=iterator.total; iterator.finished=true; return` behavior. Fresh iterator initialization uses rebuilt rule values, so no stale setter cache is needed.

## Fully proved minimal fast path

For `r.dtstart.Location() == time.UTC`, all time constructions are UTC, and no normalization or offset ambiguity exists. Compute the UTC civil start of the current period: Jan 1 for YEARLY; month day 1 for MONTHLY; current civil day for WEEKLY and all finer frequencies. Midnight is <= every possible current/future candidate. Stop only when this value is strictly After(r.until). Comparing instants handles differing UNTIL offsets. This is conservative for finer frequencies (up to a day extra scan), inclusive endpoints remain in occurrence checks, and both UTC short-horizon reported fixtures stop after about 28 periods. Default horizon stays identical. No timezone assumptions, new deps, API or Go bump. Empty timesets/daysets are handled naturally.

This fast path is the smallest provable change under all current constraints, with fallback to existing behavior for other locations. It does **not** satisfy a broad reading of FR-001 for every non-UTC location, so specification must explicitly permit conservative fallback if selected. Do not silently claim it fixes all locations. Pointer identity is a sufficient condition, not name matching; a location named UTC can have transitions/offsets.

## Ordinary timezone day-margin approach and its proof limit

For conventional IANA zones and ordinary FixedZone offsets, period-day comparison against UNTIL expressed in DTSTART location with one or two spare civil days safely handles normal one-hour DST folds/gaps and date-line changes up to 24 hours. Arbitrary **constant** FixedZone offsets are themselves fine: the offset cancels in local civil comparisons, even if many days. But Go 1.16 exposes no supported method to distinguish fixed locations from loaded transition tables or obtain their global offset range.

Consequently a naked 24h/48h margin is **not a universal proof**. `time.LoadLocationFromTZData` accepts TZif signed-int32 second offsets, including changes larger than days, and future transitions can be larger than observed DTSTART/UNTIL offsets. A test of offsets at period start, UNTIL, or candidate noon does not constrain an unseen future transition. firstyday/AddDate further complicate a simple local-date bound. Looking at names, sampling offsets, or treating equal endpoint offsets as fixed is unsound. This is a rare caller-created-location condition but falls under literal exact-compatibility claims.

Go ZoneBounds would require Go >=1.19 and transition traversal; adding it conflicts with current unchanged-version constraint and introduces machinery for exotic cases.

## Possible universal coarse bound (not recommended)

Loaded timezone offsets are signed-int32 seconds. FixedZone offsets are machine int and may exceed this. If observed DTSTART offset exceeds signed-int32 range, disable optimization; such a location must be FixedZone (or otherwise outside the documented supported construction assumptions). Otherwise use the signed-int32 absolute offset bound. `Date` can shift a requested civil value by up to twice this bound; firstyday and AddDate each normalize, then final candidate Date subtracts an offset. A roughly five-offset-bound (about 340 years), plus spare civil-day rounding, UTC period bound can therefore be justified conservatively for loaded locations. UTC gets its exact fast path. This eventually bounds non-UTC scans and preserves both repro timings, but adds opaque constants, enormous slack, and a proof burden far beyond the realistic issue. Extreme FixedZone offsets near machine-int limits also require care about Go time arithmetic wrap; fallback is preferable. Do not implement this sketch without an independent arithmetic proof.

## Recommendation

Choose explicitly between (a) tiny proved UTC optimization with non-UTC fallback and accurately narrowed coverage, or (b) ordinary-timezone day-margin optimization with an explicit compatibility domain that excludes synthetic giant-offset TZif transitions. The existing full-domain exact compatibility + every-location bound + unchanged Go1.16 + minimal no-exotic-machinery requirements pull in different directions. Do not present a day-margin assumption as a theorem.

For the prospective guard, regression cases must include YEARLY DTSTART late in year with earlier selected month/day in following period; MONTHLY DTSTART day 31 with selected day 1; partial/cross-year WEEKLY; inclusive UNTIL; differing UNTIL offset; DST midnight/date-line normalizations; arbitrary FixedZone offsets (e.g. +/-72h); no-match BYSETPOS and empty timeset; setters/default horizon and pending buffers. Test all via public paths. All seven frequencies should preserve exact baseline sequences. Model of investigation/prospective implementation: gpt-6.1-sol; independent review remains required.
