package _001__Two_Sum

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := [][]int{
		[]int{1, 2, 3},
		[]int{11, 12, 13},
		[]int{101, 102, 103, 999},
		[]int{22, 2, 1, 33, 25, 99, 66, 18},
	}

	targets := []int{
		5,
		23,
		1100,
		100,
	}

	results := [][]int{
		[]int{1, 2},
		[]int{0, 1},
		[]int{0, 3},
		[]int{2, 5},
	}

	fmt.Printf("------------------------Leetcode Problem 1. Two Sum Testing------------------------\n")
	for i := 0; i < len(targets); i++ {
		fmt.Printf("nums = %v target = %v result = %v\n", tests[i], targets[i], twoSum(tests[i], targets[i]))
		if result := twoSum(tests[i], targets[i]); result[0] != results[i][0] && result[1] != results[i][1] {
			t.Fatalf("case %d fails: %v\n", i, result)
		}
	}
}
