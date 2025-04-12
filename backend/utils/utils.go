package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"strings"
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

func BuildSSEMessage(eventType string, data string) string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("event: %s\n", eventType))
	sb.WriteString(fmt.Sprintf("data: %v\n\n", data))

	return sb.String()
}
