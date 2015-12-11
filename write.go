package main

import (
	"fmt"
	"time"
)

const separator = "----------------------------------------"

func printIrregular(recs recs, today time.Time) {
	if len(recs) == 0 {
		return
	}
	min, max := recs[0].date, recs[len(recs)-1].date
	date := min
	tomorrow := today.AddDate(0, 0, 1)
	for !date.After(max) {
		if date.Equal(tomorrow) {
			fmt.Println()
		}
		if !date.Equal(tomorrow) && date.After(today) && time.Monday == date.Weekday() {
			fmt.Println(separator)
		}
		ms := matchingDates(date, recs)
		if date.After(today) {
			fmt.Println(prettyDate(date))
			for _, m := range ms {
				fmt.Println(prettyDescOnly(m.desc))
			}
		} else {
			for _, m := range ms {
				fmt.Println(prettyRegular(m))
			}
		}
		date = date.AddDate(0, 0, 1)
	}
}

func matchingDates(date time.Time, recs recs) (matches recs) {
	for _, r := range recs {
		if r.date == date {
			matches = append(matches, r)
		}
	}
	return
}

func printRegular(recs recs, today time.Time) {
	foundTomorrow := false
	for _, r := range recs {
		if r.date.After(today) && !foundTomorrow {
			fmt.Println()
			foundTomorrow = true
		}
		fmt.Println(prettyRegular(r))
	}
}

func prettyRegular(r rec) string {
	return prettyDate(r.date) + " " + r.desc
}

func prettyDescOnly(s string) string {
	return "    " + s
}

func prettyDate(d time.Time) string {
	return fmt.Sprintf("%02d.%02d.%02d%c",
		d.Year()-2000,
		int(d.Month()),
		d.Day(),
		"umtwrfs"[d.Weekday()])
}
