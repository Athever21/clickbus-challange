package main

import (
	"clickbus/db"
	"clickbus/routes"
	"fmt"
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
	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
