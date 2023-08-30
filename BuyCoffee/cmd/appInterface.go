package Base

import "github.com/gofiber/fiber/v2"

type App interface {
	SetRouter(prefix string, application *fiber.App)
	AppName() string
}
