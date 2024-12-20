package main

type Sub struct {
	count int
	subs  []Sub
}

func subarrayCount(nums []int, left, right int) int {
	count := 0
	belows := []int{0}
	sum, sub_sum := 0, 0
	for _, n := range nums {
		if n > right {
			sub_sum = 0
			for _, b := range belows {
				sub_sum += (b + 1) * b / 2
			}
			sum += (count+1)*count/2 - sub_sum

			count = 0
			belows = belows[:1]
			belows[0] = 0
		} else if n < left {
			belows[len(belows)-1] += 1
			count += 1
		} else {
			if belows[len(belows)-1] != 0 {
				belows = append(belows, 0)
			}
			count += 1
		}
	}

	sub_sum = 0
	for _, b := range belows {
		sub_sum += (b + 1) * b / 2
	}
	return sum + (count+1)*count/2 - sub_sum
}

