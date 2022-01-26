package main

import (
	"log"
	"net/http"
)

func main() {
	handler := WebsiteStitcherServer(WebsiteStitcher)
	log.Fatal(http.ListenAndServe(":2000", handler))
}
