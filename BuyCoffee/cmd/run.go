package Base

import (
	"BuyCoffee/User"
	"github.com/gofiber/fiber/v2"
)

var (
	fiberApp *fiber.App
	BaseUrl  = ":8080"
)
var Programs = []App{
	new(User.User),
}
