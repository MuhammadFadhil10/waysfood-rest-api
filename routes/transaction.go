package routes

import (
	"go-batch2/handlers"
	"go-batch2/pkg/middleware"
	"go-batch2/pkg/mysql"
	"go-batch2/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions", middleware.Auth(h.ShowTransaction)).Methods("GET")
	r.HandleFunc("/transaction/{id}", middleware.Auth(h.GetTransactionByID)).Methods("GET")
	r.HandleFunc("/transaction", middleware.Auth(h.CreateTransaction)).Methods("POST")
	r.HandleFunc("/transaction/update/{id}", middleware.Auth(h.UpdateTransaction)).Methods("PATCH")
	r.HandleFunc("/transaction/delete/{id}", middleware.Auth(h.DeleteTransaction)).Methods("DELETE")
}
