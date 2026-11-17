# Mission Specification: Bound recurrence scans by UNTIL

**Mission Branch**: `fix/until-scan`  
**Created**: 2026-10-25  
**Status**: Scope confirmed; specification complete  
**Input**: Kent authorized issue #6 and confirmed scope in decision 01M5S0X2AW994PNVNMEXVCVENB.

## Intent Summary

For rules whose DTSTART location is exactly Go time.UTC (pointer identity), a service evaluating accepted recurrence rules should stop searching once no remaining recurrence period can produce an occurrence at or before the effective UNTIL. This applies to explicit UNTIL and the existing default. Both reported short-horizon cases must complete within 100ms of operation time. Preserve every currently correct occurrence sequence, inclusive UNTIL, DST/sub-daily behavior, setters and exported APIs. Do not add arbitrary horizon caps or change Vikunja.

## User Scenarios & Testing

### User Story 1 — Promptly finish a nonmatching scan (Priority: P1)

A caller evaluates a daily recurrence whose BYSETPOS never selects an occurrence. Despite UNTIL being weeks away, the current implementation searches to year 9999. The caller must regain control promptly with the same empty result.

**Independent Test**: Both constructors followed by public query/iterator calls; use the two cases in docs/repro/03-until-scan/main.go. DTSTART is 20260105T100000Z, UNTIL is 20260201T000000Z, BYSETPOS is 2 or every value 2 through 366. Each evaluation returns empty within 100ms excluding process startup, compilation and teardown. Include same-fixture BYSETPOS=1 positive controls and preserve timeout cleanup.

### User Story 2 — Preserve scheduled occurrences (Priority: P1)

A caller receives exactly the same ordered timestamps and exhaustion behavior for rules that already work, including occurrences equal to UNTIL and eligible dates within yearly/monthly/weekly periods.

**Independent Test**: Exact pre-fix/fixed golden sequences spanning all seven frequencies, UNTIL before/equal/after a candidate, DST gaps/folds, non-UTC offsets, cross-year weeks, BYSETPOS, COUNT and buffered results. An iterator cursor beyond UNTIL must not suppress an earlier eligible date in the same period.

### User Story 3 — Honor the current effective horizon (Priority: P1)

For UTC DTSTART, a caller omits UNTIL or changes DTSTART/UNTIL after construction. Evaluation respects the existing default horizon or updated explicit horizon without new errors or stale state.

**Independent Test**: Fresh iterators after setters, including horizon shortening/extension and zero UNTIL selecting existing default behavior. Demonstrate scan termination respects that horizon rather than MAXYEAR; the 100ms requirement applies only to the two short-horizon fixtures, not to all possible 292-year defaults.

### Edge Cases

- Inclusive endpoint and an UNTIL preceding DTSTART.
- Monthly/yearly candidates earlier than current iterator date fields.
- Weekly periods crossing year boundaries and non-default week starts.
- DST normalization, repeated/nonexistent local times, fixed offsets and differing UNTIL offsets.
- Empty timesets/daysets, BYSETPOS selecting no candidate, count exhaustion and pending buffered results.
- Very large intervals and MAXYEAR, retaining both prior missions' safety fixes.

## Requirements

### Functional Requirements

| ID | Title | User Story | Priority | Status | Delivery | No-op passable? |
|----|-------|------------|----------|--------|----------|-----------------|
| FR-001 | Horizon-bounded nonmatching scan | As a service caller, I need UTC DTSTART searches to finish when no remaining period can yield an occurrence at or before effective UNTIL, even if no candidate is emitted. | High | Open | [build] | no |
| FR-002 | Exact occurrence compatibility | As an existing caller, I need identical ordered occurrences and exhaustion, including inclusive UNTIL, full recurrence periods, DST, sub-daily rules, COUNT and buffered results. | High | Open | [ratchet] | yes — paired with FR-001 negative fixtures and exact positive controls |
| FR-003 | Updated and default horizons | As a caller using setters or omitting UNTIL, I need fresh UTC DTSTART iteration to honor the current explicit or existing default horizon without changing constructor/setter behavior. | High | Open | [build] | no |

### Non-Functional Requirements

| ID | Title | Requirement | Category | Priority | Status |
|----|-------|-------------|----------|----------|--------|
| NFR-001 | Short-horizon latency | Each reported explicit-UNTIL fixture completes within 100ms measured operation time; exclude process startup/compilation/teardown and record toolchain/platform. | Performance | High | Open |
| NFR-002 | Safe defect-sensitive validation | Record genuine pre-fix failure and post-fix success, plus evidence that search work is bounded by the effective horizon. Every timed-out isolated execution must be killed and collected. | Reliability | High | Open |

### Constraints

| ID | Title | Constraint | Category | Priority | Status |
|----|-------|------------|----------|----------|--------|
| C-001 | Compatibility and scope | Preserve exported types/functions/signatures, constructor acceptance, setter behavior, dependencies and prior overflow/reachability fixes; no arbitrary cap, default-horizon change or Vikunja work. | Technical | High | Open |
| C-002 | Quality gates | Use red-first public-path tests, exact golden controls, test/race/vet/build/format checks and a different-model reviewer with recorded identities/date. | Quality | High | Open |
| C-003 | Private release | Preserve MIT license and modified-copy notice; keep development private, with no master merge or public/upstream posting without authorization. | Release | High | Open |

### Key Entities

- **Effective UNTIL**: current explicit limit or the existing library-supplied default.
- **Recurrence period**: frequency-selected interval containing possible candidates; its earliest possible candidate need not equal the raw iterator cursor.
- **Occurrence**: emitted timestamp satisfying current filtering and inclusive end semantics.

## Success Criteria

- **SC-001**: Both documented short-horizon empty-result cases finish within 100ms per evaluation. — [build] · no-op passable: no
- **SC-002**: Every golden compatibility fixture preserves all timestamps and endpoint behavior; fixtures share public harnesses with the empty-result cases. — [ratchet] · no-op passable: yes
- **SC-003**: UTC default and setter-modified horizon cases stop according to the current horizon, with no stale state or API changes. — [build] · no-op passable: no

## Assumptions and Dependencies

Baseline is merged master 886f07adf97d59c5348bd4e64c71ca616227304b. The user explicitly selected this formerly optional performance mission. A safe conservative period bound may perform limited extra work where timezone semantics require it; correctness takes precedence over pruning every period. No universal constant-time guarantee is made. Planning must prove the bound and select defect-sensitive measurement; no product questions remain.

## Approved timezone coverage

On 2026-10-25 Kent selected option 1 in decision 01M5S1NKNPCYNJNX44929ZRZBQ: optimize only DTSTART.Location() == time.UTC, with unchanged fallback for every other location. FR-001, FR-003 and SC-003 termination improvements apply to that UTC subset; FR-002 exact compatibility applies to all locations. A location merely named UTC is not sufficient. Non-UTC never-matching scans remain a documented performance limitation. No Go-version increase or timezone-transition machinery is authorized by this choice. This replaces the broader original timezone pruning assumption.
