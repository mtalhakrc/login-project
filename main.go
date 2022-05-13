package main

import (
	"LoginProject/database"
	"LoginProject/router"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	database.InitializeDB()
	//utils.SendStaticFiles(app)
	router.SetupRoutes(app)

	log.Fatalln(app.Listen("localhost:8081"))
}
