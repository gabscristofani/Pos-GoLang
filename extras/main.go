package main

import (
	"log"
	"net/http"
)

func main() {
	server := &http.Server{Addr: ":3000"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World! \n"))
	})

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
