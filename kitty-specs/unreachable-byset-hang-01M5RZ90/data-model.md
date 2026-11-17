# Iterator state and invariants

No new domain object, exported field, serialization or persistence.

- RRule retains frequency, interval, BY-lists, DTSTART, UNTIL and count.
- rIterator retains normalized clock fields, pending day carry, total/count, remain and finished.
- A local attempt counter exists only within one inner clock search. Its maximum is the number of clock states in a day for that frequency.

Transitions: advance with existing overflow guard -> normalize -> check combined clock predicate -> continue on match; otherwise repeat until bound -> set len/finished and return. Calendar filtering remains in the outer loop. Buffered remain entries survive exhaustion and can still be popped. A fresh iterator after a setter uses rebuilt rule state without cached reachability.

Correctness invariant: no clock match is skipped before a full possible cycle has been checked. Once all states fail, advancing dates cannot change the clock predicate. General calendar unreachability is not inferred.
