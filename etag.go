package etag

import (
	"net/http"
)

// Get Get the "If-None-Match" from header.
func Get(r *http.Request) string {
	return r.Header.Get("If-None-Match")
}

// Set Sets the etag.
func Set(w http.ResponseWriter, key string) {
	w.Header().Set("Etag", key)
}

// IsStale sets the etag on the response and checks it against the client request.
// If the request doesn’t match the options provided, the request is considered stale and should be generated from scratch.
// Otherwise, it’s fresh and we don’t need to generate anything and a reply of "304 Not Modified" is sent.
func IsStale(w http.ResponseWriter, r *http.Request, key string) (isStale bool) {
	ifNoneMatch := Get(r)
	Set(w, key)
	if ifNoneMatch != "" && ifNoneMatch == key {
		isStale = false
		w.WriteHeader(304)
		w.Write([]byte(""))
	} else {
		isStale = true
	}
	return
}
