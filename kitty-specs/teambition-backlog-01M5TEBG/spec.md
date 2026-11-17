# Mission Specification: Teambition backlog

**Mission Branch**: `fix/teambition-backlog`
**Created**: 2026-10-25
**Status**: Revised for approved private metadata/v1.9.0 and two-issue scope; independent post-plan re-review required before tasks
**Input**: Kent's autonomous brief in `docs/runbooks/teambition-backlog.md` and `docs/issues/05-07`; public kentonium3/rrule-go issues #5 and #6 after Kent's2026-10-25 scope correction; contributor issue7 fix is inherited baseline only.

## Intent Summary

A developer evaluates or round-trips user-supplied recurrence rules. They need RFC5545 civil-time interpretation, preserved floating/date value types, without changing public signatures/keyed-call compatibility or unrelated results. Kent explicitly authorized private ROption parse metadata and a v1.9.0 target; formerly valid external positional ROption literals will break and must be disclosed. Kent authorized ordinary decisions autonomously, RFC corrections in this spec, red-first tests, different implementation/review models, complete local gates, a private PR to main, and no merge of that final PR. Decision01M5TEEVR2WYZ0VRC3GGRXNZ2C records original intent;01M5TFFVFY5YXYPVD4JN9SS90M records the explicit metadata/v1.9.0 exception, and01M5TFJAPGE9MN2AWDG1K4EETF records narrowed scope. Original history remains evidence, not current three-issue authorization.

## User Scenarios & Testing

### User Story 1 - Correct civil-time schedules (P1; kentonium3/rrule-go#5)

A developer schedules recurring activity through a daylight-saving transition. All frequencies YEARLY through SECONDLY skip nonexistent generated wall times and choose the first occurrence of a repeated wall time. BY selectors continue to match wall-clock fields.

Independent test: Sydney HOURLY COUNT=6 beginning 2022-10-02 00:00 yields local hours 00,01,03,04,05,06 with correct AEST/AEDT offsets; beginning 2023-04-02 00:00 yields 00,01,02 AEDT, then 03,04,05 AEST. The fold does not produce a second 02. NY and Lord Howe fixtures prove northern and half-hour transitions.

Acceptance scenarios:

1. Sydney MINUTELY;BYMINUTE=0;COUNT=4 from spring midnight yields 00,01,03,04. DAILY;BYHOUR=2,3;COUNT=4 from October 1 yields October 1 02,03; October 2 03; October 3 02.
2. DAILY;COUNT=3 at 02:30 from October 1 yields October 1,3,4 at 02:30; the spring gap consumes neither COUNT nor a BYSETPOS slot. From April 1 yields April 1/2 02:30 AEDT, April 3 02:30 AEST.
3. Named regression `BYHOUR=9` across Sydney's fold stays at 09:00; ordinary daily09 schedules stay identical.
4. Iterator, All, Between, Before, After and Set agree, results strictly increase with no duplicate instants, and UNTIL stays inclusive after skips and fold resolution.
5. Explicit parsed NY DTSTART 2007-03-11 02:30 resolves with the pre-gap offset to 03:30 EDT and counts first; a DAILY rule resumes nominal 02:30 on following days. This exception for explicit DTSTART must not allow generated gap instances. Parsed fold DTSTART resolves first occurrence. A caller-supplied time.Time is an already-resolved instant. Direct NewRRule and public DTStart setters given the later instant of a repeated local time retain that authoritative DTSTART exactly once when synchronized with selectors; only parsed ambiguous civil text chooses the first offset.

### User Story 2 - Preserve recurrence time forms (P2; kentonium3/rrule-go#6)

A developer parses DTSTART and UNTIL and serializes the same semantic rule. An explicit floating DATE-TIME stays floating; explicit UTC stays UTC; a TZID DTSTART uses a UTC UNTIL; DATE DTSTART and DATE UNTIL remain DATE.

Independent test: floating DTSTART20201228T090000 and UNTIL20201230T090000 round-trip without Z or TZID and give three daily09 occurrences. Pair each rejected form with a conformant form on the same fixture.

Acceptance scenarios:

1. Reject with the offending rule-part name: TZID or UTC DTSTART plus floating UNTIL; floating DTSTART plus UTC UNTIL; either direction of DATE/DATE-TIME mismatch; TZID applied to a Z value or DATE value (§3.2.19). Apply to StrToRRule, StrToROption and its InLocation variant, and all set parsing variants.
2. Accept TZID DTSTART plus UNTIL ending Z, UTC plus UTC, floating plus floating, DATE plus DATE (canonical DTSTART;VALUE=DATE:YYYYMMDD). Preserve semantic type on output. Keep legacy bare DATE inputs compatible where accepted, canonicalizing type where today's output is RFC-wrong.
3. Validate after collecting DTSTART and UNTIL independent of inline attribute order, two-line order and shared set DTSTART position; validate every RRULE and legacy EXRULE line even if a later RRULE overwrites an earlier one. No exclusion engine is added.
4. Bare RRULE with no DTSTART preserves today's parsing of floating UNTIL in the supplied loc. It remains a library fragment, with no absent-DTSTART type check; floating/DATE UNTIL retains legacy UTC-Z String output in this no-DTSTART fragment, rather than acquiring explicit-start metadata semantics. Programmatic NewRRule with Until in any ordinary location still serializes UTC Z. Setters and copied options preserve or deliberately replace parser semantics according to the supplied value.

### Inherited baseline — duplicate selectors (no mission work)

Kent removed issue7 from scope after the contributor fix landed on private main at76c646dde8d33ff7c3325df0613b31b058243bea (publicPR9). This mission neither assesses nor extends that fix. Its behavior is part of the golden baseline (v1.8.3 plus76c646d), and ordinary compatibility controls can include its already-existing tests.

### Edge Cases

COUNT replenishment after skipped values; UNTIL before/at/after a transition; positive/negative BYSETPOS after gaps; invalid gap DTSTART versus generated candidates; inline DTSTART following UNTIL; shared DTSTART after RRULE; UTC TZID; both DATE mismatch directions; TZID-on-DATE; shared-plus-inline DTSTART consistency; copied ROption then NewRRule; repeated setter calls; original caller slices unchanged.

## Normative interpretation and corrections

[RFC5545](https://www.rfc-editor.org/rfc/rfc5545.html) §§3.3.5,3.3.10,3.8.2.4,3.8.5.3,Appendix A.1 is normative. Generated gap instances are removed before recurrence-set position/count processing. Ambiguous DATE-TIME denotes its earlier instant. Explicit gap DTSTART uses the pre-gap offset and must preserve its civil recurrence anchor even though the represented instant displays a different local hour. DTSTART not synchronized with selectors has undefined recurrence-set semantics; no redesign is required.

Corrections/qualifications to briefs and prior art:

- PR66 elapsed-time stepping and double-fold-hour expectations are rejected; selectors are civil-time selectors. Its transition cases may be reused with RFC-derived expected values and credit.
- PR68's suffix-equality validation and local UNTIL serialization are rejected: TZID DTSTART requires UTC UNTIL. Credit any retained TZID-on-Z refusal idea.
- Brief06's blanket unchanged String statement has necessary exceptions: floating and DATE preservation repair today's RFC-wrong UTC serialization. DATE canonical syntax is required for the specified type-matching acceptance; it is not a general calendar parser expansion.
- Brief06's no-DTSTART behavior is a compatible fragment convention; the RFC's DTSTART-dependent constraint cannot be evaluated without DTSTART. Do not describe the fragment as a complete RFC component.
- The brief's DAILY;BYHOUR=2,3 example starts at midnight and is not synchronized with its selectors. RFC§3.8.5.3 leaves that set undefined; retain the example as a library compatibility fixture and add an aligned DTSTART02:00 fixture for normative proof.
- Inline DTSTART is a library extension, not an RFC RRULE rule part. Keep that syntax and validate it; shared set DTSTART cannot be bypassed by an inline DTSTART override.
- EXRULE was removed by RFC5545, and this library has no EXRULE engine. Validate its legacy UNTIL consistency while retaining historical ignore behavior as a compatibility policy, not as RFC-prescribed exclusion semantics. Multiple RRULE semantics are undefined by the RFC; retain existing last-wins while validating all lines. The old mixed set fixture containing TZID DATE-TIME DTSTART and EXRULE UNTIL=20180520 must be split into a rejection and a conformant UTC-UNTIL control.
- The earlier duplicate-selector story/FR-006/SC-003 is retired by Kent's scope decision01M5TFJAPGE9MN2AWDG1K4EETF. Identifiers are not reused; contributor behavior is baseline, not acceptance work.

## Requirements

### Functional Requirements

| ID | Title | User Story | Priority | Status | Delivery | No-op passable? |
|----|-------|------------|----------|--------|----------|-----------------|
| FR-001 | Generated civil times | Story1: all frequencies skip gap candidates before COUNT/BYSETPOS and resolve folds first, keeping civil BY matching, ordering and all public evaluation paths consistent. | High | Open | [build] | no |
| FR-002 | Explicit DTSTART | Story1: parsed ambiguous/gap DTSTART uses §3.3.5, counts first and preserves nominal recurrence cadence; programmatic time.Time remains authoritative. | High | Open | [build] | no |
| FR-003 | Time-form validation | Story2: reject every enumerated forbidden DTSTART/UNTIL or TZID-Z combination through every relevant parser, independent of order and shared DTSTART. | High | Open | [build] | no |
| FR-004 | Floating and DATE preservation | Story2: preserve floating and DATE semantic forms through parsing, option copies, construction, serialization and sets; keep conformant TZID+UTC accepted. | High | Open | [build] | no |
| FR-005 | Constructor and fragment controls | Story2: retain ordinary NewRRule UTC UNTIL serialization and no-DTSTART InLocation interpretation; setter behavior follows actual supplied values. | Medium | Open | [ratchet] | yes — paired with FR-003/004 corrected forms on same fixtures |
| FR-007 | Unaffected compatibility | Both stories: first N occurrences and String match v1.8.3 plus76c646d byte-for-byte outside enumerated RFC corrections, including BYHOUR9 fold control and no-repeat forms. | High | Open | [ratchet] | yes — paired with corrected cases FR-001/003/004 in golden corpus |
| FR-008 | Change communication | Both stories: release notes list every changed result/serialization class and newly rejected form; retained prior-art ideas/tests credited. | Medium | Open | [build] | no |

### Non-Functional Requirements

| ID | Title | Requirement | Category | Priority | Status |
|----|-------|-------------|----------|----------|--------|
| NFR-001 | Gate completion | Zero failures from go vet ./..., go test ./..., go test -race ./...; zero paths from gofmt -l .; FuzzIterate completes at least60s without failure. | Reliability | High | Open |
| NFR-002 | Golden coverage | Include existing suite/repro rules and all seven frequencies × four specified zones at spring/fall transition weeks, comparing at least16 occurrences per bounded case plus String; zero unexplained differences. | Compatibility | High | Open |
| NFR-003 | Defect evidence | Each independently defective behavior has a production-path red-before/green-after regression; independent model review records implementation/review models and date for each of two WPs. | Verification | High | Open |

### Constraints

| ID | Title | Constraint | Category | Priority | Status |
|----|-------|------------|----------|----------|--------|
| C-001 | Compatibility boundary | Go1.16 minimum, no new dependencies, existing public signatures and keyed ROption callers retained. Kent approved an additive private ROption metadata field and v1.9.0; external positional ROption literals become invalid and release notes must disclose the source break. No new exported identifier is planned. Decision01M5TFFVFY5YXYPVD4JN9SS90M. | Technical | High | Open |
| C-002 | Private delivery | Work only in private kentonium3/rrule-go-dev. Consolidate onto fix/teambition-backlog then private PR to main; stop without merging final PR. No public action, tag/release, upstream push, closing keywords, force-push or stash. | Operational | High | Open |
| C-003 | Governed arc | One mission/two WPs, one per issue5 and6; independent post-plan review before tasks; analyze, independent implementation review, strict accept, consolidate, integrated final review. Prefix commands SPEC_KITTY_NO_UPGRADE_CHECK=1; D-3 protocol and runbook progress/evidence. | Process | High | Open |

### Key Entities

A civil candidate is requested calendar fields before timezone resolution. An explicit DTSTART is the supplied recurrence anchor and first instance; generated candidates are subject to gap exclusion. A time form is DATE, floating DATE-TIME, UTC DATE-TIME or TZID DATE-TIME. An immutable parser descriptor records lexical forms, nominal civil anchor and exact supplied *time.Location without reconstructing a zone from its name.

## Success Criteria

- **SC-001**: Every specified transition schedule yields the independently derived RFC instances, with no duplicate instant and COUNT fulfilled when enough valid instances exist. — [build] · no-op passable: no
- **SC-002**: Every invalid time-form pairing is rejected and its conformant paired control succeeds, with floating/DATE semantic round-trips preserved. — [build] · no-op passable: no
- **SC-004**: Zero unexplained differences in the committed reproducible v1.8.3-plus76c646d comparison; all enumerated gates and independent reviews pass before delivery. — [ratchet] · no-op passable: yes — paired with deliberately changed golden rows for SC-001/002
- **SC-005**: Deliver one private reviewable PR to main containing the two fixes and evidence; final PR remains unmerged. — [build] · no-op passable: no

## Assumptions and non-goals

Timezone data supplied by Go/IANA is authoritative for transition offsets. A floating value has no intrinsic zone; UTC numerical storage or a caller-selected evaluation location is a library transport/evaluation convention, not permission to emit Z. No calendar component framework, new exclusion-rule engine, generalized RFC validator, leap-second redesign, public publication, release or Vikunja change is included. Preserve MIT attribution and modified-copy notice.
