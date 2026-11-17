# Research: teambition backlog

2026-10-25, gpt-6.1-sol. Current scope: public kentonium3/rrule-go issues #5 and #6 only; issue7 contributor behavior is inherited baseline76c646d and is not assessed or extended. Current approved design uses immutable private ROption metadata and targetv1.9.0; see final Approved revisions. The earlier sections retain chronological discovery/rejection evidence and do not authorize their superseded three-WP/no-field/marker/product-stream proposals.

## Historical discovery decisions — superseded where noted below

- Normative source: [RFC 5545](https://www.rfc-editor.org/rfc/rfc5545.html), consulted directly 2026-10-25, §§3.3.5, 3.3.10, 3.8.2.4, 3.8.5.3 and Appendix A.1. No dateutil result or prior PR is an oracle.
- §3.3.5 resolves ambiguous local DATE-TIME to the first occurrence; a nonexistent explicit DATE-TIME uses the pre-gap offset. §3.3.10 instead excludes nonexistent *generated* recurrence instances before COUNT/set-position selection. Keep wall-clock stepping, not elapsed-time stepping. Invalid generated local values need filtering before BYSETPOS chooses its ordinal.
- DTSTART is explicit and the initial recurrence anchor; its gap interpretation differs from subsequent generated instances. A normalized Go time passed by a caller has already lost original requested wall components. Do not invent a nonexistent wall time from such a value; string parsing can still retain/interpret the supplied components correctly. A DTSTART inconsistent with recurrence selectors gives an undefined recurrence set (§3.8.5.3); do not redesign that historical behavior here.
- UNTIL validation needs lexical kind (DATE, floating, UTC, TZID) as well as time.Time. A floating explicit DTSTART cannot be turned into UTC solely because the numerical representation uses UTC. DATE DTSTART has canonical VALUE=DATE syntax (§3.8.2.4), which current StrToDtStart does not parse. The type-matching acceptance therefore requires this narrow syntax support, with preservation of the type in output.
- A bare RRULE fragment has no DTSTART to compare. Keep its existing InLocation semantics; this is a compatible library fragment convention, not proof of complete iCalendar component conformance.
- Code Set has one RRULE, no EXRULE engine, and silently ignores EXRULE. Appendix A.1 removes EXRULE from RFC5545. Do not add an exclusion engine. The brief's consistency checking can still apply to every legacy EXRULE line before ignoring it, documenting that validation does not implement exclusions. Existing multiple-RRULE last-wins behavior remains; validate every line, including overwritten ones.
- ROption is exported and has only exported fields. Adding any field can break positional composite literals; adding an unexported field is also a source compatibility break. Preserve its exact field layout. Carry parser-only lexical kind in private time.Location tokens or another private representation preserving value-copy behavior; no process-global per-rule registry, unsafe, new exported identifier or dependency.
- Preserve caller BY lists/String output via OrigOptions. Normalize independent internal copies in buildRRule, comparing full Weekday (day plus ordinal); preserve order where observable. Do not mutate caller slices.

## Prior art

The issue briefs and docs/repro/05-07-backlog record probes of teambition PR66 head e7bed74dc0 and PR68 head d898b57a07. PR66 elapsed-time generation breaks RFC fold semantics and BYHOUR=9; reject that algorithm. Its Sydney transition fixtures are useful, but expected values come from RFC rules. PR68 rejects legal TZID+UTC UNTIL and serializes invalid local UNTIL; reject both. Its TZID-on-Z refusal is reusable as an idea; credit retained ideas/tests in the relevant commits. This research does not claim an independent new run of either patch probe.

## Risks and conditions

- Generated DST defects require recurrence candidates inside a gap or fold; ordinary scheduling in DST regions hits these annually, including daily rules. Lord Howe's half-hour change prevents a hard-coded one-hour solution.
- Floating defects require parsed no-Z input; common integrations use these forms. Wrong UNTIL type changes the endpoint by an offset, often silently. Order-sensitive validation would miss inline DTSTART after UNTIL or shared set DTSTART.
- Repeated BY values are valid and plausible machine-generated input; expansions and COUNT are vulnerable. Mask selectors may already behave as sets but need regression coverage; BYSETPOS signs and weekday ordinals matter.
- Golden comparison must distinguish intended changes from arbitrary drift and must cover String as well as instants. Fixing COUNT moves later results, so classification cannot check only individual changed timestamps.

## Historical provisional questions/choices — superseded

Keep Go1.16, no new dependencies/API fields; preserve original BY text/list order; one WP per issue; retain legacy EXRULE ignore semantics after required form validation; use private tagged date representations only if value-copy and programmatic constructor controls prove behavior. Independent post-plan review precedes tasks. If private representation cannot meet compatibility without unsafe machinery, stop as a genuine scope decision before implementation.

## Historical initial plan decisions (2026-10-25) — superseded

- Private tagged location transport selected provisionally: exact public ROption layout is preserved, unlike added fields; value-copy fidelity avoids unsafe/global registries. Implementer must prove location-observability caveat and InLocation behavior, explicit gap cadence, Set and setter handoffs. Actual code must not leak marker TZID text.
- Stream valid candidates before BYSETPOS rather than filtering selected results (wrong positions) or materializing every annual candidate (unbounded practical cost). Keep bounded requested-position/tail state.
- Sequence the three issues rather than parallel edits to shared parser/helper/build code. This is an ordinary implementation choice, not a changed three-WP scope. No fourth integration WP.
- No dependency decision/install: supply-chain challenge controls not applicable. No active SPDD/REASONS activation; generic prompt block does not activate it.
- Prior-art patch inspection supplied independently by root confirms PR66 spring count3 scenario valid and fold third instance should03AEST, not02AEST. PR68 invalid floating UNTIL acceptance and suffix check rejected. Keep credit on retained fixture/idea commits.

## Historical independent challenge (2026-10-25) — resolved by approved revisions below

The independent gpt-6-astra review found a concrete counterexample to the provisional name-based transport plan: an explicit floating InLocation caller can supply `time.FixedZone("America/New_York", -5*3600)` or custom TZif data using an IANA-looking name. Reloading that name chooses a different zone object/rule set and changes occurrences after a real IANA DST transition. Location.String is not a lossless identifier. The existing supplied *time.Location must remain authoritative. The provisional transport design is therefore not approved and must not be implemented as written.

Adding a private ROption field cannot preserve external positional literals. A private RRule field alone cannot preserve ROption.String/copy behavior. An unbounded global marker-to-original-location registry violates the selected bounded/no-registry design. Storing metadata in Wkst, subsecond bits or other unrelated fields creates ambiguity/collisions or mutates observable input and is rejected. The clean public Go APIs expose no general lossless clone/serialization of arbitrary *time.Location. Reading private time.Location internals, unsafe conversion or serializing an entire inferred timezone database is not an acceptable minimal repair. This is a genuine compatibility/design tension requiring an explicit policy decision if no clean representation is found; do not claim the existing plan meets all constraints.

The review also identified unconditional BYSETPOS candidate enumeration as a CPU regression. Gap-free periods must keep arithmetic selection. Transition periods need counts/ranking of removed local-time ranges, rather than building every Cartesian candidate. See revised plan section; compatibility is a substantive gate, not a waiver.

## Approved revisions (2026-10-25; supersedes earlier provisional choices)

Kent explicitly selected private ROption metadata and targetv1.9.0 (decision01M5TFFVFY5YXYPVD4JN9SS90M). Preserve signatures/keyed callers; disclose the external positional-literal source break and keyed migration. Immutable descriptors hold exact supplied *time.Location, nominal fields, lexical kind and per-part time snapshots; replacing a public time invalidates only its own descriptor. No marker timezone, registry or name reload. This resolves B2's design tension through an explicit authorization, not a source-compatibility claim. B1 design now preserves arithmetic fast paths and ranks around sparse gap holes; no unconditional product scan. Re-review remains required.

Kent removed issue7 after contributor fix76c646dde8d33ff7c3325df0613b31b058243bea landed on private main. Decision01M5TFJAPGE9MN2AWDG1K4EETF records scope limited to issues5/6, one WP each. Main was merged into fix branch. Contributor behavior is inherited golden baseline, neither assessed nor extended. New pristine baseline `/tmp/rrule-backlog-v183-plus-duplicates` passed modern go test ./... (.383s) and direct Go1.16.15 go test ./... (.345s), root report. Earlier three-issue/exact-layout/marker descriptions above are chronological investigation, not current design.
