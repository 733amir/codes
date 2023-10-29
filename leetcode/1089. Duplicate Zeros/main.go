package main

import "fmt"

func main() {
	tcs := [][]int{
		{1, 0, 2, 3, 0, 4, 5, 0},
		{1, 5, 2, 0, 6, 8, 0, 6, 0},
		{1, 0},
	}

	for _, tc := range tcs {
		fmt.Printf("%+v => ", tc)
		duplicateZeros(tc)
		fmt.Printf("%+v\n", tc)
	}
}

func duplicateZeros(arr []int) {
	i, j := duplicateFrom(arr)

	for i >= 0 && j >= 0 {
		if arr[i] == 0 {
			if j == len(arr) {
				arr[j-1] = 0
			} else {
				arr[j-1], arr[j] = 0, 0
			}
            j -= 1
		} else {
            arr[j] = arr[i]
        }
        i--
        j--
	}
}

func duplicateFrom(arr []int) (int, int) {
	j := -1

	for i := range arr {
		if arr[i] == 0 {
			j += 2
		} else {
			j += 1
		}

		if j >= len(arr)-1 {
			return i, j
		}
	}

	return len(arr) - 1, len(arr) - 1
}
