package handlers

import (
	"LoginProject/context"
	"LoginProject/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"html/template"
	"log"
	"time"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("./templates/*"))
	fmt.Println("parsed!")
}

func Logout(ctx *context.AppCtx) error {
	ctx.C.Cookie(&fiber.Cookie{
		Name:    "Login-session",
		Expires: time.Now().Add(time.Millisecond),
	})
	err := ctx.C.Redirect("/login", fiber.StatusSeeOther)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

func Index(ctx *context.AppCtx) error {
	user, err := utils.QuerybyUserid(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	ctx.C.Set("Content-Type", "text/html")
	err = tmp.ExecuteTemplate(ctx.C, "index.gohtml", user)
	if err != nil {
		log.Fatalln("cant execute template:", err)
	}
	return err
}
func UserJSON(ctx *context.AppCtx) error {
	user, err := utils.QuerybyUserid(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	err = ctx.C.JSON(&user)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}
