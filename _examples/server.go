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
		} else {
			// you don't need to do anything.
		}
	})
	http.ListenAndServe(":8000", nil)
}
