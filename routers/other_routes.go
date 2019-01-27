package routers

import (
	"bitbucket.org/augustoscher/API-C/controllers"
	"github.com/gorilla/mux"
)

//SetOtherRoutes adiciona outras rotas que nao requerem autenticação
func SetOtherRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/endpoint/livre", controllers.EndpointLivre).Methods("GET")
	return router
}
