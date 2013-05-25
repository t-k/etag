package main

import (
	"github.com/t-k/etag"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if etag.IsStale(w, r, "etag-key") {
			// normal response processing
			w.Write([]byte("Hello world!"))
		}
		// If the request is fresh (i.e. it's not modified) then you don't need to do anything.
		// IsStale automatically send a "304 Not Modified"
	})
	http.ListenAndServe(":8000", nil)
}
