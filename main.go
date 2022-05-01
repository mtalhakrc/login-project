package main

import (
	"LoginProject/database"
	"LoginProject/router"
	"LoginProject/utils"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	database.InitializeDB()
	utils.SendStaticFiles(app)
	router.SetupRoutes(app)
	err := app.Listen("localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
}
