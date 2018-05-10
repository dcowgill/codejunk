package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

// A dictionary is a mapping from phone numbers (consisting of digits
// only) to the words which map to that number based on wordToDigits.
type Dict map[string][]string

// Reads the dictionary from the given path.
func loadDict(path string) Dict {
	m := make(Dict)
	foreachLine(path, func(word string) {
		k := wordToDigits(word)
		m[k] = append(m[k], word)
	})
	return m
}

// Calls f for each line in the file, passing it the line with newline
// removed. Panics if an i/o error occurs.
func foreachLine(path string, f func(string)) {
	fp, err := os.Open(path)
	must(err)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		f(scanner.Text())
	}
	must(scanner.Err())
	must(fp.Close())
}

// Performs the following mapping for each letter in w to a digit:
//
//	E | J N Q | R W X | D S Y | F T | A M | C I V | B K U | L O P | G H Z
//	e | j n q | r w x | d s y | f t | a m | c i v | b k u | l o p | g h z
//	0 |   1   |   2   |   3   |  4  |  5  |   6   |   7   |   8   |   9
//
func wordToDigits(s string) string {
	t := make([]rune, 0, len(s))
	for _, c := range s {
		if d := charToDigit(c); d >= 0 {
			t = append(t, d)
		}
	}
	return string(t)
}

// Returns -1 if the rune has no corresponding digit.
func charToDigit(r rune) rune {
	switch unicode.ToLower(r) {
	case 'e':
		return '0'
	case 'j', 'n', 'q':
		return '1'
	case 'r', 'w', 'x':
		return '2'
	case 'd', 's', 'y':
		return '3'
	case 'f', 't':
		return '4'
	case 'a', 'm':
		return '5'
	case 'c', 'i', 'v':
		return '6'
	case 'b', 'k', 'u':
		return '7'
	case 'l', 'o', 'p':
		return '8'
	case 'g', 'h', 'z':
		return '9'
	}
	return -1 // skip
}

// Returns a string containing only the digit characters in p.
func justDigits(p string) string {
	t := make([]rune, 0, len(p))
	for _, c := range p {
		if c >= '0' && c <= '9' {
			t = append(t, c)
		}
	}
	return string(t)
}

// Returns each possible translation of digits into a sequence of words.
// digits should consist only of digit characters (that is, it is a
// phone number with all non-numeric characters omitted). A single digit
// may be inserted in a translation if no word can be found starting at
// its position, but consecutive digits are not allowed; if foundWord is
// true, the caller was able to add a word, or we are at the top of the
// call tree.
func translations(dict Dict, digits string, foundWord bool) [][]string {
	if len(digits) == 0 {
		return [][]string{{}}
	}
	var result [][]string
	emit := func(start string, rest []string) {
		result = append(result, append([]string{start}, rest...))
	}
	for i := 0; i < len(digits); i++ {
		for _, w := range dict[prefix(digits, i)] {
			for _, rest := range translations(dict, suffix(digits, i), true) {
				emit(w, rest)
			}
		}
	}
	if len(result) == 0 && foundWord {
		digit := string(digits[0])
		for _, rest := range translations(dict, digits[1:], false) {
			emit(digit, rest)
		}
	}
	return result
}

func prefix(s string, n int) string { return s[:len(s)-n] }
func suffix(s string, n int) string { return s[len(s)-n:] }

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: main <dict-file> <phone-file>\n")
		os.Exit(1)
	}
	dict := loadDict(os.Args[1])
	foreachLine(os.Args[2], func(phone string) {
		for _, words := range translations(dict, justDigits(phone), true) {
			fmt.Printf("%s: %s\n", phone, strings.Join(words, " "))
		}
	})
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
