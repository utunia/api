package main

import (
	"utunia_api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    routes.SetupRoutes(app)

    panic(app.Listen(":8080"))
}
