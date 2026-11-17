A combined command accidentally used pristine baseline cwd for a no-op str.go replacement and TestTimeForms|TestDST. Replacement target was absent; all existing baseline root tracked files remain byte-identical to76c646d, proven in baseline-integrity-after-cwd-error.txt. The focused run had no matching WP02 tests and is INVALID, preserved as INVALID-focused-wrong-baseline-cwd.txt. No baseline behavioral content changed. The intended baseline build used current external golden main.go to resolve imports against the pinned baseline. Commands are recorded in tool transcript. Recovery: verify bytes, mark invalid evidence, rerun intended replacement/test in allocated lane.

Exact mistaken command (exec cwd `/tmp/rrule-backlog-v183-plus-duplicates`):

```sh
python3 - <<'PY'
from pathlib import Path
p=Path('str.go');s=p.read_text().replace('expect DTSTART or RRULE but: %s','expect DTSTART but: %s');p.write_text(s)
PY
GOCACHE=/tmp/rrule-go-cache go test -run '^TestTimeForms|^TestDST' -v -count=1 > /tmp/rrule-WP02-logs/green-focused.txt 2>&1
result=$?
print "exit=$result" >> /tmp/rrule-WP02-logs/green-focused.txt
print "focused exit=$result"
GOCACHE=/tmp/rrule-go-cache go build -o /tmp/rrule-WP02-golden-baseline ~/repos/rrule-go/.worktrees/teambition-backlog-01M5TEBG-lane-b/docs/repro/05-07-backlog/golden/main.go
```

Tool output was `focused exit=0`, exec exit0. The preserved INVALID output explicitly contains `testing: warning: no tests to run`; it does not count as any gate. State before recovery: lane HEAD ef373a5872dba30063a6a706275a28e46508af77, intended lane parser diagnostic not yet restored, snapshot replacement target absent. Source bytes were verified before the correct lane repair, using `git show 76c646dde8d33ff7c3325df0613b31b058243bea:<filename>` for every snapshot file. All15 matched. The initial no-op changed only filesystem mtime. Correct lane replacement restored diagnostic wording before genuine TimeForms+DST tests. This was an agent cwd fault, not a Spec Kitty fault or a waived gate.
