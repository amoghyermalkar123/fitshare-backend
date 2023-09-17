package types

import "time"

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
	UserName string
}
