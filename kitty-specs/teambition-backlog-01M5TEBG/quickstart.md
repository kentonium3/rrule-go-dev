# Verification quickstart

Run from the applicable private mission lane or consolidated fix/teambition-backlog checkout. Use GOCACHE=/tmp/rrule-go-cache when needed. Each WP evidence directory records exact focused test names and red/green commands; do not claim nonexistent tests ran.

```
go vet ./...
gofmt -l .
go test ./...
go test -race ./...
go test -run='^$' -fuzz=FuzzIterate -fuzztime=60s
```

The committed golden driver instructions will identify corpus path, baseline code identity, bounded count, current/baseline output paths and comparison command. Both occurrence values and String/errors are compared; every changed row requires an explicit RFC/issue classification. Fixed DTSTARTs make runs reproducible. Cases outside the two corrections must be byte-identical.

Before code changes, retain regression-red evidence from unchanged baseline/preceding-WP code. After changes, retain green output plus independent review model/date and implementation model/date for each WP. Full local gates and integrated final review follow consolidation; deliver private PR to main and stop without merging.

Minimum-version gate (official temporary toolchain, no host install change):

```
GOCACHE=/tmp/rrule-go116-cache /tmp/rrule-go116-toolchain/go/bin/go test ./...
```

Root verified official archive SHA256 `77c782a633186d78c384f972fb113a43c24be0234c42fef22c2d8c4c4c8e7475`; baseline passes on Go1.16.15. Pristine code snapshot `/tmp/rrule-backlog-v183-plus-duplicates` from76c646dde8d33ff7c3325df0613b31b058243bea is available for independent baseline/red runs; confirm library blobs as described by the golden plan.

Scope/baseline correction2026-10-25: only issues5/6 remain; issue7 contributor fix is inherited at76c646d. Golden against v1.8.3 plus that commit. Test keyed ROption callers and metadata value copies; document external positional-literal source break for approved v1.9.0. Include exact custom FixedZone/IANA collision controls, custom same-name TZif controls, per-part public-field replacement, public setters and internal Set handoffs. Preserve direct arithmetic BYSETPOS performance and effective-UTC UNTIL horizon optimization.
