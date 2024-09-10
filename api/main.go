package main

import (
	"log"
	"swagger/app"
	_ "swagger/docs"

	"github.com/joho/godotenv"
)

// @title Payyer
// @version 1.0
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description This is an API for Payyer Application
// @BasePath /api
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := app.New()
	log.Fatal(app.Listen(":3003"))

}
