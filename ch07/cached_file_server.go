package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

type CacheFile struct {
	Content io.ReadSeeker
	ModTime time.Time
}

var cache map[string]*CacheFile
var mutex = new(sync.RWMutex)

func init() {
	cache = make(map[string]*CacheFile)
}

func main() {
	http.HandleFunc("/", serveFiles)
	http.ListenAndServe(":8080", nil)
}

func serveFiles(rw http.ResponseWriter, req *http.Request) {
	mutex.RLock()
	cf, found := cache[req.URL.Path]
	mutex.RUnlock()

	if !found {
		mutex.Lock()
		defer mutex.Unlock()

		fname := "./files" + req.URL.Path
		f, err := os.Open(fname)
		defer f.Close()

		if err != nil {
			http.NotFound(rw, req)
			return
		}

		var buf bytes.Buffer
		_, err = io.Copy(&buf, f)

		if err != nil {
			http.NotFound(rw, req)
			return
		}

		r := bytes.NewReader(buf.Bytes())
		info, _ := f.Stat()
		cf = &CacheFile{
			Content: r,
			ModTime: info.ModTime(),
		}
		cache[req.URL.Path] = cf
	}

	http.ServeContent(rw, req, req.URL.Path, cf.ModTime, cf.Content)
}
