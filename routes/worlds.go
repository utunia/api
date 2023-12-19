package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/utunia/api/tables"
)

func SetupWorldsRoute(store *tables.SupabaseStore, app *fiber.App) {
	app.Get("/worlds/:id<string>", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		var world tables.World

		if err := store.GetWorldById(&world, id); err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		return ctx.JSON(world)
	})
}
