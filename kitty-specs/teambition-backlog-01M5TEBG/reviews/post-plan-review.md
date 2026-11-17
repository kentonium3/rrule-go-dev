# Independent post-plan review

- Reviewer: **gpt-6-astra**
- Date: **2026-10-25**
- Implementation/planning model: gpt-6.1-sol; independent-model review, not a degraded same-model check.
- Reviewed: full `spec.md`, `plan.md`, `research.md`, `data-model.md`, `quickstart.md`; relevant construction, iteration, parsing, serialization, setters and Set code; issue briefs 05–07, runbook and AGENTS.md.
- Requested identities: specification `250222c`, plan `6fe979f`. Working checkout advanced to `551334b580f4a3b98545bf48b11eac2ee73679ef` during review; the reviewed plan includes its mandatory Go 1.16.15 validation amendment. Initial working status was clean.
- Scope: design review before tasks. No implementation, test execution, Git mutation, or public action. This review file is the only authored file.

## Verdict: CHANGES REQUIRED — do not generate tasks yet

The RFC corrections, three serial work packages, source-layout constraint, baseline comparison, red-first testing and independent review checkpoints are coherent. Two unresolved design issues block approval. B1 is an ordinary plan correction. B2 requires a concrete viable representation or an explicit compatibility decision; the present marker design does not establish the promised InLocation behavior.

## B1 — Bound BYSETPOS work, not only its retained memory

**Priority:** high. **Locations:** plan.md:21 and :39; current rrule.go:642–674.

The plan routes BYSETPOS through enumeration of every valid candidate in the period and retains only selected positive positions/a bounded negative tail. This bounds memory but changes an arithmetic selection into a potentially enormous scan even when there are no gaps. The current implementation indexes the filtered day list and sorted time list directly.

Concrete accepted case: a YEARLY rule with BYMONTHDAY=1 through 31, BYHOUR=0 through 23, BYMINUTE=0 through 59, BYSECOND=0 through 59, BYSETPOS=-1 and COUNT=1 in UTC. A synchronized DTSTART of 2026-12-31 23:59:59 selects that instant. The proposed period stream visits 31,536,000 timestamps to obtain it; current selection examines at most the annual day list plus the already-built 86,400-value time list. The same problem affects monthly/daily dense time sets and recurring calls. This is realistic with machine-generated selectors, and conflicts with the service-resource guardrail. A fast civil resolver alone does not remove the unnecessary product enumeration.

**Smallest correction:** explicitly retain direct rank/index selection for UTC/fixed-offset periods, and skip whole regular days/blocks in zoned periods. Account for missing positions only where a transition produces validity holes. Selecting a bounded number of requested positions must not require resolving every regular candidate. Preserve negative positions and overlapping-position deduplication. Require a dense-product positive/negative BYSETPOS resource regression or benchmark against baseline, plus transition fixtures proving that gaps are removed before ranking. No full annual candidate materialization or unconditional annual candidate scan.

Also remove permission to lose the existing UTC horizon optimization merely because parser metadata uses a non-UTC marker. Default-floating numerical evaluation that decodes to UTC must retain the effective UTC bound; prove this with the existing never-matching short-UNTIL controls. Ordinary non-UTC conservative fallback remains valid.

## B2 — A location name and current offset cannot transport an arbitrary caller location

**Priority:** high. **Locations:** plan.md:49–53; data-model.md:7 and :13; current str.go:173–176 and :224/:232.

The plan promises unchanged explicit InLocation evaluation while replacing its time.Location with a private fixed-offset token containing an evaluation-location name/offset. Decoding that name cannot recover the exact caller location. Go accepts arbitrary names for FixedZone and LoadLocationFromTZData; names are not unique identifiers for location behavior.

Concrete control: pass `time.FixedZone("America/New_York", -5*60*60)` to `StrToROptionInLocation`, with floating `DTSTART:20260307T090000` and `RRULE:FREQ=DAILY;COUNT=3`, then call NewRRule on the returned options. Existing occurrences are March 7, 8 and 9 at **14:00Z**. If token decoding reloads America/New_York, March 8 and 9 become **13:00Z**, although the supplied location has no transition. If decoding instead always retains only the initial fixed offset, the ordinary IANA America/New_York input breaks in the opposite direction. Comparing a few sampled offsets is not an exact discriminator. LoadLocationFromTZData with a colliding name supplies the same counterexample for caller-defined transition data.

This is a requested occurrence-compatibility failure, not a demand for general time.Time.Location pointer equality. Custom fixed offsets are a normal supported Go API use; colliding names are less common but require no malformed input. The current golden corpus does not cover this control.

**Smallest required correction:** specify and prove how the exact supplied evaluation location survives copied ROption values and NewRRule without reconstructing it from its name. Add paired custom-fixed-location and ordinary-IANA controls across the same transition, with and without UNTIL, and through Set/setter handoffs. Keep the public representation caveat precise. If fixed-offset markers remain, their effect on time.Time arithmetic should be disclosed; that observation alone is not this blocking finding.

No viable complete solution is established by the present plan. A process-global per-rule/location registry would violate its lifetime/resource constraints; unsafe or reflective reconstruction of Go's private timezone tables is not an acceptable small fix. Transporting DTSTART metadata in another option field, such as a tagged zero/nonzero Until, needs particular caution: legitimate replacement of Until can silently erase DTSTART semantics even though DTSTART was untouched. That alternative is not approved without a coherent mutation contract and evidence.

The user permits **new exported identifiers when a fix needs them**, with a plan justification and minor release. It would be wrong to characterize every exported addition as forbidden. However, adding either an exported or unexported field to ROption breaks an existing external positional composite literal. A minor-release label does not satisfy the simultaneous source-compatibility rule. A new exported type/function by itself is additive, but does not automatically make existing StrToROption/copy/NewRRule paths retain the missing metadata. If no source-compatible representation meets the existing requirements, present this precise tradeoff for a scope decision instead of silently weakening compatibility or claiming an additive field is source-compatible.

## Findings resolved by this plan

- Explicit gap DTSTART retains a nominal recurrence anchor separately from its resolved first instant; direct Go time values remain authoritative. Implementation tests should include MINUTELY at a gap anchor, where the resolved first instant can coincide with a later generated candidate, to prove no duplicate or backward emission.
- Floating/DATE serialization exceptions, absent-DTSTART fragments, inline DTSTART, and the unsynchronized brief fixture are correctly distinguished from complete component guarantees.
- Validating every legacy EXRULE before retaining historical ignore behavior is an explicit compatibility policy, not an exclusion engine. Last-rule selection remains separate from per-line validation.
- Original BY slices and serialization survive internal deduplication; weekday ordinals remain part of equality.
- Serial helper ownership avoids competing parser/DST implementations. Golden provenance, changed-row classification, minimum-Go execution, and per-defect red evidence are specified.

The normative interpretation was checked against [RFC 5545](https://www.rfc-editor.org/rfc/rfc5545.html), especially §§3.3.5, 3.3.10 and 3.8.5.3. No new RFC issue supersedes the two implementation-design blockers above.

## Re-review gate

Revise the bounded-selection design and resolve the location-transport contract in the plan/spec before tasks. Record a concrete proof or explicit scope decision for B2; prose asserting that a token preserves a location is insufficient. Re-review the changed design and associated acceptance controls. This verdict does not authorize changing public compatibility policy, creating implementation tasks, or beginning implementation.

### Checkpoint update

Root reports on 2026-10-25 that the driver also found no clean arbitrary-location marker solution under the existing constraints. Root has requested a user scope decision between permitting private ROption metadata with the external positional-literal break explicitly documented for v1.9.0, and retaining strict positional compatibility while stopping for redesign. That decision is pending; this reviewer does not infer approval. The driver is revising BYSETPOS toward gap-range rank selection. Those edits have not yet been re-reviewed. The verdict remains **CHANGES REQUIRED / COMPATIBILITY DECISION PENDING**, and tasks remain gated.
