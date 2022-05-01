package handlers

import (
	"LoginProject/context"
	"LoginProject/middleware"
	"LoginProject/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"log"
	"time"
)

func Login(ctx *context.AppCtx) error {
	user := models.User{}
	username := ctx.C.FormValue("username")
	password := ctx.C.FormValue("password")
	err := ctx.DB.Where("username = ? AND password = ? ", username, password).First(&user).Error
	if err != nil {
		err := ctx.C.Status(fiber.StatusForbidden).SendString("username or password does not correct")
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
	ctx.C.Cookie(&fiber.Cookie{
		Name:    "Login-session",
		Value:   id.String(),
		Expires: time.Now().Add(time.Minute),
	})
	//cookie valuesini ve user idsini sessions tablesine kaydet.
	session := middleware.Sessions{
		Sess_id: id.String(),
		User_id: user.Id,
	}
	err = ctx.DB.Model(&session).Create(&session).Error
	if err != nil {
		log.Println("cant create user:", err)
		return err
	}
	err = ctx.C.Redirect("/mainmenu", fiber.StatusSeeOther)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}
