package main

import (
	"github.com/hacktiv8-fp-golang/final-project-01/internal/database"
	"github.com/hacktiv8-fp-golang/final-project-01/internal/router"
)

// @title Todo Application
// @version 1.0
// @description This is a todo list management application
// @host final-project-01-production.up.railway.app
// @BasePath /
func main() {
	database.StartDB()
	router.StartServer()
}
