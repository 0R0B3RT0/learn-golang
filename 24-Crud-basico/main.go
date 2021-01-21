package main

import (
	"fmt"
	"log"
	"net/http"

	"crud/servidor"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods("POST")

	fmt.Println("Escutando a porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
