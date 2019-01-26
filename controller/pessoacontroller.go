package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/augustoscher/API-C/model"
	"github.com/gorilla/mux"
)

var pessoas []model.Pessoa

//AllPessoasEndPoint retornas todas as pessoas
func AllPessoasEndPoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pessoas)
}

//FindPessoaEndpoint busca uma pessoa por id
func FindPessoaEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range pessoas {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&model.Pessoa{})
}

//CreatePessoaEndPoint exclui uma pessoa
func CreatePessoaEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var pessoa model.Pessoa
	_ = json.NewDecoder(r.Body).Decode(&pessoa)
	pessoa.ID = params["id"]
	pessoas = append(pessoas, pessoa)
	json.NewEncoder(w).Encode(pessoas)
}

//UpdatePessoaEndPoint atualizaa uma pessoa
func UpdatePessoaEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

//DeletePessoaEndPoint exclui uma pessoa
func DeletePessoaEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range pessoas {
		if item.ID == params["id"] {
			pessoas = append(pessoas[:index], pessoas[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(pessoas)
	}
}

//AdicionarPessoas adiciona pessoas para servirem de exemplo
func AdicionarPessoas() {
	endereco := model.Endereco{
		ID:     "1",
		Rua:    "Wilhelm Budag",
		Numero: "45",
		Cep:    "89045090",
		Cidade: "Blumenau",
		Estado: "SC",
		Pais:   "Brasil",
	}
	pessoas = append(pessoas, model.Pessoa{ID: "1", Nome: "John", Cpf: "02506196013", Endereco: &endereco})
	pessoas = append(pessoas, model.Pessoa{ID: "2", Nome: "Koko", Cpf: "02506196013", Endereco: &endereco})
}
