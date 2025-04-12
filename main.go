package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	my_server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	mux.Handle("/", http.FileServer(http.Dir(".")))
	my_server.ListenAndServe()

}
