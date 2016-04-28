package main

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var recRegex = regexp.MustCompile(`^(\d\d).(\d\d).(\d\d)[mtwrfsu]?(.*)`)

// Parse some multi-line text blob into usable recs.
//
// Considers everything up to and including the "# tickler" line to be todo
// front matter.
//
// Within the calendar section very non-empty, non-whitespace-only line gets a
// date associated with it. If a date is given, it gets that one. If no date
// was given, it uses the last one. If there was no last one, it uses today.
//
// Returned recs are sorted.
//
// The blob was regular if every non-empty line had a defined date. It was
// irregular if we had to do some interpolating.
//
// Whitespace is treated without mercy. Empty lines are filtered out and
// everything is trimmed.
func parseBlob(blob string, today time.Time) (doc doc, irregular bool) {
	lines := strings.Split(blob, "\n")
	for i, line := range lines {
		if line == "# tickler" {
			doc.todos = lines[:i+1]
			lines = lines[i+1:]
			break
		}
	}
	lastDate := today
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fields := recRegex.FindStringSubmatch(line)
		if len(fields) == 0 {
			doc.recs = append(doc.recs, rec{lastDate, line})
			irregular = true
			continue
		}
		y, _ := strconv.Atoi(fields[1])
		m, _ := strconv.Atoi(fields[2])
		d, _ := strconv.Atoi(fields[3])
		lastDate = time.Date(2000+y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
		desc := strings.TrimSpace(fields[4])
		if desc == "" {
			continue
		}
		doc.recs = append(doc.recs, rec{lastDate, desc})
	}
	sort.Sort(doc.recs)
	return
}
