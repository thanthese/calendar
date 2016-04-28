package main

import (
	"fmt"
	"time"
)

func printIrregular(doc doc, today time.Time) (out string) {
	out += prettyTodos(doc)
	if len(doc.recs) == 0 {
		return
	}
	min, max := doc.recs[0].date, doc.recs[len(doc.recs)-1].date
	if today.Before(min) {
		min = today
	}
	tomorrow := today.AddDate(0, 0, 1)
	for date := min; !date.After(max); date = date.AddDate(0, 0, 1) {
		if date.Equal(tomorrow) ||
			(date.After(today) && time.Monday == date.Weekday()) {
			out += "\n"
		}
		ms := matchingDates(date, doc.recs)
		if date.After(today) {
			out += prettyDate(date) + "\n"
			for _, m := range ms {
				out += prettyDescOnly(m.desc) + "\n"
			}
		} else {
			for _, m := range ms {
				out += prettyRegular(m) + "\n"
			}
		}
	}
	return
}

func matchingDates(date time.Time, recs recs) (matches recs) {
	for _, r := range recs {
		if r.date == date {
			matches = append(matches, r)
		}
	}
	return
}

func printRegular(doc doc, today time.Time) (out string) {
	out += prettyTodos(doc)
	foundBreak := false
	for _, r := range doc.recs {
		if r.date.After(today) && !foundBreak {
			foundBreak = true
			out += "\n"
		}
		out += prettyRegular(r) + "\n"
	}
	return
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

func prettyTodos(doc doc) (out string) {
	if len(doc.todos) == 0 {
		return
	}
	for _, t := range doc.todos {
		out += t + "\n"
	}
	return
}
