package main

import "testing"

func Test_NeighborSum(t *testing.T) {
	type operation struct {
		value    int
		adjacent bool
		exp      int
	}

	testcases := []struct {
		grid [][]int
		sums []operation
	}{{
		grid: [][]int{
			{0, 1, 2},
			{3, 4, 5},
			{6, 7, 8},
		},
		sums: []operation{
			{1, true, 6},
			{4, true, 16},
			{4, false, 16},
			{8, false, 4},
		},
	}, {
		grid: [][]int{
			{1, 2, 0, 3},
			{4, 7, 15, 6},
			{8, 9, 10, 11},
			{12, 13, 14, 5},
		},
		sums: []operation{
			{15, true, 23},
			{9, false, 45},
		},
	}}

	for _, tc := range testcases {
		ns := Constructor(tc.grid)

		for _, op := range tc.sums {
			if op.adjacent {
				got := ns.AdjacentSum(op.value)

				if got != op.exp {
					t.Fatalf("expected %d got %d", op.exp, got)
				}
			}
		}
	}
}
