package main

import (
	"net/http"

	"github.com/tsaqiffatih/crud-employee-go/database"
	"github.com/tsaqiffatih/crud-employee-go/routes"
)

func main() {
	db := database.InitDatabase()

	server := http.NewServeMux()

	routes.MapRoutes(server, db)

	http.ListenAndServe(":3000", server)

}
