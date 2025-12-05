package router

import (
	"go-uts-pasien-klinik/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api");

	api.Get("/patients", handler.GetAllPasien)
	api.Get("/patients/:id", handler.GetPasienById)

	api.Post("/patients", handler.InsertPasien)

	api.Put("/patients/:id", handler.ReplacePasienById)

	api.Patch("/patients/:id", handler.UpdatePasienById)

	api.Delete("/patients/:id", handler.DeletePasienById)
}