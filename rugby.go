package main

import (
	"fmt"
)

// possible rugby scores according to current intl. rules
var scores []int = []int{3, 5, 7}

// map of scores up to now
var scoreOps []int = []int{0, -1, -1, 1, -1}

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

func main() {
	calculateMinimumScoringOps(150)
	for i := 0; i < len(scoreOps); i++ {
		fmt.Println(i, scoreOps[i])
	}

}
