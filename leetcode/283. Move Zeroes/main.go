package main

import "fmt"

func main() {
	tcs := [][]int{
		{0, 1, 0, 3, 12},
	}

    for _, tc := range tcs {
        fmt.Printf("%v => ", tc)
        moveZeroes(tc)
        fmt.Printf("%v\n", tc)
    }
}

func moveZeroes(nums []int) {
	j := 0
	for i := range nums {
        if nums[i] != 0 {
            continue
        }

		j = max(j, i+1)

		for ; j < len(nums) && nums[j] == 0; j++ {
		}

		if j >= len(nums) {
			for j := i + 1; j < len(nums); j++ {
				nums[j] = 0
			}
			return
		}

		nums[i] = nums[j]
        nums[j] = 0
	}
	return
}
