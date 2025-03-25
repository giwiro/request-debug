package database

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"request-debug/modules/request-group/dao"
	"request-debug/modules/request-group/model"
)

type RequestGroupDao interface {
	dao.RequestGroupDao
}

type requestGroupDao struct {
	db *mongo.Client
}

func NewRequestGroupDao(db *mongo.Client) RequestGroupDao {
	return &requestGroupDao{db}
}

func (r requestGroupDao) GetGroup(ctx context.Context, groupId string) (*model.RequestGroup, error) {
	//TODO implement me
	panic("implement me")
}

func (r requestGroupDao) CreateRequest(ctx context.Context, groupId string, request *model.Request) (*model.Request, error) {
	//TODO implement me
	panic("implement me")
}
