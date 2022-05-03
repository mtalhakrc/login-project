package middleware

import (
	"LoginProject/context"
	"github.com/gofiber/fiber/v2"
)

type Sessions struct {
	Id      int    `gorm:"primaryKey" json:"user_id"`
	Sess_id string `json:"sess_Id"`
	User_id int    `json:"user_Id"`
}

//New middlewaresi, cookie yok ise logine geri döndürür.
func Session(ctx *context.AppCtx) error {
	cookievalue := ctx.Ctx.Cookies("Login-session")
	if cookievalue == "" {
		return ctx.Status(fiber.StatusUnauthorized).Redirect("/login")
	}
	//cookiedeki useri sor varsa dön yoksa devam et kayıt oluştur. contexte yaz
	user := Sessions{}
	err := ctx.DB.Model(&Sessions{}).Where("sess_id = ?", cookievalue).First(&user).Error
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).Redirect("/login")
	}
	ctx.Locals(user.Sess_id, user.User_id)
	return ctx.Next()
}

//In short, PostgreSQL error 42601 occurs due to the syntax errors in the code.
