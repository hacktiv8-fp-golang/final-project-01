package main

import (
	"final-project-01/internal/database"
	"final-project-01/internal/router"
)

func main() {
	database.StartDB()
	router.StartServer()
}