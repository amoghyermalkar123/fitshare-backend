package api

import (
	"fitshare/api/types"
	dbtypes "fitshare/db/dbTypes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *Api) AddPeople(c *gin.Context) {
	user := &types.NewUserHomePageRequest{}

	if err := c.Bind(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": fmt.Errorf("failed request operation: %v", err).Error()})
		return
	}

	err := a.DB.AddPeople(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": fmt.Errorf("failed db operation: %v", err).Error()})
		return
	}
}

func (a *Api) AddWeeklySchedule(c *gin.Context) {
	schedule := &dbtypes.GymWeeklySchedule{}

	if err := c.Bind(schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": fmt.Errorf("failed request operation: %v", err).Error()})
		return
	}

	err := a.DB.AddGymWeeklySchedule(schedule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": fmt.Errorf("failed db operation: %v", err).Error()})
		return
	}
}

// TODO:
func (a *Api) UpdateWeeklySchedule(c *gin.Context) {}
