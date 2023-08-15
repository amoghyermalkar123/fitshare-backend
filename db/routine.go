package db

import (
	"context"
	"fitshare/api/types"

	"go.mongodb.org/mongo-driver/bson"
)

func (db *DB) AddRoutine(userRoutine *types.UserRoutine) error {
	coll := db.client.Database("fitshare").Collection("user_routines")

	_, err := coll.InsertOne(context.TODO(), userRoutine)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetRoutine(userEmail string) (*types.UserRoutine, error) {
	coll := db.client.Database("fitshare").Collection("user_routines")
	userResponse := &types.UserRoutine{}

	err := coll.FindOne(context.TODO(), bson.M{"email": userEmail}).Decode(userResponse)
	if err != nil {
		return nil, err
	}
	return userResponse, nil
}
