package main

import (
	"bytes"
	_ "embed"
	"reflect"
	"strings"
	"testing"
)

const exampleInput0 = `mjqjpqmgbljsphdztnvjfqwrcgsmlb`

func Test_SolveFirstPart(t *testing.T) {
	ts := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "example1",
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected: "7",
		},
		{
			name:     "example2",
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: "5",
		},
		{
			name:     "example3",
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: "6",
		},
		{
			name:     "example4",
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: "10",
		},
		{
			name:     "example5",
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: "11",
		},
	}

	for _, tt := range ts {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			in := strings.NewReader(tt.input)
			err := SolveFirstPart(in, &buf)
			if err != nil {
				t.Errorf("got unexpected error: %v", err)
			}

			out := buf.String()

			if !reflect.DeepEqual(out, tt.expected) {
				t.Errorf("out different than expected: out %v, expected %v", out, tt.expected)
			}
		})
	}
}

func Test_SolveSecondPart(t *testing.T) {
	ts := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "example1",
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected: "19",
		},
		{
			name:     "example2",
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: "23",
		},
		{
			name:     "example3",
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: "23",
		},
		{
			name:     "example4",
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: "29",
		},
		{
			name:     "example5",
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: "26",
		},
	}

	for _, tt := range ts {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			in := strings.NewReader(tt.input)
			err := SolveSecondPart(in, &buf)
			if err != nil {
				t.Errorf("got unexpected error: %v", err)
			}

			out := buf.String()

			if !reflect.DeepEqual(out, tt.expected) {
				t.Errorf("out different than expected: out %v, expected %v", out, tt.expected)
			}
		})
	}
}
