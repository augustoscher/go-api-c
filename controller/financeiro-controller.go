package controller

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/augustoscher/API-C/model"
	"github.com/gorilla/mux"
)

var movimentos []model.MovimentacoesFinanceiras

//AllMovimentoFinanceiroEndPoint retornas todos os movimentos financeiros
func AllMovimentoFinanceiroEndPoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(movimentos)
}

//FindMovimentoFinanceiroPorCpfEndpoint busca um movimento financeiro por Cpf
func FindMovimentoFinanceiroPorCpfEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var movs []model.MovimentacoesFinanceiras

	for _, item := range movimentos {
		if item.Pessoa.Cpf == params["cpf"] {
			movs = append(movs, item)
		}
	}
	json.NewEncoder(w).Encode(movs)
}

//CreateMovimentoFinanceiroEndPoint adiciona um movimento financeiro
func CreateMovimentoFinanceiroEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var movimento model.MovimentacoesFinanceiras
	_ = json.NewDecoder(r.Body).Decode(&movimento)
	movimento.ID = params["id"]
	movimentos = append(movimentos, movimento)
	json.NewEncoder(w).Encode(movimentos)
}

//AdicionarMovimentosFinanceiros adiciona alguns movimentos para servirem de exemplo
func AdicionarMovimentosFinanceiros() {
	conta := model.ContaBancaria{
		ID:          "1",
		Instituicao: "Bradesco",
		Agencia:     "7445",
		Conta:       "12284",
		Digito:      "1",
	}
	pessoa := pessoas[0]
	pessoa2 := pessoas[1]
	movimentos = append(movimentos, model.MovimentacoesFinanceiras{ID: "1", Pessoa: &pessoa, Valor: 500.00, Data: "26/01/2019", Tipo: "C", Conta: &conta, Nomeestabelecimento: ""})
	movimentos = append(movimentos, model.MovimentacoesFinanceiras{ID: "2", Pessoa: &pessoa, Valor: 125.00, Data: "26/01/2019", Tipo: "D", Conta: &conta, Nomeestabelecimento: "Shopping"})
	movimentos = append(movimentos, model.MovimentacoesFinanceiras{ID: "3", Pessoa: &pessoa2, Valor: 100.00, Data: "26/01/2019", Tipo: "C", Conta: &conta, Nomeestabelecimento: ""})
}
