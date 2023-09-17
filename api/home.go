package api

import (
	"fitshare/api/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *Api) NewUserHomePage(c *gin.Context) {
	user := &types.NewUserHomePageRequest{}

	if err := c.Bind(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": fmt.Errorf("failed request operation: %v", err).Error()})
		return
	}

	homePage, err := a.DB.LoadNewUserHomePage(user.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": fmt.Errorf("failed db operation: %v", err).Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"homepage": homePage})
}
