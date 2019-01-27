package routers

import (
	"bitbucket.org/augustoscher/API-C/authentication"
	"bitbucket.org/augustoscher/API-C/controllers"
	"github.com/gorilla/mux"
)

//SetAuthenticationRoutes rotas de autenticação
func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/test", authentication.ValidateMiddleware(controllers.TestEndpoint)).Methods("GET")

	router.HandleFunc("/pessoas", authentication.ValidateMiddleware(controllers.AllPessoasEndPoint)).Methods("GET")
	router.HandleFunc("/pessoas", authentication.ValidateMiddleware(controllers.CreatePessoaEndPoint)).Methods("POST")
	router.HandleFunc("/pessoas/{id}", authentication.ValidateMiddleware(controllers.FindPessoaEndpoint)).Methods("GET")

	router.HandleFunc("/movimentosfinanceiro", authentication.ValidateMiddleware(controllers.AllMovimentoFinanceiroEndPoint)).Methods("GET")
	router.HandleFunc("/movimentosfinanceiro", authentication.ValidateMiddleware(controllers.CreateMovimentoFinanceiroEndPoint)).Methods("POST")
	router.HandleFunc("/movimentosfinanceiro/{cpf}", authentication.ValidateMiddleware(controllers.FindMovimentoFinanceiroPorCpfEndpoint)).Methods("GET")

	router.HandleFunc("/consultascpf", authentication.ValidateMiddleware(controllers.AllConsultasCpfEndPoint)).Methods("GET")
	router.HandleFunc("/consultascpf", authentication.ValidateMiddleware(controllers.CreateConsultaCpfEndPoint)).Methods("POST")
	router.HandleFunc("/consultascpf/{cpf}", authentication.ValidateMiddleware(controllers.FindConsultasCpfPorCpfEndpoint)).Methods("GET")

	router.HandleFunc("/movimentoscartao", authentication.ValidateMiddleware(controllers.AllMovimentoCartaoEndPoint)).Methods("GET")
	router.HandleFunc("/movimentoscartao", authentication.ValidateMiddleware(controllers.CreateMovimentoCartaoEndPoint)).Methods("POST")
	router.HandleFunc("/movimentoscartao/{cpf}", authentication.ValidateMiddleware(controllers.FindMovimentoCartaoPorCpfEndpoint)).Methods("GET")

	return router
}
