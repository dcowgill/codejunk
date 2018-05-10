// https://fivethirtyeight.com/features/pirates-monkeys-and-coconuts-oh-my/
package main

import (
	"fmt"
	"log"
)

/*

x = total in round n
y = total in round n+1

(x-1)*6/7 = y
x = (y/6)*7 + 1

x7 = (((((((((((((x/6)*7 + 1)/6)*7 + 1)/6)*7 + 1)/6)*7 + 1)/6)*7 + 1)/6)*7 + 1)/6)*7 + 1
x6 = (((((((((((x/6)*7 + 1)/6)*7 + 1)/6)*7 + 1)/6)*7 + 1)/6)*7 + 1)/6)*7 + 1
x5 = (((((((((x/6)*7 + 1)/6)*7 + 1)/6)*7 + 1)/6)*7 + 1)/6)*7 + 1
x4 = (((((((x/6)*7 + 1)/6)*7 + 1)/6)*7 + 1)/6)*7 + 1
x3 = (((((x/6)*7 + 1)/6)*7 + 1)/6)*7 + 1
x2 = (((x/6)*7 + 1)/6)*7 + 1
x1 = (x/6)*7 + 1

x7 = (823543*x + 3261642) / 279936

*/

// Just brute-force it.
func try(each int) bool {
	x := each * 7
	for i := 0; i < 7; i++ {
		if x%6 != 0 {
			return false
		}
		x = 7*x/6 + 1
	}
	return true
}

// Verify by Diophantine equation.
func check(each int) bool {
	x := each * 7
	return (823543*x+3261642)%279936 == 0
}

func main() {
	i := 1
	for {
		if try(i) {
			if !check(i) {
				log.Fatalf("try(%d) succeeded but check(%d) failed", i, i)
			}
			break
		}
		if check(i) {
			log.Fatalf("try(%d) failed but check(%d) succeeded", i, i)
		}
		i++
	}
	fmt.Printf("each = %d, final pile = %d\n", i, i*7)
}
