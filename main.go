package main

import (
	"go-uts-pasien-klinik/config"
	"go-uts-pasien-klinik/model"
	"go-uts-pasien-klinik/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	config.InitDB()

	config.DB.AutoMigrate(&model.Pasien{})

	app.Use(logger.New())

	router.SetupRoutes(app)

	app.Listen(":3000")
}