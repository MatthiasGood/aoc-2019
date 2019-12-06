package main

import "testing"

type test struct {
	input []int
	want  []int
}

func TestIntcode(t *testing.T) {

	tests := []test{
		{input: []int{1, 0, 0, 0, 99}, want: []int{2, 0, 0, 0, 99}},
		{input: []int{2, 3, 0, 3, 99}, want: []int{2, 3, 0, 6, 99}},
		{input: []int{2, 4, 4, 5, 99, 0}, want: []int{2, 4, 4, 5, 99, 9801}},
		{input: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, want: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{input: []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, want: []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}},
	}

	for _, tc := range tests {
		got := intcode(tc.input)
		if !equalSlices(got, tc.want) {
			t.Fatalf("expected: %d, got: %d", tc.want, got)
		}
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
