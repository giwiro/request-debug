package utils

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

type Router interface {
	RegisterRoutes(router fiber.Router)
}

func GetEnv(key, defaultValue string) (value string) {
	value = os.Getenv(key)
	if len(value) == 0 {
		value = defaultValue
	}
	return
}
