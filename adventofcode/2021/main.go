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
	"fmt"
	"os"
)

func main() {
	d01.Run()
	d02.Run()
	d03.Run()
	d04.Run()
	d05.Run()
	d06.Run()
	d07.Run()
	d08.Run()
	d09.Run()
	d10.Run()
	d11.Run()
	d12.Run()
	d13.Run()
	d14.Run()
}

func usage(format string, args ...any) {
	fmt.Printf(format, args...)
	os.Exit(1)
}
