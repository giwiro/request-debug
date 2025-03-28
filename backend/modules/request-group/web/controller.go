package web

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"request-debug/logger"
	"request-debug/modules/error/exc"
	requestgroup "request-debug/modules/request-group"
	"request-debug/modules/request-group/model"
	"time"
)

type RequestGroupController struct {
	requestGroupUseCase requestgroup.RequestGroupUseCase
}

func (vc *RequestGroupController) CreateRequestGroup(c *fiber.Ctx) error {
	ctx := c.UserContext()

	requestGroup := &model.RequestGroup{
		Requests:  []model.Request{},
		CreateAt:  time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	rg, err := vc.requestGroupUseCase.CreateRequestGroup(
		ctx,
		requestGroup,
	)
	if err != nil {
		return exc.InternalError{
			Message: "Could not create request_group",
		}
	}

	return c.Status(fiber.StatusOK).JSON(rg)
}

func (vc *RequestGroupController) GetRequestGroup(c *fiber.Ctx) error {
	var webRequest GetRequestGroupRequestWebRequest
	ctx := c.UserContext()

	if err := c.ParamsParser(&webRequest); err != nil {
		logger.Logger.Err(err)
		return err
	}

	val := validator.New()
	if err := val.Struct(webRequest); err != nil {
		logger.Logger.Err(err)
		return err
	}

	rg, err := vc.requestGroupUseCase.GetRequestGroup(
		ctx,
		requestgroup.GetRequestGroupRequest{
			RequestGroupId: webRequest.RequestGroupId,
		})
	if err != nil {
		fmt.Println(err.Error())
		logger.Logger.Err(err)
		return exc.InternalError{
			Message: fmt.Sprintf("Could not get request_group %s", webRequest.RequestGroupId),
		}
	}
	if rg == nil {
		return exc.NotFoundError{
			Message: fmt.Sprintf("request_group %s not found", webRequest.RequestGroupId),
		}
	}

	return c.Status(fiber.StatusOK).JSON(rg)
}

func (vc *RequestGroupController) CreateRequest(c *fiber.Ctx) error {
	panic("Not implemented")
}
