package main

import (
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

func main() {
	bytes, _ := ioutil.ReadAll(os.Stdin)
	blob := string(bytes)
	today := today()
	recs, irregular := parseBlob(blob, today)
	if irregular {
		printRegular(recs, today)
	} else {
		printIrregular(recs, today)
	}
}

func today() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
}
