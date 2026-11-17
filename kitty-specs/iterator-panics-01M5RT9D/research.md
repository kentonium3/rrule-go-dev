# Design research

Date: 2026-10-25. Evidence: source rrule.go, issue 01, fixed reproductions,
docs/CONTEXT.md; no external disclosure.

- Decision: guard ordinal candidates before mask reads. The original code indexes
  wdaymask before testing the period. Rejection at ±5 was considered but would
  unnecessarily change accepted empty/mixed rules and YEARLY/WEEKLY semantics.
- Decision: compare Easter offsets against bounds before addition, skipping
  impossible mask entries. An arbitrary global offset cap can change previously
  working year-specific rules; candidate guards preserve those.
- Decision: preserve BYEASTER=257 silent-drop behavior as an explicit known limit.
  The mask includes seven spill days but YEARLY daysets stop at yearlen; repairing
  spillover selection would require a distinct semantic decision and more scope.
- Decision: avoid interval caps; compare calendar horizon and decompose carries.
  MaxInt can validly produce DTSTART then exhaustion; on 32-bit MaxInt seconds
  is approximately 68 years and can yield multiple valid dates. Raw addition,
  weekly multiplication and accumulation inside BY-filter loops all need scrutiny.
- Decision: retain Go 1.16 module declaration with a build-tagged fuzz file.
  Native fuzzing needs a newer test toolchain, not necessarily a newer library
  language minimum. No dependency decisions or supply-chain changes are proposed.
- Decision: bounded fuzz domains plus explicit adversarial regression cases.
  COUNT and UNTIL cannot guarantee a bound on one iterator call given issue 02/03.
  Exclude unreachable subdaily filter states and cap list products. No recovery
  that turns panics into passing fuzz cases.
- Baseline: Go 1.25.14 linux/amd64; timeout 180 go test ./... passed with writable
  GOCACHE in /tmp. Full review and mutation checks are implementation evidence.
