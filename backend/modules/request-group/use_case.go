package request_group

import (
	"context"
	"encoding/json"
	"request-debug/modules/request-group/dao"
	"request-debug/modules/request-group/model"
	"request-debug/modules/sse"
)

type RequestGroupUseCase interface {
	GetRequestGroup(ctx context.Context, request GetRequestGroupRequest) (*model.RequestGroup, error)
	CreateRequestGroup(ctx context.Context, requestGroup *model.RequestGroup) (*model.RequestGroup, error)
	CreateRequest(ctx context.Context, request CreateRequestRequest) (*model.RequestGroup, error)
	DeleteRequest(ctx context.Context, request DeleteRequestRequest) (*model.RequestGroup, error)
}

type requestGroupUseCase struct {
	requestGroupDao dao.RequestGroupDao
	broker          sse.Broker
}

func NewRequestGroupUseCase(requestGroupDao dao.RequestGroupDao, broker sse.Broker) RequestGroupUseCase {
	return &requestGroupUseCase{
		requestGroupDao: requestGroupDao,
		broker:          broker,
	}
}

func (r *requestGroupUseCase) GetRequestGroup(ctx context.Context, request GetRequestGroupRequest) (*model.RequestGroup, error) {
	return r.requestGroupDao.GetGroup(ctx, request.RequestGroupId)
}

func (r *requestGroupUseCase) CreateRequestGroup(ctx context.Context, requestGroup *model.RequestGroup) (*model.RequestGroup, error) {
	return r.requestGroupDao.CreateGroup(ctx, requestGroup)
}

func (r *requestGroupUseCase) CreateRequest(ctx context.Context, request CreateRequestRequest) (*model.RequestGroup, error) {
	group, err := r.requestGroupDao.CreateRequest(ctx, request.RequestGroupId, request.Request)
	if err != nil {
		return nil, err
	}

	j, err := json.Marshal(request.Request)
	if err != nil {
		return nil, err
	}

	r.broker.BroadcastGroup(group.Id, j)

	return group, nil
}

func (r *requestGroupUseCase) DeleteRequest(ctx context.Context, request DeleteRequestRequest) (*model.RequestGroup, error) {
	return r.requestGroupDao.DeleteRequest(ctx, request.RequestGroupId, request.RequestId)
}
