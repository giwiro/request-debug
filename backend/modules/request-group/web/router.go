package web

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"request-debug/config"
	requestgroup "request-debug/modules/request-group"
	"request-debug/modules/request-group/database"
	"request-debug/modules/sse"
	"request-debug/utils"
)

type RequestGroupRouter interface {
	utils.Router
}

type requestGroupRouter struct {
	db     *mongo.Client
	broker sse.Broker
}

func NewRequestGroupRouter(db *mongo.Client, broker sse.Broker) RequestGroupRouter {
	return &requestGroupRouter{db: db, broker: broker}
}

func (rgr *requestGroupRouter) RegisterRoutes(router fiber.Router) {
	requestGroupDao := database.NewRequestGroupDao(rgr.db)
	rgc := &RequestGroupController{
		requestgroup.NewRequestGroupUseCase(requestGroupDao, rgr.broker),
		rgr.broker,
	}

	r := router.Group("/group")

	r.Post("/", rgc.CreateRequestGroup)
	r.Get("/:request_group_id", rgc.GetRequestGroup)
	r.Delete("/:request_group_id/request/:request_id", rgc.DeleteRequest)

	// SSE
	r.Get("/:request_group_id/sse", rgc.HandleSSE)

	if config.Conf.Environment != "production" {
		r.Get("/:request_group_id/broker/", rgc.GetBrokerClients)
	}
}
