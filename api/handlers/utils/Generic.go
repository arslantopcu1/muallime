package utils

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
)

type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func ErrorHTTP(mes string) ResponseHTTP {
	return ResponseHTTP{
		Success: false,
		Message: mes,
		Data:    nil,
	}
}

// Protected protect routes
func Protected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(ResponseHTTP{Success: false, Message: "Missing or malformed JWT", Data: nil})

	} else {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(ResponseHTTP{Success: false, Message: "Invalid or expired JWT", Data: nil})
	}
}

func GetClaimsJWT(c *fiber.Ctx) primitive.ObjectID {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["user"].(string)
	objID, err := primitive.ObjectIDFromHex(name)
	if err != nil {
		panic(err)
	}
	return objID
}
