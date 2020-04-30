package main

// go run fs.go -port=8181 -path=site

import (
	"flag"
	"net/http"
	"os"
)

func main() {
	var dir string
	port := flag.String("port", "3000", "port to serve HTTP")
	path := flag.String("path", "site", "path to serve content")
	flag.Parse()

	if *path == "" {
		dir, _ = os.Getwd()
	} else {
		dir = *path
	}

	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(dir)))

}
