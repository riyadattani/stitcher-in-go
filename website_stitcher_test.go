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
			fmt.Fprint(w, expected1)
		}))
		defer server1.Close()

		server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, expected2)
		}))
		defer server2.Close()

		server3 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, expected3)
		}))
		defer server3.Close()

		got, _ := WebsiteStitcher(server1.URL, server2.URL, server3.URL)
		want := expected1 + expected2 + expected3

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("it returns an error, if one of the URLs is nonsense", func(t *testing.T) {
		server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Look at this shit.")
		}))
		defer server1.Close()

		_, err := WebsiteStitcher(server1.URL, "lmao")

		if err == nil {
			t.Error("expected an error, but didnt get one")
		}
	})

	t.Run("it returns an error, if one of the URLs doesnt return a response", func(t *testing.T) {
		server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Look at this shit.")
		}))
		defer server1.Close()

		failingServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			panic("wtf")
		}))
		defer failingServer.Close()

		_, err := WebsiteStitcher(server1.URL, failingServer.URL)

		if err == nil {
			t.Error("expected an error, but didnt get one")
		}
	})

}

func WebsiteStitcher(urls ...string) (string, error) {
	var resps []io.Reader

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		resps = append(resps, resp.Body)
	}

	return Stitcher(resps...), nil
}
