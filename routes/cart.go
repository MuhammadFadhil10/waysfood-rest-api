package routes

import (
	"go-batch2/handlers"
	"go-batch2/pkg/middleware"
	"go-batch2/pkg/mysql"
	"go-batch2/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/cart/add/{productID}", middleware.Auth(h.AddToCart)).Methods("POST")
	r.HandleFunc("/carts", middleware.Auth(h.GetChartByUserID)).Methods("GET")
	r.HandleFunc("/cart/update/{productID}", middleware.Auth(h.DeleteChartByQty)).Methods("PATCH")
	r.HandleFunc("/cart/delete/{productID}", middleware.Auth(h.DeleteChartByID)).Methods("DELETE")
	r.HandleFunc("/cart/delete-all", middleware.Auth(h.DeleteAllCart)).Methods("DELETE")

}
