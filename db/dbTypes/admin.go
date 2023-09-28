package dbtypes

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GymWeeklySchedule struct {
	ID       primitive.ObjectID `bson:"id" json:"id"`
	GymID    primitive.ObjectID `bson:"gym_id" json:"gym_id"`
	Schedule []Schedule         `bson:"schedule" json:"schedule"`
}

type Schedule struct {
	DateTime  time.Time `bson:"date_time" json:"date_time"`
	EventName string    `bson:"event_name" json:"event_name"`
}
