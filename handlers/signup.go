package handlers

import (
	"LoginProject/context"
	"LoginProject/models"
	"LoginProject/utils"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"log"
	"time"
)

func Signup(ctx *context.AppCtx) error {
	user := models.User{}
	err := json.Unmarshal(ctx.Body(), &user)
	err = json.Unmarshal(ctx.Body(), &user.UsersInfo)
	if err != nil {
		log.Println("cant decode the body:", err)
	}
	fmt.Println(user)

	//is username already taken?
	err = utils.IsAlreadyTaken(ctx, user)
	if err == nil {
		fmt.Println("çalıştı!")
		return ctx.Status(fiber.StatusForbidden).SendString("username is already taken!")
	}

	//create a record
	err = ctx.DB.Create(&user).Error
	if err != nil {
		log.Fatalln("cant create user:", err)
	}

	//set cookie
	id := uuid.NewV4()
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
	return ctx.DB.Model(&models.Sessions{}).Create(&session).Error
}
