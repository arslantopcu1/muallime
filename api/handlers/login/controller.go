package login

import (
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
)

// Create session
// @Summary Create Session
// @Description Get Create Session
// @Tags Authorization
// @Accept json
// @Produce json
// @Param     id    path     string  true  "Authorization id"
// @Success 200 {object} utils.ResponseHTTP{data=login.Authorization}
// @Failure 404 {object} utils.ResponseHTTP{}
// @Failure 503 {object} utils.ResponseHTTP{}
// @Router /auth/v1/api/create-session [post]
func Get(c *fiber.Ctx) error {

	localize, _ := fiberi18n.Localize(c, "Login.Success")

	auth := Authorization{
		Token:   "RA2PVJNv81H9KiWZvuMKGZ8v596JHiyl",
		Message: localize,
	}

	return c.JSON(auth)
}
