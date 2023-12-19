package routes

import (
	"net/url"

	"github.com/google/uuid"
	"github.com/utunia/api/tables"

	"github.com/gofiber/fiber/v2"
)

func SetupNationsRoute(store *tables.SupabaseStore, app *fiber.App) {
	app.Route("/nations", func(router fiber.Router) {
		router.Post("/", func(ctx *fiber.Ctx) error {
			var nation tables.Nation

			var body SignUpBody
			if err := ctx.BodyParser(body); err != nil {
				return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
			}

			nation.ID = uuid.New().String()
			nation.Name = body.Name
			nation.PassKey = body.Passkey
			nation.Flag = body.Flag
			nation.Motto = body.Motto
			nation.Currency = body.Currency
			nation.Classification = body.Classification
			nation.Ideology = body.Ideology

			EffectForStatsNumber := tables.StatsNumber(40)
			EffectForStatsPercentage := tables.StatsPercentage(2.0)

			switch body.Q1 {
			case true:
				nation.Stats.Power += EffectForStatsNumber
				nation.Stats.Army += EffectForStatsNumber
				nation.Stats.Security += EffectForStatsNumber
				nation.Stats.CivilRights -= EffectForStatsPercentage
				break
			case false:
				nation.Stats.Power -= EffectForStatsNumber
				nation.Stats.Army -= EffectForStatsNumber
				nation.Stats.Security -= EffectForStatsNumber
				nation.Stats.CivilRights += EffectForStatsPercentage
				break
			}

			switch body.Q2 {
			case true:
				nation.Stats.Sprituality -= EffectForStatsPercentage
				nation.Stats.SecularRate += EffectForStatsPercentage
				nation.Stats.CivilRights += EffectForStatsPercentage
				break
			case false:
				nation.Stats.Sprituality += EffectForStatsPercentage
				nation.Stats.SecularRate -= EffectForStatsPercentage
				nation.Stats.CivilRights -= EffectForStatsPercentage
				break
			}

			switch body.Q3 {
			case true:
				nation.Stats.Sprituality += EffectForStatsPercentage
				nation.Stats.SecularRate -= EffectForStatsPercentage
				nation.Stats.CivilRights -= EffectForStatsPercentage
				break
			case false:
				nation.Stats.Sprituality -= EffectForStatsPercentage
				nation.Stats.SecularRate += EffectForStatsPercentage
				nation.Stats.CivilRights += EffectForStatsPercentage
				break
			}

			switch body.Q4 {
			case true:
				nation.Stats.Sprituality -= EffectForStatsPercentage
				nation.Stats.CivilRights += EffectForStatsPercentage
				break
			case false:
				nation.Stats.Sprituality += EffectForStatsPercentage
				nation.Stats.CivilRights -= EffectForStatsPercentage
				nation.Stats.Corruption += EffectForStatsPercentage
				break
			}

			switch body.Q5 {
			case true:
				nation.Stats.CivilRights -= EffectForStatsPercentage
				nation.Stats.DeathRate += EffectForStatsPercentage
				break
			case false:
				nation.Stats.Education += EffectForStatsPercentage
				break
			}

			switch body.Q7 {
			case true:
				nation.Stats.TerrorismRate += EffectForStatsPercentage
				nation.Stats.Security -= EffectForStatsNumber
				nation.Stats.Population += EffectForStatsNumber
				break
			case false:
				nation.Stats.TerrorismRate -= EffectForStatsPercentage
				nation.Stats.Security += EffectForStatsNumber
				nation.Stats.Population -= EffectForStatsNumber
				break
			}

			switch body.Q8 {
			case true:
				nation.Stats.Power += EffectForStatsNumber
				nation.Stats.Army += EffectForStatsNumber
				nation.Stats.Security += EffectForStatsNumber
				nation.Stats.Corruption += EffectForStatsPercentage
				break
			case false:
				nation.Stats.Sprituality += EffectForStatsPercentage
				nation.Stats.Power -= EffectForStatsNumber
				nation.Stats.Army -= EffectForStatsNumber
				break
			}

			switch body.Q9 {
			case true:
				nation.Stats.Sprituality += EffectForStatsPercentage
				nation.Stats.CivilRights -= EffectForStatsPercentage
				nation.Stats.SecularRate -= EffectForStatsPercentage
				break
			case false:
				nation.Stats.Sprituality -= EffectForStatsPercentage
				nation.Stats.CivilRights += EffectForStatsPercentage
				nation.Stats.SecularRate += EffectForStatsPercentage
				break
			}

			switch body.Q10 {
			case true:
				nation.Stats.CivilRights += EffectForStatsPercentage
				nation.Stats.SecularRate -= EffectForStatsPercentage
				break
			case false:
				nation.Stats.CivilRights -= EffectForStatsPercentage
				nation.Stats.SecularRate += EffectForStatsPercentage
				nation.Stats.Sprituality -= EffectForStatsPercentage
				break
			}

			if err := store.InsertNation(nation); err != nil {
				return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
			}

			return ctx.SendStatus(fiber.StatusOK)
		})

		router.Get("/:name<string>", func(ctx *fiber.Ctx) error {
			name, _ := url.PathUnescape(ctx.Params("name"))

			var nation tables.Nation

			if err := store.GetNationByName(&nation, name); err != nil {
				return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
			}

			return ctx.JSON(nation)
		})

		router.Get("/:name<string>/world", func(ctx *fiber.Ctx) error {
			name, _ := url.PathUnescape(ctx.Params("name"))

			var nation tables.Nation

			if err := store.GetNationByName(&nation, name); err != nil {
				return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
			}

			var world tables.World

			if err := store.GetWorldByNationId(&world, nation.ID); err != nil {
				return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
			}

			return ctx.JSON(world)
		})

		router.Put("/:id<string>", func(ctx *fiber.Ctx) error {
			id := ctx.Params("id")

			var newNation tables.Nation

			if err := ctx.BodyParser(newNation); err != nil {
				return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
			}

			if err := store.UpdateNation(newNation, id); err != nil {
				return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
			}

			return ctx.SendStatus(fiber.StatusOK)
		})

		router.Delete("/:id<string>", func(ctx *fiber.Ctx) error {
			id := ctx.Params("id")

			if err := store.DeleteNation(id); err != nil {
				return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
			}

			var world tables.World

			if err := store.GetWorldByNationId(&world, id); err != nil {
				return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
			}

			world.NationCount -= 1
			for i, n := range world.Nations {
				if n == id {
					world.Nations = append(world.Nations[:i], world.Nations[i+1:]...)
				}
			}

			if err := store.UpdateWorld(world, world.ID); err != nil {
				return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
			}

			return ctx.SendStatus(fiber.StatusOK)
		})
	})
}
