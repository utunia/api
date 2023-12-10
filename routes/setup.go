package routes

import (
	"github.com/utunia/api/tables"

	"github.com/gofiber/fiber/v2"
)


func SetupRoutes(store *tables.SupabaseStore,app *fiber.App) {
    SetupNationsRoute(store, app)
}
