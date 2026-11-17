# Mission Specification: Prevent iterator panics on accepted RRULE input

**Mission Branch**: `fix/iterator-panics`
**Created**: 2026-10-25
**Status**: Confirmed scope; ready for planning
**Input**: Private issue kentonium3/rrule-go#4; docs/issues/01-iterator-panics.md.
**Intent confirmation**: Decision 01M5RTA7A5FAHWV04DNQBJ9V6B, confirmed by Kent.

## Intent and scope

A Go service developer accepts recurrence rules through StrToRRule or NewRRule
and asks the iterator for occurrences. Inputs accepted by those constructors
must not crash the service through the identified BYDAY, BYEASTER and INTERVAL
failure families. Construction may return an error, or iteration may safely
skip impossible candidates or finish. Valid recurrence results must stay exact.
Overflow must not produce backwards or otherwise impossible occurrences.

This mission covers the three panic families, related INTERVAL overflow output,
and focused regression/fuzz verification. Unreachable sub-daily BY-list hangs,
general never-matching UNTIL scans, Vikunja integration, and public release are
outside this mission. Fuzzing must account for those known hang paths without
silently claiming they are fixed. Planning will explicitly choose rejection
versus safe iteration, and fix or document the known BYEASTER=257 silent-drop
limit. Those are delegated design decisions within the confirmed scope.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Evaluate ordinal weekdays safely (Priority: P1)

As a service developer, I need accepted ordinal weekday rules to be evaluated
without an index panic so a user's recurrence cannot terminate the service.

**Why this priority**: Monthly ordinals outside the month pass validation today;
iteration can panic, and an unrecovered goroutine panic terminates the process.

**Independent Test**: Use equivalent text and option-based rules through both
public constructors, with the fixed DTSTART 2026-01-05T10:00:00Z. Exercise the
iterator and compare nearby valid rules against recorded expected occurrences.

**Acceptance Scenarios**:

1. **Given** MONTHLY BYDAY=-6MO, 8MO or 53MO, **when** constructed and evaluated,
   **then** construction returns an error or iteration safely yields valid
   occurrences/exhaustion, with no panic.
2. **Given** YEARLY with BYMONTH=1 and BYDAY=-6MO, or BYMONTH=12 and BYDAY=8MO,
   **when** evaluated, **then** the same safety guarantee holds.
3. **Given** nearby working ordinals 5MO and -5MO, **when** evaluated, **then**
   the exact occurrence sequence is preserved. Valid YEARLY ordinals without
   BYMONTH and WEEKLY rules with ignored ordinals retain their current behavior.

### User Story 2 - Evaluate Easter offsets safely (Priority: P1)

As a service developer, I need Easter-relative rules to avoid out-of-range
access even when an offset puts the candidate outside the supported year mask.

**Why this priority**: BYEASTER currently has no effective safety bound and an
accepted offset can panic on the first occurrence or in a later year.

**Independent Test**: Evaluate text and option forms for -100 and 300, boundary
values -81/-80 and 257/258, and machine-int extremes. Include years exposing
boundary failures and ordinary/leap-year cases; pair each refusal/absence check
with a working rule in the same fixture.

**Acceptance Scenarios**:

1. **Given** an unsafe Easter offset, **when** constructed and evaluated,
   **then** the constructor rejects it or iteration safely handles it; no panic
   occurs, including arithmetic overflow before indexing.
2. **Given** working occurrences at -80 and 257, **when** evaluated, **then**
   those correct results remain exact.
3. **Given** the known silent drop at 257, **when** the mission is delivered,
   **then** its chosen disposition is explicit and backed by a reproducible
   case; preserving that known limit must not be presented as fixing it.

### User Story 3 - Evaluate large intervals safely (Priority: P1)

As a service developer, I need large accepted intervals to advance recurrence
safely or terminate, without overflow crashes or fabricated dates.

**Why this priority**: MONTHLY arithmetic can panic; DAILY can run backwards;
HOURLY can emit an overflow-generated sequence.

**Independent Test**: Cover all frequencies with very large and boundary
intervals, representable integer extremes, and near-MAXYEAR starts, through
both public constructors. Use bounded occurrence requests and independent
expected results for nearby valid inputs.

**Acceptance Scenarios**:

1. **Given** MONTHLY INTERVAL=9223372036854775807 on a 64-bit platform, **when**
   evaluated, **then** it is rejected or handled without panic. On narrower
   platforms, cover the representable maximum and textual out-of-range rejection.
2. **Given** DAILY or HOURLY overflow reproductions, **when** evaluated,
   **then** no backwards, repeated, or overflow-generated occurrences are emitted.
3. **Given** a large interval that already evaluates correctly, **when** evaluated,
   **then** its occurrence sequence stays exact; an arbitrary cap must not reject
   a currently correct rule merely because the interval is unusual.

### Edge Cases

- Positive and negative BYDAY ordinals; YEARLY with/without BYMONTH; WEEKLY
  ordinal semantics; mixed working and impossible candidates.
- Easter offsets immediately inside/outside mask boundaries, later-year failures,
  leap years, and integer overflow when an offset is added to a day index.
- Every frequency, interval zero normalization, existing negative-interval errors,
  integer width, very large values and supported upper-year termination.
- COUNT/UNTIL controls and DTSTART normalization; requesting zero, one, or several
  occurrences must not conceal a panic on the next advance.
- Valid output may be empty. Rejection or empty output alone is not proof of safety
  if the test never reaches the real constructor and iterator.

## Requirements *(mandatory)*

### Functional Requirements

| ID | Title | User Story | Priority | Status | Delivery | No-op passable? |
|----|-------|------------|----------|--------|----------|-----------------|
| FR-001 | Safe ordinal weekdays | As a caller, I can construct/evaluate the reported MONTHLY and YEARLY-with-BYMONTH ordinal cases without panic, receiving an error or safe iteration. | High | Open | [build] | no — reproductions panic before the fix |
| FR-002 | Safe Easter offsets | As a caller, I can submit offset boundaries and extreme offsets without out-of-range or overflow-induced panic during construction or iteration. | High | Open | [build] | no — unsafe offsets panic before the fix |
| FR-003 | Safe interval advancement | As a caller, I can submit extreme intervals at every frequency without overflow panic, backwards dates, or overflow-generated occurrences; rejection or safe exhaustion is allowed. | High | Open | [build] | no — MONTHLY, DAILY and HOURLY controls expose existing failures |
| FR-004 | Exact compatibility | As an existing caller, I retain byte-identical serialized occurrences for currently correct rules, including nearby ordinal/Easter boundaries and large valid intervals. | High | Open | [ratchet] | yes — paired with FR-001/002/003 failing cases on shared fixtures, and expected nonempty occurrence controls |
| FR-005 | Public-path regression coverage | As a maintainer, I can reproduce each defect through both public construction paths and iteration, see its test fail before its individual fix, and pass afterward with committed cases/seeds. | High | Open | [build] | no — required defect-sensitive coverage and evidence are absent |
| FR-006 | Explicit Easter-limit disposition | As a maintainer, I can identify whether BYEASTER=257 silent drops were corrected or retained as a known limit, with a concrete case and documented rationale. | Medium | Open | [build] | no — mission disposition and evidence are not yet recorded |

### Non-Functional Requirements

| ID | Title | Requirement | Category | Priority | Status |
|----|-------|-------------|----------|----------|--------|
| NFR-001 | Fuzz safety evidence | Run a committed fuzz target over BY* domains, BYDAY ordinals, BYEASTER and INTERVAL for at least 60 seconds with zero observed panics/crashes; include deterministic defect seeds and document exclusions/bounds for known out-of-scope hangs. | Reliability | High | Open |
| NFR-002 | Verification gate | Existing tests, regression tests, race tests, vet and formatting checks must have zero failures; record commands, toolchain and outcomes. No unexplained baseline test loss. | Quality | High | Open |
| NFR-003 | Bounded verification | Reproduction and fuzz validation commands must have explicit finite execution limits, record timeouts as failures or scoped exclusions, and never count a timed-out run as passing. | Operability | High | Open |

### Constraints

| ID | Title | Constraint | Category | Priority | Status |
|----|-------|------------|----------|----------|--------|
| C-001 | Exported API | Preserve exported types, functions, methods and signatures; preserve currently correct recurrence sequences. | Compatibility | High | Open |
| C-002 | Focused dependency-free repair | Add no dependencies; prefer the smallest root-cause repair that satisfies safety and compatibility. | Design | High | Open |
| C-003 | Test-first and separate models | Each independent defect must have red/green evidence; each WP must be reviewed by a model different from its implementer, with model identities and date recorded. | Process | High | Open |
| C-004 | Toolchain decision | A minimum Go version increase is authorized if needed for fuzzing; select and justify it during planning. | Compatibility | Medium | Open |
| C-005 | Private delivery | Keep work private, use only private origin for the PR, preserve MIT LICENSE/notice, and stop before master merge or any public/upstream posting. | Disclosure | High | Open |
| C-006 | Scope boundary | Do not implement the separate unreachable-BY-list hang or general UNTIL-scan optimization. Document fuzz limits and the BYEASTER=257 disposition. | Scope | High | Open |
| C-007 | Autonomous fault discipline | Follow docs/runbooks/autonomous-iterator-panics.md: preserve fault evidence, diagnose, validate and record safe recovery; never bypass substantive gates. | Process | High | Open |

### Key Entities

- **Accepted rule**: a rule for which StrToRRule or NewRRule returns no error.
- **Occurrence**: a timestamp yielded by the iterator; expected sequence includes
  ordering, count, dates, times and location/offset serialization.
- **Impossible candidate**: a selection with no date in the relevant period.
- **Correct existing rule**: one whose existing occurrences are consistent with
  its recurrence; known overflow artifacts and documented silent drops are not
  automatically promoted to compatibility requirements.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Every reported panic case is rejected or evaluated without panic through both construction paths, and each defect has an independently failing pre-fix test. — [build] · no-op passable: no
- **SC-002**: Large-interval regression fixtures emit zero backwards or overflow-generated occurrences. — [build] · no-op passable: no
- **SC-003**: All compatibility fixtures preserve exact expected occurrences; nonempty positive controls distinguish preservation from returning no results. — [ratchet] · no-op passable: yes; paired with the failing safety cases on shared fixtures
- **SC-004**: At least 60 seconds of bounded fuzzing completes with zero observed panics and all declared verification checks pass. — [build] · no-op passable: no; committed defect seeds expose unfixed failures
- **SC-005**: The private PR contains review evidence, validation results and explicit remaining limitations, including the Easter-offset disposition. — [build] · no-op passable: no

## Assumptions and planning decisions

The issue and reproductions are the primary problem evidence. Exploitability
requires a service to evaluate attacker-controlled or otherwise untrusted rules;
no population-wide likelihood estimate is claimed. Triggering values are simple
and accepted today, so the crash risk is realistic for exposed callers.

Planning decides the precise error/skip/exhaustion behavior per family while
preserving correct inputs, the interval arithmetic strategy, Easter spillover
scope, and the smallest adequate toolchain/fuzz arrangement. No unanswered
operator scope question remains. The known sub-daily hang is not fixed by this
mission; the verification design must avoid confusing a bound on requested
occurrences with a bound on one iterator call.
