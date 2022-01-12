package main

import (
	"adventofcode2021/d01"
	"adventofcode2021/d02"
	"adventofcode2021/d03"
	"adventofcode2021/d04"
	"adventofcode2021/d05"
	"adventofcode2021/d06"
	"adventofcode2021/d07"
	"adventofcode2021/d08"
	"adventofcode2021/d09"
	"adventofcode2021/d10"
	"adventofcode2021/d11"
	"adventofcode2021/d12"
	"adventofcode2021/d13"
	"adventofcode2021/d14"
	"adventofcode2021/d15"
	"adventofcode2021/d16"
	"adventofcode2021/d17"
	"adventofcode2021/d18"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	var (
		cpuProf = flag.String("cpuprofile", "", "write cpu profile to `file`")
		// memProf = flag.String("memprofile", "", "write memory profile to `file`")
		day = flag.Int("day", 0, "which day (1-25, or 0 to run all days)")
	)
	flag.Parse()

	if *day < 0 || *day > 25 {
		usage("-day must be between 1-25 inclusive (or 0)")
	}

	if *cpuProf != "" {
		f, err := os.Create(*cpuProf)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer func() {
			if err := f.Close(); err != nil {
				panic(fmt.Sprintf("i/o error closing CPU profile: %v"))
			}
		}()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	if *day == 0 {
		for i := 1; i <= 25; i++ {
			runDay(i)
		}
	} else {
		runDay(*day)
	}
}

func runDay(day int) {
	switch day {
	case 1:
		d01.Run()
	case 2:
		d02.Run()
	case 3:
		d03.Run()
	case 4:
		d04.Run()
	case 5:
		d05.Run()
	case 6:
		d06.Run()
	case 7:
		d07.Run()
	case 8:
		d08.Run()
	case 9:
		d09.Run()
	case 10:
		d10.Run()
	case 11:
		d11.Run()
	case 12:
		d12.Run()
	case 13:
		d13.Run()
	case 14:
		d14.Run()
	case 15:
		d15.Run()
	case 16:
		d16.Run()
	case 17:
		d17.Run()
	case 18:
		d18.Run()
		// case 19:
		// 	d19.Run()
		// case 20:
		// 	d20.Run()
		// case 21:
		// 	d21.Run()
		// case 22:
		// 	d22.Run()
		// case 23:
		// 	d23.Run()
		// case 24:
		// 	d24.Run()
		// case 25:
		// 	d25.Run()
	}
}

func usage(format string, args ...any) {
	fmt.Printf(format, args...)
	os.Exit(1)
}
