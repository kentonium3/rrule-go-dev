# Mission Specification: Bound unreachable sub-daily recurrence iteration

**Mission Branch**: `fix/unreachable-byset-hang`  
**Created**: 2026-10-25  
**Status**: Scope confirmed; specification complete  
**Input**: Kent confirmed the proposed issue #5 scope on 2026-10-25.

## Intent Summary

A service evaluating accepted hourly, minutely or secondly rules must regain control promptly when the interval cannot reach the requested clock filters. Preserve correct recurrence sequences and exported APIs, including setters. Add timeout-safe red-first regression coverage and independent model review. General never-matching calendar scans and Vikunja adoption remain outside scope. Confirmation: decision `01M5RZ9H1HV8HPSBD3VE3YF6WE`.

## User Scenarios & Testing

### User Story 1 — Evaluate unreachable clock filters (Priority: P1)

A service accepts a recurrence rule whose interval never reaches its requested clock values. Evaluation returns no occurrences, or construction rejects that rule, instead of occupying the service indefinitely.

**Why this priority**: Small ordinary inputs can consume a service worker; the trigger is realistic for user-authored recurrences and easy to submit deliberately.

**Independent Test**: Exercise both public constructors and the public iteration/query operations with isolated, bounded execution.

**Acceptance Scenarios**:

1. Given DTSTART `20260105T100000Z`, interval 2 and hourly BYHOUR=1, minutely BYMINUTE=1, or secondly BYSECOND=1, construction rejects the rule or `After` and `All` finish with no occurrences within one second per fixture.
2. Adding the reported explicit UNTIL to the hourly fixture still satisfies the same bound.
3. Combined clock filters that are jointly unreachable finish promptly; a filter containing both reachable and unreachable values still produces the correct reachable occurrences.

### User Story 2 — Edit an existing recurrence (Priority: P1)

A caller changes DTSTART or UNTIL using existing setters and evaluates a fresh iterator. Reachability reflects the current rule without changing setter signatures.

**Independent Test**: Change a shared fixture from reachable to unreachable and back; assert bounded empty results in the unreachable state and exact occurrences in the reachable state. Exercise both setters and both constructors where applicable.

### User Story 3 — Keep correct schedules (Priority: P1)

A caller of a currently correct rule receives identical timestamps in identical order, with identical exhaustion behavior.

**Independent Test**: Pair each negative fixture with a reachable control using the same constructor/query harness. The hourly interval-2 BYHOUR=2 control first returns `2026-01-06T02:00:00Z`. Compare exact bounded sequences for minute and second rules, day carries, mixed lists and filtered days.

### Edge Cases

- Interval equal to or larger than a clock cycle, coprime intervals, and extreme accepted intervals.
- Multiple BY-lists that are individually possible but jointly unreachable.
- Missing clock filters, mixed reachable/unreachable values, midnight/year carries and non-UTC locations.
- A current occurrence buffered before later exhaustion; COUNT and UNTIL boundaries.
- Setters used before a fresh iterator. Concurrent mutation of an active iterator is not newly supported.

## Requirements

### Functional Requirements

| ID | Title | User Story | Priority | Status | Delivery | No-op passable? |
|----|-------|------------|----------|--------|----------|-----------------|
| FR-001 | Unreachable clock termination | As a service caller, I need hourly, minutely and secondly rules with unreachable clock BY-lists to reject at construction or exhaust with no occurrences in bounded time through both public constructors and After/All. Include combined filters and explicit UNTIL. | High | Open | [build] | no |
| FR-002 | Setter-safe evaluation | As a caller editing a rule, I need fresh evaluation after DTSTART/UNTIL setters to satisfy FR-001 and restore correct occurrences when the rule becomes reachable. | High | Open | [build] | no |
| FR-003 | Correct sequence compatibility | As an existing caller, I need all currently correct occurrence sequences preserved, including mixed lists, carries, filtered days, buffered occurrences and exhaustion. | High | Open | [ratchet] | yes — paired reachable/unreachable fixtures exercise FR-001 and FR-002 |

### Non-Functional Requirements

| ID | Title | Requirement | Category | Priority | Status |
|----|-------|-------------|----------|----------|--------|
| NFR-001 | Prompt response | Every reported unreachable fixture completes construction rejection or evaluation in at most one second of measured operation time. This is not a universal runtime limit on arbitrary calendar filters. | Performance | High | Open |
| NFR-002 | Safe regression execution | Every timeout test terminates its isolated execution and waits for cleanup; zero timed-out workers remain after the test. Record pre-fix failure and post-fix success for all three frequency families. | Reliability | High | Open |

### Constraints

| ID | Title | Constraint | Category | Priority | Status |
|----|-------|------------|----------|----------|--------|
| C-001 | Compatibility | Preserve exported types, functions, methods and signatures; add no dependencies. Preserve the first mission's overflow safety. | Technical | High | Open |
| C-002 | Scope | Exclude general UNTIL scan optimization, unrelated calendar semantics and Vikunja adoption. Prefer the smallest root-cause fix. | Scope | High | Open |
| C-003 | Review and verification | Different models implement and review each work package, recording identities/date. Require red-first tests, test/race/vet/format checks and exact reachable controls. | Quality | High | Open |
| C-004 | Private delivery | Keep development private; no upstream/public publication or master merge without authorization. Preserve MIT license and modified-copy notice. | Release | High | Open |

### Key Entities

- **Clock filter**: BYHOUR/BYMINUTE/BYSECOND constraints applied at the rule's frequency.
- **Unreachable clock state**: No interval step from the current starting phase satisfies the applicable clock filters together.
- **Occurrence sequence**: Ordered timestamps produced by public iteration/query APIs.

## Success Criteria

- **SC-001**: All reported unreachable fixtures return within one second and produce no impossible occurrences. — [build] · no-op passable: no
- **SC-002**: Reachable-to-unreachable-to-reachable setter scenarios terminate and restore the exact expected sequences. — [build] · no-op passable: no
- **SC-003**: All paired reachable controls match the pre-fix sequences exactly; each shares its harness with a negative case from SC-001 or SC-002. — [ratchet] · no-op passable: yes

## Assumptions and Dependencies

The baseline is merged master `1a62a24b9460201d16e43b693c5ea9cfc3700198`. The panic repair bounds arithmetic but does not make unreachable clock search prompt. Go remains at its existing minimum unless a demonstrated test need justifies the already-authorized version change. Planning decides whether to reject construction or safely exhaust iteration, while meeting setter compatibility. No unresolved product decisions remain.
