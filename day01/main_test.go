package main

import (
	_ "embed"
	"reflect"
	"testing"
)

//go:embed example_input.txt
var exampleInput string

var summedExampleInput []int

func must(f func() (interface{}, error)) interface{} {
	res, err := f()
	if err != nil {
		panic(err)
	}
	return res
}

func init() {
	summedExampleInput = sumGroups(must(func() (interface{}, error) {
		return parseInput(exampleInput)
	}).([][]int))
}

func Test_solveFirstPart(t *testing.T) {
	ts := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "example",
			input:    summedExampleInput,
			expected: 24000,
		},
	}

	for _, tt := range ts {
		t.Run(tt.name, func(t *testing.T) {
			got := solveFirstPart(summedExampleInput)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got different than expected: got %v, expected %v", got, tt.expected)
			}
		})
	}
}

func Test_solveSecondPart(t *testing.T) {
	ts := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "example",
			input:    summedExampleInput,
			expected: 45000,
		},
	}

	for _, tt := range ts {
		t.Run(tt.name, func(t *testing.T) {
			got := solveSecondPart(summedExampleInput)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got different than expected: got %v, expected %v", got, tt.expected)
			}
		})
	}
}
