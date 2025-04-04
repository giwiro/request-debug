package web

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
	requestgroup "request-debug/modules/request-group"
	"request-debug/modules/request-group/database"
	"request-debug/utils"
)

type RequestGroupRouter interface {
	utils.Router
}

type requestGroupRouter struct {
	db *mongo.Client
}

func NewRequestGroupRouter(db *mongo.Client) RequestGroupRouter {
	return &requestGroupRouter{db: db}
}

func (rgr *requestGroupRouter) RegisterRoutes(router fiber.Router) {
	requestGroupDao := database.NewRequestGroupDao(rgr.db)
	rgc := &RequestGroupController{
		requestgroup.NewRequestGroupUseCase(requestGroupDao),
	}

	r := router.Group("/request")
	r.Post("/", rgc.CreateRequest)
	r.Post("/group", rgc.CreateRequestGroup)
	r.Use("/group/:request_group_id/request", rgc.CreateRequest)
	r.Get("/group/:request_group_id", rgc.GetRequestGroup)
}
