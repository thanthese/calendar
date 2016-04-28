package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// A rec is a single entry in the calendar, like "16.04.26t some text".
type rec struct {
	date time.Time
	desc string
}

// Many recs, like a whole calendar. This struct just exists so I can sort
// recs.
type recs []rec

func (rs recs) Len() int      { return len(rs) }
func (rs recs) Swap(i, j int) { rs[i], rs[j] = rs[j], rs[i] }
func (rs recs) Less(i, j int) bool {
	if !rs[i].date.Equal(rs[j].date) {
		return rs[i].date.Before(rs[j].date)
	}
	return rs[i].desc < rs[j].desc
}

// The whole document, including the todo front matter. "# tickler" is included
// in the "todo" section.
type doc struct {
	todos []string
	recs  recs
}

// options from the command line
type opt int

const (
	toggleOpt opt = iota // default
	regOpt
	irrOpt
	sameOpt
)

func parseCommandLineArgs() opt {
	reg := flag.Bool("regular", false, "Use regular format.")
	irr := flag.Bool("irregular", false, "Use irregular format.")
	same := flag.Bool("same", false, "Use input's format.")
	flag.Parse()

	if (*irr && *reg) || (*irr && *same) || (*reg && *same) {
		fmt.Println("ERROR: Cannot set more than one flag.")
		os.Exit(1)
	}

	if *irr {
		return irrOpt
	}
	if *reg {
		return regOpt
	}
	if *same {
		return sameOpt
	}
	return toggleOpt
}

func main() {
	opts := parseCommandLineArgs()
	bytes, _ := ioutil.ReadAll(os.Stdin)
	fmt.Print(transform(string(bytes), today(), opts))
}

// The heart of the program, changes a blob of text to another blob of text.
func transform(blob string, today time.Time, opt opt) string {
	doc, irrInput := parseBlob(blob, today)
	if opt == regOpt ||
		(opt == toggleOpt && irrInput) ||
		(opt == sameOpt && !irrInput) {
		return printRegular(doc, today)
	}
	return printIrregular(doc, today)
}

func today() time.Time {
	n := time.Now()
	return time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, time.UTC)
}
