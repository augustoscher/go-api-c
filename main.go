package main

import (
	"log"
	"net/http"

	"bitbucket.org/augustoscher/API-C/controller"
	"github.com/gorilla/mux"
)

func main() {
	controller.AdicionarPessoas()
	r := mux.NewRouter()
	r.HandleFunc("/pessoas", controller.AllPessoasEndPoint).Methods("GET")
	r.HandleFunc("/pessoas", controller.CreatePessoaEndPoint).Methods("POST")
	r.HandleFunc("/pessoas", controller.UpdatePessoaEndPoint).Methods("PUT")
	r.HandleFunc("/pessoas", controller.DeletePessoaEndPoint).Methods("DELETE")
	r.HandleFunc("/pessoas/{id}", controller.FindPessoaEndpoint).Methods("GET")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
