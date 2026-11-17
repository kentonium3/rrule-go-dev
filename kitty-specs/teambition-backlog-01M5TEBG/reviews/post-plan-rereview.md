# Independent post-plan re-review

- Reviewer: **gpt-6-astra**
- Date: **2026-10-25**
- Planning/implementation model: gpt-6.1-sol. Independent-model review.
- Reviewed revision: `3b776c5405bb032d53edd1a1d49988f3b136b792`, including the working artifacts read immediately before that commit completed.
- Reviewed in full: spec.md, plan.md, research.md, data-model.md, quickstart.md. Also checked the explicit metadata decision and current AGENTS.md/runbook policy amendments. Rechecked the requested follow-up edits before issuing this verdict.
- Scope: only public kentonium3/rrule-go issues **#5 and #6**, two serial work packages. Contributor #7 at `76c646dde8d33ff7c3325df0613b31b058243bea` is inherited baseline; this review does not assess, extend, or endorse that contributor implementation.

## Verdict: APPROVED — post-plan gate satisfied

No blocking plan findings remain. The two initial blockers are resolved in the revised design and the explicitly approved compatibility scope. Task generation may proceed for exactly the two remaining work packages, followed by analyze and the planned implementation/review gates. This is design approval, not implementation acceptance or permission for public publication or final-PR merge.

## Resolution of prior blockers

**B1 — BYSETPOS CPU cost: resolved in plan.** Gap-free periods retain arithmetic selection. Transition periods subtract sparse nonexistent civil ranges from per-day counts and select signed ranks without enumerating the annual day×time product. The plan explicitly requires dense-product controls and benchmark evidence, including positive/negative positions and gap-affected ordinals. It also preserves effective-UTC short-UNTIL termination through parsed default-floating paths. Bounded allocation alone is no longer substituted for bounded selection work.

**B2 — exact supplied location: resolved by approved representation and scope decision.** Decision `01M5TFFVFY5YXYPVD4JN9SS90M` records Kent's explicit approval of private ROption metadata and a v1.9.0 target. Immutable descriptors retain the exact supplied *time.Location and nominal fields; ordinary public time.Time values retain their actual location. There is no marker location, name reload, global registry, reflection, unsafe conversion, or metadata hidden in unrelated fields. The FixedZone/IANA-name collision and custom same-name TZif controls now test the precise earlier counterexample.

The plan states the source break accurately: existing signatures and keyed callers remain compatible; previously valid external positional ROption literals become invalid. Release notes must document keyed-literal migration and v1.9.0. This is a user-authorized exception, not a claim that private fields preserve every source caller. No broader API break follows from this approval.

## Follow-up corrections verified before approval

1. Research clearly identifies the prior three-WP, exact-layout, location-marker and unconditional-stream proposals as historical/superseded. Its current approved decisions match the active plan.
2. Direct NewRRule and public DTStart setters receiving the later physical instant of a fold preserve that authoritative initial instant exactly once when aligned with selectors. Parsed ambiguous civil text still selects the first fold. Explicit constructor/setter regressions are required, separately from ordinary generated-fold rules.
3. Bare RRULE fragments without DTSTART preserve legacy floating/DATE UNTIL interpretation and UTC-Z String output. They do not accidentally acquire explicit-start lexical serialization semantics.
4. Per-part snapshots include location identity. Replacing Dtstart invalidates only DTSTART metadata; replacing Until invalidates only UNTIL metadata. Copies share immutable descriptors, with no in-place mutation. Public time-only setters replace the relevant semantics authoritatively; private Set/shared-start handoffs retain descriptors deliberately.
5. The golden baseline is the inherited v1.8.3-plus-contributor state at `76c646dde8d33ff7c3325df0613b31b058243bea`, with local snapshot `/tmp/rrule-backlog-v183-plus-duplicates`. Difference classifications no longer authorize new duplicate-selector fixes. Transition coverage explicitly avoids missing boundaries with the first sixteen sub-daily occurrences.

## Constraints carried into tasks and implementation review

- Exactly two WPs: civil-time/DTSTART first, parsing/serialization second. Carry shared helper ownership and all live FRs into tasks; do not resurrect retired FR-006/SC-003 or contributor assessment.
- Keep gap-anchor nominal cadence separate from resolved first instant, with COUNT, inclusive UNTIL, BYSETPOS and monotonic output checked through query and Set paths. Include a sub-daily explicit-gap anchor where a later generated candidate can share the initial resolved instant; the existing uniqueness requirement must be demonstrated, not assumed.
- State and test the civil resolver's bounded-search invariant. No hard-coded one-hour correction or unsupported claim of arbitrary historical-transition coverage.
- Preserve the custom-location, per-part mutation, copy, ordinary constructor, bare-fragment and public/private setter distinctions specified by the revised plan.
- Complete red-before/green-after evidence, dense BYSETPOS performance checks, the classified golden comparison, modern Go tests/race/vet/format/fuzz, and direct Go 1.16.15 validation. Record actual implementation/review models and dates for each WP.
- Deliver the authorized private PR to main and stop before merging it. No public pushes, comments, tags or releases are authorized by this review.

No implementation tests were run for this design checkpoint; this reviewer authored only this review file and made no Git mutations. The prior review remains historical evidence of the rejected design.

## Reviewed artifact fingerprints

SHA256 captured after the final corrective edits:

| File | SHA256 |
|---|---|
| spec.md | aa12c974c11f0594e4406b21bc30b9f34e966e5956de7dbe7d064f42c67cb2da |
| plan.md | 47146cac019f6a194232cbe573768279cd722248f87f5f3f452227c496a5d660 |
| research.md | 22d86d0aa2c71969a6e13df01c3a373b3e8d06095ffe4908f279ab4a9ef1bfb1 |
| data-model.md | 15fa1f57c6d1ea775a826f4dee06599acc2866361bf061a585c097ca5f7bf360 |
| quickstart.md | 1e61e78ca1a8ec7e968f7a62bc4f8dcda258256773c0f19b2a00aeddf904c69d |
