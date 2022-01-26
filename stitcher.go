package main

import (
	"bytes"
	"io"
)

func Stitcher(readers ...io.Reader) string {
	result := bytes.Buffer{}

	for _, reader := range readers {
		io.Copy(&result, reader)
	}
	return result.String()
}
