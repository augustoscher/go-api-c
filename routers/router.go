package routers

import "github.com/gorilla/mux"

//InitRoutes inicializa rotas
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetOtherRoutes(router)
	router = SetAuthenticationRoutes(router)
	return router
}
