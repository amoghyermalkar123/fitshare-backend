package db

import (
	"context"
	"fitshare/api/types"
	dbtypes "fitshare/db/dbTypes"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *DB) AddPeople(user *types.NewUserHomePageRequest) error {
	coll := db.client.Database("fitshare").Collection("gym_members")
	objID, err := primitive.ObjectIDFromHex(user.GymID)
	if err != nil {
		return err
	}

	_, err = coll.UpdateOne(context.TODO(), bson.M{"gym_id": objID}, bson.M{"$push": bson.M{"members": user.UserName}}, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) AddGymWeeklySchedule(schedule *dbtypes.GymWeeklySchedule) error {
	coll := db.client.Database("fitshare").Collection("gym_schedule")

	_, err := coll.InsertOne(context.TODO(), schedule)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) UpdateGymWeeklySchedule() {}
