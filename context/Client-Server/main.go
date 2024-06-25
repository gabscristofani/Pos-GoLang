package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
	ctx := r.Context()
	log.Println("request iniciado")
	defer log.Println("request finalizado")
	select {
	case <-time.After(5 * time.Second):
		//imprime no console
		log.Println("request processada")
		//imprime no browser
		w.Write([]byte("request processada"))
	case <-ctx.Done():
		log.Println("request cancelada")
		http.Error(w, "request cancelada", http.StatusRequestTimeout)
	}
}
