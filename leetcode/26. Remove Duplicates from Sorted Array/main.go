package main

import "fmt"

func main() {

	tcs := []struct{
        n []int
        r int
    }{
        {n: []int{0}, r: 1},
        {n: []int{0, 0}, r: 1},
        {n: []int{0, 1}, r: 2},
        {n: []int{0, 0, 0}, r: 1},
        {n: []int{0, 0, 1}, r: 2},
        {n: []int{0, 1, 1}, r: 2},
        {n: []int{0, 1, 2}, r: 3},
        {n: []int{0, 0, 0, 0}, r: 1},
        {n: []int{0, 0, 0, 1}, r: 2},
        {n: []int{0, 0, 1, 1}, r: 2},
        {n: []int{0, 0, 1, 2}, r: 3},
        {n: []int{0, 1, 1, 1}, r: 2},
        {n: []int{0, 1, 1, 2}, r: 3},
        {n: []int{0, 1, 2, 2}, r: 3},
        {n: []int{0, 1, 2, 3}, r: 4},
        {n: []int{0, 1, 2, 2, 2, 3, 4}, r: 5},
        {n: []int{0, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 5, 5, 5}, r: 6},
        {n: []int{0, 1, 2, 3, 2, 2, 3, 4, 5, 5, 5}, r: 6},
	}

	for _, tc := range tcs {
        s := fmt.Sprintf("%+v", tc.n)
		k := removeElement(tc.n)
        if k != tc.r {
            fmt.Printf("%+v => %+v (%+v)\n\n", s, tc.r, k)
        }
	}
}

func removeElement(nums []int) int {
    // fmt.Printf("n: %v\n", nums)
	j := 0
	for i := 0; i < len(nums)-1; i++ {
        // fmt.Printf("i: %v, j: %v\n", i, j)
        if i != 0 && nums[i-1] >= nums[i] {
            i--
        }
		if nums[i] < nums[i+1] {
			continue
		}

		j = max(i+2, j)
		for ; j < len(nums); j++ {
            // fmt.Printf("\tj: %v\n", j)
			if nums[i] >= nums[j] {
				continue
			}

			nums[i+1] = nums[j]
            i++
			j++
			break
		}

		if j >= len(nums) {
			return i + 1
		}
	}

	return len(nums)
}
