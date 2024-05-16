package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nurettintopal/leaderboard"
	"nurettintopal/leaderboard/config"
	"nurettintopal/leaderboard/request"
	"strconv"
)

var rs = leaderboard.RedisSettings{
	Host:     config.Config("REDIS_ADDR"),
	Password: config.Config("REDIS_PASS"),
}

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Your request has been completed successfully.",
	})
}

func Show(c *fiber.Ctx) error {

	leaderboardId := c.Params("id")

	scores, _ := leaderboard.New(rs).Show(leaderboardId)

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Your request has been completed successfully.",
		"data":    scores,
	})
}

func Players(c *fiber.Ctx) error {

	leaderboardId := c.Params("id")

	start, _ := strconv.ParseInt(c.Query("start", "0"), 10, 64)
	end, _ := strconv.ParseInt(c.Query("end", "10"), 10, 64)

	scores, _ := leaderboard.New(rs).Players(leaderboardId, start, end)

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Your request has been completed successfully.",
		"data":    scores,
	})
}

func Player(c *fiber.Ctx) error {
	leaderboardId := c.Params("id")
	playerId := c.Params("playerId")

	scores, _ := leaderboard.New(rs).Player(leaderboardId, playerId)

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Your request has been completed successfully.",
		"data": &fiber.Map{
			"rank": scores,
		},
	})
}

func Create(c *fiber.Ctx) error {
	newItem := new(request.Player)
	if err := c.BodyParser(&newItem); err != nil {
		fmt.Println("validation error: ", err)
	}

	newPlayer := leaderboard.Player{
		ID:    c.Params("playerId"),
		Score: newItem.Score,
	}

	createdPlayer := leaderboard.New(rs).Create(c.Params("id"), newPlayer)

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "Your request has been completed successfully.",
		"data":    createdPlayer,
	})
}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")

	_ = leaderboard.New(rs).Delete(id)

	return c.Status(fiber.StatusNoContent).JSON(&fiber.Map{
		"status":  "success",
		"message": "Your request has been completed successfully.",
		"data":    nil,
	})
}
