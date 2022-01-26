package main

import (
	"net/http"
	"strings"
	"testing"
)

func TestWebsiteStitcher(t *testing.T) {
	t.Run("given some urls, when we stitch them together, I expect to see the response bodies of all urls", func(t *testing.T) {
		url1 := "https://motherfuckingwebsite.com"
		url2 := "https://quii.dev"

		got := WebsiteStitcher(url1, url2)

		if !strings.Contains(got, "Look at this shit.") {
			t.Error("it didnt have the first websites body")
		}

		if !strings.Contains(got, "Chris") {
			t.Error("it doesnt have the cj stuff")
		}
	})
}

func WebsiteStitcher(url1 string, url2 string) string {
	resp1, _ := http.Get(url1)
	defer resp1.Body.Close()

	resp2, _ := http.Get(url2)
	defer resp2.Body.Close()

	stitched := Stitcher(resp1.Body, resp2.Body)

	return stitched
}
