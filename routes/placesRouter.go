package routes

import (
	"clickbus/services"

	"github.com/gofiber/fiber/v2"
)

func PlacesRouter(app *fiber.App) {
	app.Get("/api/places", services.GetAllPlaces)
	app.Post("/api/places", services.CreatePlace)
	app.Get("/api/places/:id", services.GetPlace)
	app.Put("/api/places/:id", services.UpdatePlace)
	app.Delete(("/api/places/:id"), services.DeletePlace)
}
