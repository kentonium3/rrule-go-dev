# Focused design review — unreachable-byset-hang-01M5RZ90

Reviewer: gpt-6-astra. Implementer: gpt-6.1-sol. Date: 2026-10-25.
Scope: read-only design review of AGENTS.md, docs/CONTEXT.md, docs/issues/02-unreachable-byset-hang.md, mission spec.md/plan.md/tasks.md and baseline rrule.go. No lifecycle action or source change.

Verdict: design ready; no blockers. Actual implementation and evidence require separate fixed-commit review.

## Finite-cycle proof

For each inner loop, use clock coordinate x in Z/MZ, where M is respectively 24, 1440 or 86400. advanceClock advances x by the positive accepted interval modulo M. Its unnormalized civil day carry does not enter clock predicates. The orbit period is M/gcd(M, interval), including intervals divisible by or larger than M. Therefore M tested candidates include every reachable clock state. All failed predicates prove that continuing the inner search cannot succeed. Missing filters and mixed lists retain the same predicate; conjunction correctly covers jointly impossible lists. Retaining the previous horizon check can terminate earlier without introducing impossible dates or overflowing arithmetic.

A filtered-day shortcut adds k*interval to a raw clock component. The combined minute/second coordinate may exceed the corresponding component range, but remains a legitimate integer starting coordinate for modulo-M advancement. The next advanceClock normalizes it. The same finite proof holds; raw component equality as a cycle sentinel would be inferior.

This establishes clock-search termination, not a universal one-second limit for arbitrary calendar filters. Keeping calendar-only scans excluded is accurate and avoids broadening the change.

## Concrete implementation hazards

1. Check a successful predicate on attempt M before concluding exhaustion. A realistic off-by-one change silently drops once-per-day schedules. Strong controls use interval 1, only DTSTART's clock value allowed at each applicable precision, and COUNT=2: the second occurrence requires the full M attempts.
2. generate may already have appended results before advancing. New exhaustion must assign r.len = iterator.total, set finished, and return without clearing remain. next calls generate then Pop even when finished; buffered occurrences must still drain. A clock match already buffered implies its phase is reachable again, so horizon exhaustion is the practical direct test of buffered termination; preserve existing huge-interval/MAXYEAR regressions as well.
3. Preserve advanceClock's check before addition and the existing order of field assignments and combined predicates. Its int64 arithmetic and deferred civil-day normalization are part of the previous panic repair and 32-bit compatibility.
4. Keep attempt state local to each search. Persistent reachability caches would require setter invalidation and are unnecessary.

## Verification expectations

- Both public constructors; After and All for each negative frequency, explicit UNTIL, mixed lists and joint impossibility.
- Setter transitions reachable/unreachable/reachable, fresh iterators, unchanged exact sequences after restoration; Until rebuild coverage.
- Exact reachable controls for final candidate, day and year carry, filtered-day shortcut, non-UTC behavior, COUNT/UNTIL boundaries, and repeated exhaustion.
- Child-process watchdog kills and waits; assert operation time inside child separately from startup/race teardown overhead. No surviving spinning goroutines in the parent.
- Independently captured red results for all three frequencies on baseline; green on fixed code; removal of each individual frequency guard makes that regression fail.
- Full tests, race, vet, build, formatting and linux/386 checks; retain prior overflow tests.

No runtime tools were invoked and no Spec Kitty fault was encountered during this read-only review.
