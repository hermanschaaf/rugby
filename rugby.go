package main

import (
	"fmt"
)

// possible rugby scores according to current intl. rules
var scores []int = []int{3, 5, 7}

// map of scores up to now
var scoreOps []int = []int{0, -1, -1, 1, -1}

var scoreCombinations map[int][][3]int = map[int][][3]int{
	0: [][3]int{[3]int{0, 0, 0}},
	1: [][3]int{},
	2: [][3]int{},
	3: [][3]int{[3]int{1, 0, 0}},
	4: [][3]int{},
}

// calculateMinimumScoringOps uses dynamic programming to
// build up an array of minimum scoring opportunities needed
// for specific scores, and stores it in the global variable
// scoreOps.
func calculateMinimumScoringOps(max int) {
	// loop through all scores up to the maximum score we want
	// to calculate
	for i := len(scoreOps); i <= max; i++ {
		// set minValue to maximum 32 bit integer
		minValue := 1<<31 - 1
		// go through all the scoring opportunity totals
		for _, s := range scores {
			// if the opportunity is smaller than the total
			// we are considering, see if one plus the previously
			// calculated minimum is smaller than what we already
			// have.
			if i-s >= 0 && scoreOps[i-s] >= 0 && scoreOps[i-s]+1 < minValue {
				minValue = scoreOps[i-s] + 1
			}
		}
		scoreOps = append(scoreOps, minValue)
	}
}

// MinimumScoringOps calculates the minimum number of scoring
// opportunities needed for a team to reach a given total score.
// For example, given an input of score = 8, this function will
// return 2, as this can only be achieved by scoring a try and
// a penalty/drop kick, for a total of two scoring opportunities.
func MinimumScoringOps(score int) (minOps int) {
	if len(scoreOps) > score {
		return scoreOps[score]
	} else {
		calculateMinimumScoringOps(score)
		return scoreOps[score]
	}
}

// calculateScoringCombinations uses dynamic programming to
// build up an array of scoring opportunity combinations
// for specific scores, and stores it in the global variable
// scoreCombinations.
func calculateScoringCombinations(max int) {
	// loop through all scores up to the maximum score we want
	// to calculate
	for i := len(scoreCombinations); i <= max; i++ {
		// set combs to empty list of combinations
		combs := [][3]int{}
		// we use a map of found combinations to do O(1) uniqueness lookups
		combMap := make(map[string]bool)
		// go through all the scoring opportunities
		for cnt, s := range scores {
			if i-s >= 0 && len(scoreCombinations[i-s]) > 0 {
				for _, cmb := range scoreCombinations[i-s] {
					a := [3]int{}
					for u, cmbEl := range cmb {
						a[u] = cmbEl
						if u == cnt {
							a[u] += 1
						}
					}
					_, has := combMap[fmt.Sprintf("%d %d %d", a[0], a[1], a[2])]
					if !has {
						combs = append(combs, a)
						combMap[fmt.Sprintf("%d %d %d", a[0], a[1], a[2])] = true
					}
				}
			}
		}
		scoreCombinations[i] = combs
	}
}

// ScoringCombinations calculates the possible scoring
// opportunities combinations for a team to reach a given total score.
// For example, given an input of score = 10, this function returns
// a slice of length 2. (try+try, try+conversion+penalty)
func ScoringCombinations(score int) (combinations [][3]int) {
	if len(scoreCombinations) > score {
		return scoreCombinations[score]
	} else {
		calculateScoringCombinations(score)
		return scoreCombinations[score]
	}
}

func main() {
	// calculateMinimumScoringOps(150)
	// for i := 0; i < len(scoreOps); i++ {
	// 	fmt.Println(i, scoreOps[i])
	// }

	calculateScoringCombinations(50)
	for i := 0; i < len(scoreCombinations); i++ {
		fmt.Println(i, scoreCombinations[i])
	}
}
