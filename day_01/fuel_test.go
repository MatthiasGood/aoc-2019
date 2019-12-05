package main

import (
	"testing"
)

type test struct {
	input int
	want  int
}

func TestFuelNeeded(t *testing.T) {

	tests := []test {
		{input: 12, want: 2},
		{input: 14, want: 2},
		{input: 1969, want: 654},
		{input: 100756, want: 33583},
	}

	for _, tc := range tests {
		got := findFuelForModule(tc.input)
		if got != tc.want {
			t.Fatalf("expected: %d, got: %d", tc.want, got)
		}
	}
}

func TestFuelNeededWithFuelMass(t *testing.T) {

	tests := []test {
		{input: 12, want: 2},
		{input: 14, want: 2},
		{input: 1969, want: 966},
		{input: 100756, want: 50346},
	}

	for _, tc := range tests {
		got := findFuelForModuleWithFuelMass(tc.input)
		if got != tc.want {
			t.Fatalf("expected: %d, got: %d", tc.want, got)
		}
	}
}