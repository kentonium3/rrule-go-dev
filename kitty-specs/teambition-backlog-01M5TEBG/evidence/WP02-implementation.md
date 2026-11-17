# WP02 implementation handoff

Implementation model **gpt-6.1-sol**, actual date **2026-10-25**. Independent review model must be **gpt-6-astra** and is root-owned; this handoff does not claim approval. Allocated lane-b `~/repos/rrule-go/.worktrees/teambition-backlog-01M5TEBG-lane-b`; dependency HEAD verified clean at approved WP01 `ef373a5872dba30063a6a706275a28e46508af77`. Implementation commit `62b2527458d7790bf277bfc23245be56093ce195` (`fix(WP02): preserve and validate recurrence time forms`). Final documentation-only head recorded below. Claims, status, mission evidence, accept and final delivery remain root-owned. No push, PR, public mutation, final merge, sibling edit, installation/host change or extra agent occurred.

## Fix and boundary

Root cause: parser and serializers discarded UNTIL lexical form, inferred DTSTART form from a location name, and validated/attached shared anchors only sequentially. The implementation gathers lexical DTSTART/UNTIL, validates their forms after collection, and carries WP01's existing immutable descriptor transport through option copies and Set handoffs. No rrule.go or civil_time.go modification is needed.

`strToROption` is a private shared-anchor adapter. It supports either line order and inline attribute order, validates UTC/TZID versus floating UNTIL and DATE mismatches with named errors, and checks an inline start against a separate/shared anchor. Canonical VALUE=DATE and VALUE=DATE-TIME parsing is deliberately narrow; TZID-on-DATE/Z is refused independently. Form-aware formatting consults per-part valid snapshots, using nominal text for parsed gap anchors. Explicit floating gap UNTIL is resolved using the pre-gap offset, reusing WP01's civil resolver; fragments without DTSTART keep legacy local normalization and UTC-Z output.

Sets collect shared DTSTART before RRULE processing and retain a private descriptor. Each RRULE and legacy EXRULE time part is checked before last-wins/ignore behavior. EXRULE still has no evaluation engine. Its unrelated properties remain ignored (explicit control), avoiding broader validation. Internal `Set.dtStart` carries descriptors, while public time-only `DTStart` clears its part. `Set.RRule` preserves the adopted parsed rule or transfers the stored parsed shared anchor into a rule with no explicit start. Descriptors remain immutable; exact caller FixedZone/TZif pointers are never reloaded by name. Programmatic constructor/setter replacement behavior remains per-part and deliberate.

Scope remains public kentonium3/rrule-go#6 dependent on approved#5. Contributor#7 at76c646d is inherited baseline only. No new exported identifier/dependency, manifest or Go minimum change. Private ROption metadata and the external positional-literal break were already approved forv1.9.0; release notes specify keyed migration. MIT LICENSE/README notice unchanged.

## Files and narrow ownership adjustments

Owned production: str.go and rruleset.go. New tests str_timeforms_test.go and external-package str_timeforms_api_test.go. New docs/releases/v1.9.0-timeforms.md, golden/{README.md,corpus-forms.json,baseline-forms.json}.

Authorized out-of-map edits:

- str_test.go: old mixed EXRULE DATE UNTIL migrated to conformant `20180520T040000Z`; its original invalid input remains an explicit named rejection in new tests.
- rruleset_test.go: root explicitly approved migration of LA setter control from floating `20181117T235959` to equivalent UTC `20181118T075959Z`, plus checking parse error before dereferencing. Original invalid input retained in rejection tests.
- golden/main.go: optional FixedZone/custom minimal TZif and copy/Set/setter/programmatic actions extend the same existing public-API driver. Old fixtures follow the same path. Exactly the same driver is compiled against baseline and current modules, without editing baseline or committing module replacements.
- golden/corpus-dst.json:19 individual current-row updates relative to approved WP01:17 nominal-gap String-only changes, literal087 floating String, literal088 forbidden TZID/floating UNTIL. All other current rows and all approved WP01 occurrence lists are unchanged, except literal088 deliberately becoming an error. Full delta proof in logs. Old baseline-dst.json bytes remain untouched.

## Red/green evidence

`/tmp/rrule-WP02-logs/red-commands.md` records actual commands/cwd/state. Main red and additional red were both on untouched approved ef373a5 source before any production edit. Logs red-timeforms.txt and red-additional.txt exit1; legal controls pass, and no already-correct control is misrepresented as a defect. Coverage spans all12 form pairs; six production parser adapters; canonical and forbidden parameters; line/attribute/shared-start placement; ignored/overwritten rules; floating/DATE text and round-trips; immutable per-part copies/setters; exact FixedZone/IANA/customTZif locations; public keyed API; parsed gap text and floating gap UNTIL; combined gap/BYSETPOS/inclusive UNTIL. Tests assert production outputs/instants/location pointer identity, not only helper state.

A supplemental conflicting separate-plus-inline DTSTART case was added and shown red on the initial uncommitted parser fix before its narrow anchor check. red-separate-inline.txt exit1; standalone adapters fail, already-correct set adapters and matching-inline controls pass. This chronology is explicit, not claimed as an untouched-baseline red. Final green-focused.txt runs TimeForms and all WP01 DST tests, exit0. gate-modern.txt is the final full-suite pass (0.505s) on the final source, followed by all gates below. The later parsed dense/horizon assertions are compatibility/resource controls, not invented new defects.

## Golden comparison

Existing811-row DST corpus: **116 changed,0 unexplained,0 timeouts**. Its pinned baseline SHA256 remains **68fbd11d9cf3ca187dc473a543876196c23552ea5c2bb1b45f1ca561f75bd226**. Source provenance remains304 static options included/39 individually listed dynamic options skipped,92 string literals,392 transition matrix,23 targeted rows. Skipped dynamic fixtures still execute in the full suite. Exact19-delta proof relative to approved WP01 current output is golden-dst-delta-proof.json.

New158-row form corpus: **141 changed,0 unexplained,17 unchanged controls,0 timeouts,0 skipped**. Deterministic matrix/source descriptions are in corpus-forms.json provenance. Baseline SHA256 **1cc3299bf76264f6bcaf653ed3a6b2509765c62c4639880b081775667b1ae934**. The same newly extended driver compiled against pristine76c646d generated baseline-forms.json, then was rerun into a temporary file and cmp proved byte-identical reproducibility. All custom-location/form/copy/setter/default-fragment and integrated rows have explicit individual rationale and exact current-row hashes. No blanket allowlist or rebaseline.

Current driver SHA256 **ef590e56baca5005ad173824a905225cb4821bfa14261a97c4cb6d8c9ea19b89**. Built binaries `/tmp/rrule-WP02-golden-current` and `/tmp/rrule-WP02-golden-baseline`. Raw current outputs `/tmp/rrule-WP02-current-{dst,forms}.json`; full classified row logs golden-{dst,forms}-final.txt (exit0). README provides commands. Sensitivity checks in negative-{forms,dst}-{empty,corruption}.txt all correctly fail exit1: baseline self-comparison has0 changes; a corrupted temporary unchanged UTC row yields exactly1 unexplained change. No pristine baseline was corrupted for these checks.

## Integrated gates

All logs reside under **/tmp/rrule-WP02-logs**. Companion .command.json files pin exact argv/cwd/HEAD, source SHA256 and actual elapsed wall time for gated runner commands. Modern toolchain Go1.25.14; direct minimum Go1.16.15. Separate caches /tmp/rrule-go-cache and /tmp/rrule-go116-cache.

- `GOCACHE=/tmp/rrule-go-cache go test -count=1 ./...`:exit0, package0.505s, gate-modern.txt.
- `GOCACHE=/tmp/rrule-go116-cache /tmp/rrule-go116-toolchain/go/bin/go test -count=1 ./...`:exit0, package0.409s, gate-go116.txt.
- `GOCACHE=/tmp/rrule-go-cache go vet ./...`:exit0, zero diagnostics, gate-vet.txt.
- `gofmt -l .`:exit0, zero paths, gate-gofmt.txt.
- `GOMAXPROCS=2 GOCACHE=/tmp/rrule-go-cache go test -run=^$ -fuzz=FuzzIterate -fuzztime=60s`:exit0, actual fuzz60.077s,2,980,037 executions, gate-fuzz.txt. Fuzz setup/build excluded from actual fuzz duration.
- Resource command in gate-resource.command.json: parsed six-adapter dense/horizon tests plus existing TestDSTDenseAnnualArithmetic and floating short-UNTIL control; BenchmarkDSTAnnualBysetpos1s. Exit0,1,034,487ns/op,2,226,026B/op,47allocs/op. Dense selection proves ranks1,366,-366,-1 across31,536,000 nominal candidates through all six parser adapters, with no annual candidate enumeration. Existing WP01 resource behavior remains bounded.
- `GOCACHE=/tmp/rrule-go-cache go test -race -count=1 ./...`:exit0, package148.174s, actual runner149.316s, gate-race.txt and companion command metadata. No retry, timeout waiver or focused-only substitute.
- `git diff --check`:exit0 before commit.

## Agent cwd fault and safe recovery (D-3 evidence)

One combined exec accidentally used pristine baseline cwd while attempting a diagnostic text replacement and focused test. The replacement target was absent, so str.go was rewritten with identical bytes; only mtime changed. The focused command matched no WP02 tests and its output is explicitly **INVALID**, preserved in INVALID-focused-wrong-baseline-cwd.txt and never counted as a gate. Root was informed immediately. Exact command/cwd/output/state and safe recovery are in fault-cwd.md.

Before correct lane repair, every existing baseline root file was compared to git show76c646d; all15 Go/module/LICENSE files were byte-identical. baseline-integrity-after-cwd-error.txt records per-file hashes. A later full required-file enumeration also proves all15 canonical root Go/module/LICENSE files are present and unchanged (baseline-final-integrity.json). Correct lane diagnostic replacement and genuine TimeForms+DST test were then run and preserved separately. This was an agent cwd mistake, not a Spec Kitty fault or substantive gate bypass. No source restoration from altered behavior was needed; no invalid evidence was reused.

Profile implementer-ivan and action context were resolved through prefixed Spec Kitty commands. Requested code-review-checklist and terminology-canon sections are not authored in the charter; CLI reported that advisory without mutation. No new terminology policy or host alteration was inferred.

## Golden determinism correction before commit

The first final reproducibility cmp failed for exactly one new form fixture, shared-ignored-unknown-pos1: its COUNT3 lacked a first-line shared start recognized by the old parser, exposing time.Now output. The initial text replacement missed the DTSTART inserted between RRULE and EXRULE in that middle-placement variant. Original baseline/manifest and exact differing rows are preserved as pre-determinism-* and golden-forms-determinism-failure.json; the initial cmp failure is not a pass. Corrected only this control to explicit UNTIL20201230T090000, matching its other placement variants. The same unchanged driver against unchanged76c646d generated corrected baseline-forms.json, and a separate regeneration/cmp now passes exactly. Final158/141/0 counts remain unchanged. See golden-forms-determinism.md and corrected/repro/cmp command logs. No production/test code or old811 baseline changed, and no broad test gate was rerun.

The reproducibility context is2026-10-25: late-shared-start baseline controls may internally default to time.Now, but each has explicit2020 UNTIL and therefore deterministically yields no baseline occurrence in this recorded context. README states this convention explicitly; no claim that the old fallback itself is fixed is made. A documentation-only follow-up records that context after the implementation safe-commit.

## Final committed head

**Final head `0f7154e7bdb98b7e2de3e22ed2db78b8bfe0ce53`**, documentation-only golden context clarification following implementation62b2527458d7790bf277bfc23245be56093ce195. Both safe-commit calls succeeded with explicit lane-b destination and scoped files; safe-commit.json and safe-commit-docs.json preserve results. Working tree clean. final-source-and-gate-proof.json verifies every root Go source/test SHA256 from minimum-Go, vet, full race, fuzz, resource and formatting command metadata still matches the final committed head. The final docs commit changes only README context; no production gate evidence is stale. Independent gpt-6-astra review and all status transitions remain root-owned.
