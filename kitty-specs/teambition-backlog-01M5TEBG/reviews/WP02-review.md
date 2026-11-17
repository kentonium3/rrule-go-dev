# WP02 independent implementation review

- Reviewer: **gpt-6-astra**, **2026-10-25**. Implementation model: **gpt-6.1-sol**, 2026-10-25.
- Exact reviewed head: **0f7154e7bdb98b7e2de3e22ed2db78b8bfe0ce53**. Production commit: **62b2527458d7790bf277bfc23245be56093ce195**. The final commit changes documentation only.
- Dependency: approved WP01 **ef373a5872dba30063a6a706275a28e46508af77**.
- Lane: `~/repos/rrule-go/.worktrees/teambition-backlog-01M5TEBG-lane-b`, clean before and after review.
- Official review invocation: `78d1a3e71d984c91bcfbdb93dc7012a4`. The generated prompt's implementer identity does not identify the actual reviewer; the independent identity above does. Root claimed this review and owns its canonical transition. No production, Git, workflow-status or public mutation was made by this reviewer.

## Verdict: APPROVE

No blocking correctness, compatibility or scope finding was identified. WP02 meets its approved time-form and integrated-gate requirements at the exact head above. Root may record the canonical approval and advance to the remaining mission gates. This is work-package approval, not final mission acceptance or authorization to merge the final private PR.

## Findings and scope assessment

The private parser collects DTSTART and UNTIL before checking their lexical forms. The accepted matrix is DATE/DATE, floating/floating, UTC/UTC and TZID/UTC. Numeric UTC storage does not turn a floating value into UTC syntax. TZID-on-DATE and TZID-on-Z are refused independently. These rules follow [RFC 5545 §§3.2.19, 3.3.5 and 3.3.10](https://www.rfc-editor.org/rfc/rfc5545.html); the primary RFC text was checked during review.

Separate/inline DTSTART consistency and shared Set collection prevent an inline override from bypassing the component anchor. All RRULE time parts are checked before historical last-wins handling; legacy EXRULE time parts are checked before historical ignore handling. Unrelated ignored EXRULE properties remain ignored. No exclusion engine or broad general-purpose validator was introduced.

Formatting uses valid per-part immutable snapshots. Copies, independent public-field assignments and public time-only setters preserve or invalidate only their respective descriptors. Internal Set transfers preserve explicit floating/DATE form and nominal gap cadence. Exact supplied location pointers survive; same-name fixed/custom zones are not reconstructed from their names. Bare fragments retain their historical local parsing and UTC-Z UNTIL serialization, as do ordinary programmatic constructors. Independent mutations may intentionally yield mixed forms; the approved plan does not require automatic constructor/setter pair validation.

The production changes are limited to str.go and rruleset.go. The approved WP01 resolver/iteration code is unchanged. The two old invalid parser fixtures were narrowly migrated to conformant UTC UNTIL controls, with original invalid inputs retained in rejection tests. Driver extensions are optional public-API location/action variants. Release notes disclose all specified rejection/serialization classes, retained PR68 attribution, the user-approved private ROption field and external positional-literal break, keyed-literal migration and v1.9.0 target. No new exported identifier, dependency, module-version change or contributor issue #7 extension was found. Baseline remains v1.8.3 plus inherited76c646d.

## Independent behavior evidence

The prepared probe, written before inspection of WP02 implementation, passes **296 checks** at the reviewed head. It covers all12 DTSTART/UNTIL pairs through six parser entry points and both line orders, error-part names, inline order, overwritten RRULE and ignored EXRULE validation, floating/DATE forms, location identity, metadata copies/setters and constructor/fragment controls.

Independent supplemental probes pass **19 checks**, covering:

- Explicit TZID=UTC still requires UTC UNTIL for RRULE and legacy EXRULE.
- Shared DATE rejects inline DATE-TIME or a different date; a matching inline DATE succeeds.
- A normalized physical gap display cannot replace the nominal shared DTSTART, even when the instants coincide.
- Parsing a shared DTSTART alone and subsequently installing a no-start rule preserves floating form, exact location and the expected recurrence. Public Set.DTStart intentionally replaces that form.
- New York floating daily02:30 over March8 2026 yields March7 07:30Z, March9 06:30Z and March10 06:30Z; String/reparse in the same evaluation location preserves this result.
- Clearing one public time removes only that part; the original option is unchanged.
- RDATE/EXDATE before a late shared TZID start use its location consistently.

The location controls distinguish real New York (March7-9 2026 at14:00Z,13:00Z,13:00Z), same-name fixed−05 (14:00Z each day), and same-name custom TZif+02 (07:00Z each day), with and without UNTIL. DATE local midnights are05:00Z,05:00Z,04:00Z. The unchanged independent WP01 R1 collision probe also passes all12 assertions on the integrated head.

Sources and raw outputs are durable under `evidence/WP02-review-probes.go.txt`, `WP02-review-probes.txt`, `WP02-review-extra.go.txt`, `WP02-review-extra.txt` and `WP02-review-WP01-control.txt`.

## Verification actually rerun

All commands below used lane-b unless explicitly stated otherwise; all exited0.

| Command | Result |
|---|---|
| `GOCACHE=/tmp/rrule-go-cache go run /tmp/rrule-wp02-review-probes.go` | 296 checks pass |
| `GOCACHE=/tmp/rrule-go-cache go run /tmp/rrule-wp02-review-extra.go` | 19 checks pass |
| `GOCACHE=/tmp/rrule-go-cache go run /tmp/rrule-wp01-rank-collision-review.go` | 12 assertions pass |
| `GOCACHE=/tmp/rrule-go-cache go test -count=1 ./...` | Fresh full-suite pass,0.509s |
| `GOCACHE=/tmp/rrule-go116-cache /tmp/rrule-go116-toolchain/go/bin/go test -count=1 ./...` | Fresh minimum-Go full-suite pass,0.397s |
| `GOCACHE=/tmp/rrule-go-cache go test -race -count=1 -run 'TestTimeForms\|TestDSTGapAnchor\|TestDSTAuthoritative' ./...` | Fresh affected-path race pass,1.096s |
| `GOCACHE=/tmp/rrule-go-cache go vet ./...` | Zero diagnostics |
| `gofmt -l .` | Zero paths |

The full suite includes parsed dense BYSETPOS arithmetic and effective-UTC short-UNTIL controls. The supplied full-race log passes at148.174s and the fuzz log records60.077s with2,980,037 executions. I inspected the actual logs and independently recomputed all16 source hashes from each of the full-race, fuzz, minimum-Go, resource, vet and formatting command records: every hash equals the reviewed head. Those complete gates were not redundantly rerun. No hook-based success is inferred from the workflow claim.

## Golden provenance and independent rerun

Built a new current executable from the reviewed driver, evaluated both corpora and ran their pinned-baseline comparators:

```sh
GOCACHE=/tmp/rrule-go-cache go build -o /tmp/rrule-review-wp02-golden docs/repro/05-07-backlog/golden/main.go
/tmp/rrule-review-wp02-golden -corpus docs/repro/05-07-backlog/golden/corpus-forms.json -out /tmp/rrule-review-wp02-forms.json
/tmp/rrule-review-wp02-golden -corpus docs/repro/05-07-backlog/golden/corpus-forms.json -baseline docs/repro/05-07-backlog/golden/baseline-forms.json -current /tmp/rrule-review-wp02-forms.json
```

Repeated the latter two commands for corpus-dst/baseline-dst with separate output. **Forms:158 rows,141 classified changes,0 unexplained,0 timeouts. DST:811 rows,116 classified changes,0 unexplained,0 timeouts.** Durable comparison outputs are `evidence/WP02-review-golden-forms.txt` and `WP02-review-golden-dst.txt`.

Additionally built the exact same current driver by absolute path from the pristine `/tmp/rrule-backlog-v183-plus-duplicates` cwd, evaluated forms into `/tmp/rrule-review-wp02-baseline-forms.json`, and compared it with the committed baseline using `cmp`: byte-identical, exit0. No baseline source was edited. The documented late-shared-start convention uses explicit2020 UNTIL so the old time.Now fallback yields empty output in the recorded2026 context; this bounded convention is stated rather than concealed as a timeless clock freeze.

Compared the fresh DST result with this reviewer's independently saved approved-WP01 result: exactly19 rows differ. Eighteen change only String (17 nominal-gap strings and literal087 floating form); literal088 becomes the deliberately rejected TZID/floating-UNTIL error. Every other occurrence list remains identical. Reviewed the per-row classifications and supplied negative empty/corruption controls. There is no new broad zone allowlist or unexplained rebaseline.

## Evidence integrity and runtime checklist

The red logs precede production edits except the explicitly labeled supplemental separate/inline check, which was red on the initial intermediate parser before its correction. Already-correct legal controls are not represented as defects. The implementer's wrong-cwd diagnostic is explicitly excluded from valid evidence; baseline byte-integrity and the separate genuine rerun are preserved. Its one nondeterministic golden control was repaired and regenerated before commit, with original failure evidence retained; the independent baseline cmp above confirms the final result.

| Item | Verdict | Basis |
|---|---|---|
| Dead code | PASS | New private parsing, validation, formatting and Set transfer helpers have live production callers. |
| Synthetic-fixture tests | PASS | Tests invoke public parsers, constructors, setters, iteration, serialization and Set APIs. |
| Silent empty return | PASS | Errors are propagated with relevant time-part names; missing metadata/ignored EXRULE are deliberate compatibility branches. |
| FR coverage | PASS | FR-003/004/005 have production assertions; FR-007 has paired controls and golden evidence; FR-008 is covered by the release note. |
| Frozen surface | PASS | Inherited contributor normalization, module minimum, dependencies and approved WP01 algorithm unchanged. |
| Locked decision | PASS | Lexical forms, exact locations, immutable per-part metadata, no marker/registry and bounded iteration are retained. |
| Shared-file ownership | PASS | Sequential shared parser/set work and narrow test/driver/corpus adaptations are disclosed and authorized. |
| Production fragility | PASS | No new panic or global mutable transport; explicit invalid forms return named errors. |

No required corrective edit remains from this review. Root owns canonical verdict capture and the subsequent mission acceptance/consolidation/final-review steps.
