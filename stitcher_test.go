package main

import (
	"io"
	"strings"
	"testing"
)

func TestStitcher(t *testing.T) {
	tests := []struct{
		description string
		inputs []io.Reader
		expected string
	}{
		{
			description: "stitching multiple readers",
			inputs:   []io.Reader{strings.NewReader("Salt"), strings.NewReader("Pay")},
			expected: "SaltPay",
		},
		{
			description: "stitching one reader",
			inputs:   []io.Reader{strings.NewReader("Salt")},
			expected: "Salt",
		},
	}
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := Stitcher(tc.inputs...)
			want := tc.expected
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}

