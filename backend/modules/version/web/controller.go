package web

import (
	"github.com/gofiber/fiber/v2"
	"request-debug/modules/version"
	"request-debug/types"
)

type VersionController struct{}

func (vc *VersionController) GetVersion(c *fiber.Ctx) error {
	err := c.JSON(types.M{"version": version.Version})
	if err != nil {
		return err
	}

	return nil
}
