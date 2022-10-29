package routes

import (
	"go-batch2/handlers"
	"go-batch2/pkg/middleware"
	"go-batch2/pkg/mysql"
	"go-batch2/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/user/create", h.CreateUser).Methods("POST")
	r.HandleFunc("/users", h.GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", h.FindUserById).Methods("GET")
	r.HandleFunc("/profile", middleware.Auth(h.GetProfile)).Methods("GET")
	r.HandleFunc("/user/update/{id}", middleware.Auth(middleware.UploadFile(h.UpdateUser))).Methods("PATCH")
	r.HandleFunc("/user/delete/{id}", middleware.Auth(h.DeleteUser)).Methods("DELETE")
}
