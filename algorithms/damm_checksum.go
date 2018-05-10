// https://en.wikipedia.org/wiki/Damm_algorithm
package main

import "fmt"

var quasi = [10][10]int8{
	{0, 3, 1, 7, 5, 9, 8, 6, 4, 2},
	{7, 0, 9, 2, 1, 5, 4, 8, 6, 3},
	{4, 2, 0, 6, 8, 7, 1, 3, 5, 9},
	{1, 7, 5, 0, 9, 8, 3, 4, 2, 6},
	{6, 1, 2, 3, 0, 4, 5, 9, 7, 8},
	{3, 6, 7, 4, 2, 0, 9, 5, 8, 1},
	{5, 8, 6, 9, 7, 2, 0, 1, 3, 4},
	{8, 9, 4, 5, 3, 6, 2, 0, 1, 7},
	{9, 4, 3, 8, 6, 1, 7, 2, 0, 5},
	{2, 5, 8, 1, 4, 3, 6, 7, 9, 0},
}

// ChecksumDigit returns the checksum digit for n, as a string. Returns
// the empty string if n does not consist entirely of digits.
func Checksum(n string) string {
	var c int8
	for i := 0; i < len(n); i++ {
		x := n[i] - '0'
		if x < 0 || x > 9 {
			return ""
		}
		c = quasi[c][x]
	}
	return string(c + '0')
}

// IsValid returns true if n's checksum digit is correct.
func IsValid(n string) bool {
	return Checksum(n) == "0"
}

func main() {
	for i, s := range []string{"0", "1", "572", "5724", "hello, world", "", "1111", "2398412098410293840182401924"} {
		fmt.Println(i, s, Checksum(s), IsValid(s+Checksum(s)))
	}
	fmt.Println(IsValid("23984120984102938401824019249"))
}
