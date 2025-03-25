package web

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"request-debug/logger"
	"request-debug/modules/error/exc"
	requestgroup "request-debug/modules/request-group"
)

type RequestGroupController struct {
	requestGroupUseCase requestgroup.RequestGroupUseCase
}

func (vc *RequestGroupController) GetRequestGroup(c *fiber.Ctx) error {
	var webRequest GetRequestGroupRequestWebRequest
	ctx := c.UserContext()

	if err := c.ParamsParser(&webRequest); err != nil {
		logger.Error(ctx, err.Error())
		return err
	}

	val := validator.New()
	if err := val.Struct(webRequest); err != nil {
		logger.Error(ctx, err.Error())
		return err
	}

	rg, err := vc.requestGroupUseCase.GetRequestGroup(
		ctx,
		requestgroup.GetRequestGroupRequest{
			RequestGroupId: webRequest.RequestGroupId,
		})
	if err != nil {
		return exc.InternalError{
			Message: fmt.Sprintf("Could not get request_group %s", webRequest.RequestGroupId),
		}
	}
	if rg == nil {
		return exc.NotFoundError{
			Message: fmt.Sprintf("request_group %s not found", webRequest.RequestGroupId),
		}
	}

	fmt.Println("3")

	return c.Status(fiber.StatusOK).JSON(rg)
}
