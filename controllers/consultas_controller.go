package controllers

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/augustoscher/API-C/models"
	"github.com/gorilla/mux"
)

var consultas []models.HistoricoConsultaCpf

//AllConsultasCpfEndPoint retornas todas as consultas realizadas ao CPF
func AllConsultasCpfEndPoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(consultas)
}

//FindConsultasCpfPorCpfEndpoint retorna todas as consultas de determinado cpf
func FindConsultasCpfPorCpfEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var cons []models.HistoricoConsultaCpf

	for _, item := range consultas {
		if item.Pessoa.Cpf == params["cpf"] {
			cons = append(cons, item)
		}
	}
	json.NewEncoder(w).Encode(cons)
}

//CreateConsultaCpfEndPoint adiciona um registro no histórico de consultas a cpf
func CreateConsultaCpfEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var consulta models.HistoricoConsultaCpf
	_ = json.NewDecoder(r.Body).Decode(&consulta)
	consulta.ID = params["id"]
	consultas = append(consultas, consulta)
	json.NewEncoder(w).Encode(consultas)
}

//AdicionarConsultas adiciona pessoas para servirem de exemplo
func AdicionarConsultas() {
	endereco := models.Endereco{
		ID:     "1",
		Rua:    "Wilhelm Budag",
		Numero: "45",
		Cep:    "89045090",
		Cidade: "Blumenau",
		Estado: "SC",
		Pais:   "Brasil",
	}
	estabelecimento := models.Estabelecimento{
		ID:       "1",
		Nome:     "Meu Estabelecimento",
		Cnpj:     "21555582000142",
		Ie:       "687516546",
		Endereco: &endereco,
	}
	estabelecimento2 := models.Estabelecimento{
		ID:       "2",
		Nome:     "Estabelecimento2",
		Cnpj:     "21123482000142",
		Ie:       "79798798",
		Endereco: &endereco,
	}

	pessoa := pessoas[0]
	pessoa2 := pessoas[1]
	consultas = append(consultas, models.HistoricoConsultaCpf{ID: "1", Pessoa: &pessoa, Data: "26/01/2019", Situacaocpf: "Adimplente", Consultadopor: &estabelecimento})
	consultas = append(consultas, models.HistoricoConsultaCpf{ID: "2", Pessoa: &pessoa, Data: "14/01/2019", Situacaocpf: "Adimplente", Consultadopor: &estabelecimento2})
	consultas = append(consultas, models.HistoricoConsultaCpf{ID: "3", Pessoa: &pessoa2, Data: "20/07/2018", Situacaocpf: "Negativado", Consultadopor: &estabelecimento2})
}
