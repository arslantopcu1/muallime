package login

import (
	"github.com/gofiber/fiber/v2"
)

func Router(v1 fiber.Router) {
	v1.Get("/create-session", Get)
}
