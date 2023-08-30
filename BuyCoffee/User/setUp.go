package User

import "github.com/gofiber/fiber/v2"

type User struct {
}

func (u *User) SetRouter(url string, application *fiber.App) {
	application.Post(url+"/login", Login)

	application.Post(url+"/logout", func(ctx *fiber.Ctx) error {
		return nil
	})

}

func (u *User) AppName() string {
	return "user"
}
