package main

import (
	"fmt"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"request-debug/config"
	"request-debug/database"
	"request-debug/logger"
	requestGroupWeb "request-debug/modules/request-group/web"
	"request-debug/modules/sse"
	versionWeb "request-debug/modules/version/web"
	"request-debug/utils"
	"strings"
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
	sseBroker := sse.NewBroker()

	app := setUpApp(config.NewFiberConfiguration())
	baseRouter := app.Group("/")
	mainRouter := app.Group(config.Conf.Server.BasePath)

	// CORS
	if len(config.Conf.Cors.Url) > 0 {
		allowedOrigins := strings.Join(config.Conf.Cors.Url, ",")
		logger.Logger.Info().Msgf("CORS: %s", allowedOrigins)

		mainRouter.Use(cors.New(cors.Config{
			AllowOrigins: allowedOrigins,
			AllowHeaders: "Cache-Control",
		}))
	}

	versionRouter := versionWeb.NewVersionRouter()
	versionRouter.RegisterRoutes(mainRouter)

	baseRequestGroupRouter := requestGroupWeb.NewBaseRequestGroupRouter(mongoDB, sseBroker)
	baseRequestGroupRouter.RegisterRoutes(baseRouter)

	requestGroupRouter := requestGroupWeb.NewRequestGroupRouter(mongoDB, sseBroker)
	requestGroupRouter.RegisterRoutes(mainRouter)

	go sseBroker.Handler()

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
