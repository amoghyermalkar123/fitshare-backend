package types

type RoutineCategory string

const (
	UpperRoutine  RoutineCategory = "upper"
	LowerRoutine  RoutineCategory = "lower"
	CoreRoutine   RoutineCategory = "core"
	CardioRoutine RoutineCategory = "cardio"
	OtherRoutine  RoutineCategory = "other"
)

type Set struct {
	Count int    `json:"count" bson:"count"`
	Reps  string `json:"reps" bson:"reps"`
}

type Exercise struct {
	Name string `json:"name" bson:"name"`
	Sets []Set  `json:"sets" bson:"sets"`
}

type Routine struct {
	Category  RoutineCategory `json:"category" bson:"category"`
	Exercises []Exercise      `json:"exercises" bson:"exercises"`
}

type UserRoutineCreation struct {
	Username string  `json:"username" bson:"username"`
	Routine  Routine `json:"routine" bson:"routine"`
}

type UserRoutineRequest struct {
	UserEmail string `json:"email"`
}
