package api

import (
	"fitshare/api/types"
	"fitshare/auth"
	"fitshare/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Api struct {
	DB *db.DB
}

func (a *Api) Signup(c *gin.Context) {
	userDetails := &types.UserDetails{}
	if err := c.Bind(userDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": fmt.Errorf("failed request operation: %v", err).Error()})
		return
	}

	profileEmail := userDetails.Email
	if err := a.DB.AddUserDetails(userDetails); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": fmt.Errorf("failed db operation: %v", err).Error()})
		return
	}

	tokenString, err := auth.GenerateJWT(profileEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (a *Api) Login(c *gin.Context) {
	userDetails := &types.UserLogin{}
	if err := c.Bind(userDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": fmt.Errorf("failed request operation: %v", err).Error()})
		return
	}

	profileEmail := userDetails.Email
	userInfo, err := a.DB.GetUserDetails(userDetails)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "failed", "error": fmt.Errorf("failed db operation: %v", err).Error()})
		return
	}

	tokenString, err := auth.GenerateJWT(profileEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_details": userInfo, "token": tokenString})
}
