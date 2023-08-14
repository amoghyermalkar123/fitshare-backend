package main

import (
	"fitshare/api"
	"fitshare/api/middlewares"
	"fitshare/db"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	db, err := db.GetDB("db")
	if err != nil {
		panic(err)
	}
	api := api.Api{
		DB: db,
	}

	r.POST("/login", api.Login)
	r.POST("/signup", api.Signup)

	r.POST("/protectedAPI", middlewares.Auth())
	godotenv.Load()
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
