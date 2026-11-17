# Human retrospective notes

2026-10-25. This companion qualifies the automatically captured retrospective without rewriting its runtime record or proposing governance changes.

The WP01 rejection was useful independent review: gpt-6-astra found a real explicit-gap-anchor/BYSETPOS rank collision missed by the existing tests and golden corpus. Red reproduction, the scoped correction and independent re-review prevented shipping the defect. An automatic `not_helpful` bucket must not be read as an assessment that this rejection had no value.

Independent post-plan review also prevented two design regressions: a memory-bounded BYSETPOS stream still scanned the entire annual candidate product, and a name-based timezone marker could not retain a caller's custom same-name location. The plan preserved arithmetic ranks and Kent explicitly authorized private immutable metadata with the disclosed positional-literal break and v1.9.0 target.

Evidence provenance mattered. The wrong-cwd focused run was preserved as invalid and excluded from gates; baseline integrity was checked file by file. Source hashes tied full race/fuzz/minimum-Go/resource evidence to the final reviewed code. A separate baseline regeneration/cmp detected one nondeterministic golden fixture, whose bounded correction and recorded clock context were then independently reproduced. Passing a comparator or a test command name alone would not have established these facts.

Keep these lessons with the mission. The automatic artifact remains intact; no governance proposal or additional mission work is implied.
