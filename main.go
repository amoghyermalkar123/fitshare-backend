package main

import (
	"fitshare/api"
	"fitshare/db"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	db, err := db.GetDB("127.0.0.1")
	if err != nil {
		panic(err)
	}
	api := api.Api{
		DB: db,
	}

	r.Use(cors.Default())

	r.POST("/login", api.Login)
	r.POST("/signup", api.Signup)

	userApi := r.Group("/user")
	{
		userApi.POST("/:username/routine", api.CreateRoutine)
		userApi.GET("/:username/routine", api.GetRoutine)
		userApi.GET("/:username/homepage", api.HomePage)
	}

	gymAdminApi := r.Group("/gym-admin")
	{
		// add people to gym
		gymAdminApi.POST("/people", api.AddPeople)
		// add event schedule for the week
		gymAdminApi.POST("/schedule", api.AddWeeklySchedule)
		// update event schedule for the week
		gymAdminApi.PUT("/schedule", api.UpdateWeeklySchedule)
	}

	godotenv.Load()
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
