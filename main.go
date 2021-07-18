package main

import (
	"clickbus/db"
	"clickbus/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := fiber.New()

	routes.PlacesRouter(app)

	defer db.CloseDb()
	log.Fatal(app.Listen(os.Getenv("PORT")))
}
