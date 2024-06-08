package main

import "net/http" // Importa o pacote http

func main() {
	http.ListenAndServe(":8080", nil)
}
