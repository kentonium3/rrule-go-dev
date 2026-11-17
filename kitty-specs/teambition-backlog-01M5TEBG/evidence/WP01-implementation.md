# WP01 implementation handoff

Implementation model: **gpt-6.1-sol**, date **2026-10-25**. Implementer profile: resolver-loaded `implementer-ivan`. Independent review has not been performed by this implementer; root dispatches the required **gpt-6-astra** review and records its actual date/verdict.

Code commit: **af862491333c8a06fdb30bb60a58619e65e9baa6** (`fix(WP01): resolve recurrence civil times across DST`). Author/committer Kent Gale, 2026-10-25 17:58:30 -0400. Lane `kitty/mission-teambition-backlog-01M5TEBG-lane-a`, workspace `~/repos/rrule-go/.worktrees/teambition-backlog-01M5TEBG-lane-a`. Commit verified with git log/show; working tree clean. No push, PR, merge, status transition, sibling work, installation/host/security modification, or public action. Root owns mission evidence and task transitions.

Source baseline: **76c646dde8d33ff7c3325df0613b31b058243bea**, v1.8.3 plus inherited contributor duplicate-selector work. Initial lane HEAD a0ea3af6f1e7896334e50189610c393c2a68b0b5. Exact git-show bytes for all14 Go/module files in the pristine snapshot `/tmp/rrule-backlog-v183-plus-duplicates` match the pinned baseline; hashes in `baseline-provenance.txt`. That original snapshot was not mutated. `/tmp/rrule-WP01-baseline-testing` is a separate copy used for additional public-path regression replay and identical benchmark test source. Contributor issue7 was not assessed or extended.

## Changed files and ownership rationale

Owned implementation: `rrule.go`, new `civil_time.go`, new `rrule_dst_test.go`.
Owned harness/data: `docs/repro/05-07-backlog/golden/main.go`, `corpus-dst.json`, `baseline-dst.json`.
Owned release note: `docs/releases/v1.9.0-dst.md`.

Explicitly justified narrow adapters outside the original map:

- `str.go`: parse nominal DTSTART calendar fields before Go normalization; standalone DTSTART uses RFC pre-gap/first-fold interpretation; option parsing carries an immutable start descriptor. Shared-start Set parsing transfers that descriptor into an otherwise no-DTSTART rule before construction. This fixes public production paths for WP01 and provides WP02 foundations, without implementing UNTIL form validation, DATE syntax expansion, line-order redesign, or lexical serialization. Existing inline/shared overwrite behavior remains until WP02 validates it.
- `rrule_until_test.go`: update exactly3 obsolete expected-value rows after full-suite failure: NY generated gap, Sao Paulo midnight gap, Apia skipped civil date. Only missing candidates/duplicate normalized instants are removed per RFC5545 section3.3.10. No UNTIL algorithm or unrelated golden change. Root authorized this adapter after inspecting the exact failure.
- The following six existing ignored repro programs gain only the equivalent `// +build ignore` line below `//go:build ignore`: `docs/repro/01-iterator-panics/main.go`, `02-unreachable-byset-hang/main.go`, `03-until-scan/main.go`, `05-07-backlog/main.go`, `edge-cases/main.go`, `matrix/main.go`. Actual Go1.16.15 whole-checkout `go test ./...` failed because these tags were missing; root authorized the exact build-compatibility repairs. No program behavior, CI, manifest, or dependency changed. The new golden program already has both tags.

## Implementation and limits

`civilDay` resolves nominal civil tuples through actual location pointers and exact field round trips. A normal day gets one candidate offset; a transition gets both offsets and resolves the earliest exact match. No exact generated match means a skipped gap. Parsed explicit gap DTSTART instead uses the pre-gap offset and keeps nominal cadence separately in `RRule.anchor`. A caller-supplied later-fold time is authoritative at the matching first nominal slot; ordinary generated folds still use the first offset. Both output branches share resolution and `emit`, which rejects physical duplicates/backwards overlaps without consuming COUNT and preserves inclusive UNTIL.

Bounded-search invariant: a UTC civil-day interval extends from midnight-24h to midnight+48h, a72-hour lookup window. Realistic supported modern transitions are separated by more than72h and have zone offsets within24h; at most one transition occurs in the window. Its exact second is located with at most18 binary-search zone lookups after the two endpoint lookups. No hard-coded one-hour shift, ZoneBounds, global cache, registry, sentinel location, name reload, reflection, unsafe, generics, new dependency or Go API newer than1.16. Profiles are iterator-local and reused within a civil day. Constant FixedZone offsets work directly, including unusual offsets exercised by the inherited suite. Verified realistic northern/southern/half-hour transitions: New York, Sydney, Lord Howe. Sao Paulo/Apia inherited controls also pass after RFC gap updates; this is not a claim of arbitrary historical discontinuity support or rapidly repeated transition support.

BYSETPOS does not scan the annual day×time product. It counts selected civil days, binary-searches sorted nominal timeset entries at gap boundaries, subtracts sparse index holes, and maps signed ranks around holes. Gap-free periods use direct division/modulo by timeset length. Only selected candidates are constructed, with selected-instant deduplication. `emit` retains the existing effective-UTC short-UNTIL horizon (`r.dtstart.Location()==time.UTC`); new default-floating UTC control proves the parser metadata does not turn an empty short-horizon rule into a MAXYEAR scan.

## WP02 private contracts

`ROption.parse` points to immutable `parseMetadata{dtstart, until *parsedTime}`. `parsedTime` fields: `form timeForm` (`timeFormDate/Floating/UTC/TZID`), `nominal civilFields`, exact `location *time.Location`, `snapshot time.Time`. `parsedStart()` / `parsedUntil()` independently validate current public value by instant **and pointer identity**. `setParsedStart` / `setParsedUntil` replace the container, preserve the other valid part, and do not mutate a shared descriptor. Ordinary no-metadata options stay nil when their public setters have no descriptor to retain.

`RRule.dtStart(time.Time, *parsedTime)` is the private shared-start handoff. Public `DTStart` calls it with nil, intentionally replacing nominal semantics even when passed a time identical to the snapshot. Public `Until` clears only its descriptor. The narrow Set parser adapter retains nominal starts without adding Set lexical-form storage; WP02 may extend Set storage and validation as planned. Detached public StrToDtStart returns only its correct physical time and cannot preserve nominal metadata through arbitrary time copies.

WP02 still owns complete form parsing, DATE/VALUE validation, order independence, shared/inline validation, Until descriptors and lexical serialization. WP01 parsed gap String currently follows the resolved public time; WP02 lexical gap serialization will deliberately change a few anchor/set golden rows. Root explicitly permits precise row hash/rationale updates for those RFC corrections. Do not add a broad allowlist.

## Test-first evidence

All logs below are in **`/tmp/rrule-WP01-logs`** and include exact command, exit, and output unless explicitly marked raw. Red tests use public production entry points; every named subtest runs even when a sibling fails.

Before any main production edit, on untouched lane source inherited from76c646d:

- `red-TestDSTSydney.txt`: all specified Sydney hourly/minutely/daily gap/fold defects fail; `BYHOUR9WallClockControl` passes as an already-correct control.
- `red-TestDSTAllFrequencies.txt`: all7 frequencies fail independently for generated gap and first-fold behavior.
- `red-TestDSTOtherZones.txt`: NY gap and Lord Howe fold fail; NY first-fold and the original Lord Howe gap fixture already pass and are controls.
- `red-final-TestDSTBysetposRanks.txt`: positive3/negative-3 daily ranks, signed-overlap dedup and half-hour positive3/negative-3 all fail. Same-fixture half-hour1/2/-1/-2 controls already pass.
- `red-TestDSTPublicQueries.txt`: Iterator, Between, After, Before, Set and UNTIL before/at/after boundary expose incorrect generated gap output. UNTIL tests clear COUNT before applying UNTIL.
- `red-final-TestDSTParsedAnchors.txt`: gap nominal cadence, NY pre-gap offset, parsed first-fold, standalone gap/fold and NY subdaily overlap fail. Sydney subdaily overlap was already-correct baseline control and is not represented as a defect.
- `red-TestDSTAuthoritativeLaterFold.txt`: NY constructor and setter independently expose the missing authoritative later-fold first instance. Sydney/Lord Howe later-fold controls already pass.
- `red-TestDSTMetadataCopiesAndLocation.txt`: copied gap descriptor/cadence and Until-setter retaining start semantics fail; per-part invalidation/location controls are baseline-correct.

Subsequent strengthening used the **untouched source copy**, never edits to production to manufacture red failures:

- `expanded-baseline-red.txt`: all7 frequencies×3DST zones×both transitions plus yearly/monthly signed rank fixtures. Some NY fold controls correctly pass baseline; all other reported failures are retained. This is supplemental baseline replay after the main resolver edit, not a claim those extra variants preceded it. The independent defect families above already had red-first evidence.
- `red-parsed-set-baseline.txt`: all3 parsed Set anchor cases fail on baseline.
- `red-parsed-set-adapter.txt`: after the main resolver but **before** the shared Set descriptor adapter, Sydney/NY gap cadence still fails while corrected Sydney fold passes. The narrow adapter then makes all cases green in `green-parsed-set.txt`.

Test-authoring corrections are disclosed: initial BYSETPOS2/-2 fixtures happened to normalize to the same physical result as correct ranking, so they were strengthened to3/-3 and separately rerun red **before implementation**. An initial yearly Sydney expected2023-10-01 was wrong because that date is also a gap; corrected to2023-10-02. The initial yearly test also independently failed its earlier real gap candidate, but that mistaken expected date is not claimed as a library defect. The initially attempted dense benchmark retained YEARLY's implicit Jan1 restriction; corrected to all7 BYDAY values before final measurements and added explicit Jan1/Dec31 assertions.

## Final gates (all exit0)

Modern toolchain: `go version go1.25.14 linux/amd64`.
Minimum toolchain: direct `/tmp/rrule-go116-toolchain/go/bin/go version` reports `go1.16.15 linux/amd64`.

- `GOCACHE=/tmp/rrule-go-cache go test -run TestDST -v -count=1`:0.007s, `gate-focused-final.txt`.
- `GOCACHE=/tmp/rrule-go-cache go test ./...`:0.478s, `gate-modern-final.txt`.
- `GOCACHE=/tmp/rrule-go116-cache /tmp/rrule-go116-toolchain/go/bin/go test ./...`:0.345s, `gate-go116-final.txt`.
- `GOCACHE=/tmp/rrule-go-cache go vet ./...`:zero diagnostics, `gate-vet-final.txt`.
- `gofmt -l .`:zero paths, `gate-gofmt-final.txt`.
- `GOCACHE=/tmp/rrule-go-cache go test -race -run TestDST ./...`:1.033s, `gate-race-focused-final.txt`.
- `git diff --check`:zero diagnostics before commit.

The integrated full race and>=60s fuzz gates remain WP02/root closeout responsibilities as explicitly planned, not claimed completed here. No new tests or benchmarks are needed merely to repeat already-passed work absent review findings.

## True dense annual benchmark and arithmetic proof

Identical final `rrule_dst_test.go` source in baseline testing copy and current lane:

`GOCACHE=/tmp/rrule-go-cache go test -run '^$' -bench BenchmarkDSTAnnualBysetpos -benchmem -benchtime=1s -count=3`

- Baseline:1,210,873 /1,153,926 /997,929ns/op;2,111,810–2,111,816B/op;54allocs/op. `benchmark-dense-baseline.txt`.
- Current:1,223,499 /1,134,124 /1,148,099ns/op;2,226,026–2,226,037B/op;47allocs/op. `benchmark-dense-current.txt`.
- Median approximately1.154ms baseline /1.148ms current; modest extra day-profile memory (~114KB), fewer allocations. Runtime measurements are noisy; no speedup claim.
- `TestDSTDenseAnnualArithmetic` passes identically on baseline/current and asserts ranks1,366,-366,-1 produce Jan1T00:00:00/00:06:05 and Dec31T23:53:54/23:59:59. All7 BYDAY values suppress default month/day restriction, so365×86,400=31,536,000 nominal candidates are represented. Source branch uses arithmetic rank selection and day-level holes, never annual Cartesian enumeration. Logs `dense-baseline-control.txt`, `dense-current-control.txt`.

## Golden corpus, provenance, classifications

**799 rows**, first16 instances or exhaustion, fixed DTSTARTs, every String/error/instant/exhaustion flag compared. Driver is compiled against baseline/current module directories with the exact same source file, no go.mod replace or network operation. Every row runs in a separate killable2-second subprocess; **zero timeouts** in both runs.

-304 static ROption expressions from baseline tests (Go AST extraction, public equivalent of private weekday literals, fixed missing starts).
-92 string literals containing FREQ from baseline test files plus existing matrix/edge-cases/backlog repro sources, including invalid/error controls and inherited duplicate-selector controls.
-392 explicit transition-week fixtures:7frequencies×4zones×spring/fall×7days. Subdaily clocks/intervals are near transition boundaries; coarser selectors align to transition days.
-11 explicit parsed-anchor/shared-Set/signed-rank sensitivity fixtures.
-39 dynamic ROption expressions individually listed with exact source line and unresolved identifier in corpus provenance; they remain covered by the full source suite. No silent omission or pathological exclusion. Static extraction total343=304included+39listed-skipped.

**396 extracted baseline fixture rows remain byte-identical**, as do all UTC matrix rows. **102 changed rows, zero unexplained**. Every changed row carries its own source/regression/RFC rationale and exact approved-current-row SHA256; no zone allowlist. Full before/after row output is in `golden-classified-final.txt`. Classes: generated gap/fold with downstream COUNT; signed BYSETPOS gap ranking/period replenishment; explicit gap/fold anchor interpretation/nominal cadence; shared-start descriptor transfer. Only NY explicit-gap anchor/Set public-time String reflects its corrected resolved instant at this stage. No unrelated parse error or source fixture difference.

Committed baseline artifact SHA256: **5b28a5252eecce332e9f418ad2dff53731883fd00ed26fb25b04aae8fc688c90**. Comparison verifies this file hash before checking IDs/order/count, and hashes exact classified current rows.

Negative controls: baseline compared with itself fails exit1 because changed=0; corrupting an otherwise unchanged UTC fixture fails exit1 with changed103/unexplained1. Logs `sensitivity-unchanged-final.txt`, `sensitivity-corrupt-final.txt`. An empty or broadly ignored comparison cannot pass.

Repro commands (all actual successful exit0 unless the negative controls above):

```
# Current lane cwd
GOCACHE=/tmp/rrule-go-cache go build -o /tmp/rrule-dst-golden-current docs/repro/05-07-backlog/golden/main.go
/tmp/rrule-dst-golden-current -out /tmp/rrule-current-dst.json
# Pristine snapshot cwd; exact same driver source imports that module
GOCACHE=/tmp/rrule-go-cache go build -o /tmp/rrule-dst-golden-baseline ~/repos/rrule-go/.worktrees/teambition-backlog-01M5TEBG-lane-a/docs/repro/05-07-backlog/golden/main.go
/tmp/rrule-dst-golden-baseline -corpus ~/repos/rrule-go/.worktrees/teambition-backlog-01M5TEBG-lane-a/docs/repro/05-07-backlog/golden/corpus-dst.json -out ~/repos/rrule-go/.worktrees/teambition-backlog-01M5TEBG-lane-a/docs/repro/05-07-backlog/golden/baseline-dst.json
# Current lane cwd
/tmp/rrule-dst-golden-current -baseline docs/repro/05-07-backlog/golden/baseline-dst.json -current /tmp/rrule-current-dst.json
```

Additional useful temporary source/data: `/tmp/extract_options.go`, `/tmp/rrule-static-options.json`, `/tmp/rrule-extract-static.go`, `/tmp/rrule-static-rules.json`, `/tmp/extract_strings.go`, `/tmp/rrule-static-strings.json`, `/tmp/create_corpus.py`. The committed corpus itself is deterministic and sufficient to rerun without extraction. No public fetch or dependency download.

## Fault handling and scope discipline

Initial full modern test failure is preserved in `green-full-modern.txt`: only the3 obsolete DST gap expected rows failed. Minimum-Go pre-repair output is preserved in `gate-go116-tests.txt`: six `//go:build comment without // +build comment` diagnostics. Diagnosis found code-only baseline snapshot lacks private docs programs, explaining why its prior Go1.16 pass did not establish whole-checkout compatibility. Root explicitly authorized equivalent build tags and exact expected-value updates; no gate bypass, installation change, CI change, or UNTIL algorithm edit.

A read-only extraction attempt assumed the code-only snapshot contained docs/repro paths and failed opening its matrix/main.go. The corpus was then extracted from snapshot test files plus the unchanged tracked lane docs source, with individual source-line provenance. Initial compile-time extraction authoring mistakes (private Weekday literal conversion, missing primitive int allowlist) were corrected in temporary extractor scripts only. These were local harness-authoring errors, not Spec Kitty/runtime faults. No blind mutation retry or original baseline mutation.

Spec Kitty read-only governance/profile and safe-commit commands all used `SPEC_KITTY_NO_UPGRADE_CHECK=1`. Terminology/checklist section lookups report those sections are unauthored; no invented doctrine or charter change. Safe-commit was run once with the15 explicit implementation/adapter paths and exact `--to-branch`, using the authorized Git metadata escalation boundary; success exit0. No status commands were run by this implementer. The commit body credits retained PR66 Sydney scenarios and RFC-rederived expectations, without colliding issue-closing keywords.
