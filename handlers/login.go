package handlers

import (
	"LoginProject/context"
	"LoginProject/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"log"
	"time"
)

func Login(ctx *context.AppCtx) error {
	user := models.User{}
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	err := ctx.DB.Where("username = ? AND password = ? ", username, password).First(&user).Error
	if err != nil {
		err := ctx.Status(fiber.StatusForbidden).SendString("username or password does not correct")
		if err != nil {
			log.Fatalln("cant send message: ", err)
		}
		return err
	}
	//set cookie
	id, err := uuid.NewV4()
	if err != nil {
		log.Println("cant create uuid:", err)
	}
	ctx.Cookie(&fiber.Cookie{
		Name:    "Login-session",
		Value:   id.String(),
		Expires: time.Now().Add(time.Minute),
	})
	//cookie valuesini ve user idsini sessions tablesine kaydet.
	session := models.Sessions{
		SessionId: id.String(),
		UserId:    user.Id,
	}
	err = ctx.DB.Model(&models.Sessions{}).Create(&session).Error
	if err != nil {
		log.Println("cant create user:", err)
		return err
	}
	return ctx.Redirect("/mainmenu", fiber.StatusSeeOther)
}
