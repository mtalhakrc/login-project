package router

import (
	"LoginProject/context"
	"LoginProject/handlers"
	"LoginProject/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8080",
		AllowCredentials: true,
	}))
	api := app.Group("/api")
	api.Post("/login", context.CtxWrap(handlers.Login))
	api.Post("/signup", context.CtxWrap(handlers.Signup))

	api.Use(context.CtxWrap(middleware.Session))
	api.Post("/settings", context.CtxWrap(handlers.Settings))
	api.Get("/user", context.CtxWrap(handlers.Index))
	api.Get("/logout", context.CtxWrap(handlers.Logout))
	//api.Get("/json", context.CtxWrap(handlers.UserJSON))

}
