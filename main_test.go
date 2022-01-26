package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestStitcher(t *testing.T) {
	var(
		a = "Salt"
		b = "Pepper"
		c = "Pay"
	)
	t.Run("given two io.Readers, when we stitch it up, we expect a joint string", func(t *testing.T) {
		got := Stitcher(strings.NewReader(a), strings.NewReader(c))
		want := "SaltPay"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})


	t.Run("given 3 io.Readers, when we stitch it up, we expect a joint string", func(t *testing.T) {
		got := Stitcher(strings.NewReader(a), strings.NewReader(b), strings.NewReader(c))
		want := "SaltPepperPay"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func Stitcher(readers ...io.Reader) string {
	result := bytes.Buffer{}

	for _, reader := range readers {
		data, _ := io.ReadAll(reader)
		result.Write(data)

	}
	return result.String()
}
