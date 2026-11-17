# WP03 migration verification — private evidence

Author: Implementer Ivan, actual model **gpt-6.1-sol**, 2026-10-28 America/New_York (execution may extend into 2026-10-29 UTC). Independent reviewer is separately assigned gpt-6-astra; this report is implementation evidence, not that review.

## Source authority and scope

Verified approved lane-b candidate `9bf894f4b899d0c1cb814e2f30753eeec6e01dab`, including WP01 `5de6b56adb8c1208749ff1d4d56c436d731db68e`. Executed lane HEAD was `9a4f64dc2169d96969f03c12ee5a9fa24bfe1154`; its only delta from approved implementation commit is runtime-owned status.events.jsonl/status.json. All tested source fingerprints therefore identify the approved implementation. Candidate cwd is `~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b`. Root remains the runtime-assigned planning/artifact workspace on `rrule-go-dev-xfer`; its old module declaration is intentional until consolidation. No manual merge or worktree was created. Pristine baseline `737ae01be40048d1d2932dbafb623230efa8d361` was exported using `git archive` and tar extraction into a disposable source directory. This baseline already includes approved DST/time-form corrections.

All Go commands used process-scoped `GOWORK=off GOPROXY=off GOCACHE=/tmp/rrule-module-path-wp03-cache`. Global Go settings were unchanged. No public retrieval, proxy, registry, shim, dependency, or permanent test framework was added.

## T011–T012: declaration, consumers, and dependency safety

The identical explicit declaration assertion printed the old identity and exited 1 in the pristine baseline export, then printed the intended identity and exited 0 in candidate cwd. This independently reconstructs WP01's preserved historical red/green transcript. It is a declaration mismatch, not an unresolved dependency/tool failure. Candidate `go.mod` contains exactly the new module declaration, blank line and `go 1.16`; no dependency added.

Three disposable modules preserve original input and final text in the transcript below. Each began with a working old-import test against pristine baseline source. The documented public fork v1.8.3 replacement was normalized to a local source location only for offline execution; this is not a test of the historical release download. The direct original-dependency shape used the same local old-source scaffolding. The baseline's extra already-approved corrections are immaterial to fixture shape and DAILY positive control.

Each fixture changes imports, removes the exact old directive, adds the new v1.9.0 requirement and fixture-only new-path replacement, tidies, inspects selected module JSON and graph, and passes tests/build. The actual removal arguments are the guide's bare old path for all-version and old path `@v1.8.2` for version-specific. Their alternate forms are never mixed. Tidy eliminates unused old direct requirements. Selected replacement directory is the approved lane-b candidate. Test source below exercises UTC DAILY/COUNT=3, all three timestamps, exact String output and keyed ROption construction. The final direct fixture also supplies the external-consumer positive check.

The retained dependency's public function returns only int; internally it imports the old package. `go mod graph`, `go mod why -m` and package-import inspection all reveal it. Tests pass while the protective old-path replacement remains. The preflight therefore stops: coordinate that dependency's migration or remain on the existing pinned v1.8.3 arrangement. No replacement is removed in this case. Restoring original upstream selection would lose panic/hang fixes even with no shared named types. The production guide explicitly covers this safety rule, build tags/platforms, distinct named types, and separate inherited keyed-ROption compatibility migration.

WP02's original package-query VCS-stamping failure remains recorded in WP02-review.md; it was not a passing command. This WP uses the authorized process-scoped `GOFLAGS=-buildvcs=false` only for the disposable retained-consumer package query. Ordinary candidate tests/builds and fixtures use no such override. Production instructions are unchanged.

## T013: recurrence preservation and sensitivity

Each source snapshot builds its own driver from its own cwd and proper import identity; binary hashes appear below. Both binaries read the identical absolute unchanged corpus path, emitting distinct files. Counts, matching IDs/order, nonempty output and byte equality were asserted. `cmp` exits 0 for each pristine pair. Appending a newline to a separate temporary copy makes `cmp` exit 1 for each corpus. Historical classified comparator is intentionally not used because its changes==0 assertion is unsuitable here.

There are 39 DST rejection rows and 60 forms rejection rows. Inspected messages are expected validation/parser rejections: BY* bounds, invalid rule/property/value syntax, DTSTART/UNTIL type-form pairing, invalid TZID/VALUE usage and shared DTSTART conflicts. There are no panic or timeout errors and no unexplained infrastructure failure rows. Positive UTC consumer behavior provides an additional healthy control. Full rejection ID/message inventory is included in the JSON summary below.

## T014: required checks and preserved surfaces

Broad `git diff --check 737ae01 HEAD` exited 2 for inherited planning Markdown whitespace in checklists/requirements.md, spec.md and a blank line in the inherited transition draft. Exact diagnostics are retained below. The separately scoped candidate source/documentation check excluding kitty-specs exits 0. This is not a claim that the broad check passed; unowned planning files were not edited. Current owned drafts were whitespace-checked separately.

All required checks were rerun here against final lane-b candidate rather than relying on earlier reports: `go test ./...`, `go test -race ./...`, `go vet ./...`, `gofmt -l .` (zero filenames), and `go run ./tools/changelog check`; all pass. All seven changed standalone repro drivers compile individually; intentionally dangerous repros were not run unbounded. Exact commands/outcomes/toolchain follow.

Programmatic body-equivalence check passes for every changed Go file after the single identity substitution and authorized golden header exclusion. Thus exported declarations, signatures, package name and executable bodies are unchanged. Go minimum/dependencies and LICENSE are unchanged. Every inherited golden JSON file and released CHANGELOG section starting at 1.8.3 remains byte-identical. Candidate golden README/header correctly separates source-specific drivers. Changed source/doc files match the approved occurrence-map exceptions: go.mod, two external tests, seven drivers, golden README, README, CHANGELOG, MIGRATION. Intentional old references remain in migration before-state commands, attribution, historical outputs and authority records. Private README preamble/badges are not public-port copy.

## T015–T016: packets and deferred boundaries

Both private draft packets are reconciled with MIGRATION.md and the matching Unreleased additions. The deprecation proposal now points directly to the future public MIGRATION.md rather than an unverified README anchor. Public excerpts use neutral separately-maintained-fork language, preserve original module identity and maintainer decision authority, both removal forms and the pre-removal dependency gate. Release notifications remain opt-in; latest-selection deprecation and unconditional replacement limitations remain explicit. The dated maintenance audit is historical evidence, not current permissions or independently verified private conversations.

The unsent [Vikunja handoff](vikunja-handoff.md) contains the exact candidate and future allowlist, integration scenarios and release gates. All public availability/version-use checks remain deferred. Kent must approve exact public copy/actions. No PR, tag, release, index request, message, Vikunja edit, recipient choice, automation or main merge occurred.

## Reproduction and durable fingerprints

Temporary paths in the transcript are execution locations, not promised durable artifacts. Reconstruct source exports from the pinned commits using `git archive <sha>` into new temporary directories. Recreate the consumer files from preserved initial/final text below, substitute your actual baseline/candidate absolute directories, and use the recorded process environment. Build each source's own `docs/repro/05-07-backlog/golden/main.go`; run `-corpus <same absolute corpus> -out <distinct path>` and compare with `cmp`. Binary hashes record this execution; different paths/toolchain/VCS metadata can change binaries without altering recurrence output. Output/corpus/source hashes and exact rows provide the reproducible behavior evidence.

One temporary verification-script path collision (`baseline` directory versus intended binary output) caused an IsADirectoryError while hashing the build output before golden execution; corrected by distinct `baseline-golden`/`candidate-golden` binary names and reran the entire disposable verification from fresh exports. This was an authoring error in disposable scaffolding, not a workflow fault or candidate failure. No production file was touched.

### Source, binary/output/corpus identities and rejection inventory

```json
{
  "temp": "/tmp/wp03-jxscekdx",
  "baseline": "737ae01be40048d1d2932dbafb623230efa8d361",
  "candidate": "9a4f64dc2169d96969f03c12ee5a9fa24bfe1154",
  "fingerprints": {
    "go.mod": "ac0b9b1758244ea5eec28b475292c2f6e2be94a05c96d1f7e7826d0659615930",
    "rrule.go": "ca821beee860da6019b94a40c5acde18e8d01cf3f97cc1f6ac94dea303886f08",
    "str.go": "74a1b537e02b9ce9d03f8d01ac7763cdb4af88218143735006167d4efd184c59",
    "LICENSE": "2c5aa8001e37c2ba4488e1c809772c9372623f2eaa0605b9abd8863cf7a3470f",
    "MIGRATION.md": "7e1e1b076b28107d4ac6a22963d5439d6eec4a251575dfe78c0268442f15d55f",
    "README.md": "95e0f31ad9957756b9ee565a0ba1e4a844b7a1beef1596bd552046f371f95b89",
    "CHANGELOG.md": "7e9af33c625aa3aa8df3d9d81bfbdc0a0cf1b2519e13d9d48da6dac616fd83de",
    "docs/repro/05-07-backlog/golden/main.go": "238b1d2dfe162302d3d009712ae4d988b799b52f9c0a40824778800b81335203",
    "docs/repro/05-07-backlog/golden/README.md": "b24950f1533ef97050cb7b414412386a5f3a96ad697060c6b7665fb0c350f370"
  },
  "golden": [
    {
      "kind": "dst",
      "count": 811,
      "corpus_sha256": "297d1c73215be4cabb3cb0b769350306ce346522816f2429f717ad37a5b1bf28",
      "output_sha256": "1fca954470ccc972e6eadd40faaf2c9c4ebe99b9f779897c92d6e6f0a5c01a4b",
      "bytes": 402552,
      "errors": [
        [
          "option-001",
          "bysetpos must be between 1 and 366 or -1 and -366"
        ],
        [
          "option-002",
          "bysetpos must be between 1 and 366 or -1 and -366"
        ],
        [
          "option-006",
          "bysecond must be between 0 and 59"
        ],
        [
          "option-007",
          "bysecond must be between 0 and 59"
        ],
        [
          "option-008",
          "byminute must be between 0 and 59"
        ],
        [
          "option-009",
          "byminute must be between 0 and 59"
        ],
        [
          "option-010",
          "byhour must be between 0 and 23"
        ],
        [
          "option-011",
          "byhour must be between 0 and 23"
        ],
        [
          "option-012",
          "bymonthday must be between 1 and 31 or -1 and -31"
        ],
        [
          "option-013",
          "bymonthday must be between 1 and 31 or -1 and -31"
        ],
        [
          "option-014",
          "bymonthday must be between 1 and 31 or -1 and -31"
        ],
        [
          "option-015",
          "byyearday must be between 1 and 366 or -1 and -366"
        ],
        [
          "option-016",
          "byyearday must be between 1 and 366 or -1 and -366"
        ],
        [
          "option-017",
          "byyearday must be between 1 and 366 or -1 and -366"
        ],
        [
          "option-018",
          "byweekno must be between 1 and 53 or -1 and -53"
        ],
        [
          "option-019",
          "byweekno must be between 1 and 53 or -1 and -53"
        ],
        [
          "option-020",
          "byweekno must be between 1 and 53 or -1 and -53"
        ],
        [
          "option-021",
          "bymonth must be between 1 and 12"
        ],
        [
          "option-022",
          "bymonth must be between 1 and 12"
        ],
        [
          "option-023",
          "bysetpos must be between 1 and 366 or -1 and -366"
        ],
        [
          "option-024",
          "bysetpos must be between 1 and 366 or -1 and -366"
        ],
        [
          "option-025",
          "bysetpos must be between 1 and 366 or -1 and -366"
        ],
        [
          "option-026",
          "byday must be between 1 and 53 or -1 and -53"
        ],
        [
          "option-027",
          "byday must be between 1 and 53 or -1 and -53"
        ],
        [
          "option-028",
          "interval must be greater than 0"
        ],
        [
          "literal-001",
          "invalid RRULE string"
        ],
        [
          "literal-032",
          "BYSETPOS option has no value"
        ],
        [
          "literal-033",
          "invalid RRULE string"
        ],
        [
          "literal-034",
          "invalid RRULE string"
        ],
        [
          "literal-042",
          "undefined frequency: HELLO"
        ],
        [
          "literal-043",
          "unknown RRULE property: HELLO"
        ],
        [
          "literal-044",
          "strconv.Atoi: parsing \"I\": invalid syntax"
        ],
        [
          "literal-045",
          "undefined weekday: M"
        ],
        [
          "literal-046",
          "undefined weekday: MQ"
        ],
        [
          "literal-047",
          "strconv.Atoi: parsing \"+\": invalid syntax"
        ],
        [
          "literal-055",
          "unknown RRULE property: EXRULE:FREQ"
        ],
        [
          "literal-088",
          "UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "literal-089",
          "expect DTSTART but: CONTROL"
        ],
        [
          "literal-090",
          "BYSECOND option has no value"
        ]
      ]
    },
    {
      "kind": "forms",
      "count": 158,
      "corpus_sha256": "a064b4fcdbbbf40c4380cad9ef9fe1655493f432ba8387f51470581845e6895d",
      "output_sha256": "c632fe0db2b700ff2acdca385d2e74f9a3392a432982020898e85c0b5503a236",
      "bytes": 40951,
      "errors": [
        [
          "pair-date-floating-reverse0-set0",
          "UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-date-floating-reverse0-set1",
          "StrToROption failed: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-date-floating-reverse1-set0",
          "UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-date-floating-reverse1-set1",
          "StrToROption failed: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-date-utc-reverse0-set0",
          "UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-date-utc-reverse0-set1",
          "StrToROption failed: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-date-utc-reverse1-set0",
          "UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-date-utc-reverse1-set1",
          "StrToROption failed: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-floating-date-reverse0-set0",
          "UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-floating-date-reverse0-set1",
          "StrToROption failed: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-floating-date-reverse1-set0",
          "UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-floating-date-reverse1-set1",
          "StrToROption failed: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-floating-utc-reverse0-set0",
          "UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-floating-utc-reverse0-set1",
          "StrToROption failed: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-floating-utc-reverse1-set0",
          "UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-floating-utc-reverse1-set1",
          "StrToROption failed: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-utc-date-reverse0-set0",
          "UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-utc-date-reverse0-set1",
          "StrToROption failed: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-utc-date-reverse1-set0",
          "UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-utc-date-reverse1-set1",
          "StrToROption failed: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-utc-floating-reverse0-set0",
          "UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-utc-floating-reverse0-set1",
          "StrToROption failed: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-utc-floating-reverse1-set0",
          "UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-utc-floating-reverse1-set1",
          "StrToROption failed: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-tzid-date-reverse0-set0",
          "UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-tzid-date-reverse0-set1",
          "StrToROption failed: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-tzid-date-reverse1-set0",
          "UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-tzid-date-reverse1-set1",
          "StrToROption failed: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-tzid-floating-reverse0-set0",
          "UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-tzid-floating-reverse0-set1",
          "StrToROption failed: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-tzid-floating-reverse1-set0",
          "UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "pair-tzid-floating-reverse1-set1",
          "StrToROption failed: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "parameter-05-set0",
          "DTSTART: TZID must not be applied to DATE or UTC DTSTART"
        ],
        [
          "parameter-05-set1",
          "DTSTART: TZID must not be applied to DATE or UTC DTSTART"
        ],
        [
          "parameter-06-set0",
          "DTSTART: TZID must not be applied to DATE or UTC DTSTART"
        ],
        [
          "parameter-06-set1",
          "DTSTART: TZID must not be applied to DATE or UTC DTSTART"
        ],
        [
          "parameter-07-set0",
          "DTSTART: TZID must not be applied to DATE or UTC DTSTART"
        ],
        [
          "parameter-07-set1",
          "DTSTART: TZID must not be applied to DATE or UTC DTSTART"
        ],
        [
          "parameter-08-set0",
          "DTSTART: TZID must not be applied to DATE or UTC DTSTART"
        ],
        [
          "parameter-08-set1",
          "DTSTART: TZID must not be applied to DATE or UTC DTSTART"
        ],
        [
          "parameter-09-set0",
          "DTSTART: DTSTART VALUE parameter does not match value type"
        ],
        [
          "parameter-09-set1",
          "DTSTART: DTSTART VALUE parameter does not match value type"
        ],
        [
          "parameter-10-set0",
          "DTSTART: DTSTART VALUE parameter does not match value type"
        ],
        [
          "parameter-10-set1",
          "DTSTART: DTSTART VALUE parameter does not match value type"
        ],
        [
          "inline-reverse0-bad1",
          "UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "inline-reverse1-bad1",
          "UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "shared-overwritten-pos0",
          "StrToROption failed: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "shared-overwritten-pos1",
          "StrToROption failed: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "shared-overwritten-pos2",
          "StrToROption failed: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "shared-ignored-pos0",
          "EXRULE: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "shared-ignored-pos1",
          "EXRULE: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "shared-ignored-pos2",
          "EXRULE: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "shared-inline-bypass-pos0",
          "StrToROption failed: DTSTART conflicts with shared DTSTART"
        ],
        [
          "shared-inline-bypass-pos1",
          "StrToROption failed: DTSTART conflicts with shared DTSTART"
        ],
        [
          "shared-inline-bypass-pos2",
          "StrToROption failed: DTSTART conflicts with shared DTSTART"
        ],
        [
          "shared-inline-anchor-pos0",
          "StrToROption failed: DTSTART conflicts with shared DTSTART"
        ],
        [
          "shared-inline-anchor-pos1",
          "StrToROption failed: DTSTART conflicts with shared DTSTART"
        ],
        [
          "shared-inline-anchor-pos2",
          "StrToROption failed: DTSTART conflicts with shared DTSTART"
        ],
        [
          "old-mixed-exrule",
          "EXRULE: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ],
        [
          "old-LA-invalid",
          "StrToROption failed: UNTIL value type/time form must match DTSTART (UTC UNTIL required for UTC/TZID DTSTART)"
        ]
      ]
    }
  ]
}
```

### Exact command transcript

```text
cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ go version
exit=0
go version go1.25.14 linux/amd64


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ git rev-parse HEAD
exit=0
9a4f64dc2169d96969f03c12ee5a9fa24bfe1154


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ git status --short
exit=0


cwd=/tmp/wp03-jxscekdx/baseline
$ sh -c 'go list -m; test "$(go list -m)" = github.com/kentonium3/rrule-go'
exit=1
github.com/teambition/rrule-go


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ sh -c 'go list -m; test "$(go list -m)" = github.com/kentonium3/rrule-go'
exit=0
github.com/kentonium3/rrule-go


INITIAL direct
module example.invalid/direct

go 1.16
require github.com/teambition/rrule-go v1.8.2

package consumer
import("testing";"time"; rrule "github.com/teambition/rrule-go")
func TestDaily(t *testing.T){r,e:=rrule.StrToRRule("DTSTART:20260105T100000Z\nRRULE:FREQ=DAILY;COUNT=3");if e!=nil{t.Fatal(e)};a:=r.All();if len(a)!=3{t.Fatal(a)};for i,x:=range a{want:=time.Date(2026,1,5+i,10,0,0,0,time.UTC);if !x.Equal(want){t.Fatal(x,want)}};if r.String()!="DTSTART:20260105T100000Z\nRRULE:FREQ=DAILY;COUNT=3"{t.Fatal(r.String())};if _,e=rrule.NewRRule(rrule.ROption{Freq:rrule.DAILY,Count:3});e!=nil{t.Fatal(e)}}


cwd=/tmp/wp03-jxscekdx/direct
$ go mod edit -replace=github.com/teambition/rrule-go=/tmp/wp03-jxscekdx/baseline
exit=0


cwd=/tmp/wp03-jxscekdx/direct
$ go mod tidy
exit=0


cwd=/tmp/wp03-jxscekdx/direct
$ go test ./...
exit=0
ok  	example.invalid/direct	0.002s


cwd=/tmp/wp03-jxscekdx/direct
$ go mod edit -dropreplace=github.com/teambition/rrule-go
exit=0


cwd=/tmp/wp03-jxscekdx/direct
$ go mod edit -json
exit=0
{
	"Module": {
		"Path": "example.invalid/direct"
	},
	"Go": "1.16",
	"Require": [
		{
			"Path": "github.com/teambition/rrule-go",
			"Version": "v1.8.2"
		}
	],
	"Exclude": null,
	"Replace": null,
	"Retract": null,
	"Tool": null,
	"Ignore": null
}


cwd=/tmp/wp03-jxscekdx/direct
$ go mod edit -require=github.com/kentonium3/rrule-go@v1.9.0
exit=0


cwd=/tmp/wp03-jxscekdx/direct
$ go mod edit -replace=github.com/kentonium3/rrule-go=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
exit=0


cwd=/tmp/wp03-jxscekdx/direct
$ go mod tidy
exit=0


cwd=/tmp/wp03-jxscekdx/direct
$ go list -m -json all
exit=0
{
	"Path": "example.invalid/direct",
	"Main": true,
	"Dir": "/tmp/wp03-jxscekdx/direct",
	"GoMod": "/tmp/wp03-jxscekdx/direct/go.mod",
	"GoVersion": "1.16"
}
{
	"Path": "github.com/kentonium3/rrule-go",
	"Version": "v1.9.0",
	"Replace": {
		"Path": "~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b",
		"Dir": "~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b",
		"GoMod": "~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b/go.mod",
		"GoVersion": "1.16"
	},
	"Dir": "~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b",
	"GoMod": "~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b/go.mod",
	"GoVersion": "1.16"
}


cwd=/tmp/wp03-jxscekdx/direct
$ go test ./...
exit=0
ok  	example.invalid/direct	0.002s


cwd=/tmp/wp03-jxscekdx/direct
$ go build ./...
exit=0


cwd=/tmp/wp03-jxscekdx/direct
$ go mod graph
exit=0
example.invalid/direct github.com/kentonium3/rrule-go@v1.9.0
example.invalid/direct go@1.16


FINAL direct
module example.invalid/direct

go 1.16

require github.com/kentonium3/rrule-go v1.9.0

replace github.com/kentonium3/rrule-go => ~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b

package consumer
import("testing";"time"; rrule "github.com/kentonium3/rrule-go")
func TestDaily(t *testing.T){r,e:=rrule.StrToRRule("DTSTART:20260105T100000Z\nRRULE:FREQ=DAILY;COUNT=3");if e!=nil{t.Fatal(e)};a:=r.All();if len(a)!=3{t.Fatal(a)};for i,x:=range a{want:=time.Date(2026,1,5+i,10,0,0,0,time.UTC);if !x.Equal(want){t.Fatal(x,want)}};if r.String()!="DTSTART:20260105T100000Z\nRRULE:FREQ=DAILY;COUNT=3"{t.Fatal(r.String())};if _,e=rrule.NewRRule(rrule.ROption{Freq:rrule.DAILY,Count:3});e!=nil{t.Fatal(e)}}


INITIAL all
module example.invalid/all

go 1.16
require github.com/teambition/rrule-go v1.8.2
replace github.com/teambition/rrule-go => github.com/kentonium3/rrule-go v1.8.3

package consumer
import("testing";"time"; rrule "github.com/teambition/rrule-go")
func TestDaily(t *testing.T){r,e:=rrule.StrToRRule("DTSTART:20260105T100000Z\nRRULE:FREQ=DAILY;COUNT=3");if e!=nil{t.Fatal(e)};a:=r.All();if len(a)!=3{t.Fatal(a)};for i,x:=range a{want:=time.Date(2026,1,5+i,10,0,0,0,time.UTC);if !x.Equal(want){t.Fatal(x,want)}};if r.String()!="DTSTART:20260105T100000Z\nRRULE:FREQ=DAILY;COUNT=3"{t.Fatal(r.String())};if _,e=rrule.NewRRule(rrule.ROption{Freq:rrule.DAILY,Count:3});e!=nil{t.Fatal(e)}}


cwd=/tmp/wp03-jxscekdx/all
$ go mod edit -replace=github.com/teambition/rrule-go=/tmp/wp03-jxscekdx/baseline
exit=0


cwd=/tmp/wp03-jxscekdx/all
$ go mod tidy
exit=0


cwd=/tmp/wp03-jxscekdx/all
$ go test ./...
exit=0
ok  	example.invalid/all	0.002s


cwd=/tmp/wp03-jxscekdx/all
$ go mod edit -dropreplace=github.com/teambition/rrule-go
exit=0


cwd=/tmp/wp03-jxscekdx/all
$ go mod edit -json
exit=0
{
	"Module": {
		"Path": "example.invalid/all"
	},
	"Go": "1.16",
	"Require": [
		{
			"Path": "github.com/teambition/rrule-go",
			"Version": "v1.8.2"
		}
	],
	"Exclude": null,
	"Replace": null,
	"Retract": null,
	"Tool": null,
	"Ignore": null
}


cwd=/tmp/wp03-jxscekdx/all
$ go mod edit -require=github.com/kentonium3/rrule-go@v1.9.0
exit=0


cwd=/tmp/wp03-jxscekdx/all
$ go mod edit -replace=github.com/kentonium3/rrule-go=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
exit=0


cwd=/tmp/wp03-jxscekdx/all
$ go mod tidy
exit=0


cwd=/tmp/wp03-jxscekdx/all
$ go list -m -json all
exit=0
{
	"Path": "example.invalid/all",
	"Main": true,
	"Dir": "/tmp/wp03-jxscekdx/all",
	"GoMod": "/tmp/wp03-jxscekdx/all/go.mod",
	"GoVersion": "1.16"
}
{
	"Path": "github.com/kentonium3/rrule-go",
	"Version": "v1.9.0",
	"Replace": {
		"Path": "~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b",
		"Dir": "~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b",
		"GoMod": "~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b/go.mod",
		"GoVersion": "1.16"
	},
	"Dir": "~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b",
	"GoMod": "~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b/go.mod",
	"GoVersion": "1.16"
}


cwd=/tmp/wp03-jxscekdx/all
$ go test ./...
exit=0
ok  	example.invalid/all	0.002s


cwd=/tmp/wp03-jxscekdx/all
$ go build ./...
exit=0


cwd=/tmp/wp03-jxscekdx/all
$ go mod graph
exit=0
example.invalid/all github.com/kentonium3/rrule-go@v1.9.0
example.invalid/all go@1.16


FINAL all
module example.invalid/all

go 1.16

require github.com/kentonium3/rrule-go v1.9.0

replace github.com/kentonium3/rrule-go => ~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b

package consumer
import("testing";"time"; rrule "github.com/kentonium3/rrule-go")
func TestDaily(t *testing.T){r,e:=rrule.StrToRRule("DTSTART:20260105T100000Z\nRRULE:FREQ=DAILY;COUNT=3");if e!=nil{t.Fatal(e)};a:=r.All();if len(a)!=3{t.Fatal(a)};for i,x:=range a{want:=time.Date(2026,1,5+i,10,0,0,0,time.UTC);if !x.Equal(want){t.Fatal(x,want)}};if r.String()!="DTSTART:20260105T100000Z\nRRULE:FREQ=DAILY;COUNT=3"{t.Fatal(r.String())};if _,e=rrule.NewRRule(rrule.ROption{Freq:rrule.DAILY,Count:3});e!=nil{t.Fatal(e)}}


INITIAL version
module example.invalid/version

go 1.16
require github.com/teambition/rrule-go v1.8.2
replace github.com/teambition/rrule-go v1.8.2 => github.com/kentonium3/rrule-go v1.8.3

package consumer
import("testing";"time"; rrule "github.com/teambition/rrule-go")
func TestDaily(t *testing.T){r,e:=rrule.StrToRRule("DTSTART:20260105T100000Z\nRRULE:FREQ=DAILY;COUNT=3");if e!=nil{t.Fatal(e)};a:=r.All();if len(a)!=3{t.Fatal(a)};for i,x:=range a{want:=time.Date(2026,1,5+i,10,0,0,0,time.UTC);if !x.Equal(want){t.Fatal(x,want)}};if r.String()!="DTSTART:20260105T100000Z\nRRULE:FREQ=DAILY;COUNT=3"{t.Fatal(r.String())};if _,e=rrule.NewRRule(rrule.ROption{Freq:rrule.DAILY,Count:3});e!=nil{t.Fatal(e)}}


cwd=/tmp/wp03-jxscekdx/version
$ go mod edit -replace=github.com/teambition/rrule-go@v1.8.2=/tmp/wp03-jxscekdx/baseline
exit=0


cwd=/tmp/wp03-jxscekdx/version
$ go mod tidy
exit=0


cwd=/tmp/wp03-jxscekdx/version
$ go test ./...
exit=0
ok  	example.invalid/version	0.001s


cwd=/tmp/wp03-jxscekdx/version
$ go mod edit -dropreplace=github.com/teambition/rrule-go@v1.8.2
exit=0


cwd=/tmp/wp03-jxscekdx/version
$ go mod edit -json
exit=0
{
	"Module": {
		"Path": "example.invalid/version"
	},
	"Go": "1.16",
	"Require": [
		{
			"Path": "github.com/teambition/rrule-go",
			"Version": "v1.8.2"
		}
	],
	"Exclude": null,
	"Replace": null,
	"Retract": null,
	"Tool": null,
	"Ignore": null
}


cwd=/tmp/wp03-jxscekdx/version
$ go mod edit -require=github.com/kentonium3/rrule-go@v1.9.0
exit=0


cwd=/tmp/wp03-jxscekdx/version
$ go mod edit -replace=github.com/kentonium3/rrule-go=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
exit=0


cwd=/tmp/wp03-jxscekdx/version
$ go mod tidy
exit=0


cwd=/tmp/wp03-jxscekdx/version
$ go list -m -json all
exit=0
{
	"Path": "example.invalid/version",
	"Main": true,
	"Dir": "/tmp/wp03-jxscekdx/version",
	"GoMod": "/tmp/wp03-jxscekdx/version/go.mod",
	"GoVersion": "1.16"
}
{
	"Path": "github.com/kentonium3/rrule-go",
	"Version": "v1.9.0",
	"Replace": {
		"Path": "~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b",
		"Dir": "~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b",
		"GoMod": "~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b/go.mod",
		"GoVersion": "1.16"
	},
	"Dir": "~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b",
	"GoMod": "~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b/go.mod",
	"GoVersion": "1.16"
}


cwd=/tmp/wp03-jxscekdx/version
$ go test ./...
exit=0
ok  	example.invalid/version	0.002s


cwd=/tmp/wp03-jxscekdx/version
$ go build ./...
exit=0


cwd=/tmp/wp03-jxscekdx/version
$ go mod graph
exit=0
example.invalid/version github.com/kentonium3/rrule-go@v1.9.0
example.invalid/version go@1.16


FINAL version
module example.invalid/version

go 1.16

require github.com/kentonium3/rrule-go v1.9.0

replace github.com/kentonium3/rrule-go => ~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b

package consumer
import("testing";"time"; rrule "github.com/kentonium3/rrule-go")
func TestDaily(t *testing.T){r,e:=rrule.StrToRRule("DTSTART:20260105T100000Z\nRRULE:FREQ=DAILY;COUNT=3");if e!=nil{t.Fatal(e)};a:=r.All();if len(a)!=3{t.Fatal(a)};for i,x:=range a{want:=time.Date(2026,1,5+i,10,0,0,0,time.UTC);if !x.Equal(want){t.Fatal(x,want)}};if r.String()!="DTSTART:20260105T100000Z\nRRULE:FREQ=DAILY;COUNT=3"{t.Fatal(r.String())};if _,e=rrule.NewRRule(rrule.ROption{Freq:rrule.DAILY,Count:3});e!=nil{t.Fatal(e)}}


cwd=/tmp/wp03-jxscekdx/retained
$ go mod tidy
exit=0


cwd=/tmp/wp03-jxscekdx/retained
$ go test ./...
exit=0
?   	example.invalid/retained	[no test files]


cwd=/tmp/wp03-jxscekdx/retained
$ go mod graph
exit=0
example.invalid/retained example.invalid/dependency@v0.0.0
example.invalid/retained go@1.16
example.invalid/dependency@v0.0.0 github.com/teambition/rrule-go@v1.8.2


cwd=/tmp/wp03-jxscekdx/retained
$ go mod why -m github.com/teambition/rrule-go
exit=0
# github.com/teambition/rrule-go
example.invalid/retained
example.invalid/dependency
github.com/teambition/rrule-go


cwd=/tmp/wp03-jxscekdx/retained
$ go list -deps -test -f '{{.ImportPath}}: {{join .Imports " "}}' ./...
exit=0
internal/goarch: 
unsafe: 
internal/abi: internal/goarch unsafe
internal/unsafeheader: unsafe
internal/cpu: unsafe
internal/bytealg: internal/cpu unsafe
internal/byteorder: 
internal/chacha8rand: internal/byteorder internal/cpu internal/goarch unsafe
internal/coverage/rtcov: unsafe
internal/godebugs: 
internal/goexperiment: 
internal/goos: 
internal/profilerecord: 
internal/runtime/atomic: internal/goarch unsafe
internal/runtime/math: internal/goarch
internal/runtime/strconv: internal/runtime/math
internal/runtime/syscall: internal/goarch unsafe
internal/runtime/cgroup: internal/bytealg internal/runtime/strconv internal/runtime/syscall unsafe
internal/runtime/exithook: internal/runtime/atomic unsafe
internal/runtime/gc: internal/goarch
internal/asan: unsafe
internal/msan: unsafe
internal/race: internal/abi unsafe
internal/runtime/sys: internal/goarch internal/goos
internal/runtime/maps: internal/abi internal/asan internal/goarch internal/msan internal/race internal/runtime/math internal/runtime/sys unsafe
internal/stringslite: internal/bytealg unsafe
internal/trace/tracev2: 
runtime: internal/abi internal/bytealg internal/byteorder internal/chacha8rand internal/coverage/rtcov internal/cpu internal/goarch internal/godebugs internal/goexperiment internal/goos internal/profilerecord internal/runtime/atomic internal/runtime/cgroup internal/runtime/exithook internal/runtime/gc internal/runtime/maps internal/runtime/math internal/runtime/strconv internal/runtime/sys internal/runtime/syscall internal/stringslite internal/trace/tracev2 unsafe
internal/reflectlite: internal/abi internal/goarch internal/unsafeheader runtime unsafe
errors: internal/reflectlite unsafe
cmp: 
internal/itoa: 
iter: internal/race runtime unsafe
math/bits: unsafe
math: internal/cpu math/bits unsafe
unicode/utf8: 
strconv: errors internal/bytealg internal/stringslite math math/bits unicode/utf8
sync/atomic: unsafe
internal/sync: internal/abi internal/goarch internal/race sync/atomic unsafe
internal/synctest: internal/abi unsafe
sync: internal/race internal/sync internal/synctest runtime sync/atomic unsafe
unicode: 
reflect: errors internal/abi internal/bytealg internal/goarch internal/itoa internal/race internal/runtime/maps internal/runtime/sys internal/unsafeheader iter math runtime strconv sync unicode unicode/utf8 unsafe
slices: cmp iter math/bits unsafe
internal/fmtsort: cmp reflect slices
io: errors sync
internal/oserror: errors
path: errors internal/bytealg unicode/utf8
internal/bisect: runtime sync sync/atomic
internal/godebug: internal/bisect internal/godebugs sync sync/atomic unsafe
syscall: errors internal/asan internal/bytealg internal/byteorder internal/goarch internal/itoa internal/msan internal/oserror internal/race internal/runtime/syscall runtime slices sync sync/atomic unsafe
time: errors internal/bytealg internal/godebug internal/stringslite math/bits runtime sync syscall unsafe
io/fs: errors internal/bytealg internal/oserror io path slices time unicode/utf8
internal/filepathlite: errors internal/bytealg internal/stringslite io/fs slices
internal/syscall/unix: internal/itoa runtime sync/atomic syscall unsafe
internal/poll: errors internal/itoa internal/syscall/unix io runtime sync sync/atomic syscall time unsafe
internal/syscall/execenv: syscall
internal/testlog: sync sync/atomic unsafe
os: errors internal/bytealg internal/byteorder internal/filepathlite internal/goarch internal/itoa internal/poll internal/stringslite internal/syscall/execenv internal/syscall/unix internal/testlog io io/fs runtime slices sync sync/atomic syscall time unsafe
fmt: errors internal/fmtsort io math os reflect slices strconv sync unicode/utf8
sort: internal/reflectlite math/bits slices
strings: errors internal/abi internal/bytealg internal/stringslite io iter math/bits sync unicode unicode/utf8 unsafe
github.com/teambition/rrule-go: errors fmt math sort strconv strings time
example.invalid/dependency: github.com/teambition/rrule-go
example.invalid/retained: example.invalid/dependency fmt


SAFETY STOP: retained dependency imports original path while exposing only int. No protective replacement removed; coordinate migration or stay pinned v1.8.3. Local source substitution represents pinned fork semantics, not retrieval.

cwd=/tmp/wp03-jxscekdx/baseline
$ go build -o /tmp/wp03-jxscekdx/baseline-golden ./docs/repro/05-07-backlog/golden/main.go
exit=0


baseline binary SHA256 1b22da0eb37d8e6bab255c00242fca7d24a56c902ddc1b0e30323136d55785ae

cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ go build -o /tmp/wp03-jxscekdx/candidate-golden ./docs/repro/05-07-backlog/golden/main.go
exit=0


candidate binary SHA256 6708fa26ee66360d930253759474795d360089cef773739a06788e6528685ab1

cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ /tmp/wp03-jxscekdx/baseline-golden -corpus ~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b/docs/repro/05-07-backlog/golden/corpus-dst.json -out /tmp/wp03-jxscekdx/baseline-golden-dst.json
exit=0


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ /tmp/wp03-jxscekdx/candidate-golden -corpus ~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b/docs/repro/05-07-backlog/golden/corpus-dst.json -out /tmp/wp03-jxscekdx/candidate-golden-dst.json
exit=0


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ cmp /tmp/wp03-jxscekdx/baseline-golden-dst.json /tmp/wp03-jxscekdx/candidate-golden-dst.json
exit=0


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ cmp /tmp/wp03-jxscekdx/baseline-golden-dst.json /tmp/wp03-jxscekdx/corrupt-dst
exit=1
cmp: EOF on /tmp/wp03-jxscekdx/baseline-golden-dst.json after byte 402552, line 14337


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ /tmp/wp03-jxscekdx/baseline-golden -corpus ~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b/docs/repro/05-07-backlog/golden/corpus-forms.json -out /tmp/wp03-jxscekdx/baseline-golden-forms.json
exit=0


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ /tmp/wp03-jxscekdx/candidate-golden -corpus ~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b/docs/repro/05-07-backlog/golden/corpus-forms.json -out /tmp/wp03-jxscekdx/candidate-golden-forms.json
exit=0


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ cmp /tmp/wp03-jxscekdx/baseline-golden-forms.json /tmp/wp03-jxscekdx/candidate-golden-forms.json
exit=0


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ cmp /tmp/wp03-jxscekdx/baseline-golden-forms.json /tmp/wp03-jxscekdx/corrupt-forms
exit=1
cmp: EOF on /tmp/wp03-jxscekdx/baseline-golden-forms.json after byte 40951, line 1486


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ go test ./...
exit=0
ok  	github.com/kentonium3/rrule-go	(cached)
ok  	github.com/kentonium3/rrule-go/tools/changelog	(cached)


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ go test -race ./...
exit=0
ok  	github.com/kentonium3/rrule-go	(cached)
ok  	github.com/kentonium3/rrule-go/tools/changelog	(cached)


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ go vet ./...
exit=0


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ gofmt -l .
exit=0


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ go run ./tools/changelog check
exit=0
CHANGELOG.md: valid


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ go build -o /tmp/wp03-jxscekdx/compile-01-iterator-panics ~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b/docs/repro/01-iterator-panics/main.go
exit=0


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ go build -o /tmp/wp03-jxscekdx/compile-02-unreachable-byset-hang ~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b/docs/repro/02-unreachable-byset-hang/main.go
exit=0


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ go build -o /tmp/wp03-jxscekdx/compile-03-until-scan ~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b/docs/repro/03-until-scan/main.go
exit=0


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ go build -o /tmp/wp03-jxscekdx/compile-golden ~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b/docs/repro/05-07-backlog/golden/main.go
exit=0


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ go build -o /tmp/wp03-jxscekdx/compile-05-07-backlog ~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b/docs/repro/05-07-backlog/main.go
exit=0


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ go build -o /tmp/wp03-jxscekdx/compile-edge-cases ~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b/docs/repro/edge-cases/main.go
exit=0


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ go build -o /tmp/wp03-jxscekdx/compile-matrix ~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b/docs/repro/matrix/main.go
exit=0


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ git diff --check 737ae01be40048d1d2932dbafb623230efa8d361 HEAD
exit=2
kitty-specs/module-path-v1-9-0-01M62K6W/checklists/requirements.md:3: trailing whitespace.
+**Purpose**: Validate specification completeness before planning.  
kitty-specs/module-path-v1-9-0-01M62K6W/checklists/requirements.md:4: trailing whitespace.
+**Created**: 2026-10-28 (America/New_York)  
kitty-specs/module-path-v1-9-0-01M62K6W/drafts/teambition-transition-pr.md:40: trailing whitespace.
+ 
kitty-specs/module-path-v1-9-0-01M62K6W/spec.md:3: trailing whitespace.
+**Mission Branch**: `rrule-go-dev-xfer`  
kitty-specs/module-path-v1-9-0-01M62K6W/spec.md:4: trailing whitespace.
+**Created**: 2026-10-28 (America/New_York)  
kitty-specs/module-path-v1-9-0-01M62K6W/spec.md:5: trailing whitespace.
+**Status**: Specified; implementation and publication pending  


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ git diff --check 737ae01be40048d1d2932dbafb623230efa8d361 HEAD -- . ':(exclude)kitty-specs'
exit=0


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ git diff --name-status 737ae01be40048d1d2932dbafb623230efa8d361 HEAD
exit=0
M	CHANGELOG.md
A	MIGRATION.md
M	README.md
M	docs/repro/01-iterator-panics/main.go
M	docs/repro/02-unreachable-byset-hang/main.go
M	docs/repro/03-until-scan/main.go
M	docs/repro/05-07-backlog/golden/README.md
M	docs/repro/05-07-backlog/golden/main.go
M	docs/repro/05-07-backlog/main.go
M	docs/repro/edge-cases/main.go
M	docs/repro/matrix/main.go
M	example_test.go
M	go.mod
A	kitty-specs/module-path-v1-9-0-01M62K6W/acceptance-matrix.json
A	kitty-specs/module-path-v1-9-0-01M62K6W/checklists/requirements.md
A	kitty-specs/module-path-v1-9-0-01M62K6W/contracts/module-migration.md
A	kitty-specs/module-path-v1-9-0-01M62K6W/data-model.md
A	kitty-specs/module-path-v1-9-0-01M62K6W/decisions/DM-01M62K7ADP565BH1469Q0VEJ12.md
A	kitty-specs/module-path-v1-9-0-01M62K6W/decisions/DM-01M62N0BZ44HJDCVQPNTC7D0F8.md
A	kitty-specs/module-path-v1-9-0-01M62K6W/decisions/index.json
A	kitty-specs/module-path-v1-9-0-01M62K6W/decisions/index.json.lock
A	kitty-specs/module-path-v1-9-0-01M62K6W/drafts/release-communication.md
A	kitty-specs/module-path-v1-9-0-01M62K6W/drafts/teambition-transition-pr.md
A	kitty-specs/module-path-v1-9-0-01M62K6W/evidence/occurrence-map-approval.md
A	kitty-specs/module-path-v1-9-0-01M62K6W/evidence/occurrence-map-validation.txt
A	kitty-specs/module-path-v1-9-0-01M62K6W/evidence/rrule-module-mission-base-check.txt
A	kitty-specs/module-path-v1-9-0-01M62K6W/evidence/rrule-module-mission-create-failure.txt
A	kitty-specs/module-path-v1-9-0-01M62K6W/evidence/task-generation-check.md
A	kitty-specs/module-path-v1-9-0-01M62K6W/evidence/upstream-maintenance-audit.md
A	kitty-specs/module-path-v1-9-0-01M62K6W/meta.json
A	kitty-specs/module-path-v1-9-0-01M62K6W/occurrence_map.yaml
A	kitty-specs/module-path-v1-9-0-01M62K6W/plan.md
A	kitty-specs/module-path-v1-9-0-01M62K6W/quickstart.md
A	kitty-specs/module-path-v1-9-0-01M62K6W/reasons-canvas.md
A	kitty-specs/module-path-v1-9-0-01M62K6W/research.md
A	kitty-specs/module-path-v1-9-0-01M62K6W/reviews/post-plan-disposition.md
A	kitty-specs/module-path-v1-9-0-01M62K6W/reviews/post-plan-prompt.txt
A	kitty-specs/module-path-v1-9-0-01M62K6W/reviews/post-plan-rereview-prompt.txt
A	kitty-specs/module-path-v1-9-0-01M62K6W/reviews/post-plan-rereview.md
A	kitty-specs/module-path-v1-9-0-01M62K6W/reviews/post-plan-review.md
A	kitty-specs/module-path-v1-9-0-01M62K6W/spec.md
A	kitty-specs/module-path-v1-9-0-01M62K6W/status.events.jsonl
A	kitty-specs/module-path-v1-9-0-01M62K6W/status.json
A	kitty-specs/module-path-v1-9-0-01M62K6W/tasks/.gitkeep
A	kitty-specs/module-path-v1-9-0-01M62K6W/tasks/README.md
A	kitty-specs/module-path-v1-9-0-01M62K6W/tasks/WP01-module-identity.md
A	kitty-specs/module-path-v1-9-0-01M62K6W/tasks/WP01-module-identity/review-cycle-1.md
A	kitty-specs/module-path-v1-9-0-01M62K6W/tasks/WP02-consumer-migration.md
A	kitty-specs/module-path-v1-9-0-01M62K6W/tasks/WP03-verification-handoff.md
M	str_timeforms_api_test.go


cwd=~/repos/rrule-go-dev/.worktrees/module-path-v1-9-0-01M62K6W-lane-b
$ git rev-parse HEAD
exit=0
9a4f64dc2169d96969f03c12ee5a9fa24bfe1154

```

### Disposable runner source

```python
import os,subprocess,tempfile,pathlib,json,hashlib,tarfile,io,shlex
R=pathlib.Path('~/repos/rrule-go-dev'); C=R/'.worktrees/module-path-v1-9-0-01M62K6W-lane-b'; B='737ae01be40048d1d2932dbafb623230efa8d361'; T=pathlib.Path(tempfile.mkdtemp(prefix='wp03-')); S=T/'baseline'; S.mkdir()
E=dict(os.environ,GOWORK='off',GOPROXY='off',GOCACHE='/tmp/rrule-module-path-wp03-cache'); logs=[]
def run(args,cwd=C,env=None,ok=(0,)):
 p=subprocess.run(args,cwd=cwd,env=env or E,text=True,stdout=subprocess.PIPE,stderr=subprocess.STDOUT); logs.append('cwd='+str(cwd)+'\n$ '+shlex.join(args)+'\nexit='+str(p.returncode)+'\n'+p.stdout); assert p.returncode in ok,logs[-1]; return p.stdout
raw=subprocess.check_output(['git','archive',B],cwd=R); tarfile.open(fileobj=io.BytesIO(raw)).extractall(S,filter='data')
def sha(p):return hashlib.sha256(pathlib.Path(p).read_bytes()).hexdigest()
run(['go','version']); run(['git','rev-parse','HEAD']); run(['git','status','--short']);
for d in [S,C]: run(['sh','-c','go list -m; test "$(go list -m)" = github.com/kentonium3/rrule-go'],d,ok=(0,) if d==C else (1,))
source='''package consumer
import("testing";"time"; rrule "IMPORT")
func TestDaily(t *testing.T){r,e:=rrule.StrToRRule("DTSTART:20260105T100000Z\\nRRULE:FREQ=DAILY;COUNT=3");if e!=nil{t.Fatal(e)};a:=r.All();if len(a)!=3{t.Fatal(a)};for i,x:=range a{want:=time.Date(2026,1,5+i,10,0,0,0,time.UTC);if !x.Equal(want){t.Fatal(x,want)}};if r.String()!="DTSTART:20260105T100000Z\\nRRULE:FREQ=DAILY;COUNT=3"{t.Fatal(r.String())};if _,e=rrule.NewRRule(rrule.ROption{Freq:rrule.DAILY,Count:3});e!=nil{t.Fatal(e)}}
'''
old='github.com/teambition/rrule-go'; new='github.com/kentonium3/rrule-go'
for case in ['direct','all','version']:
 d=T/case;d.mkdir();initial='module example.invalid/'+case+'\n\ngo 1.16\nrequire '+old+' v1.8.2\n';initial+=('replace '+old+(' v1.8.2' if case=='version' else '')+' => '+new+' v1.8.3\n') if case!='direct' else '';logs.append('INITIAL '+case+'\n'+initial+'\n'+source.replace('IMPORT',old));(d/'go.mod').write_text(initial);(d/'consumer_test.go').write_text(source.replace('IMPORT',old));left=old+('@v1.8.2' if case=='version' else '');run(['go','mod','edit','-replace='+left+'='+str(S)],d);run(['go','mod','tidy'],d);run(['go','test','./...'],d)
 # Exact production removal form applied to normalized corresponding left side.
 run(['go','mod','edit','-dropreplace='+left],d);assert not json.loads(run(['go','mod','edit','-json'],d)).get('Replace');(d/'consumer_test.go').write_text(source.replace('IMPORT',new));run(['go','mod','edit','-require='+new+'@v1.9.0'],d);run(['go','mod','edit','-replace='+new+'='+str(C)],d);run(['go','mod','tidy'],d);run(['go','list','-m','-json','all'],d);run(['go','test','./...'],d);run(['go','build','./...'],d);run(['go','mod','graph'],d);assert old not in (d/'go.mod').read_text();logs.append('FINAL '+case+'\n'+(d/'go.mod').read_text()+'\n'+(d/'consumer_test.go').read_text())
# Retained dependency exposes no rrule types: preflight must still stop.
d=T/'retained';d.mkdir();dep=T/'dependency';dep.mkdir();(dep/'go.mod').write_text('module example.invalid/dependency\ngo 1.16\nrequire '+old+' v1.8.2\n');(dep/'dep.go').write_text('package dependency\nimport rrule "'+old+'"\nfunc Count()int{r,_:=rrule.StrToRRule("FREQ=DAILY;COUNT=1");return len(r.All())}\n');(d/'go.mod').write_text('module example.invalid/retained\ngo 1.16\nrequire example.invalid/dependency v0.0.0\nreplace example.invalid/dependency => '+str(dep)+'\nreplace '+old+' => '+str(S)+'\n');(d/'main.go').write_text('package main\nimport("fmt";"example.invalid/dependency")\nfunc main(){fmt.Println(dependency.Count())}\n');run(['go','mod','tidy'],d);run(['go','test','./...'],d);run(['go','mod','graph'],d);run(['go','mod','why','-m',old],d);out=run(['go','list','-deps','-test','-f','{{.ImportPath}}: {{join .Imports " "}}','./...'],d,env=dict(E,GOFLAGS='-buildvcs=false'));assert 'example.invalid/dependency: '+old in out;assert old in (d/'go.mod').read_text();logs.append('SAFETY STOP: retained dependency imports original path while exposing only int. No protective replacement removed; coordinate migration or stay pinned v1.8.3. Local source substitution represents pinned fork semantics, not retrieval.')
g=C/'docs/repro/05-07-backlog/golden';binaries=[]
for label,src in [('baseline',S),('candidate',C)]:
 binary=T/(label+"-golden");run(['go','build','-o',str(binary),'./docs/repro/05-07-backlog/golden/main.go'],src);binaries.append(binary);logs.append(label+' binary SHA256 '+sha(binary))
summary=[]
for kind,count in [('dst',811),('forms',158)]:
 corpus=g/('corpus-'+kind+'.json');outs=[]
 for binary in binaries:
  out=T/(binary.name+'-'+kind+'.json');run([str(binary),'-corpus',str(corpus),'-out',str(out)]);outs.append(out)
 rows=[json.loads(x.read_text()) for x in outs];assert all(len(x)==count for x in rows);assert [x['id'] for x in rows[0]]==[x['id'] for x in rows[1]];run(['cmp',str(outs[0]),str(outs[1])]);errors=[(x['id'],x['error']) for x in rows[1] if x['error']];assert not any('timeout' in e.lower() or 'panic:' in e.lower() for _,e in errors);summary.append(dict(kind=kind,count=count,corpus_sha256=sha(corpus),output_sha256=sha(outs[0]),bytes=outs[0].stat().st_size,errors=errors));bad=T/('corrupt-'+kind);bad.write_bytes(outs[0].read_bytes()+b'\n');run(['cmp',str(outs[0]),str(bad)],ok=(1,))
for args in [['go','test','./...'],['go','test','-race','./...'],['go','vet','./...'],['gofmt','-l','.'],['go','run','./tools/changelog','check']]:
 out=run(args);assert args[0]!='gofmt' or not out
for p in sorted((C/'docs/repro').rglob('main.go')):
 if str(p.relative_to(C)) in subprocess.check_output(['git','diff','--name-only',B,'HEAD'],cwd=C,text=True).splitlines():run(['go','build','-o',str(T/('compile-'+p.parent.name)),str(p)])
run(['git','diff','--check',B,'HEAD'],ok=(0,2));run(['git','diff','--check',B,'HEAD','--','.',':(exclude)kitty-specs']);diff=run(['git','diff','--name-status',B,'HEAD']);fingerprints={str(p.relative_to(C)):sha(p) for p in [C/'go.mod',C/'rrule.go',C/'str.go',C/'LICENSE',C/'MIGRATION.md',C/'README.md',C/'CHANGELOG.md',g/'main.go',g/'README.md']};(T/'summary.json').write_text(json.dumps(dict(temp=str(T),baseline=B,candidate=run(['git','rev-parse','HEAD']).strip(),fingerprints=fingerprints,golden=summary),indent=2));(T/'transcript.txt').write_text('\n\n'.join(logs));print(T)

```
