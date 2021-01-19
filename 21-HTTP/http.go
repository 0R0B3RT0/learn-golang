package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Primeiro servidor http!!!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Cerregar a página de usuários!!!"))
}

func main() {

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
