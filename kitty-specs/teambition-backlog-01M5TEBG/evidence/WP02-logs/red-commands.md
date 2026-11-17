# Red and initial-green command provenance

All valid commands below used lane-b cwd `~/repos/rrule-go/.worktrees/teambition-backlog-01M5TEBG-lane-b`; HEAD was approved dependency `ef373a5872dba30063a6a706275a28e46508af77`. Implementation model gpt-6.1-sol, date2026-10-25. Logs include actual go test output and final exit.

1. Before any production edit, after adding str_timeforms_test.go and str_timeforms_api_test.go:
`GOCACHE=/tmp/rrule-go-cache go test -run '^TestTimeForms' -v -count=1`
Saved red-timeforms.txt, exit1. Every parser pair/parameter/form/order/shared adapter and immutable/custom-location fixture was exercised. Existing legal UTC/TZID, fragment/programmatic and keyed-API controls pass. Assertions about source serialization in copy tests fail on old serialization; their "source changed" message does not claim an observed old mutable-pointer bug.
2. Still before any production edit, after explicit-gap-text/UNTIL and DATE-copy tests were added:
`GOCACHE=/tmp/rrule-go-cache go test -run '^TestTimeForms(ExplicitGapTextAndUntil|LegacyExruleUnknownControl|DateCopyAndSet)$' -v -count=1`
Saved red-additional.txt, exit1. Gap text, pre-gap UNTIL inclusivity, DATE transport fail; unrelated ignored EXRULE property control passes.
3. Initial parser/set implementation:
`GOCACHE=/tmp/rrule-go-cache go test -run '^TestTimeForms' -v -count=1`
Saved green-timeforms-first.txt, exit0. Scoped old invalid fixture migrations followed.
4. `GOCACHE=/tmp/rrule-go-cache go test -count=1 ./...`
Saved green-all-first.txt, exit0. This was an initial whole-suite check, not claimed as the final exact-source gate.
5. A supplemental consistency case found separate-line DTSTART plus conflicting inline DTSTART still overriding the standalone parser anchor. Before adding the corresponding `shared = start` check, on the uncommitted intermediate parser above:
`GOCACHE=/tmp/rrule-go-cache go test -run '^TestTimeFormsSeparateStartCannotBeOverridden$' -v -count=1`
Saved red-separate-inline.txt, exit1. Three standalone adapters fail; set adapters already pass, and all matching-inline controls pass. This is supplemental red evidence on the intermediate fix, not evidence falsely claimed to precede all production editing.
6. After that narrow correction and diagnostic restoration, correct lane-b command:
`GOCACHE=/tmp/rrule-go-cache go test -run '^TestTimeForms|^TestDST' -v -count=1`
Saved green-focused.txt, exit0. INVALID-focused-wrong-baseline-cwd.txt is explicitly excluded; see fault-cwd.md and all-file baseline integrity proofs.
7. Final production source (only a resource/horizon control test and documentation added since6):
`GOCACHE=/tmp/rrule-go-cache go test -count=1 ./...`
Saved gate-modern.txt, exit0, package0.505s. All later gates include exact argv/cwd/HEAD/source hashes in companion .command.json files.
