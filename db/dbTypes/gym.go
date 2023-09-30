package dbtypes

import "go.mongodb.org/mongo-driver/bson/primitive"

type DbGymMembers struct {
	GymID           primitive.ObjectID `bson:"gym_id"`
	MemberUsernames []string           `bson:"members"`
}
