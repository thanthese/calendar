package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type rec struct {
	date time.Time
	desc string
}

type recs []rec

func (rs recs) Len() int      { return len(rs) }
func (rs recs) Swap(i, j int) { rs[i], rs[j] = rs[j], rs[i] }
func (rs recs) Less(i, j int) bool {
	if !rs[i].date.Equal(rs[j].date) {
		return rs[i].date.Before(rs[j].date)
	}
	return rs[i].desc < rs[j].desc
}

type opt int

const (
	regOpt opt = iota
	irrOpt
	sameOpt
	defaultOpt
)

func main() {
	opts := getOptions()
	bytes, _ := ioutil.ReadAll(os.Stdin)
	fmt.Print(toggle(string(bytes), today(), opts))
}

func getOptions() opt {
	irr := flag.Bool("i", false, "Force irregular printout.")
	reg := flag.Bool("r", false, "Force regular printout.")
	same := flag.Bool("s", false, "Force same as input style.")
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
	return defaultOpt
}

func toggle(blob string, today time.Time, opt opt) string {
	recs, irrInput := parseBlob(blob, today)
	if opt == regOpt ||
		(opt == defaultOpt && irrInput) ||
		(opt == sameOpt && !irrInput) {
		return printRegular(recs, today)
	}
	return printIrregular(recs, today)
}

func today() time.Time {
	n := time.Now()
	return time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, time.UTC)
}
