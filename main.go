package main

import (
	"Users-CRUD-API/config"
	"Users-CRUD-API/routes"
	"log"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	r := routes.SetupRouter()
	log.Println("Server is running on port 8080")
	r.Run(":8080")
}
