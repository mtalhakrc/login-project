package router

import (
	"LoginProject/context"
	"LoginProject/handlers"
	"LoginProject/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/login", context.CtxWrap(handlers.Login))
	app.Post("/signup", context.CtxWrap(handlers.Signup))
	app.Use(context.CtxWrap(middleware.Session))
	api := app.Group("/mainmenu")
	api.Get("/", context.CtxWrap(handlers.Index))
	api.Get("/json", context.CtxWrap(handlers.UserJSON))
	api.Post("/settings", context.CtxWrap(handlers.Settings))
	api.Get("/logout", context.CtxWrap(handlers.Logout))
}
