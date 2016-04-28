package main

import (
	"strings"
	"testing"
	"time"
)

func TestToggle(t *testing.T) {
	commonDay := time.Date(2015, 12, 11, 0, 0, 0, 0, time.UTC)
	regularExample := `15.12.09w bears

15.12.12s beets
15.12.13u battlestar
15.12.14m galactica`
	irregularExample := `15.12.09w bears

15.12.12s
    beets
15.12.13u
    battlestar

15.12.14m
    galactica`

	cases := []struct {
		before   string
		expected string
		today    time.Time
		opt      opt
	}{
		// ============================================================
		// test the command line options

		{regularExample, irregularExample, commonDay, toggleOpt},
		{irregularExample, regularExample, commonDay, toggleOpt},

		{regularExample, regularExample, commonDay, sameOpt},
		{irregularExample, irregularExample, commonDay, sameOpt},

		{regularExample, regularExample, commonDay, regOpt},
		{irregularExample, regularExample, commonDay, regOpt},

		{regularExample, irregularExample, commonDay, irrOpt},
		{irregularExample, irregularExample, commonDay, irrOpt},

		// ============================================================
		// test today logic

		{`15.12.09w bears

15.12.12s beets
15.12.13u battlestar
15.12.14m galactica`,
			`15.12.09w bears
15.12.12s beets
15.12.13u battlestar

15.12.14m galactica`,
			time.Date(2015, 12, 13, 0, 0, 0, 0, time.UTC),
			regOpt,
		},
		{`15.12.09w bears

15.12.12s beets
15.12.13u battlestar
15.12.14m galactica`,
			`15.12.09w bears
15.12.12s beets
15.12.13u battlestar
15.12.14m galactica`,
			time.Date(2015, 12, 15, 0, 0, 0, 0, time.UTC),
			regOpt,
		},
		{`15.12.09w bears

15.12.12s beets
15.12.13u battlestar
15.12.14m galactica`,
			`
15.12.09w bears
15.12.12s beets
15.12.13u battlestar
15.12.14m galactica`,
			time.Date(2015, 12, 8, 0, 0, 0, 0, time.UTC),
			regOpt,
		},

		// ============================================================
		// complicated irregular printing

		{`16.12.09f bears
16.12.10s another
16.12.12m beets test
16.12.13t battlestar
16.12.14w galactica
16.12.14w test
16.12.14w sort me 1
16.12.16f sort me 4
16.12.16f sort me 2
16.12.16f sort me 5
16.12.16f sort me 3`,
			`16.12.09f bears
16.12.10s another

16.12.12m
    beets test
16.12.13t
    battlestar
16.12.14w
    galactica
    sort me 1
    test
16.12.15r
16.12.16f
    sort me 2
    sort me 3
    sort me 4
    sort me 5`,
			time.Date(2016, 12, 11, 0, 0, 0, 0, time.UTC),
			irrOpt,
		},

		// ============================================================
		// complicated irregular reading

		{`16.12.09f bears
16.12.10s another

16.12.12m
    beets test
16.12.13t
    battlestar
16.12.14w
    galactica
    sort me 1
    test
16.12.15r
16.12.16f
    sort me 2
    sort me 3
    sort me 4
    sort me 5`,
			`16.12.09f bears
16.12.10s another

16.12.12m beets test
16.12.13t battlestar
16.12.14w galactica
16.12.14w sort me 1
16.12.14w test
16.12.16f sort me 2
16.12.16f sort me 3
16.12.16f sort me 4
16.12.16f sort me 5`,
			time.Date(2016, 12, 11, 0, 0, 0, 0, time.UTC),
			regOpt,
		},

		// ============================================================
		// todos

		{`test of
things to
come and such

# tickler
15.12.09w bears

15.12.12s beets
15.12.13u battlestar
15.12.14m galactica`,
			`test of
things to
come and such

# tickler

15.12.09w bears
15.12.12s beets
15.12.13u battlestar
15.12.14m galactica`,
			time.Date(2015, 12, 8, 0, 0, 0, 0, time.UTC),
			regOpt,
		},
		{`test of
things to
come and such


# tickler




15.12.09w bears

15.12.12s beets
15.12.13u battlestar
15.12.14m galactica`,
			`test of
things to
come and such


# tickler
15.12.09w bears
15.12.12s beets

15.12.13u battlestar
15.12.14m galactica`,
			time.Date(2015, 12, 12, 0, 0, 0, 0, time.UTC),
			regOpt,
		},
		{`test of
things to
come and such
# tickler




15.12.09w bears

15.12.12s beets
15.12.13u battlestar
15.12.14m galactica`,
			`test of
things to
come and such
# tickler
15.12.09w bears
15.12.12s beets

15.12.13u battlestar
15.12.14m galactica`,
			time.Date(2015, 12, 12, 0, 0, 0, 0, time.UTC),
			regOpt,
		},
	}

	for _, c := range cases {
		got := transform(c.before, c.today, c.opt)
		if strings.TrimSpace(c.expected) != strings.TrimSpace(got) {
			t.Errorf("=== After (%s) ===\n\n%s \n\n=== was expecing ===\n\n%s\n\n=== but got ===\n\n%s",
				prettyOption(c.opt), c.before, c.expected, got)
		}
	}
}

func prettyOption(opt opt) string {
	switch opt {
	case toggleOpt:
		return "toggle"
	case regOpt:
		return "reg"
	case irrOpt:
		return "irr"
	case sameOpt:
		return "same"
	}
	return "bad opt found"
}
