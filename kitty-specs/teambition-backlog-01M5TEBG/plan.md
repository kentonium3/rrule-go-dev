# Implementation Plan: Teambition backlog

**Branch**: `fix/teambition-backlog` | **Date**: 2026-10-25 | **Status**: Revised after explicit metadata decision and two-issue scope; independent re-review pending | **Spec**: [spec.md](spec.md)

## Summary

Repair civil-time candidate resolution and lexical DTSTART/UNTIL preservation at their existing seams. Keep public signatures/keyed callers and Go1.16, use only the standard library, and prove the behavioral boundary with red-first fixtures and a reusable v1.8.3-plus76c646d comparison. Exactly two WPs own public issues5 and6. Kent explicitly approved private ROption parse metadata and a v1.9.0 target; external positional ROption literals become invalid. No new exported identifier. Contributor issue7 fix76c646d is inherited baseline only, with no assessment or extension in this mission.

## Engineering Alignment

Kent supplied the implementation boundaries and authorized ordinary choices autonomously. Questions concerning API growth, dependencies, baseline, publication, branch target, minimum Go, models and acceptance are already answered by the user/runbook. Current branch, planning branch and consolidation target all equal `fix/teambition-backlog` (canonical create/setup-plan output; match=true). Final delivery is a private PR from that branch to main, not a consolidation/merge onto main. Step0 charter PR12 was separately authorized and merged at 6ac755d.

## Technical Context

**Language/Version**: Go1.16 minimum; host validation Go1.25.14 or installed latest; preserve go.mod go1.16.
**Primary Dependencies**: Go standard library only; time/IANA zone data supplies transition offsets. No supply-chain installation or dependency change.
**Storage**: Existing Go values and committed mission/test fixtures; no database or service.
**Testing**: Plain Go production-path regressions, race/vet/format checks, current FuzzIterate for60s, v1.8.3-plus76c646d golden harness. Fuzz test's existing version build constraints preserve minimum Go.
**Target Platform**: Existing supported Go platforms; linux/amd64 host, with Go1.16-compatible production code and tests.
**Project Type**: Root-layout library; rrule.go, str.go, rruleset.go and root *_test.go.
**Performance Goals**: Preserve streaming iteration and arithmetic BYSETPOS selection for gap-free periods. Civil resolution uses bounded zone/field work per ordinary emitted candidate; transition-period BYSETPOS ranks across removed local-time ranges using per-day counts/binary searches, never unconditional year×time candidate enumeration.
**Constraints**: Existing public signatures and keyed callers, no new dependencies/exported identifiers; approved private ROption metadata-field exception makes external positional literals invalid and targets v1.9.0. Preserve supplied BY slices and unrelated serialization. Avoid Go APIs newer than1.16 (notably Time.ZoneBounds, slices and generics).
**Scale/Scope**: Two fixes, two WPs; transition fixtures Sydney/NY/LordHowe, all seven frequencies plus UTC controls. No new RFC component/exclusion engine.

## Charter Check

Pre/post design: RRULE-001 applied with Kent's explicit private-field exception (decision01M5TFFVFY5YXYPVD4JN9SS90M), keyed/signature compatibility, no dependencies and RFC-limited golden differences; RRULE-002 by red-first tests and gpt-6.1-sol implementation/gpt-6-astra independent reviews with per-WP date; RRULE-003 by private destination/no public action; RRULE-004 by two WPs/full arc/D-3 after explicit narrowed scope. The approved ROption layout exception and v1.9.0 target are recorded rather than claiming universal source compatibility. LICENSE and existing README modified-copy notice remain. Supply-chain controls: no installation, dependency, manifest or lockfile change; dependency challenge not applicable. Independent post-plan review is still required before tasks; this document alone does not pass that checkpoint.

## Design

### Civil time resolution and explicit DTSTART

Add a small unexported civil resolver, in a separate root Go file if that keeps rrule.go readable. Its inputs are requested year/month/day/hour/minute/second and an actual location; output identifies zero valid instants (gap), earliest valid instant (normal/fold), and pre-gap interpretation for explicit DTSTART. Re-derive expected instants from RFC and known zone offsets, never from the helper being tested.

Derive candidate offsets from Go's zone lookup around the target civil day with a bounded window spanning both sides of the supported transitions; construct candidate instants from nominal UTC fields minus each candidate offset, convert back to the actual location, and retain exact civil-field matches. Pick the earliest match. A no-match is a generated gap and is discarded. For explicit gap DTSTART use the offset before the transition. Account for Sydney, NY and Lord Howe; no hard-coded one-hour adjustment or elapsed-time sub-daily stepping. Keep a normal/UTC fast path. The implementer must state the chosen bounded-search invariant and test both hemispheres/half-hour shifts; never silently widen support claims to arbitrary historical timezone discontinuities.

Preserve original civil fields for parsed gap DTSTART before Go normalization. Store nominal fields, exact actual/evaluation *time.Location, form and snapshots of parsed public time values in immutable private ROption metadata. Public time.Time values use the actual supplied location and RFC-resolved instant; there are no marker locations or name reconstruction. RRule may privately keep nominal anchor fields separate from the resolved lower bound after reading valid metadata. First explicit DTSTART counts once when aligned with the recurrence; subsequent generated candidates retain nominal hour/minute and use gap filtering. NewRRule supplied an ordinary Go time treats its instant/fields as authoritative, not guessed original input. In particular a programmatic later-fold DTSTART (constructed via a UTC instant converted into the zone) is included exactly once when aligned with selectors, even though normal generated ambiguous civil values resolve to the first offset. Add constructor and public DTStart-setter regressions proving this initial-instant exception without emitting both folds for ordinary generated values.

Route both generate output paths through the resolver. For no-BYSETPOS keep direct streaming. Preserve the existing arithmetic BYSETPOS selection fast path for gap-free periods, including UTC and fixed zones. For periods with gaps, count valid candidates per selected civil day by subtracting timeset entries inside each nonexistent local range; use binary searches over the already-sorted timeset to find gap boundaries. Resolve each requested signed ordinal by day counts and rank/select across removed ranges, then construct only selected candidates. Fold values remain one civil candidate and resolve to the earliest instant. Do not enumerate the year×time Cartesian stream merely to select positions. Verify with large UTC/normal-period BYSETPOS fixtures and mixed transition-period ordinal controls; record bounded work and a benchmark comparison. Preserve inclusive UNTIL and lower-bound handling. Retain the existing effective-UTC short-circuit UNTIL bound, including default-floating evaluation that uses UTC numerically; metadata must not make never-matching short-UNTIL cases scan to MAXYEAR. Ordinary non-UTC conservative fallback remains allowed. Test the existing never-matching short-UNTIL controls through parsed floating paths as well as UTC.

```text
calendar candidates → civil resolver → valid ordered period stream → BYSETPOS (if present)
                                 ↘ gap: discard                      ↓
explicit DTSTART → pre-gap/first-fold interpretation → first/lower bound → UNTIL + COUNT → output
```

### Parser form transport with the approved private ROption field

Kent explicitly selected “Allow private metadata; target v1.9.0” on2026-10-25. Decision01M5TFFVFY5YXYPVD4JN9SS90M authorizes one private immutable metadata field in ROption, preserving its existing exported fields/signatures/keyed callers while breaking formerly valid external positional ROption literals. Release notes must state that limitation and the v1.9.0 target; a minor label is not a claim that positional callers remain compatible. This replaces the rejected location-marker design.

Use a private pointer to immutable parse descriptors (or equivalent immutable value) carrying each part's lexical kind, exact supplied *time.Location, original civil fields, and snapshots of parsed Dtstart/Until values. Copying ROption safely shares immutable metadata. Keep the exact caller location pointer; never reconstruct arbitrary FixedZone/TZif behavior from Location.String or a sampled offset. No global registry, unsafe, reflection, unrelated-field encoding, sentinel timezone, new exported identifier or dependency.

Before String/build consumes a part's metadata, compare its current public time value with that part's snapshot including location identity (not Equal-instant alone). Replacing Dtstart invalidates only DTSTART metadata; replacing Until invalidates only UNTIL metadata. Retain the unaffected part and copy-on-replace descriptors when an internal operation changes interpretation. Never mutate shared metadata in place. An assignment of an identical public time value is observationally identical and may retain its metadata. NewRRule on ordinary manually constructed options has no parser metadata and retains current programmatic UTC-UNTIL serialization.

Bare RRULE fragments with no DTSTART retain legacy UNTIL parsing and UTC-Z String output (including DATE/floating UNTIL interpreted in supplied loc); do not attach explicit-start lexical serialization semantics until an explicit DTSTART is present.

Explicit floating default parsing stores numerical UTC without claiming a zone binding and serializes no Z/TZID. InLocation parsing keeps the exact supplied location and current evaluation convention while recording floating syntax separately. Pair custom FixedZone("America/New_York",-5h) and real IANA NY controls across the same spring transition, with/without UNTIL and via copied options/Set/setter handoffs; custom TZif same-name control confirms names never substitute for pointers. Preserve physical instants and supplied Location observability in parsed public time values; nominal gap fields stay private.

A standalone StrToDtStart returns only time.Time: apply correct §3.3.5 instant resolution there, but carry lexical/nominal metadata through private parser helpers when building options/sets. Do not add an out-of-band registry to preserve metadata through arbitrary detached time.Time copies. Internal shared-DTSTART handoffs use a private helper that carries immutable descriptors, rather than discarding them through the public time-only setter. Public DTStart/Until setter calls treat their supplied ordinary time as authoritative, replacing the relevant part's old parser descriptor while preserving the other part. Set's private storage may hold shared DTSTART descriptors; no new public setter/API is needed. Tests must pin these replacement semantics and copies.

Parse lexical DTSTART and UNTIL into local descriptors first, validate after collecting attributes/lines, then construct ROption. Support canonical VALUE=DATE and VALUE=DATE-TIME narrowly, rejecting TZID on DATE or Z. Preserve old bare DATE syntax as a compatibility input. Validate all enumerated form mismatches with named errors. In two-line RRule parsing accept DTSTART/RRULE order independent; inline DTSTART is an existing extension and may precede/follow UNTIL. In sets collect shared DTSTART before processing rules; validate every RRULE/legacy EXRULE before last-wins/ignore behavior. A supplied inline DTSTART cannot evade shared DTSTART type/UNTIL checks or silently replace the shared semantic anchor. Do not implement EXRULE evaluation, change general unknown-property policy, or redesign multiple-RRULE behavior.

### Golden compatibility harness

Current baseline is v1.8.3 plus the inherited contributor duplicate-selector fix, private main76c646dde8d33ff7c3325df0613b31b058243bea. Local pristine snapshot: `/tmp/rrule-backlog-v183-plus-duplicates`. Do not evaluate or extend contributor issue7 completeness; compare its inherited behavior as baseline.

Commit a standalone ignored program or equivalent test utility under docs/repro/05-07-backlog with a deterministic JSON/text manifest. Run exactly the same driver in a temporary copy of baseline code from private commit76c646dde8d33ff7c3325df0613b31b058243bea (v1.8.3 plus contributor fix; verify blobs against this exact baseline) and current code. Do not fetch/push/tag the public repo. Avoid a replace directive committed to go.mod. Driver imports the module normally in each tree and has bounded iteration/timeouts for existing pathological rules.

Corpus sources: extract valid rule fixtures from existing *_test.go plus docs/repro/matrix and edge-cases; include explicit fixed DTSTART for otherwise-now rules; add a deterministic matrix all seven frequencies × UTC/NY/Sydney/LordHowe × starts in the week around each relevant transition (UTC uses corresponding dates). Take first16 values or exhaustion, and String/error status. Include explicit near-boundary starts/intervals for sub-daily frequencies and aligned day/month selectors for coarser frequencies: the first16 MINUTELY/SECONDLY values from midnight will otherwise miss a transition. Use longer targeted sequences when needed to cross the affected boundary. Add both issue fixtures, time-form order variants and inherited duplicate-selector controls. Save corpus provenance/count and a stable baseline output artifact/hash. Before claiming coverage, report how many source fixtures were included/skipped and why. Known nonterminating/max-bound controls must use bounded iterator calls and explain exclusion or timeout rather than silently disappear.

Diff every row and classify each difference against explicit categories: generated gap removal/fold-first and downstream COUNT/BYSETPOS movement; explicit gap/fold DTSTART resolution/cadence; floating/DATE serialization preservation; enumerated forbidden-form errors. Preserve original BY String. Each changed row cites a regression and RFC rationale. Include deliberately changed rows as sensitivity controls so an empty/incorrect comparison cannot pass. No broad ignore or blanket DST-zone allowlist. Integrated cross-fix cases combine DST, BYSETPOS and floating/date parsing; inherited duplicate inputs are compatibility controls only.

## Project Structure

```
rrule.go                 # construction and generation
str.go                   # lexical parsing and serialization
rruleset.go              # existing set handoff/representation if needed
civil_time.go            # proposed unexported resolver/transport helpers
*_test.go                # regression files per issue; compatibility API compile fixture
fuzz_test.go             # existing FuzzIterate; retain minimum-Go constraints
docs/repro/05-07-backlog/ # reusable golden harness and deterministic corpus
kitty-specs/teambition-backlog-01M5TEBG/
  spec.md, plan.md, research.md, data-model.md, quickstart.md
  reviews/, evidence/    # durable checkpoints, red/green, golden and gate evidence
```

No new endpoint or external interface is defined; meta.json records contracts=none with rationale. Existing exported Go API is the contract to preserve, verified by an external-package compile fixture using keyed ROption initialization and existing method signatures. The formerly valid positional ROption fixture is documented as the user-approved source break, not expected to compile in this mission.

## Implementation Concern Map

### IC-01 — Civil generation and DTSTART

- Purpose: solve public kentonium3/rrule-go#5 at shared candidate construction with bounded BYSETPOS filtering and exact explicit-anchor handling.
- Requirements: FR-001, FR-002, FR-007, FR-008; SC-001; NFR-001/002/003.
- Surfaces: rrule.go, private civil helper, narrow StrToDtStart parsing, dedicated DST tests; baseline golden harness/corpus.
- Dependencies: none. Owns helper names/contracts so parser work can extend them without conflicting independent designs.
- Risks: fold earlier instant, gap explicit-anchor exception, half-hour transitions, positional selection, all-frequency runtime cost.

### IC-02 — Parser lexical forms

- Purpose: solve public kentonium3/rrule-go#6 with form-aware validation/serialization using copied-value-compatible private transport.
- Requirements: FR-003/004/005/007/008; SC-002; NFR-001/002/003.
- Surfaces: str.go, private transport helpers, narrow rrule.go/rruleset.go handoffs, parser/set tests, golden corpus extensions.
- Dependencies: IC-01 helper design and explicit gap handling. Sequential implementation is preferred despite conceptual independence: metadata and StrToDtStart overlap.
- Risks: document approved positional-literal break; preserve keyed callers and exact InLocation pointers, stale-part detection, setters/Set overwrites, each shared RRULE/EXRULE checked, DATE canonical form.

## Verification and execution sequence

1. Independent gpt-6-astra post-plan review. Resolve findings in spec/plan before tasks. Root owns orchestration; gpt-6.1-sol is implementation model and gpt-6-astra review model, record each with date per WP.
2. Tasks create exactly two WPs in order IC-01→IC-02 (one lane is acceptable). Map all FRs with shared compatibility/release requirements explicitly owned; run analyze before implementation.
3. For each defect family author tests and save red output on untouched baseline/preceding-WP state. Pair forbidden/absence tests with same-fixture conformant controls. Independent parts of compound fixes must each have a red test. Inherited contributor duplicate-selector tests are baseline controls, not mission defects or extra acceptance work.
4. Implement scoped fix and release-note entry, green production-path tests, full go test plus focused tests, then independent reviewer. Keep complete review artifacts, actual model/date and code commit identity.
5. Integrated gates on consolidated fix branch: go vet ./..., gofmt -l . (zero output), go test ./..., go test -race ./..., FuzzIterate>=60s and golden comparison with zero unexplained differences. Run the verified temporary Go1.16.15 toolchain directly (`/tmp/rrule-go116-toolchain/go/bin/go test ./...`, cache `/tmp/rrule-go116-cache`) in addition to modern-toolchain gates. Root verified its official archive SHA256 and baseline passes. Do not change go.mod.
6. Strict accept, consolidate onto fix/teambition-backlog, independent integrated final gpt-6-astra review; address findings and rerun affected gates. Deliver private PR to main; no merge. Check actual origin visibility before any publication. No closing keywords for colliding issue numbers.

## D-3 and risk disposition

Preserve exact fault command/output/build/state before recovery; no blind retry of mutations. Known read-only .git boundary uses authorized escalated execution for writes. Capture substantive runtime faults in mission evidence and runbook before diagnosing/retrying. No install/source edits to Spec Kitty, host changes, bypassed gates, public actions or expanded scope. The private-metadata compatibility exception is explicitly authorized by decision01M5TFFVFY5YXYPVD4JN9SS90M and targets v1.9.0; no broader API break is authorized. Scope is explicitly narrowed by01M5TFJAPGE9MN2AWDG1K4EETF. Independent plan re-review still precedes tasks.
