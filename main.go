package main

import (
	"log"
	"net/http"

	"bitbucket.org/augustoscher/API-C/controller"
	"github.com/gorilla/mux"
)

func main() {
	controller.AdicionarPessoas()
	controller.AdicionarMovimentosFinanceiros()
	controller.AdicionarConsultas()
	controller.AdicionarMovimentosCartaoCredito()

	r := mux.NewRouter()
	r.HandleFunc("/pessoas", controller.AllPessoasEndPoint).Methods("GET")
	r.HandleFunc("/pessoas", controller.CreatePessoaEndPoint).Methods("POST")
	r.HandleFunc("/pessoas/{id}", controller.FindPessoaEndpoint).Methods("GET")

	r.HandleFunc("/movimentosfinanceiro", controller.AllMovimentoFinanceiroEndPoint).Methods("GET")
	r.HandleFunc("/movimentosfinanceiro", controller.CreateMovimentoFinanceiroEndPoint).Methods("POST")
	r.HandleFunc("/movimentosfinanceiro/{cpf}", controller.FindMovimentoFinanceiroPorCpfEndpoint).Methods("GET")

	r.HandleFunc("/consultascpf", controller.AllConsultasCpfEndPoint).Methods("GET")
	r.HandleFunc("/consultascpf", controller.CreateConsultaCpfEndPoint).Methods("POST")
	r.HandleFunc("/consultascpf/{cpf}", controller.FindConsultasCpfPorCpfEndpoint).Methods("GET")

	r.HandleFunc("/movimentoscartao", controller.AllMovimentoCartaoEndPoint).Methods("GET")
	r.HandleFunc("/movimentoscartao", controller.CreateMovimentoCartaoEndPoint).Methods("POST")
	r.HandleFunc("/movimentoscartao/{cpf}", controller.FindMovimentoCartaoPorCpfEndpoint).Methods("GET")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
