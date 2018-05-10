// What's the best way to attack in the boardgame "Risk: Legacy"?
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
)

func d6() int { return 1 + rand.Intn(6) }

func applyMod(rolls []int, mod int) {
	sort.Ints(rolls)
	if mod != 0 {
		i := len(rolls) - 1
		rolls[i] += mod
		if rolls[i] > 6 {
			rolls[i] = 6
		} else if rolls[i] < 1 {
			rolls[i] = 1
		}
		for i > 0 && rolls[i] < rolls[i-1] {
			rolls[i], rolls[i-1] = rolls[i-1], rolls[i]
			i--
		}
	}
}

func sim(numAttackers, numDefenders, attackerMod, defenderMod int) (attackerWins bool, attackersKilled, defendersKilled int) {
	var (
		attackerRolls [3]int
		defenderRolls [2]int
	)
	if numAttackers < 1 || numAttackers > 3 || numDefenders < 1 || numDefenders > 2 || numAttackers < numDefenders {
		panic(fmt.Sprintf("illegal arguments: numAttackers=%d numDefenders=%d", numAttackers, numDefenders))
	}
	for i := 0; i < numAttackers; i++ {
		attackerRolls[i] = d6()
	}
	for i := 0; i < numDefenders; i++ {
		defenderRolls[i] = d6()
	}
	applyMod(attackerRolls[:numAttackers], attackerMod)
	applyMod(defenderRolls[:numDefenders], defenderMod)
	for i, j := numAttackers-1, numDefenders-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if attackerRolls[i] > defenderRolls[j] {
			defendersKilled++
			if defendersKilled == numDefenders {
				attackerWins = true
				break
			}
		} else {
			attackersKilled++
			if attackersKilled == numAttackers {
				break
			}
		}
	}
	return
}

func main() {
	var (
		seed         = flag.Int64("seed", 0, "PRNG seed")
		numTrials    = flag.Int("trials", 1000, "number of simulation trials")
		numAttackers = flag.Int("atk", 1, "number of attackers")
		numDefenders = flag.Int("def", 1, "number of defenders")
		attackerMod  = flag.Int("atkMod", 0, "attacker die modifier")
		defenderMod  = flag.Int("defMod", 0, "defender die modifier")
	)
	flag.Parse()

	if *seed != 0 {
		rand.Seed(*seed)
	}

	var attackerWins, attackersLeft, defendersLeft, roundsWon int
	for n := 0; n < *numTrials; n++ {
		// fmt.Printf("=== TRIAL %d ===\n", n)
		numAttackers := *numAttackers
		numDefenders := *numDefenders
		for numDefenders > 0 && numAttackers > 0 {
			numAtk := intMin(numAttackers, 3)
			numDef := intMin(numDefenders, intMin(numAttackers, 2))
			won, ak, dk := sim(numAtk, numDef, *attackerMod, *defenderMod)
			// fmt.Printf("attack: won?=%v ak=%d dk=%d\n", won, ak, dk)
			numAttackers -= ak
			numDefenders -= dk
			if won {
				roundsWon++
			}
		}
		if numDefenders == 0 {
			attackerWins++
		}
		attackersLeft += numAttackers
		defendersLeft += numDefenders
	}

	n := float64(*numTrials)
	fmt.Printf("pAttackerWins = %.2f%%\n", 100*float64(attackerWins)/n)
	fmt.Printf("avg. successful attacks per war: %.1f\n", float64(roundsWon)/n)
	fmt.Printf("avg. attackers left = %.1f\n", float64(attackersLeft)/n)
	fmt.Printf("avg. defenders left = %.1f\n", float64(defendersLeft)/n)
}

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
