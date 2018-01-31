package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	listenAddr := flag.String("http", ":9099", "The host:port to which the server should bind")
	rootDir := flag.String("dir", "/app/www", "The directory to serve")
	flag.Parse()

	log.Printf("Serving %q at %q", *rootDir, *listenAddr)

	http.Handle("/", http.FileServer(http.Dir(*rootDir)))

	if err := http.ListenAndServe(*listenAddr, nil); err != nil {
		log.Fatalf("error listening: %+v", err)
	}
}
