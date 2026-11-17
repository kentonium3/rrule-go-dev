# Data model and invariants

Kent approved one private ROption metadata field and targetv1.9.0 (decision01M5TFFVFY5YXYPVD4JN9SS90M). Existing exported fields/signatures/keyed literals remain; formerly valid external positional ROption literals break and must be documented. No new exported identifier.

- CivilFields: private calendar/time tuple, no instant until zone interpretation.
- TimeForm: private DATE/floating/UTC/TZID kind; source syntax determines form, not Location()==UTC.
- ParseMetadata: immutable descriptors for DTSTART/UNTIL, exact original *time.Location pointers, nominal civil fields and public time snapshots. Value copies may share an immutable pointer safely; builders/setters never mutate it in place. No name reload or sentinel location.
- ParsedAnchor: nominal tuple plus RFC-resolved physical first instant in the actual location; RRule keeps nominal cadence separate from the lower bound.
- Candidate: nominal tuple plus earliest exact-match instant, or invalid generated gap. Only valid generated candidates contribute BYSETPOS positions. Explicit DTSTART is interpreted separately.
- PeriodRanks: regular-day counts and sparse removed gap ranges; arithmetic selection for gap-free periods, binary-search/rank mapping around holes; no full annual candidate enumeration.

Flow: collect lexical descriptors → validate pair/shared DTSTART → construct ordinary time values plus immutable metadata → NewRRule validates snapshots and initializes private evaluation state → select valid candidates/BYSETPOS ranks → inclusive UNTIL/COUNT → query/Set adapters. Serialization consults valid part metadata.

Before use, compare each current public Dtstart/Until to its own snapshot including location identity; invalidation is per-part. A new ordinary Dtstart/Until replaces only that descriptor; other part remains. Internal shared-DTSTART transfers carry metadata through private helpers; public time-only setters intentionally treat supplied values as authoritative. Standalone StrToDtStart resolves its instant but cannot carry lexical/cadence metadata through detached time.Time copies; option/set parsers retain it directly. Direct NewRRule on manually constructed options keeps historical programmatic serialization. An authoritative later-fold DTSTART from a constructor or public setter counts once when aligned; parsed ambiguous civil DTSTART chooses first. No-DTSTART fragments retain legacy UTC-Z UNTIL output and parsing without explicit-start descriptors.

Scope: issues5 and6 only. Contributor duplicate-selector behavior at76c646d is inherited golden baseline, not an entity or implementation concern for this mission. Legacy EXRULE only undergoes required pair validation and remains historically ignored; multiple RRULE lines retain last-wins after all validate.
