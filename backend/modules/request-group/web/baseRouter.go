package web

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
	requestgroup "request-debug/modules/request-group"
	"request-debug/modules/request-group/database"
	"request-debug/modules/sse"
	"request-debug/utils"
)

type BaseRequestGroupRouter interface {
	utils.Router
}

type baseRequestGroupRouter struct {
	db     *mongo.Client
	broker sse.Broker
}

func NewBaseRequestGroupRouter(db *mongo.Client, broker sse.Broker) BaseRequestGroupRouter {
	return &baseRequestGroupRouter{db: db, broker: broker}
}

func (rgr *baseRequestGroupRouter) RegisterRoutes(router fiber.Router) {
	requestGroupDao := database.NewRequestGroupDao(rgr.db)
	rgc := &RequestGroupController{
		requestgroup.NewRequestGroupUseCase(requestGroupDao, rgr.broker),
		rgr.broker,
	}

	r := router.Group("/group")
	r.Use("/:request_group_id/*", rgc.CreateRequest)
}
