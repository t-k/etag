package etag

import (
	"bufio"
	"github.com/bmizerany/assert"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
)

const (
	ETAG_KEY = "foobar"
)

func init() {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if IsStale(w, r, ETAG_KEY) {
				w.Write([]byte("hello"))
			} else {
				w.Write([]byte(""))
			}
		})
		http.ListenAndServe(":4567", nil)
	}()
}

func get(etag string) (r *http.Response, err error) {
	var conn net.Conn
	if conn, err = net.Dial("tcp", "localhost:4567"); err == nil {
		req, _ := http.NewRequest("GET", "http://localhost:4567/", nil)
		req.Header.Set("If-None-Match", etag)
		req.Write(conn)
		buf := bufio.NewReader(conn)
		r, err = http.ReadResponse(buf, req)
	}
	return
}

func TestEtagIsStaleWithoutEtag(t *testing.T) {
	res, err := get("")
	if err == nil {
		contents, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t, string(contents), "hello")
		assert.Equal(t, res.Status, "200 OK")
		assert.Equal(t, string(res.Header.Get("Etag")), ETAG_KEY)
	}
}

func TestEtagIsStaleWithEtag(t *testing.T) {
	res, err := get(ETAG_KEY)
	if err == nil {
		contents, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t, string(contents), "")
		assert.Equal(t, res.Status, "304 Not Modified")
		assert.Equal(t, string(res.Header.Get("Etag")), ETAG_KEY)
	}
}
