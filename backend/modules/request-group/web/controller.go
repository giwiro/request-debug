package web

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/url"
	"request-debug/config"
	"request-debug/logger"
	"request-debug/modules/error/exc"
	requestgroup "request-debug/modules/request-group"
	"request-debug/modules/request-group/model"
	"request-debug/modules/sse"
	"request-debug/utils"
	"time"
)

type RequestGroupController struct {
	requestGroupUseCase requestgroup.RequestGroupUseCase
	broker              sse.Broker
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
		logger.Logger.Err(err).Msg(fmt.Sprintf("Could not get request_group=%s", webRequest.RequestGroupId))
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

	val := validator.New()
	if err := val.Struct(webRequest); err != nil {
		logger.Logger.Err(err).Msg(fmt.Sprintf("Could not create request for request_group=%s", webRequest.RequestGroupId))
		return err
	}

	fullUrl, err := url.JoinPath(c.BaseURL(), c.OriginalURL())
	if err != nil {
		logger.Logger.Err(err).Msg(fmt.Sprintf("Could not build url for new request at request_group=%s", webRequest.RequestGroupId))
	}
	if fullUrl == "" {
		fullUrl = fmt.Sprintf("%s%s", c.BaseURL(), c.OriginalURL())
	}

	bodyRaw := c.BodyRaw()

	headers := map[string]string{}

	var headersCount uint = 0
	for k, v := range c.GetReqHeaders() {
		headersCount++
		headers[k] = v[0]

		if headersCount >= config.Conf.App.MaxHeaders {
			break
		}
	}

	queryParams := map[string]string{}

	var queryParamsCount uint = 0
	for k, v := range c.Queries() {
		queryParamsCount++
		queryParams[k] = v

		if queryParamsCount >= config.Conf.App.MaxQueries {
			break
		}
	}

	form := map[string][]string{}
	files := map[string][]model.RequestFile{}

	mp, err := c.MultipartForm()
	if err == nil {
		for k, v := range mp.Value {
			form[k] = v
		}

		for k, v := range mp.File {
			var fs []model.RequestFile

			for _, f := range v {
				fs = append(fs, model.RequestFile{
					Filename: f.Filename,
					Size:     f.Size,
				})
			}

			files[k] = fs
		}
	}

	request := &model.Request{
		Id:          uuid.New().String(),
		Method:      c.Method(),
		Host:        c.Hostname(),
		Url:         fullUrl,
		BodySize:    uint(len(bodyRaw)),
		BodyRaw:     string(bodyRaw),
		Date:        time.Now().UTC(),
		Ip:          c.IP(),
		Form:        form,
		Files:       files,
		QueryParams: c.Queries(),
		Headers:     headers,
	}

	rg, err := vc.requestGroupUseCase.CreateRequest(
		ctx,
		requestgroup.CreateRequestRequest{
			RequestGroupId: webRequest.RequestGroupId,
			Request:        request,
		})
	if err != nil {
		logger.Logger.Err(err).Msg(fmt.Sprintf("Could not create request for request_group=%s", webRequest.RequestGroupId))

		if errors.Is(err, bson.ErrInvalidHex) {
			return exc.ValidationError{
				Message: fmt.Sprintf("Could not create request for request_group=%s", webRequest.RequestGroupId),
			}
		}

		if errors.Is(err, mongo.ErrNoDocuments) {
			return exc.NotFoundError{
				Message: fmt.Sprintf("Could not create request for request_group=%s", webRequest.RequestGroupId),
			}
		}

		return exc.InternalError{
			Message: fmt.Sprintf("Could not create request for request_group=%s", webRequest.RequestGroupId),
		}
	}

	// return c.Status(fiber.StatusOK).JSON(rg)
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

	href := fmt.Sprintf("%s/dashboard/%s/%s", config.Conf.App.DashboardBaseUrl, rg.Id, request.Id)
	html := fmt.Sprintf("Request created! <a href=\"%s\">Check the request details at request-debug</a>", href)

	return c.Status(fiber.StatusOK).SendString(html)
}

func (vc *RequestGroupController) DeleteRequest(c *fiber.Ctx) error {
	var webRequest DeleteRequestWebRequest
	ctx := c.UserContext()

	if err := c.ParamsParser(&webRequest); err != nil {
		logger.Logger.Err(err).Msg(fmt.Sprintf("Could not delete request for request_group=%s", webRequest.RequestGroupId))
		return err
	}

	val := validator.New()
	if err := val.Struct(webRequest); err != nil {
		logger.Logger.Err(err).Msg(fmt.Sprintf("Could not delete request for request_group=%s", webRequest.RequestGroupId))
		return err
	}

	rg, err := vc.requestGroupUseCase.DeleteRequest(
		ctx,
		requestgroup.DeleteRequestRequest{
			RequestGroupId: webRequest.RequestGroupId,
			RequestId:      webRequest.RequestId,
		})
	if err != nil {
		logger.Logger.Err(err).Msg(fmt.Sprintf("Could not delete request for request_group=%s", webRequest.RequestGroupId))

		if errors.Is(err, bson.ErrInvalidHex) {
			return exc.ValidationError{
				Message: fmt.Sprintf("Could not delete request for request_group=%s", webRequest.RequestGroupId),
			}
		}

		if errors.Is(err, mongo.ErrNoDocuments) {
			return exc.NotFoundError{
				Message: fmt.Sprintf("Could not delete request for request_group=%s", webRequest.RequestGroupId),
			}
		}

		return exc.InternalError{
			Message: fmt.Sprintf("Could not delete request for request_group=%s", webRequest.RequestGroupId),
		}
	}

	return c.Status(fiber.StatusOK).JSON(rg)
}

func (vc *RequestGroupController) GetBrokerClients(c *fiber.Ctx) error {
	clients := vc.broker.GetClients()
	return c.Status(fiber.StatusOK).JSON(clients)
}

func (vc *RequestGroupController) HandleSSE(c *fiber.Ctx) error {
	var webRequest ConnectSseWebRequest

	if err := c.ParamsParser(&webRequest); err != nil {
		logger.Logger.Err(err).Msg(fmt.Sprintf("Could stablish SSE connection for request_group=%s", webRequest.RequestGroupId))
		return err
	}

	val := validator.New()
	if err := val.Struct(webRequest); err != nil {
		logger.Logger.Err(err).Msg(fmt.Sprintf("Could stablish SSE connection for request_group=%s", webRequest.RequestGroupId))
		return err
	}

	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")

	c.Status(fiber.StatusOK).Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		ch := vc.broker.AddNewClient(webRequest.RequestGroupId)

		keepAliveTickler := time.NewTicker(8 * time.Second)
		keepAliveMsg := ":keepalive\n"

		for {
			select {
			case <-keepAliveTickler.C:
				_, err := fmt.Fprintf(w, keepAliveMsg)
				if err != nil {
					logger.Logger.Err(err).Msg(fmt.Sprintf("Error writing SSE data for request_group=%s", webRequest.RequestGroupId))
					continue
				}

				err = w.Flush()
				if err != nil {
					logger.Logger.Err(err).Msg(fmt.Sprintf("Error writing SSE data for request_group=%s", webRequest.RequestGroupId))
					vc.broker.RemoveClient(webRequest.RequestGroupId, ch)
					keepAliveTickler.Stop()
					break
				}

			// Listen for incoming messages from messageChan
			case msg := <-ch:
				sseMessage := utils.BuildSSEMessage("sse-requests", string(msg))

				_, err := fmt.Fprintf(w, "%s", sseMessage)
				if err != nil {
					logger.Logger.Err(err).Msg(fmt.Sprintf("Error writing SSE data for request_group=%s", webRequest.RequestGroupId))
					continue
				}
				// Flush the data immediately instead of buffering it for later.
				err = w.Flush()
				if err != nil {
					logger.Logger.Err(err).Msg(fmt.Sprintf("Error flushing SSE data for request_group=%s", webRequest.RequestGroupId))
					vc.broker.RemoveClient(webRequest.RequestGroupId, ch)
					keepAliveTickler.Stop()
					break
				}
			default:
				logger.Logger.Info().Msgf("No handler")
			}
		}
	})

	return nil
}
