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
// Every non-empty, non-whitespace-only line gets a date associated with it. If
// a date is given, it gets that one. If no date was given, it uses the last
// one. If there was no last one, it uses today.
//
// Returned recs are sorted.
//
// The blob was regular if every non-empty line had a defined date. It was
// irregular if we had to do some interpolating.
//
// Whitespace is treated without mercy. Empty lines are filtered out and
// everything is trimmed.
func parseBlob(blob string, today time.Time) (recs recs, irregular bool) {
	lastDate := today
	for _, line := range strings.Split(blob, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fields := recRegex.FindStringSubmatch(line)
		if len(fields) == 0 {
			recs = append(recs, rec{lastDate, line})
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
		recs = append(recs, rec{lastDate, desc})
	}
	sort.Sort(recs)
	return recs, irregular
}
