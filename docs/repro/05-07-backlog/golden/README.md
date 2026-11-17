# Reproducible compatibility comparison

The bounded public-API driver records first16 occurrences or exhaustion, String
and error status. Each fixture runs in a killable two-second subprocess; timeouts
are explicit rows. No module replacement or network operation is needed.
Candidate source declares the new module identity, so its driver must be built
from candidate cwd. Historical source declares the old identity and requires a
frozen old-import driver. Never build the current new-import driver from an
old-declaration source checkout.

## Source and driver selection

From candidate cwd, export pristine source snapshots into fresh temporary
folders. These are source exports, not Git worktrees. The pre-migration snapshot
contains the old-import driver; export that driver separately for the older
76c646d source, which need not contain it.

```sh
candidate_dir="$PWD"
comparison_dir=$(mktemp -d /tmp/rrule-golden.XXXXXX)
mkdir "$comparison_dir/pre-migration" "$comparison_dir/historical"
git archive 737ae01be40048d1d2932dbafb623230efa8d361 | tar -x -C "$comparison_dir/pre-migration"
git archive 76c646dde8d33ff7c3325df0613b31b058243bea | tar -x -C "$comparison_dir/historical"
git show 737ae01be40048d1d2932dbafb623230efa8d361:docs/repro/05-07-backlog/golden/main.go > "$comparison_dir/old-import-driver.go"
# New-import driver, built from candidate cwd.
GOWORK=off GOPROXY=off go build -o "$comparison_dir/current" ./docs/repro/05-07-backlog/golden/main.go
# Each old-import driver is built from the source cwd it must evaluate.
(cd "$comparison_dir/pre-migration" && GOWORK=off GOPROXY=off go build -o "$comparison_dir/before" ./docs/repro/05-07-backlog/golden/main.go)
(cd "$comparison_dir/historical" && GOWORK=off GOPROXY=off go build -o "$comparison_dir/historical-driver" "$comparison_dir/old-import-driver.go")
```

## Zero-change module migration comparison

The migration baseline is pre-migration source
`737ae01be40048d1d2932dbafb623230efa8d361`, including the approved DST/time-form
fixes. Collect fresh outputs for both binaries using the same unchanged absolute
corpus path, then compare their bytes:

```sh
"$comparison_dir/before" -corpus "$candidate_dir/docs/repro/05-07-backlog/golden/corpus-forms.json" -out "$comparison_dir/before-forms.json"
"$comparison_dir/current" -corpus "$candidate_dir/docs/repro/05-07-backlog/golden/corpus-forms.json" -out "$comparison_dir/current-forms.json"
cmp "$comparison_dir/before-forms.json" "$comparison_dir/current-forms.json"
```

Repeat with `corpus-dst.json` and distinct before/current DST outputs. Require
158 forms rows and 811 DST rows with matching IDs/order, zero differences and no
unexpected error or timeout rows. For a sensitivity control, corrupt only a
temporary output copy and require `cmp` to fail. Never alter inherited artifacts.

## Historical classified comparison

The historical baseline is
`76c646dde8d33ff7c3325df0613b31b058243bea` (v1.8.3 plus inherited contributor
fix). This comparison intentionally includes the later DST/time-form corrections;
it is separate from the zero-change module migration proof.

```sh
"$comparison_dir/historical-driver" -corpus "$candidate_dir/docs/repro/05-07-backlog/golden/corpus-forms.json" -out "$comparison_dir/regenerated-baseline-forms.json"
cmp "$comparison_dir/regenerated-baseline-forms.json" "$candidate_dir/docs/repro/05-07-backlog/golden/baseline-forms.json"
"$comparison_dir/current" -corpus "$candidate_dir/docs/repro/05-07-backlog/golden/corpus-forms.json" -baseline "$candidate_dir/docs/repro/05-07-backlog/golden/baseline-forms.json" -current "$comparison_dir/current-forms.json"
```

Repeat with `corpus-dst.json` / `baseline-dst.json` / the distinct current DST
output. The historical comparator pins the baseline SHA256, fixture
IDs/order/counts, and an individual exact current-row hash plus rationale for
every difference. It deliberately rejects zero changed rows and any unclassified
change, so it cannot serve as this migration's zero-change assertion.

The inherited DST corpus has811 rows:304 static ROption fixtures,92 FREQ string
literals,392 transition-matrix rows,23 targeted rows. The39 skipped dynamic
ROption fixtures are individually listed in its provenance and still execute in
the full test suite. WP02 retains the exact baseline artifact SHA256
`68fbd11d9cf3ca187dc473a543876196c23552ea5c2bb1b45f1ca561f75bd226`.
Only19 approved-WP01 current rows change in WP02:17 explicit gap-anchor strings
now preserve nominal source fields, one floating round-trip string, and one newly
forbidden TZID/floating UNTIL rule. Every occurrence outside the latter rejection
is unchanged relative to approved WP01. Against old76c646d,116 rows differ and all
are individually classified.

The separate forms corpus adds158 deterministic rows,141 classified changes and
17 unchanged controls; none are skipped or timed out. It covers all12 form pairs,
both line orders, rule/set paths, canonical VALUE parameters, illegal TZID forms,
inline/shared/overwritten/ignored rules, copies and per-part setters, ordinary
programmatic constructors, bare fragments, real NY versus same-name FixedZone
and custom TZif locations, and combined DST/BYSETPOS/UNTIL cases. Its baseline
SHA256 is `1cc3299bf76264f6bcaf653ed3a6b2509765c62c4639880b081775667b1ae934`.
For late-shared-DTSTART placement controls, the old parser may internally default
the start to `time.Now()`. Every such rule has an explicit2020 UNTIL, deliberately
past in the recorded2026-10-25 evaluation context, so the old result is empty and
reproducible. The harness does not claim that the old internal fallback is fixed.
The driver adds only optional location/action fields; old fixtures retain their
original execution path. The unit suite additionally checks every parser adapter,
location pointer identity, metadata isolation, query APIs and dense arithmetic.

Rationales distinguish newly rejected forms, newly supported canonical/order
forms, floating/DATE serialization, and issue5 explicit/generation corrections.
They are acceptance-specific, with no blanket zone allowlist. Contributor issue7
behavior is inherited compatibility baseline; this mission does not assess or
extend that work.
