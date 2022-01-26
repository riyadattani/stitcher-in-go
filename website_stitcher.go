package main

import (
	"io"
	"net/http"
	"sort"
	"sync"
)

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
	// google how to sort a slice
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
