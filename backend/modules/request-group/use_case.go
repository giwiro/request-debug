package request_group

import (
	"context"
	"request-debug/modules/request-group/dao"
	"request-debug/modules/request-group/model"
)

type RequestGroupUseCase interface {
	GetRequestGroup(ctx context.Context, request GetRequestGroupRequest) (*model.RequestGroup, error)
}

type requestGroupUseCase struct {
	requestGroupDao dao.RequestGroupDao
}

func NewRequestGroupUseCase(requestGroupDao dao.RequestGroupDao) RequestGroupUseCase {
	return &requestGroupUseCase{
		requestGroupDao: requestGroupDao,
	}
}

func (r requestGroupUseCase) GetRequestGroup(ctx context.Context, request GetRequestGroupRequest) (*model.RequestGroup, error) {
	return nil, nil
}
