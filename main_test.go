package main

import "testing"

func TestTwoSumSorted(t *testing.T) {
	cases := []struct {
		in     []int
		target int
		i, j   int
		ok     bool
	}{
		{[]int{2, 7, 11, 15}, 9, 1, 2, true},
		{[]int{2, 3, 4}, 6, 1, 3, true},
		{[]int{-1, 0}, -1, 1, 2, true},
		{[]int{1, 2, 3}, 100, 0, 0, false},
		{[]int{}, 5, 0, 0, false},
		{[]int{3, 3}, 6, 1, 2, true},
		{[]int{1, 2, 3, 4, 4}, 8, 4, 5, true},
	}
	for _, c := range cases {
		i, j, ok := TwoSumSorted(c.in, c.target)
		if i != c.i || j != c.j || ok != c.ok {
			t.Errorf("TwoSumSorted(%v,%d) = (%d,%d,%v), want (%d,%d,%v)",
				c.in, c.target, i, j, ok, c.i, c.j, c.ok)
		}
	}
}
