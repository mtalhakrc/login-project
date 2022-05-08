package middleware

import (
	"LoginProject/context"
	"LoginProject/models"
	"github.com/gofiber/fiber/v2"
)

//New middlewaresi, cookie yok ise logine geri döndürür.
func Session(ctx *context.AppCtx) error {
	cookievalue := ctx.Cookies("Login-session")
	if cookievalue == "" {
		return ctx.Status(fiber.StatusUnauthorized).Redirect("/login")
	}
	//cookiedeki useri sor varsa dön yoksa devam et kayıt oluştur. contexte yaz
	user := models.Sessions{}
	err := ctx.DB.Model(&models.Sessions{}).Where("session_id = ?", cookievalue).First(&user).Error
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).Redirect("/login")
	}
	ctx.Locals(user.SessionId, user.UserId)
	return ctx.Next()
}

//In short, PostgreSQL error 42601 occurs due to the syntax errors in the code.
