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
	defaultOpt
)

func main() {
	opts := getOptions()
	bytes, _ := ioutil.ReadAll(os.Stdin)
	fmt.Print(toggle(string(bytes), today(), opts))
}

func getOptions() opt {
	irr := flag.Bool("i", false, "Force irregular printout.")
	reg := flag.Bool("w", false, "Force regular printout.")
	flag.Parse()

	if *irr && *reg {
		fmt.Println("ERROR: Can't set both -i and -r flags.")
		os.Exit(1)
	}
	if *irr {
		return irrOpt
	}
	if *reg {
		return regOpt
	}
	return defaultOpt
}

func toggle(blob string, today time.Time, opt opt) string {
	recs, irregular := parseBlob(blob, today)
	if opt == irrOpt || (opt == defaultOpt && irregular) {
		return printRegular(recs, today)
	}
	return printIrregular(recs, today)
}

func today() time.Time {
	n := time.Now()
	return time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, time.UTC)
}
