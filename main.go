package main

import (
	"log"
	"net/http"
)

func main() {
	const rootDir = "."
	const port = "8080"

	mux := http.NewServeMux()
	dir := http.Dir(rootDir)
	fileServer := http.FileServer(dir)
	mux.Handle("/", fileServer)

	server := &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n",rootDir, port)
	log.Fatal(server.ListenAndServe())
}