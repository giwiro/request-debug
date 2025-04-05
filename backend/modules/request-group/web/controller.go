package web

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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
	var webRequest GetRequestGroupWebRequest
	ctx := c.UserContext()

	if err := c.ParamsParser(&webRequest); err != nil {
		logger.Logger.Err(err).Msg(fmt.Sprintf("Could not get request_group=%s", webRequest.RequestGroupId))
		return err
	}

	val := validator.New()
	if err := val.Struct(webRequest); err != nil {
		logger.Logger.Err(err).Msg(fmt.Sprintf("Could not get request_group %s", webRequest.RequestGroupId))
		return err
	}

	rg, err := vc.requestGroupUseCase.GetRequestGroup(
		ctx,
		requestgroup.GetRequestGroupRequest{
			RequestGroupId: webRequest.RequestGroupId,
		})
	if err != nil {
		logger.Logger.Err(err).Msg(fmt.Sprintf("Could not get request_group=%s", webRequest.RequestGroupId))

		if errors.Is(err, bson.ErrInvalidHex) {
			return exc.ValidationError{
				Message: fmt.Sprintf("Could not get request_group=%s", webRequest.RequestGroupId),
			}
		}

		if errors.Is(err, mongo.ErrNoDocuments) {
			return exc.NotFoundError{
				Message: fmt.Sprintf("Could not get request_group=%s", webRequest.RequestGroupId),
			}
		}

		return exc.InternalError{
			Message: fmt.Sprintf("Could not get request_group=%s", webRequest.RequestGroupId),
		}
	}
	if rg == nil {
		return exc.NotFoundError{
			Message: fmt.Sprintf("request_group=%s not found", webRequest.RequestGroupId),
		}
	}

	return c.Status(fiber.StatusOK).JSON(rg)
}

func (vc *RequestGroupController) CreateRequest(c *fiber.Ctx) error {
	var webRequest CreateRequestWebRequest
	ctx := c.UserContext()

	if err := c.ParamsParser(&webRequest); err != nil {
		logger.Logger.Err(err).Msg(fmt.Sprintf("Could not create request for request_group=%s", webRequest.RequestGroupId))
		return err
	}

	request := &model.Request{
		Id:          uuid.New().String(),
		Method:      c.Method(),
		Host:        c.Hostname(),
		Date:        time.Now().UTC(),
		Ip:          c.IP(),
		QueryParams: c.Queries(),
	}

	rg, err := vc.requestGroupUseCase.CreateRequest(
		ctx,
		requestgroup.CreateRequestRequest{
			RequestGroupId: webRequest.RequestGroupId,
			Request:        request,
		})
	if err != nil {
		logger.Logger.Err(err).Msg(fmt.Sprintf("Could not get request_group Could not create request for request_group=%s", webRequest.RequestGroupId))

		if errors.Is(err, bson.ErrInvalidHex) {
			return exc.ValidationError{
				Message: fmt.Sprintf("Could not get request_group=%s", webRequest.RequestGroupId),
			}
		}

		if errors.Is(err, mongo.ErrNoDocuments) {
			return exc.NotFoundError{
				Message: fmt.Sprintf("Could not get request_group=%s", webRequest.RequestGroupId),
			}
		}

		return exc.InternalError{
			Message: fmt.Sprintf("Could not create request for request_group=%s", webRequest.RequestGroupId),
		}
	}

	return c.Status(fiber.StatusOK).JSON(rg)
}
