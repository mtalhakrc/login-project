package middleware

import (
	"LoginProject/context"
	"LoginProject/models"
	"github.com/gofiber/fiber/v2"
)

//New middlewaresi, cookie yok ise logine geri döndürür.
func Session(ctx *context.AppCtx) error {
	cookieValue := ctx.Cookies("Login-session")
	if cookieValue == "" {
		return ctx.Status(fiber.StatusForbidden).SendString("user has no auth")
	}
	//cookiedeki useri sor varsa dön yoksa devam et kayıt oluştur. contexte yaz
	user := models.Sessions{}
	err := ctx.DB.Model(&models.Sessions{}).Where("session_id = ?", cookieValue).First(&user).Error
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).SendString("username or password not correct")
	}
	ctx.Locals(user.SessionId, user.UserId)
	return ctx.Next()
}

//In short, PostgreSQL error 42601 occurs due to the syntax errors in the code.
