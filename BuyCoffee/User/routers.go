package User

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

var privateKey *rsa.PrivateKey

func Login(ctx *fiber.Ctx) error {
	request := RequestLogin{}
	var t string
	rng := rand.Reader
	var err error
	privateKey, err = rsa.GenerateKey(rng, 2048)
	err = ctx.BodyParser(&request)
	if err != nil {
		return err
	}

	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	t, err = token.SignedString(privateKey)
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.JSON(fiber.Map{"token": t})

}
