package main

import (
	"github.com/hacktiv8-fp-golang/final-project-01/internal/database"
	"github.com/hacktiv8-fp-golang/final-project-01/internal/router"
)

// @Title Todo Application
// @version 1.0
// @description This is a todo list management application
// @host localhost:8080
// @BasePath /
func main() {
	database.StartDB()
	router.StartServer()
}