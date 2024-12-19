package main

import "fmt"

func numWays(s string) int {
	countOfOne := 0
	for _, c := range s {
		if c == '1' {
			countOfOne += 1
		}
	}

	if countOfOne%3 != 0 {
		return 0
	}

	if countOfOne == 0 {
		return int(((int64(len(s)) - 1) * (int64(len(s)) - 2) / 2) % 1_000_000_007)
	}

	first, second := -1, -1
	oneThird := countOfOne / 3
	countOfOne = 0
	for i := 0; i <= len(s); i++ {
		if s[i] == '1' {
			countOfOne += 1
		}

		if countOfOne == oneThird {
			first = consecutiveZeroCount(s[i+1:])
			i += first
		}

		if countOfOne == 2*oneThird {
			second = consecutiveZeroCount(s[i+1:])
			break
		}
	}

	return int((int64(first) + 1) * (int64(second) + 1) % 1_000_000_007)
}

func consecutiveZeroCount(s string) int {
	for i, c := range s {
		if c != '0' {
			return i
		}
	}
	return len(s)
}

func main() {
	//                             f  s
	fmt.Println(numWays("11001010011"))
}
