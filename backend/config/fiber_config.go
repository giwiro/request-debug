package config

import (
	"github.com/gofiber/fiber/v2"
	errorWeb "request-debug/modules/error/web"
)

func NewFiberConfiguration() fiber.Config {
	return fiber.Config{
		AppName:        Conf.App.Name,
		ErrorHandler:   errorWeb.ErrorHandler,
		ReadBufferSize: 256 * 1024,
		BodyLimit:      20 * 1024 * 1024,
	}
}
