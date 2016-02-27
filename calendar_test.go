package main

import (
	"strings"
	"testing"
	"time"
)

func TestToggle(t *testing.T) {
	today := time.Date(2015, 12, 11, 0, 0, 0, 0, time.UTC)
	cases := []struct {
		a string
		b string
	}{
		{`15.12.09w bears

15.12.12s beets
15.12.13u battlestar

15.12.14m galactica`, `15.12.09w bears

15.12.12s
    beets
15.12.13u
    battlestar

15.12.14m
    galactica`}}

	for _, c := range cases {
		caseTest(c.a, c.b, today, t)
		caseTest(c.b, c.a, today, t)
	}
}

func caseTest(before string, expected string, today time.Time, t *testing.T) {
	got := toggle(before, today, defaultOpt)
	if strings.TrimSpace(expected) != strings.TrimSpace(got) {
		t.Errorf("=== After ===\n\n%s \n\n=== was expecing ===\n\n%s\n\n=== but got ===\n\n%s",
			before, expected, got)
	}
}
