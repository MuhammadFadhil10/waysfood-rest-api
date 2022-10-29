package routes

import (
	"go-batch2/handlers"
	"go-batch2/pkg/middleware"
	"go-batch2/pkg/mysql"
	"go-batch2/repositories"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(userRepository)

	r.HandleFunc("/products", h.GetProducts).Methods("GET")
	r.HandleFunc("/product/{id}", h.GetProductByID).Methods("GET")
	r.HandleFunc("/products/{userId}", middleware.Auth(h.GetProductByPartner)).Methods("GET")
	r.HandleFunc("/product/create", middleware.Auth(middleware.UploadFile(h.CreateProduct))).Methods("POST")
	r.HandleFunc("/product/update/{id}", middleware.Auth(h.UpdateProduct)).Methods("PATCH")
	r.HandleFunc("/product/delete/{id}", middleware.Auth(h.DeleteProduct)).Methods("DELETE")

}
