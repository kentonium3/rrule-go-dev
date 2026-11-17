# Verification quickstart

Run from ~/repos/rrule-go, or the implementation checkout assigned by Spec Kitty. Planned tests are not implemented yet.

1. Record go version and baseline go test ./... on master 1a62a24b9460201d16e43b693c5ea9cfc3700198.
2. Add timeout-safe public-constructor regressions. Record independent pre-fix failure for each sub-daily frequency; kill and wait for every timed-out child.
3. Add bounded search, rerun targeted tests and exact reachable controls, then the full suite.
4. Run go test -race ./..., go vet ./..., go build ./..., GOARCH=386 go test ./..., and gofmt -l on changed Go files.
5. Record measured operation durations, setter transitions, final-cycle candidate controls, overflow overlap tests and independent review.

Keep all evidence private. Consolidation target is fix/unreachable-byset-hang; the later private PR targets master and awaits user merge authorization.
