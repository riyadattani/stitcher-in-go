package main

import (
	"io"
	"strings"
	"testing"
)

func TestStitcher(t *testing.T) {
	t.Run("given two io.Readers, when we stitch it up, we expect a joint string", func(t *testing.T) {
		a := strings.NewReader("Salt")
		b := strings.NewReader("Pay")

		got := Stitcher(a, b)
		want := "SaltPay"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("given two other io.Readers, when we stitch it up, we expect a joint string", func(t *testing.T) {
		a := strings.NewReader("Salt")
		b := strings.NewReader("Lame")

		got := Stitcher(a, b)
		want := "SaltLame"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func Stitcher(a, b io.Reader) string {
	aData, _ := io.ReadAll(a)
	bData, _ := io.ReadAll(b)

	return string(aData) + string(bData)
}
