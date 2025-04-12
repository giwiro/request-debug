package database

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"request-debug/config"
	"request-debug/database"
	"request-debug/modules/request-group/dao"
	"request-debug/modules/request-group/model"
	"time"
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

func (r *requestGroupDao) GetGroup(ctx context.Context, groupId string) (*model.RequestGroup, error) {
	var requestGroup model.RequestGroup

	coll := database.GetCollection(r.db, "requestGroups")

	objID, err := bson.ObjectIDFromHex(groupId)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", objID}}

	err = coll.FindOne(ctx, filter).Decode(&requestGroup)
	if err != nil {
		return nil, err
	}

	return &requestGroup, nil
}

func (r *requestGroupDao) CreateGroup(ctx context.Context, requestGroup *model.RequestGroup) (*model.RequestGroup, error) {
	coll := database.GetCollection(r.db, "requestGroups")

	result, err := coll.InsertOne(ctx, requestGroup)
	if err != nil {
		return nil, err
	}

	requestGroup.Id, err = database.GetStringId(result.InsertedID)
	if err != nil {
		return nil, err
	}

	return requestGroup, nil
}

func (r *requestGroupDao) CreateRequest(ctx context.Context, groupId string, request *model.Request) (*model.RequestGroup, error) {
	var requestGroup model.RequestGroup
	objID, err := bson.ObjectIDFromHex(groupId)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", objID}}
	update := bson.M{
		"$set": bson.M{"updatedAt": time.Now().UTC()},
		"$push": bson.M{
			"requests": bson.M{
				"$each":  []model.Request{*request},
				"$slice": -1 * int64(config.Conf.App.MaxRequests),
			},
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	coll := database.GetCollection(r.db, "requestGroups")

	err = coll.FindOneAndUpdate(ctx, filter, update, opts).Decode(&requestGroup)
	if err != nil {
		return nil, err
	}

	return &requestGroup, nil
}

func (r *requestGroupDao) DeleteRequest(ctx context.Context, groupId string, requestId string) (*model.RequestGroup, error) {
	var requestGroup model.RequestGroup
	objID, err := bson.ObjectIDFromHex(groupId)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", objID}}
	update := bson.M{
		"$set": bson.M{
			"updatedAt": time.Now().UTC(),
		},
		"$pull": bson.M{
			"requests": bson.M{"id": requestId},
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	coll := database.GetCollection(r.db, "requestGroups")

	err = coll.FindOneAndUpdate(ctx, filter, update, opts).Decode(&requestGroup)
	if err != nil {
		return nil, err
	}

	return &requestGroup, nil
}
