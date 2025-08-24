package main

import (
	"reflect"
	"testing"
)

func Test_getSkyline(t *testing.T) {
	testcases := []struct {
		buildings []Building
		skyline   []Skyline
	}{{
		buildings: []Building{
			{start: 2, end: 9, height: 10},
			{start: 3, end: 7, height: 15},
			{start: 5, end: 12, height: 12},
			{start: 15, end: 20, height: 10},
			{start: 19, end: 24, height: 8},
		},
		skyline: []Skyline{
			{start: 2, height: 10},
			{start: 3, height: 15},
			{start: 7, height: 12},
			{start: 12, height: 0},
			{start: 15, height: 10},
			{start: 20, height: 8},
			{start: 24, height: 0},
		},
	}, {
		// Input: buildings = [[0,2,3],[2,5,3]]
		// Output: [[0,3],[5,0]]
		buildings: []Building{
			{start: 0, end: 2, height: 3},
			{start: 2, end: 5, height: 3},
		},
		skyline: []Skyline{
			{start: 0, height: 3},
			{start: 5, height: 0},
		},
	}}

	for _, tc := range testcases {
		got := getSkyline(tc.buildings)

		if !reflect.DeepEqual(got, tc.skyline) {
			t.Errorf("For %v expected %v but got %v", tc.buildings, tc.skyline, got)
		}
	}
}

func Test_merger(t *testing.T) {
	testcases := []struct {
		a    Building
		b    Building
		expB []Building
		expC bool
	}{{
		a: Building{start: 1, end: 2, height: 3},
		b: Building{start: 3, end: 4, height: 5},
		expB: []Building{
			{start: 1, end: 2, height: 3},
			{start: 3, end: 4, height: 5},
		},
		expC: true,
	}, {
		a: Building{start: 1, end: 3, height: 3},
		b: Building{start: 3, end: 4, height: 5},
		expB: []Building{
			{start: 1, end: 3, height: 3},
			{start: 3, end: 4, height: 5},
		},
		expC: true,
	}, {
		a: Building{start: 1, end: 3, height: 3},
		b: Building{start: 2, end: 4, height: 5},
		expB: []Building{
			{start: 1, end: 2, height: 3},
			{start: 2, end: 3, height: 5},
			{start: 3, end: 4, height: 5},
		},
		expC: true,
	}, {
		a: Building{start: 1, end: 3, height: 5},
		b: Building{start: 2, end: 4, height: 3},
		expB: []Building{
			{start: 1, end: 2, height: 5},
			{start: 2, end: 3, height: 5},
			{start: 3, end: 4, height: 3},
		},
		expC: true,
	}, {
		a: Building{start: 1, end: 3, height: 3},
		b: Building{start: 2, end: 3, height: 5},
		expB: []Building{
			{start: 1, end: 2, height: 3},
			{start: 2, end: 3, height: 5},
		},
		expC: false,
	}, {
		a: Building{start: 1, end: 3, height: 5},
		b: Building{start: 2, end: 3, height: 3},
		expB: []Building{
			{start: 1, end: 2, height: 5},
			{start: 2, end: 3, height: 5},
		},
		expC: false,
	}, {
		a: Building{start: 1, end: 4, height: 3},
		b: Building{start: 2, end: 3, height: 5},
		expB: []Building{
			{start: 1, end: 2, height: 3},
			{start: 2, end: 3, height: 5},
			{start: 3, end: 4, height: 3},
		},
		expC: false,
	}, {
		a: Building{start: 1, end: 4, height: 5},
		b: Building{start: 2, end: 3, height: 3},
		expB: []Building{
			{start: 1, end: 2, height: 5},
			{start: 2, end: 3, height: 5},
			{start: 3, end: 4, height: 5},
		},
		expC: false,
	}, {
		a: Building{start: 1, end: 2, height: 3},
		b: Building{start: 1, end: 3, height: 5},
		expB: []Building{
			{start: 1, end: 2, height: 5},
			{start: 2, end: 3, height: 5},
		},
		expC: true,
	}, {
		a: Building{start: 1, end: 2, height: 5},
		b: Building{start: 1, end: 3, height: 3},
		expB: []Building{
			{start: 1, end: 2, height: 5},
			{start: 2, end: 3, height: 3},
		},
		expC: true,
	}, {
		a: Building{start: 1, end: 3, height: 3},
		b: Building{start: 1, end: 3, height: 5},
		expB: []Building{
			{start: 1, end: 3, height: 5},
		},
		expC: false,
	}, {
		a: Building{start: 1, end: 3, height: 5},
		b: Building{start: 1, end: 3, height: 3},
		expB: []Building{
			{start: 1, end: 3, height: 5},
		},
		expC: false,
	}, {
		a: Building{start: 1, end: 3, height: 3},
		b: Building{start: 1, end: 2, height: 5},
		expB: []Building{
			{start: 1, end: 2, height: 5},
			{start: 2, end: 3, height: 3},
		},
		expC: false,
	}, {
		a: Building{start: 1, end: 3, height: 5},
		b: Building{start: 1, end: 2, height: 3},
		expB: []Building{
			{start: 1, end: 2, height: 5},
			{start: 2, end: 3, height: 5},
		},
		expC: false,
	}}

	for _, tc := range testcases {
		gotB, gotC := merger(tc.a, tc.b)

		if gotC != tc.expC {
			t.Errorf("building a: %+v, b: %+v merger should be %v but got %v", tc.a, tc.b, tc.expC, gotC)
			continue
		}

		if len(gotB) != len(tc.expB) {
			t.Errorf("building a: %+v, b: %+v merger should be %v but got %v", tc.a, tc.b, tc.expB, gotB)
			continue
		}

		for i := range gotB {
			if gotB[i].start != tc.expB[i].start ||
				gotB[i].end != tc.expB[i].end ||
				gotB[i].height != tc.expB[i].height {
				t.Errorf("building a: %+v, b: %+v merger should be %v but got %v", tc.a, tc.b, tc.expB, gotB)
			}
		}
	}
}

func Test_collision(t *testing.T) {
	testcases := []struct {
		a   Building
		b   Building
		exp bool
	}{{
		a:   Building{start: 1, end: 2},
		b:   Building{start: 3, end: 4},
		exp: false,
	}, {
		a:   Building{start: 1, end: 3},
		b:   Building{start: 3, end: 4},
		exp: false,
	}, {
		a:   Building{start: 1, end: 3},
		b:   Building{start: 2, end: 4},
		exp: true,
	}, {
		a:   Building{start: 1, end: 3},
		b:   Building{start: 2, end: 3},
		exp: true,
	}, {
		a:   Building{start: 1, end: 4},
		b:   Building{start: 2, end: 3},
		exp: true,
	}, {
		a:   Building{start: 1, end: 2},
		b:   Building{start: 1, end: 3},
		exp: true,
	}, {
		a:   Building{start: 1, end: 2},
		b:   Building{start: 1, end: 2},
		exp: true,
	}}

	for _, tc := range testcases {
		got := collision(tc.a, tc.b)

		if got != tc.exp {
			t.Errorf("building a: %+v, b: %+v collision should be %v but got %v", tc.a, tc.b, tc.exp, got)
		}

		// Changing the order of building should not have any effect on the result.

		got = collision(tc.b, tc.a)

		if got != tc.exp {
			t.Errorf("building b: %+v, a: %+v collision should be %v but got %v", tc.b, tc.a, tc.exp, got)
		}
	}
}
