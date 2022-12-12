package main

import (
	_ "embed"
	"io"
	"reflect"
	"strings"
	"testing"
)

const exampleInput = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func Test_SolveFirstPart(t *testing.T) {
	ts := []struct {
		name     string
		in       io.Reader
		out      io.Writer
		expected int
	}{
		{
			name:     "example",
			in:       strings.NewReader(exampleInput),
			out:      io.Discard,
			expected: 24000,
		},
	}

	for _, tt := range ts {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolveFirstPart(tt.in, tt.out)
			if err != nil {
				t.Errorf("got unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got different than expected: got %v, expected %v", got, tt.expected)
			}
		})
	}
}

func Test_SolveSecondPart(t *testing.T) {
	ts := []struct {
		name     string
		in       io.Reader
		out      io.Writer
		expected int
	}{
		{
			name:     "example",
			in:       strings.NewReader(exampleInput),
			out:      io.Discard,
			expected: 45000,
		},
	}

	for _, tt := range ts {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolveSecondPart(tt.in, tt.out)
			if err != nil {
				t.Errorf("got unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got different than expected: got %v, expected %v", got, tt.expected)
			}
		})
	}
}
