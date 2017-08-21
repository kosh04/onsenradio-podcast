package main

import (
	"log"
	"net/http"
)

func newServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, baseURL, http.StatusMovedPermanently)
	})
	mux.HandleFunc("/podcast", func(w http.ResponseWriter, r *http.Request) {
		podcastWriter(w)
	})
	return mux
}

func serve(addr string) {
	err := http.ListenAndServe(addr, newServeMux())

	if err != nil {
		log.Fatal("http.ListenAndServe():", err)
	}
}
