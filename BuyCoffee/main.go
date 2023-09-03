package main

import (
	Base "BuyCoffee/cmd"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fiberApp := fiber.New()
	var err error
	for _, app := range Base.Programs {
		err = app.Migration()
		if err != nil {
			panic(err)
		}
		app.SetRouter(Base.BaseUrl+"/"+app.AppName(), fiberApp)
	}

	fmt.Print(fiberApp.Listen(Base.BaseUrl))
}
