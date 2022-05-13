package handlers

import (
	"LoginProject/context"
	"LoginProject/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

//var tmp *template.Template

type deneme struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func Index(ctx *context.AppCtx) error {
	user, err := utils.QuerybyUserid(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	gonderilecek := deneme{
		Username:  user.Username,
		Password:  user.Password,
		Firstname: user.UsersInfo.Firstname,
		Lastname:  user.UsersInfo.Lastname,
	}
	fmt.Println(gonderilecek)
	return ctx.JSON(gonderilecek)
}
func Logout(ctx *context.AppCtx) error {
	ctx.Cookie(&fiber.Cookie{
		Name:    "Login-session",
		Expires: time.Now().Add(time.Millisecond),
	})
	return nil
}

func UserJSON(ctx *context.AppCtx) error {
	user, err := utils.QuerybyUserid(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	return ctx.JSON(&user)
}
