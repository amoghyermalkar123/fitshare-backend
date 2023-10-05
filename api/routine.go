package api

import (
	"fitshare/api/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *Api) CreateRoutine(c *gin.Context) {
	userRoutine := &types.UserRoutineCreation{}
	if err := c.Bind(userRoutine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": fmt.Errorf("failed request operation: %v", err).Error()})
		return
	}

	if err := a.DB.AddRoutine(userRoutine); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": fmt.Errorf("failed db operation: %v", err).Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "added routine successfully"})
}

func (a *Api) GetRoutine(c *gin.Context) {
	username := c.Param("username")

	routine, err := a.DB.GetRoutine(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": fmt.Errorf("failed db operation: %v", err).Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"workouts": routine})
}
