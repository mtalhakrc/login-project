package handlers

import (
	"LoginProject/context"
	"LoginProject/middleware"
	"LoginProject/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"log"
	"time"
)

func Signup(ctx *context.AppCtx) error {
	user, err := utils.FormValue(ctx)
	if err != nil {
		log.Println("cant form value:", err)
	}
	//is username already taken?
	err = utils.IsAlreadyTaken(ctx, user)
	if err == nil {
		fmt.Println(err)
		return err
	}
	//create a record
	err = ctx.DB.Model(&user).Create(&user).Error
	if err != nil {
		log.Fatalln("cant create user:", err)
	}
	err = ctx.DB.Model(&user).Where("username = ?", user.Username).First(&user).Error

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
	}
	return err
}
