package routes

import (
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/utunia/api/tables"

	"github.com/gofiber/fiber/v2"
)

type SignUpBody struct {
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

type SignInBody struct {
	Name    string `json:"nation_name"`
	Passkey string `json:"nation_passkey"`
}

func SetupAuthRoute(store *tables.SupabaseStore, app *fiber.App) {
	app.Route("/auth", func(router fiber.Router) {
		router.Post("/signup", func(ctx *fiber.Ctx) error {
			var body SignUpBody
			if err := ctx.BodyParser(body); err != nil {
				return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
			}

			hashedPassword, err := bcrypt.GenerateFromPassword(
				[]byte(body.Passkey),
				bcrypt.DefaultCost,
			)
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
			}

			body.Passkey = string(hashedPassword)

			agent := fiber.Post(ctx.BaseURL() + "/nations").JSON(body)
			if statusCode, body, errs := agent.Bytes(); len(errs) > 0 {
				return ctx.Status(statusCode).Send(body)
			}

			return ctx.Status(fiber.StatusPermanentRedirect).
				Redirect(os.Getenv("FRONT_END_URL") + "/signin")
		})

		router.Post("/signin", func(ctx *fiber.Ctx) error {
			var body SignInBody
			if err := ctx.BodyParser(body); err != nil {
				return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
			}

			var nation tables.Nation
			if err := store.GetNationByName(&nation, body.Name); err != nil {
				return ctx.Status(fiber.StatusBadRequest).SendString("Invalid nation name")
			}

			if err := bcrypt.CompareHashAndPassword([]byte(nation.PassKey), []byte(body.Passkey)); err != nil {
				return ctx.Status(fiber.StatusBadRequest).SendString("Invalid passkey")
			}

			tokenByte := jwt.New(jwt.SigningMethodHS256)

			now := time.Now().UTC()

			claims := tokenByte.Claims.(jwt.MapClaims)

			claims["sub"] = nation.ID
			claims["exp"] = now.Add(time.Hour * 24 * 7).Unix()
			claims["iat"] = now.Unix()
			claims["nbf"] = now.Unix()

			tokenString, err := tokenByte.SignedString([]byte(os.Getenv("JWT_SECRET")))
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
			}

			ctx.Cookie(&fiber.Cookie{
				Name:     "token",
				Value:    tokenString,
				Domain:   os.Getenv("FRONT_END_DOMAIN"),
				Path:     "/",
				MaxAge:   int(time.Hour * 24 * 7),
				Secure:   false,
				HTTPOnly: true,
			})

			return ctx.Status(fiber.StatusPermanentRedirect).
				Redirect(os.Getenv("FRONT_END_URL") + "/nations/" + nation.Name)
		})
	})
}
