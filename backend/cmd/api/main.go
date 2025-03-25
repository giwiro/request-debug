package main

import (
	"fmt"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"request-debug/config"
	"request-debug/database"
	"request-debug/logger"
	requestGroupWeb "request-debug/modules/request-group/web"
	versionWeb "request-debug/modules/version/web"
	"request-debug/utils"
)

func main() {
	// Config
	configPath := utils.GetEnv("REQUEST_DEBUG_CONFIG_PATH", "./config.yaml")
	err := config.ReadConfig(configPath)
	if err != nil {
		panic(err)
	}

	logger.Logger.Info().Msg("Using following environment variables: ")
	logger.Logger.Info().Msgf("  └── REQUEST_DEBUG_CONFIG_PATH=%s", configPath)

	mongoDB := database.NewMongoDB()

	app := setUpApp(config.NewFiberConfiguration())
	mainRouter := app.Group(config.Conf.Server.BasePath)

	versionRouter := versionWeb.NewVersionRouter()
	versionRouter.RegisterRoutes(mainRouter)

	requestGroupRouter := requestGroupWeb.NewRequestGroupRouter(mongoDB)
	requestGroupRouter.RegisterRoutes(mainRouter)

	addr := fmt.Sprintf("%s:%s", config.Conf.Server.Address, config.Conf.Server.Port)
	logger.Logger.Info().Msgf("Listening: %s%s", addr, config.Conf.Server.BasePath)
	err = app.Listen(addr)
	if err != nil {
		logger.Logger.Fatal().Msg(err.Error())
		return
	}
}

func setUpApp(appConfig fiber.Config) *fiber.App {
	fiberZerologConfig := fiberzerolog.Config{
		Logger: logger.Logger,
		Fields: []string{
			fiberzerolog.FieldLatency,
			fiberzerolog.FieldStatus,
			fiberzerolog.FieldMethod,
			fiberzerolog.FieldURL,
			fiberzerolog.FieldError,
		},
	}

	app := fiber.New(appConfig)
	app.Use(fiberzerolog.New(fiberZerologConfig))
	app.Use(recover.New())

	return app
}
