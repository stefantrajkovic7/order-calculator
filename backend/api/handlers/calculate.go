package handlers

import (
	"backend/internal/calculator"
	"backend/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CalculateRequest struct {
	Order int `json:"order" binding:"required,gte=0"`
}

type CalculateResponse struct {
	Packs map[int]int `json:"packs"`
}

func CalculateHandler(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CalculateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if len(cfg.PackSizes) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No pack sizes available."})
			return
		}

		packs := calculator.CalculatePacks(req.Order, cfg.PackSizes)
		c.JSON(http.StatusOK, CalculateResponse{Packs: packs})
	}
}
