package main

import (
	"bytes"
	_ "embed"
	"io"
	"reflect"
	"strings"
	"testing"
)

const exampleInput = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

func Test_SolveFirstPart(t *testing.T) {
	ts := []struct {
		name     string
		in       io.Reader
		expected []byte
	}{
		{
			name:     "example",
			in:       strings.NewReader(exampleInput),
			expected: []byte("2"),
		},
	}

	for _, tt := range ts {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			err := SolveFirstPart(tt.in, &buf)
			if err != nil {
				t.Errorf("got unexpected error: %v", err)
			}

			out := buf.Bytes()

			if !reflect.DeepEqual(out, tt.expected) {
				t.Errorf("out different than expected: out %v, expected %v", out, tt.expected)
			}
		})
	}
}

func Test_SolveSecondPart(t *testing.T) {
	ts := []struct {
		name     string
		in       io.Reader
		expected []byte
	}{
		{
			name:     "example",
			in:       strings.NewReader(exampleInput),
			expected: []byte("4"),
		},
	}

	for _, tt := range ts {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			err := SolveSecondPart(tt.in, &buf)
			if err != nil {
				t.Errorf("got unexpected error: %v", err)
			}

			out := buf.Bytes()

			if !reflect.DeepEqual(out, tt.expected) {
				t.Errorf("out different than expected: out %v, expected %v", out, tt.expected)
			}
		})
	}
}
