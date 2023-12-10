package routes

import (
	"net/url"

	"github.com/google/uuid"
	"github.com/utunia/api/tables"

	"github.com/gofiber/fiber/v2"
)

func SetupNationsRoute(store *tables.SupabaseStore, app *fiber.App) {
	app.Post("/nations", func(ctx *fiber.Ctx) error {
		var nation tables.Nation

		var body struct {
			Name           string `json:"nation_name"`
			Passkey        string `json:"nation_passkey"`
			Flag           string `json:"nation_flag"`
			Motto          string `json:"nation_motto"`
			Currency       string `json:"nation_currency"`
			Classification string `json:"nation_classification"`
			Ideology       string `json:"nation_ideology"`
			Q1             bool   `json:"q1"`
			Q2             bool   `json:"q2"`
			Q3             bool   `json:"q3"`
			Q4             bool   `json:"q4"`
			Q5             bool   `json:"q5"`
			Q6             bool   `json:"q6"`
			Q7             bool   `json:"q7"`
			Q8             bool   `json:"q8"`
			Q9             bool   `json:"q9"`
			Q10            bool   `json:"q10"`
		}

		if err := ctx.BodyParser(body); err != nil {
			return ctx.Status(400).SendString(err.Error())
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
			return ctx.Status(400).SendString(err.Error())
		}

		return ctx.SendStatus(200)
	})

	app.Get("/nations/:name<string>", func(ctx *fiber.Ctx) error {
		name, _ := url.PathUnescape(ctx.Params("name"))

		var nation tables.Nation

		if err := store.GetNationByName(&nation, name); err != nil {
			return ctx.Status(400).SendString(err.Error())
		}

		return ctx.JSON(nation)
	})

	app.Put("/nations/:id<string>", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		var newNation tables.Nation

		if err := ctx.BodyParser(newNation); err != nil {
			return ctx.Status(400).SendString(err.Error())
		}

		if err := store.UpdateNation(newNation, id); err != nil {
			return ctx.Status(400).SendString(err.Error())
		}

		return ctx.SendStatus(200)
	})

	app.Delete("/nations/:id<string>", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		if err := store.DeleteNation(id); err != nil {
			return ctx.Status(400).SendString(err.Error())
		}

		return ctx.SendStatus(200)
	})
}
