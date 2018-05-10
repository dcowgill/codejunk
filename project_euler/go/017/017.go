package p017

import (
	"unicode"
)

var dictionary = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

// Only works for numbers between 1 and 999 inclusive
func numberToEnglish(n int) string {
	if n < 1 || n > 9999 {
		return ""
	}

	var (
		hundreds = n - n%100
		tens     = n%100 - n%10
		ones     = n % 10
	)

	var a, b string

	if hundreds != 0 {
		a = dictionary[hundreds/100] + " hundred"
		if tens != 0 || ones != 0 {
			a += " and "
		}
	}

	if tens == 0 {
		b = dictionary[ones] // 0-9
	} else if tens == 10 {
		b = dictionary[tens+ones] // 10-19
	} else if ones == 0 {
		b = dictionary[tens] // 20, 30, ...
	} else {
		b = dictionary[tens] + " " + dictionary[ones] // 21-29, 31-39, ...
	}

	return a + b
}

func countLetters(s string) int {
	n := 0
	for _, r := range s {
		if unicode.IsLetter(r) {
			n++
		}
	}
	return n
}

func solve() int {
	n := countLetters("one thousand")
	for i := 1; i <= 999; i++ {
		n += countLetters(numberToEnglish(i))
	}
	return n
}
