package main

import (
	"fitshare/api"
	"fitshare/db"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	db, err := db.GetDB("db")
	if err != nil {
		panic(err)
	}
	api := api.Api{
		DB: db,
	}

	r.Use(CORSMiddleware())

	r.POST("/login", api.Login)
	r.POST("/signup", api.Signup)

	userApi := r.Group("/user")
	{
		userApi.POST("/routine", api.CreateRoutine)
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
