/*

Using names.txt (right click and 'Save Link/Target As...'), a 46K text file
containing over five-thousand first names, begin by sorting it into
alphabetical order. Then working out the alphabetical value for each name,
multiply this value by its alphabetical position in the list to obtain a name
score.

For example, when the list is sorted into alphabetical order, COLIN, which is
worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN
would obtain a score of 938 x 53 = 49714.

What is the total of all the name scores in the file?

*/

package p022

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strings"
)

func readNames() []string {
	var (
		err     error
		file    *os.File
		records [][]string
		names   []string
	)
	file, err = os.Open("022.txt")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(file)
	records, err = reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range records {
		for _, name := range record {
			names = append(names, name)
		}
	}
	return names
}

func alphaValue(s string) int {
	n := 0
	for _, r := range strings.ToUpper(s) {
		n += int(r-'A') + 1
	}
	return n
}

func solve() int64 {
	names := readNames()
	sort.Strings(names)
	var sum int64
	for i, name := range names {
		sum += int64((i + 1) * alphaValue(name))
	}
	return sum
}
