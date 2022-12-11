package main

import (
	_ "embed"
	"reflect"
	"testing"
)

//go:embed example_input.txt
var exampleInput string

func Test_solveFirstPart(t *testing.T) {
	ts := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "first part",
			input:    exampleInput,
			expected: 15,
		},
	}

	for _, tt := range ts {
		t.Run(tt.name, func(t *testing.T) {
			got := SolveFirstPart(tt.input)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got different than expected: got %v, expected %v", got, tt.expected)
			}
		})
	}
}

func Test_solveSecondPart(t *testing.T) {
	ts := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "example",
			input:    exampleInput,
			expected: 12,
		},
	}

	for _, tt := range ts {
		t.Run(tt.name, func(t *testing.T) {
			got := SolveSecondPart(tt.input)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got different than expected: got %v, expected %v", got, tt.expected)
			}
		})
	}
}
