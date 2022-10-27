package main

import (
	"fmt"
	"go-batch2/database"
	"go-batch2/pkg/mysql"
	"go-batch2/routes"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	mysql.DatabaseInit()

	database.RunMigration()

	r := mux.NewRouter()

	routes.RoutesInit(r.PathPrefix("/api/v1").Subrouter())

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	// uploads path prefix
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	// cors
	var allowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var allowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"})
	var allowedOrigins = handlers.AllowedOrigins([]string{"*"})

	fmt.Println("server running on port 8000")
	http.ListenAndServe(":8000", handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(r))
}
