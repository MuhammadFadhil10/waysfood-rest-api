package main

import (
	"fmt"
	"go-batch2/database"
	"go-batch2/pkg/mysql"
	"go-batch2/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	
	mysql.DatabaseInit()

	database.RunMigration()

	r := mux.NewRouter()

	routes.RoutesInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running on port 8000")
	http.ListenAndServe(":8000", r)
}
