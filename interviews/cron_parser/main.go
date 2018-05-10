package main

import (
	"fmt"
	"os"
	"strconv"

	"./crontab"
)

func main() {
	// We expect exactly 1 command-line argument.
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: parsecron <crontab-entry>")
		os.Exit(1)
	}

	// Parse the crontab line.
	job, err := crontab.ParseJob(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Display the result.
	printField := func(label, value string) { fmt.Printf("%-14s %s\n", label, value) }
	printField("minute", joinInts(job.Minute, " "))
	printField("hour", joinInts(job.Hour, " "))
	printField("day of month", joinInts(job.MonthDay, " "))
	printField("month", joinInts(job.Month, " "))
	printField("day of week", joinInts(job.Weekday, " "))
	printField("command", job.Command)
}

// Joins the ints in xs with sep.
func joinInts(xs []int, sep string) string {
	var s, t string
	for _, x := range xs {
		s += t
		s += strconv.Itoa(x)
		t = sep
	}
	return s
}
