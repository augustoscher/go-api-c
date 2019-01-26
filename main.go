package main

import (
	"log"
	"net/http"

	"bitbucket.org/augustoscher/API-C/controllers"
	"github.com/gorilla/mux"
)

func main() {
	controllers.AdicionarPessoas()
	controllers.AdicionarMovimentosFinanceiros()
	controllers.AdicionarConsultas()
	controllers.AdicionarMovimentosCartaoCredito()

	r := mux.NewRouter()
	r.HandleFunc("/pessoas", controllers.AllPessoasEndPoint).Methods("GET")
	r.HandleFunc("/pessoas", controllers.CreatePessoaEndPoint).Methods("POST")
	r.HandleFunc("/pessoas/{id}", controllers.FindPessoaEndpoint).Methods("GET")

	r.HandleFunc("/movimentosfinanceiro", controllers.AllMovimentoFinanceiroEndPoint).Methods("GET")
	r.HandleFunc("/movimentosfinanceiro", controllers.CreateMovimentoFinanceiroEndPoint).Methods("POST")
	r.HandleFunc("/movimentosfinanceiro/{cpf}", controllers.FindMovimentoFinanceiroPorCpfEndpoint).Methods("GET")

	r.HandleFunc("/consultascpf", controllers.AllConsultasCpfEndPoint).Methods("GET")
	r.HandleFunc("/consultascpf", controllers.CreateConsultaCpfEndPoint).Methods("POST")
	r.HandleFunc("/consultascpf/{cpf}", controllers.FindConsultasCpfPorCpfEndpoint).Methods("GET")

	r.HandleFunc("/movimentoscartao", controllers.AllMovimentoCartaoEndPoint).Methods("GET")
	r.HandleFunc("/movimentoscartao", controllers.CreateMovimentoCartaoEndPoint).Methods("POST")
	r.HandleFunc("/movimentoscartao/{cpf}", controllers.FindMovimentoCartaoPorCpfEndpoint).Methods("GET")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
