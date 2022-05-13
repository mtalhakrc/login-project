package utils

import (
	"LoginProject/context"
	"LoginProject/models"
)

func IsAlreadyTaken(ctx *context.AppCtx, user models.User) error {
	return ctx.DB.Model(&user).Where("username = ?", user.Username).First(&user).Error
}

func QuerybyUserid(ctx *context.AppCtx) (models.User, error) {
	user := models.User{}
	sessionId := ctx.Cookies("Login-session")
	err := ctx.DB.Preload("UsersInfo").Where("id = ?", ctx.Locals(sessionId)).First(&user).Error
	return user, err
}
