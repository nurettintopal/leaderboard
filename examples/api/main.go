package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"nurettintopal/leaderboard/config"
	"nurettintopal/leaderboard/router"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func init() {
	log.Info("Current time in UTC:", time.Now().UTC().Format(time.DateTime))
}

func main() {
	run()
}

func run() {
	app := fiber.New(fiber.Config{
		ServerHeader: config.Config("APP_NAME"),
		AppName:      config.Config("APP_NAME") + " - " + config.Config("APP_VERSION"),
	})

	app.Use(
		logger.New(logger.Config{
			TimeFormat: "2006-01-02 15:04:05",
			TimeZone:   "UTC",
			Format:     "${time} | ${locals:requestid} | ${latency} | ${status} | ${method} | ${path}\n",
		}),
		cors.New(),
		recover.New(),
		helmet.New(),
		requestid.New(),
	)

	router.InitRoutes(app)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	var serverShutdown sync.WaitGroup
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		serverShutdown.Add(1)
		defer serverShutdown.Done()
		_ = app.ShutdownWithTimeout(60 * time.Second)
	}()

	if err := app.Listen(":" + config.Config("APP_PORT")); err != nil {
		log.Panic(err)
	}
	serverShutdown.Wait()
}
