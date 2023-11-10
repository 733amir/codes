package main

import "fmt"

func main() {
    tcs := [][]int{
        {3,1,2,4},
    }

    for _, tc := range tcs {
        fmt.Printf("%v => ", tc)
        fmt.Printf("%v\n", sortArrayByParity(tc))
    }
}

func sortArrayByParity(nums []int) []int {
	i, j := 0, len(nums)-1
	for {
		for ; i < len(nums) && (nums[i]&1) == 0; i++ {
		}
		for ; 0 <= j && (nums[j]&1) == 1; j-- {
		}

        if i >= j {
            return nums
        }

        nums[i], nums[j] = nums[j], nums[i]
	}
}
