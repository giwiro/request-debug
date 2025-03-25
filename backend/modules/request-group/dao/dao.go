package dao

import (
	"context"
	"request-debug/modules/request-group/model"
)

type RequestGroupDao interface {
	GetGroup(ctx context.Context, groupId string) (*model.RequestGroup, error)
	CreateRequest(ctx context.Context, groupId string, request *model.Request) (*model.Request, error)
}
