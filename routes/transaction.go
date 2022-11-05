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

	r.HandleFunc("/transaction", middleware.Auth(h.CreateTransaction)).Methods("POST")
	r.HandleFunc("/partner/transaction/{partnerId}", middleware.Auth(h.GetTransactionByPartner)).Methods("GET")
	r.HandleFunc("/notification", h.Notification).Methods("POST")
	// r.HandleFunc("/transactions", middleware.Auth(h.GetAllTransaction)).Methods("GET")
	// r.HandleFunc("/transaction/{id}", middleware.Auth(h.GetTransactionByID)).Methods("GET")
	// r.HandleFunc("/transaction/update/{id}", middleware.Auth(h.UpdateTransaction)).Methods("PATCH")
	// r.HandleFunc("/transaction/delete/{id}", middleware.Auth(h.DeleteTransaction)).Methods("DELETE")
}
