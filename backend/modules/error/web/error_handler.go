package web

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"reflect"
	"request-debug/modules/error/exc"
	"request-debug/modules/error/model"
	"time"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	httpError := model.ApiError{
		Error:     getType(err),
		Message:   err.Error(),
		Timestamp: time.Now().UTC(),
	}

	var fiberError *fiber.Error
	if errors.As(err, &fiberError) {
		httpError.Status = fiberError.Code

		if fiberError.Code != fiber.StatusInternalServerError {
			return ctx.Status(httpError.Status).JSON(httpError)
		}
	}

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		var subErrors []model.ApiSubError

		for _, subError := range validationErrors {
			subErrors = append(subErrors, model.ApiSubError{
				Error:   getType(subError),
				Field:   subError.Field(),
				Message: subError.Error(),
			})
		}
		httpError.Message = ""
		httpError.SubErrors = subErrors
		return ctx.Status(fiber.StatusBadRequest).JSON(httpError)
	}

	var validationError exc.ValidationError
	if errors.As(err, &validationError) {
		httpError.Status = fiber.StatusBadRequest
		return ctx.Status(fiber.StatusBadRequest).JSON(httpError)
	}

	var notFoundError exc.NotFoundError
	if errors.As(err, &notFoundError) {
		httpError.Status = fiber.StatusBadRequest
		return ctx.Status(fiber.StatusBadRequest).JSON(httpError)
	}

	var unauthorizedError exc.UnauthorizedError
	if errors.As(err, &unauthorizedError) {
		httpError.Status = fiber.StatusUnauthorized
		return ctx.Status(fiber.StatusUnauthorized).JSON(httpError)
	}

	var forbiddenError exc.ForbiddenError
	if errors.As(err, &forbiddenError) {
		httpError.Status = fiber.StatusForbidden
		return ctx.Status(fiber.StatusForbidden).JSON(httpError)
	}

	httpError.Status = fiber.StatusInternalServerError
	return ctx.Status(fiber.StatusInternalServerError).JSON(httpError)
}

func getType(v interface{}) string {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
