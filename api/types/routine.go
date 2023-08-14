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
	Count int    `json:"count"`
	Reps  string `json:"reps"`
}

type Exercise struct {
	Name string `json:"name"`
	Sets []Set  `json:"sets"`
}

type Routine struct {
	Name      string          `json:"name"`
	Category  RoutineCategory `json:"category"`
	Exercises []Exercise      `json:"exercises"`
}

type UserRoutine struct {
	Email   string  `json:"email"`
	Routine Routine `json:"routine"`
}
