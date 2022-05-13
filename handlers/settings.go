package handlers

import (
	"LoginProject/context"
	"LoginProject/models"
	"LoginProject/utils"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Settings(ctx *context.AppCtx) error {
	//form the value
	updatedUser := models.User{}
	err := json.Unmarshal(ctx.Body(), &updatedUser)
	err = json.Unmarshal(ctx.Body(), &updatedUser.UsersInfo)
	if err != nil {
		log.Println("cant decode the body:", err)
	}
	fmt.Println(updatedUser)
	gonnaUpdate, err := utils.QuerybyUserid(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	//check is username is alrady taken?
	err = utils.IsAlreadyTaken(ctx, updatedUser)
	if err == nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("username is already taken!")
	}
	err = ctx.DB.Model(&models.User{}).Where("users.username = ?", gonnaUpdate.Username).Updates(updatedUser).Error
	err = ctx.DB.Model(&models.UsersInfo{}).Where("users_infos.firstname = ?", gonnaUpdate.UsersInfo.Firstname).Updates(updatedUser.UsersInfo).Error
	return err
}
