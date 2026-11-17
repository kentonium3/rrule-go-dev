# Public behavior contract

NewRRule(ROption) and StrToRRule(string) retain signatures and existing error
behavior for invalid bounds. Iterator() and Next retain signatures. No exported
symbol is added, removed or changed.

Impossible ordinal and Easter candidates are skipped with overflow-safe bounds
checks. Safe valid candidates in the same rule remain eligible. Large intervals
retain mathematical advancement when representable and exhaust when the next
cursor would exceed MAXYEAR. Existing correct occurrences stay exact.

The implementation must not silently swallow a panic with recover, return empty
for every rule, add a blanket interval cap, or alter unrelated recurrence semantics.
