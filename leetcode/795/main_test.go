package main

import "testing"

func Test_subarrayCount(t *testing.T) {
	testcases := []struct {
		nums        []int
		left, right int
		count       int
	}{
		{nums: []int{2, 1, 4, 3}, left: 2, right: 3, count: 3},
		{nums: []int{2, 9, 2, 5, 6}, left: 2, right: 8, count: 7},
	}

	for _, tc := range testcases {
		got := subarrayCount(tc.nums, tc.left, tc.right)

		if got != tc.count {
			t.Errorf("For %v expected %d but got %d", tc.nums, tc.count, got)
		}
	}
}
