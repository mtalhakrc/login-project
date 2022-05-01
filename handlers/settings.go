package handlers

import (
	"LoginProject/context"
	"LoginProject/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Settings(ctx *context.AppCtx) error {
	//form the value
	updatedUser, err := utils.FormValue(ctx)
	if err != nil {
		log.Println("cant form value:", err)
		return err
	}
	gonnaUpdate, err := utils.QuerybyUserid(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	//check is username is alrady taken?
	err = utils.IsAlreadyTaken(ctx, updatedUser)
	if err == nil {
		err = ctx.C.Status(fiber.StatusBadRequest).SendString("username is already taken!")
		if err != nil {
			log.Println("cant send message:", err)
			return err
		}
		return err
	}

	err = ctx.DB.Model(&gonnaUpdate).Where("username = ?", gonnaUpdate.Username).Updates(&updatedUser).Error
	if err != nil {
		log.Println("cant update user: ", err)
		return err
	}
	fmt.Println("user is updated:", updatedUser)
	//go back mainmenu
	err = ctx.C.Redirect("/mainmenu", fiber.StatusSeeOther)
	if err != nil {
		log.Println(err)
	}
	return err
}
