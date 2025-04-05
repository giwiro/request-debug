package request_group

import (
	"context"
	"request-debug/modules/request-group/dao"
	"request-debug/modules/request-group/model"
)

type RequestGroupUseCase interface {
	GetRequestGroup(ctx context.Context, request GetRequestGroupRequest) (*model.RequestGroup, error)
	CreateRequestGroup(ctx context.Context, requestGroup *model.RequestGroup) (*model.RequestGroup, error)
	CreateRequest(ctx context.Context, request CreateRequestRequest) (*model.RequestGroup, error)
}

type requestGroupUseCase struct {
	requestGroupDao dao.RequestGroupDao
}

func NewRequestGroupUseCase(requestGroupDao dao.RequestGroupDao) RequestGroupUseCase {
	return &requestGroupUseCase{
		requestGroupDao: requestGroupDao,
	}
}

func (r *requestGroupUseCase) GetRequestGroup(ctx context.Context, request GetRequestGroupRequest) (*model.RequestGroup, error) {
	return r.requestGroupDao.GetGroup(ctx, request.RequestGroupId)
}

func (r *requestGroupUseCase) CreateRequestGroup(ctx context.Context, requestGroup *model.RequestGroup) (*model.RequestGroup, error) {
	return r.requestGroupDao.CreateGroup(ctx, requestGroup)
}

func (r *requestGroupUseCase) CreateRequest(ctx context.Context, request CreateRequestRequest) (*model.RequestGroup, error) {
	return r.requestGroupDao.CreateRequest(ctx, request.RequestGroupId, request.Request)
}
