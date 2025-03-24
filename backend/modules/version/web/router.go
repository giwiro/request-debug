package web

import (
	"github.com/gofiber/fiber/v2"
	"request-debug/utils"
)

type VersionRouter interface {
	utils.Router
}

type versionRouter struct{}

func NewVersionRouter() VersionRouter {
	return &versionRouter{}
}

func (*versionRouter) RegisterRoutes(router fiber.Router) {
	vc := &VersionController{}

	r := router.Group("/version")
	r.Get("", vc.GetVersion)
}
