package main

import (
	"log"
	"net/http"
)

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func main() {
	const rootDir = "."
	const port = "8080"

	mux := http.NewServeMux()

	

	dir := http.Dir(rootDir)
	fileServer := http.FileServer(dir)
	mux.Handle("/app/", http.StripPrefix("/app", fileServer))
	mux.HandleFunc("/healthz", readinessHandler)

	server := &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n",rootDir, port)
	log.Fatal(server.ListenAndServe())
}