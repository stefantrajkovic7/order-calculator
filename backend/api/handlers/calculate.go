package handlers

import (
	"backend/api/types"
	"backend/internal/calculator"
	"backend/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CalculateRequest represents the input for calculation
// @Description Input data for calculation
// @Tags Calculate
// @Accept json
// @Produce json
// @Param request body CalculateRequest true "Order input"
// @Success 200 {object} CalculateResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /calculate [post]
type CalculateRequest struct {
	// Number of items to be packed
	// Required: true
	// minimum: 0
	Order int `json:"order" binding:"required,gte=0" example:"123"`
}

// CalculateResponse represents the response from calculation
// @Description Map of pack sizes to their quantities
// @Tags Calculate
// @Accept json
// @Produce json
// @Success 200 {object} CalculateResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /calculate [post]
type CalculateResponse struct {
	// Map of pack sizes to their quantities
	// Example: {"250":1,"500":2}
	Packs map[int]int `json:"packs"`
}

// @Summary Calculate optimal pack sizes
// @Description Calculate the optimal combination of pack sizes for a given order
// @Tags Calculate
// @Accept json
// @Produce json
// @Param request body CalculateRequest true "Order input"
// @Success 200 {object} CalculateResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /calculate [post]
func CalculateHandler(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CalculateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
			return
		}

		if len(cfg.PackSizes) == 0 {
			c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "No pack sizes available."})
			return
		}

		packs := calculator.CalculatePacks(req.Order, cfg.PackSizes)
		c.JSON(http.StatusOK, CalculateResponse{Packs: packs})
	}
}
