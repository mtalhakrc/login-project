package handlers

import (
	"LoginProject/context"
	"LoginProject/models"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"log"
	"time"
)

func Login(ctx *context.AppCtx) error {
	user := models.User{}
	err := json.Unmarshal(ctx.Body(), &user)
	if err != nil {
		log.Println("cant decode the body:", err)
	}
	err = ctx.DB.Where("username = ? AND password = ? ", user.Username, user.Password).First(&user).Error
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).SendString("username or password not correct")
	}
	//set cookie
	id, err := uuid.NewV4()
	if err != nil {
		log.Println("cant create uuid:", err)
	}
	ctx.Cookie(&fiber.Cookie{
		Name:    "Login-session",
		Value:   id.String(),
		Expires: time.Now().Add(time.Hour),
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
	return ctx.JSON("succesfull!")
}
