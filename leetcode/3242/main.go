package main

type NeighborSum struct {
	grid [][]int
	idx  []int8
}

func Constructor(grid [][]int) NeighborSum {
	ns := NeighborSum{
		grid: grid,
		idx:  make([]int8, len(grid)*len(grid[0])),
	}

	for i := range grid {
		for j := range grid[i] {
			ns.idx[grid[i][j]] = int8(i*len(grid[0]) + j)
		}
	}
	return ns
}

func (ns NeighborSum) AdjacentSum(value int) int {
	n := int8(len(ns.grid))
	idx := ns.idx[value]
	i, j := idx/n, idx%n

	t := 0
	if j := j - 1; 0 <= j {
		t += ns.grid[i][j]
	}
	if j := j + 1; j < n {
		t += ns.grid[i][j]
	}
	if i := i + 1; i < n {
		t += ns.grid[i][j]
	}
	if i := i - 1; 0 <= i {
		t += ns.grid[i][j]
	}
	return t
}

func (ns NeighborSum) DiagonalSum(value int) int {
	n := int8(len(ns.grid))
	idx := ns.idx[value]
	i, j := idx/n, idx%n

	t := 0
	if i := i - 1; 0 <= i {
		if j := j - 1; 0 <= j {
			t += ns.grid[i][j]
		}
		if j := j + 1; j < n {
			t += ns.grid[i][j]
		}
	}
	if i := i + 1; i < n {
		if j := j - 1; 0 <= j {
			t += ns.grid[i][j]
		}
		if j := j + 1; j < n {
			t += ns.grid[i][j]
		}
	}
	return t
}
