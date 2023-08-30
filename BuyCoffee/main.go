package BuyCoffee

import (
	Base "BuyCoffee/cmd"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fiberApp := fiber.New()
	for _, app := range Base.Programs {
		app.SetRouter(Base.BaseUrl+"/"+app.AppName(), fiberApp)
	}

	fmt.Print(fiberApp.Listen(Base.BaseUrl))
}
