# Iterator invariants

No public type or persistent schema changes.

ROption -> validation -> RRule -> Iterator -> buffered occurrences or exhausted.
An error remains a constructor result; the iterator signature has no error return.

- Ordinal candidate: must be within selected period before wdaymask access.
- Easter candidate: offset must fit mask bounds before index addition.
- Calendar cursor: year/month/day and clock components must stay representable;
  stop before crossing MAXYEAR rather than overflowing and repairing afterwards.
- Buffered occurrences: already-computed valid results remain observable when
  advancement discovers exhaustion.
- Correct existing rule: exact occurrence sequence remains unchanged.
- Known-limit state: unreachable subdaily filtering and general never-matching
  scans are outside the repair and explicitly excluded from unconstrained fuzzing.
