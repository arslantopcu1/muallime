package app

import (
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"golang.org/x/text/language"
	"swagger/handlers/login"
	"time"
)

// New create an instance of Book app routes
func New() *fiber.App {
	app := fiber.New()

	app.Get("/metrics", monitor.New())
	app.Static("/static", "./public")

	app.Use(healthcheck.New())
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 30 * time.Second,
	}))

	app.Use(cors.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Use(
		fiberi18n.New(&fiberi18n.Config{
			RootPath:        "./i18n",
			AcceptLanguages: []language.Tag{language.English, language.Turkish},
			DefaultLanguage: language.English,
		}),
	)

	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${locals:requestid}  ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "UTC",
	}))

	// Middleware
	api := app.Group("/auth")
	v1 := api.Group("/v1")

	login.Router(v1.Group("/api"))

	return app
}
