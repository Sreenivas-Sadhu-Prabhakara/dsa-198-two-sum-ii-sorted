package main

import "fmt"

// TwoSumSorted returns 1-based indices (i < j) of the two entries of a
// SORTED array summing to target, or (0, 0, false). Sortedness replaces the
// hash map of plain Two Sum with converging pointers and O(1) space.
func TwoSumSorted(numbers []int, target int) (int, int, bool) {
	l, r := 0, len(numbers)-1
	for l < r {
		switch s := numbers[l] + numbers[r]; {
		case s == target:
			return l + 1, r + 1, true
		case s < target:
			l++
		default:
			r--
		}
	}
	return 0, 0, false
}

func main() {
	fmt.Println(TwoSumSorted([]int{2, 7, 11, 15}, 9))
}
