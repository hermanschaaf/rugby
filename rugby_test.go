package main

import "testing"

func TestScoringCombinations(t *testing.T) {
	combs := map[int][][3]int{
		0:  [][3]int{[3]int{0, 0, 0}},
		1:  [][3]int{},
		3:  [][3]int{[3]int{1, 0, 0}},
		5:  [][3]int{[3]int{0, 1, 0}},
		7:  [][3]int{[3]int{0, 0, 1}},
		8:  [][3]int{[3]int{1, 1, 0}},
		9:  [][3]int{[3]int{3, 0, 0}},
		10: [][3]int{[3]int{0, 2, 0}, [3]int{1, 0, 1}},
	}
	for k, expected := range combs {
		ans := ScoringCombinations(k)
		if len(expected) != len(ans) {
			t.Errorf("Expected ScoringCombinations(%v)=%v, but got %v", k, expected, ans)
		}
	}
}
