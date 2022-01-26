package main

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestWebsiteStitcher(t *testing.T) {
	/*
	 Why does testing against real URLs suck?
	1.Slow
	2.Unreliable, site could be down, the site's content could change
	3. It's hard to write the test. We have to test string contains, which is naff
	 */
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

	t.Run("given some more urls, when we stitch them together, I expect to see the response bodies of all urls", func(t *testing.T) {
		url1 := "https://motherfuckingwebsite.com"
		url2 := "https://quii.dev"
		url3 := "https://www.riyadattani.com"

		got := WebsiteStitcher(url1, url2, url3)

		if !strings.Contains(got, "Look at this shit.") {
			t.Error("it didnt have the first websites body")
		}

		if !strings.Contains(got, "Chris") {
			t.Error("it doesnt have the cj stuff")
		}

		if !strings.Contains(got, "pair programming") {
			t.Error("it doesnt have the riya stuff")
		}
	})
}

func WebsiteStitcher(urls ...string) string {
	var resps []io.Reader

	for _, url := range urls {
		resp, _ := http.Get(url)
		defer resp.Body.Close()

		resps = append(resps, resp.Body)
	}

	stitched := Stitcher(resps...)

	return stitched
}
