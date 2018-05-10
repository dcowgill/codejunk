package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	sort.Slice(lines, func(i, j int) bool {
		return len(lines[i]) < len(lines[j])
	})
	for _, line := range lines {
		fmt.Println(line)
	}
}
