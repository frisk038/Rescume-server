package handler

import (
	"context"
	"fmt"
	"net/http"
	"rescueme-server/business/models"

	"github.com/gin-gonic/gin"
)

type storer interface {
	InsertUserCoordinates(ctx context.Context, userID string, coordinate models.Coordinates) error
}

type Coordinates struct {
	Latitude  string `json:"lt" binding:"required"`
	Longitude string `json:"lg" binding:"required"`
}

func POSTCoordinates(s storer) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userid")
		if len(userID) == 0 {
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("user is required"))
			return
		}

		var cdt Coordinates
		err := c.ShouldBindJSON(&cdt)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		err = s.InsertUserCoordinates(c.Request.Context(), userID, models.Coordinates{
			Longitude: cdt.Longitude,
			Latitiude: cdt.Latitude,
		})
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusOK)
	}
}
