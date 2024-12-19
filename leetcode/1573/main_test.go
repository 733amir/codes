package main

import "testing"

func Test_consecutiveZeroCount(t *testing.T) {
	testcases := []struct {
		count int
		num   string
	}{
		{count: 0, num: ""},
		{count: 0, num: "1"},
		{count: 0, num: "10"},
		{count: 0, num: "11"},
		{count: 1, num: "0"},
		{count: 1, num: "01"},
		{count: 1, num: "010"},
		{count: 1, num: "011"},
		{count: 2, num: "00"},
		{count: 2, num: "001"},
		{count: 2, num: "0010"},
	}

	for _, tc := range testcases {
		got := consecutiveZeroCount(tc.num)

		if got != tc.count {
			t.Errorf("for %s expected %d but got %d", tc.num, tc.count, got)
		}
	}
}

func Test_numWays(t *testing.T) {
	testcases := []struct {
		ways int
		num  string
	}{
		{ways: 0, num: "1001"},
		{ways: 1, num: "000"},
		{ways: 3, num: "0000"},
		{ways: 4, num: "10101"},
	}

	for _, tc := range testcases {
		got := numWays(tc.num)

		if got != tc.ways {
			t.Errorf("for %s expected %d but got %d", tc.num, tc.ways, got)
		}
	}
}
