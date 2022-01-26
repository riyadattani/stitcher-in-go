package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
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
		expected1 := "Look at this shit."
		expected2 := "Chris"
		expected3 := "pair programming"

		server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, expected1)
		}))
		defer server1.Close()

		server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, expected2)
		}))
		defer server2.Close()

		server3 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, expected3)
		}))
		defer server3.Close()

		got := WebsiteStitcher(server1.URL, server2.URL, server3.URL)
		want := expected1 + expected2 + expected3

		if got != want {
			t.Errorf("got %q, want %q", got, want)
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

	return Stitcher(resps...)
}
