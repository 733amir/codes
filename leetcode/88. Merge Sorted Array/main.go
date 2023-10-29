package main

import "fmt"

func main() {
	tcs := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}{
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
		},
		{
			nums1: []int{3, 3, 4, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
		},
	}

	for _, tc := range tcs {
		fmt.Printf("%+v\n", tc)
		merge(tc.nums1, tc.m, tc.nums2, tc.n)
		fmt.Printf("%+v\n\n", tc)
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
    n, m = n - 1, m - 1

	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[n + m + 1] = nums1[m]
			m--
		} else {
			nums1[n + m + 1] = nums2[n]
			n--
		}
	}

	for n >= 0 {
		nums1[n] = nums2[n]
        n--
	}
}
