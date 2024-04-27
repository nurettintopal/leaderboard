package router

import (
	"nurettintopal/leaderboard/handler"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	app.Get("/", handler.Home).Name("home")

	api := app.Group("/api")
	api.Get("/leaderboards/:id", handler.Show).Name("Show")
	api.Delete("/leaderboards/:id", handler.Delete).Name("Delete")
	api.Get("/leaderboards/:id/players", handler.Players).Name("Players")
	api.Get("/leaderboards/:id/players/:playerId", handler.Player).Name("Player")
	api.Post("/leaderboards/:id/players/:playerId", handler.Create).Name("Score")
}
