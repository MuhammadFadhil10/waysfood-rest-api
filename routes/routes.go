package routes

import "github.com/gorilla/mux"

func RoutesInit(r *mux.Router) {
	UserRoutes(r)
	AuthRoutes(r)
	ProductRoutes(r)
	TransactionRoutes(r)
}
