package handlers

import (
	"backend/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdatePackSizesRequest struct {
	PackSizes []int `json:"pack_sizes" binding:"required"`
}

func UpdatePackSizesHandler(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req UpdatePackSizesRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if len(req.PackSizes) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Pack sizes cannot be empty."})
			return
		}

		cfg.PackSizes = req.PackSizes
		c.JSON(http.StatusOK, gin.H{"message": "Pack sizes updated successfully"})
	}
}
