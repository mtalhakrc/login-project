package utils

import (
	"LoginProject/context"
	"LoginProject/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

var err error

func SendStaticFiles(app *fiber.App) {
	app.Static("/login", "./staticFiles/login.html")
	app.Static("/signup", "./staticFiles/signup.html")
	app.Static("/mainmenu/settings", "./staticFiles/settings.html")
}
func FormValue(ctx *context.AppCtx) (models.User, error) {
	//get handlers information
	//user := User{}
	var user models.User
	user.Username = ctx.FormValue("username")
	user.Password = ctx.FormValue("password")
	user.State = ctx.FormValue("state")
	age, err := strconv.Atoi(ctx.FormValue("age"))
	if err != nil {
		log.Println("cant convert:", err)
		return user, err
	}
	user.Age = age
	return user, err
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
	sessionId := ctx.Cookies("Login-session")
	user := models.User{}
	//get user values by querying user id
	err = ctx.DB.Model(&user).Where("id = ?", ctx.Locals(sessionId)).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, err
}
