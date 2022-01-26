package main

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"sync"
)

type WebsiteStitcherServer func(urls ...string) (string, error)

func (ws WebsiteStitcherServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urls := r.URL.Query()["url"]
	results, _ := ws(urls...)
	fmt.Fprint(w, results)
}

func WebsiteStitcher(urls ...string) (string, error) {
	type getResult struct {
		index int
		body  io.ReadCloser
		err   error
	}

	respChannel := make(chan getResult, len(urls))
	var wg sync.WaitGroup

	for i, url := range urls {
		wg.Add(1)
		go func(index int, u string) {
			defer wg.Done()
			resp, err := http.Get(u)
			if err != nil {
				respChannel <- getResult{
					index: index,
					body:  nil,
					err:   err,
				}
			} else {
				respChannel <- getResult{
					index: index,
					body:  resp.Body,
					err:   nil,
				}
			}
		}(i, url)
	}
	wg.Wait()

	// read all the results into a slice
	var results []getResult
	for i := 0; i < len(urls); i++ {
		results = append(results, <-respChannel)
	}

	// sort them
	sort.SliceStable(results, func(i, j int) bool {
		return results[i].index < results[j].index
	})

	// then do this loop, but over the sorted versions rather than the channel

	var bodies []io.Reader
	for _, res := range results {
		if res.err != nil {
			return "", res.err
		}
		defer res.body.Close()
		bodies = append(bodies, res.body)
	}

	return Stitcher(bodies...), nil
}
