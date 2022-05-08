package handlers

import (
	"LoginProject/context"
	"LoginProject/models"
	"LoginProject/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"log"
	"time"
)

func Signup(ctx *context.AppCtx) error {
	user := utils.FormValue(ctx)
	//is username already taken?
	err := utils.IsAlreadyTaken(ctx, user)
	if err == nil {
		fmt.Println(err)
		return err
	}
	//create a record
	err = ctx.DB.Create(&user).Error
	if err != nil {
		log.Fatalln("cant create user:", err)
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
	err = ctx.Redirect("/mainmenu", fiber.StatusSeeOther)
	if err != nil {
		log.Println(err)
	}
	return err
}
