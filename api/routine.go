package api

import (
	"fitshare/api/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *Api) CreateRoutine(c *gin.Context) {
	userRoutine := &types.UserRoutine{}
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
	userEmail := c.Param("user-email")

	routine, err := a.DB.GetRoutine(userEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": fmt.Errorf("failed db operation: %v", err).Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"routine": routine})
}
