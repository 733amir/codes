package main

import "fmt"

func main() {

	tcs := []struct {
		nums []int
		val  int
	}{
		{
			nums: []int{3, 2, 2, 3},
			val:  3,
		},
		{
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
		},
		{
			nums: []int{2},
			val:  3,
		},
	}

	for _, tc := range tcs {
		fmt.Printf("%+v => ", tc)
		k := removeElement(tc.nums, tc.val)
		fmt.Printf("%+v %v\n", tc.nums, k)
	}
}

func removeElement(nums []int, val int) int {
	j := len(nums) - 1

	for i := range nums {
		if nums[i] == val {
			for i < j {
				if nums[j] != val {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
				j--
			}
		}

        if nums[i] == val {
            return i
        }

		if i == len(nums) - 1 {
			return i + 1
		}
	}

	return len(nums)
}
