//go:build ignore
// +build ignore

// Golden DST compatibility driver. Build this new-import driver from candidate cwd:
//
//	GOWORK=off GOPROXY=off go build -o /tmp/dst-golden ./docs/repro/05-07-backlog/golden/main.go
//
// Historical sources require the frozen old-import driver exported from snapshot
// 737ae01be40048d1d2932dbafb623230efa8d361, built from the historical source cwd.
// See README.md beside this driver for exact source/driver build commands.
// For this module migration, collect fresh -out files for the same absolute corpus
// path from pre-migration and candidate binaries, then compare their bytes with cmp.
// The historical -baseline/-current comparator requires changed rows and rejects
// zero differences; use it only for the classified comparison against 76c646d.
// Each changed row must have an explicit RFC rationale and exact approved
// current-row SHA256 in the corpus. Unknown changes fail; changed sensitivity
// controls, matching IDs/counts, and nonempty outputs guard a broken comparison.
// Each fixture gets its own killable subprocess (2 seconds); first16 values or
// exhaustion only. A timeout is a recorded result, never a silent omission.
package main

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	rrule "github.com/kentonium3/rrule-go"
)

type fixture struct {
	ID              string `json:"id"`
	Source          string `json:"source"`
	Rule            string `json:"rule"`
	Location        string `json:"location,omitempty"`
	Set             bool   `json:"set,omitempty"`
	LocationKind    string `json:"location_kind,omitempty"`
	Offset          int    `json:"offset,omitempty"`
	Action          string `json:"action,omitempty"`
	Start           string `json:"start,omitempty"`
	Rationale       string `json:"rationale,omitempty"`
	ApprovedCurrent string `json:"approved_current_sha256,omitempty"`
}
type corpus struct {
	Baseline       string          `json:"baseline"`
	BaselineSHA256 string          `json:"baseline_sha256"`
	Provenance     json.RawMessage `json:"provenance"`
	Cases          []fixture       `json:"cases"`
}
type row struct {
	ID          string   `json:"id"`
	String      string   `json:"string"`
	Error       string   `json:"error"`
	Occurrences []string `json:"occurrences"`
	Exhausted   bool     `json:"exhausted"`
}

func evaluate(c fixture) (out row) {
	out.ID = c.ID
	out.Occurrences = []string{}
	defer func() {
		if v := recover(); v != nil {
			out.Error = fmt.Sprintf("panic: %v", v)
		}
	}()
	loc := time.UTC
	if c.LocationKind == "fixed" {
		loc = time.FixedZone(c.Location, c.Offset)
	} else if c.LocationKind == "tzif" {
		data := make([]byte, 54)
		copy(data, "TZif")
		binary.BigEndian.PutUint32(data[36:40], 1)
		binary.BigEndian.PutUint32(data[40:44], 4)
		binary.BigEndian.PutUint32(data[44:48], uint32(c.Offset))
		copy(data[50:], "XTZ\x00")
		var e error
		loc, e = time.LoadLocationFromTZData(c.Location, data)
		if e != nil {
			out.Error = e.Error()
			return
		}
	} else if c.Location != "" {
		var e error
		loc, e = time.LoadLocation(c.Location)
		if e != nil {
			out.Error = e.Error()
			return
		}
	}
	if c.Set {
		set, e := rrule.StrSliceToRRuleSetInLoc(strings.Split(c.Rule, "\n"), loc)
		if e != nil {
			out.Error = e.Error()
			return
		}
		if c.Action == "set-replace-rule" {
			replacement, e := rrule.StrToRRule("FREQ=DAILY;COUNT=3")
			if e != nil {
				out.Error = e.Error()
				return
			}
			set.RRule(replacement)
		}
		out.String = set.String()
		next := set.Iterator()
		for i := 0; i < 16; i++ {
			v, ok := next()
			if !ok {
				out.Exhausted = true
				break
			}
			out.Occurrences = append(out.Occurrences, v.UTC().Format(time.RFC3339))
		}
		return
	}
	o, e := rrule.StrToROptionInLocation(c.Rule, loc)
	if e != nil {
		out.Error = e.Error()
		return
	}
	if o.Dtstart.IsZero() {
		start := c.Start
		if start == "" {
			start = "20260105T100000"
		}
		o.Dtstart, e = time.ParseInLocation(rrule.LocalDateTimeFormat, start, loc)
		if e != nil {
			out.Error = e.Error()
			return
		}
	}
	if c.Action == "programmatic" {
		o = &rrule.ROption{Freq: o.Freq, Dtstart: o.Dtstart, Until: o.Until, Count: o.Count}
	}
	if c.Action == "copy" || c.Action == "set-after-copy" {
		copied := *o
		o = &copied
	}
	r, e := rrule.NewRRule(*o)
	if e != nil {
		out.Error = e.Error()
		return
	}
	if c.Action == "start-setter" {
		r.DTStart(o.Dtstart)
	}
	if c.Action == "until-setter" {
		r.Until(o.Until)
	}
	out.String = r.String()
	next := r.Iterator()
	if c.Action == "set-after-copy" {
		set := rrule.Set{}
		set.RRule(r)
		out.String, next = set.String(), set.Iterator()
	}
	for i := 0; i < 16; i++ {
		v, ok := next()
		if !ok {
			out.Exhausted = true
			break
		}
		out.Occurrences = append(out.Occurrences, v.UTC().Format(time.RFC3339))
	}
	return
}
func read(path string, out interface{}) {
	b, e := os.ReadFile(path)
	if e != nil {
		panic(e)
	}
	if e = json.Unmarshal(b, out); e != nil {
		panic(e)
	}
}
func canonical(v interface{}) []byte {
	b, e := json.Marshal(v)
	if e != nil {
		panic(e)
	}
	return b
}
func digest(v interface{}) string { h := sha256.Sum256(canonical(v)); return hex.EncodeToString(h[:]) }
func main() {
	corpusPath := flag.String("corpus", "docs/repro/05-07-backlog/golden/corpus-dst.json", "corpus path")
	outPath := flag.String("out", "", "output JSON path")
	one := flag.Int("one", -1, "internal fixture index")
	basePath := flag.String("baseline", "", "compare baseline JSON")
	currentPath := flag.String("current", "", "compare current JSON")
	flag.Parse()
	var c corpus
	read(*corpusPath, &c)
	if *one >= 0 {
		if e := json.NewEncoder(os.Stdout).Encode(evaluate(c.Cases[*one])); e != nil {
			panic(e)
		}
		return
	}
	if *basePath != "" {
		var base, current []row
		read(*basePath, &base)
		raw, e := os.ReadFile(*basePath)
		if e != nil {
			panic(e)
		}
		hash := sha256.Sum256(raw)
		if hex.EncodeToString(hash[:]) != c.BaselineSHA256 {
			panic("baseline artifact SHA256 mismatch")
		}
		read(*currentPath, &current)
		if len(c.Cases) == 0 || len(base) != len(c.Cases) || len(current) != len(c.Cases) {
			panic("corpus/output size mismatch")
		}
		changes, unexplained := 0, 0
		for i, f := range c.Cases {
			if base[i].ID != f.ID || current[i].ID != f.ID {
				panic("row ID/order mismatch")
			}
			if string(canonical(base[i])) == string(canonical(current[i])) {
				continue
			}
			changes++
			if f.Rationale == "" || f.ApprovedCurrent != digest(current[i]) {
				unexplained++
				fmt.Printf("UNEXPLAINED %s baseline=%s current=%s\n", f.ID, canonical(base[i]), canonical(current[i]))
				continue
			}
			fmt.Printf("CLASSIFIED %s: %s\n  baseline=%s\n  current=%s\n", f.ID, f.Rationale, canonical(base[i]), canonical(current[i]))
		}
		fmt.Printf("rows=%d changed=%d unexplained=%d\n", len(base), changes, unexplained)
		if changes == 0 || unexplained != 0 {
			os.Exit(1)
		}
		return
	}
	executable, e := os.Executable()
	if e != nil {
		panic(e)
	}
	results := make([]row, 0, len(c.Cases))
	for i, f := range c.Cases {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		cmd := exec.CommandContext(ctx, executable, "-corpus", *corpusPath, "-one", fmt.Sprint(i))
		b, err := cmd.Output()
		timeout := ctx.Err() != nil
		cancel()
		var r row
		if timeout {
			r = row{ID: f.ID, Error: "timeout after 2s (bounded pathological fixture)", Occurrences: []string{}}
		} else if err != nil {
			panic(fmt.Sprintf("%s: %s: %s", f.ID, err, b))
		} else if e = json.Unmarshal(b, &r); e != nil {
			panic(e)
		}
		results = append(results, r)
	}
	b, e := json.MarshalIndent(results, "", "  ")
	if e != nil {
		panic(e)
	}
	if *outPath == "" {
		fmt.Println(string(b))
	} else if e = os.WriteFile(*outPath, append(b, '\n'), 0644); e != nil {
		panic(e)
	}
}
