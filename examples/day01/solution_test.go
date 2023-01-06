package main

import (
	"bytes"
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
		expected []byte
	}{
		{
			name:     "example",
			in:       strings.NewReader(exampleInput),
			expected: []byte("24000"),
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
			expected: []byte("45000"),
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
