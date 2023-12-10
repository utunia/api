package main

import (
	"os"

	"github.com/utunia/api/routes"
	"github.com/utunia/api/tables"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := fiber.New()

	store := tables.NewSupabaseStore(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))

	routes.SetupRoutes(store, app)

	panic(app.Listen(":8080"))
}
