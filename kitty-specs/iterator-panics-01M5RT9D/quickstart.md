# Validation quickstart

From the lane workspace after implementation, using a writable GOCACHE:

```sh
timeout 180 go test ./...
timeout 180 go test -race ./...
timeout 180 go vet ./...
gofmt -l .
GOARCH=386 timeout 180 go test ./...
timeout 120 go test -run='^$' -fuzz=FuzzIterate -fuzztime=60s -parallel=2
```

Formatting must emit no paths. Fuzz timeout is failure, not success. Read the
fuzz target's documented restrictions before interpreting its coverage. Compare
the existing reproduction program against its recorded unfixed output. Preserve
red/green and individual-fix sensitivity evidence, and reviewer model/date.

Consolidation target is fix/iterator-panics. Open a private PR to master only;
no master merge, release or upstream/public posting is authorized in this run.
