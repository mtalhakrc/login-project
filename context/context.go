package context

import (
	"LoginProject/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AppCtx struct {
	*fiber.Ctx
	DB *gorm.DB
}

func CtxWrap(h func(ctx *AppCtx) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return h(&AppCtx{
			Ctx: c,
			DB:  database.DB(),
		})
	}
}
