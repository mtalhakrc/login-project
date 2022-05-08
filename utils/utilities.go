package utils

import (
	"LoginProject/context"
	"LoginProject/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func SendStaticFiles(app *fiber.App) {
	app.Static("/login", "./staticFiles/login.html")
	app.Static("/signup", "./staticFiles/signup.html")
	app.Static("/mainmenu/settings", "./staticFiles/settings.html")
}
func FormValue(ctx *context.AppCtx) models.User {
	user := models.User{}
	user.Username = ctx.FormValue("username")
	user.Password = ctx.FormValue("password")
	user.UsersInfo.Firstname = ctx.FormValue("firstname")
	user.UsersInfo.Lastname = ctx.FormValue("lastname")
	return user
}

func IsAlreadyTaken(ctx *context.AppCtx, user models.User) error {
	err := ctx.DB.Model(&user).Where("username = ?", user.Username).First(&user).Error
	if err == nil {
		fmt.Println("username is already taken!")
		return err
	}
	return err
}

func QuerybyUserid(ctx *context.AppCtx) (models.User, error) {
	user := models.User{}
	sessionId := ctx.Cookies("Login-session")
	//err := ctx.DB.Model(&models.User{}).Where("id = ?", ctx.Locals(sessionId)).First(&user).Error
	err := ctx.DB.Preload("UsersInfo").Where("id = ?", ctx.Locals(sessionId)).First(&user).Error
	if err != nil {
		log.Println(err)
	}
	return user, err
}
