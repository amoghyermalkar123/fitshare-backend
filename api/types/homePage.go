package types

import (
	"time"
)

type NewUserHomePage struct {
	DiscoverPeople []People
	TodaysEvent    []GymEvent
	TodaysExercise string
}

type People struct {
	Name             string
	ProfileImageLink string
}

type GymEvent struct {
	Name string
	Time time.Time
}

type NewUserHomePageRequest struct {
	GymID    string `json:"gym_id"`
	UserName string `json:"username"`
}

type DiscoverPeople struct {
	Username         string `json:"username"`
	ProfileImageLink string `json:"profile_image_link"`
}

type TodaysEvent struct {
	EventName string    `json:"event_name"`
	EventTime time.Time `json:"event_time"`
}

type TodaysExercise struct {
	ExerciseName string `json:"exercise_name"`
}

type UserHomePage struct {
	UserName       string           `json:"user_name"`
	DiscoverPeople []DiscoverPeople `json:"discover_people"`
	TodaysEvent    []TodaysEvent    `json:"todays_event"`
	TodaysExercise []TodaysExercise `json:"todays_exercise"`
}
