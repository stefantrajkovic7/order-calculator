package handlers

import (
	"backend/api/types"
	"backend/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdatePackSizesRequest represents the input for updating pack sizes
// @Description Input data for updating pack sizes
// @Tags Packs
// @Accept json
// @Produce json
// @Param request body UpdatePackSizesRequest true "Pack sizes input"
// @Success 200 {object} types.SuccessResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /packs [post]
type UpdatePackSizesRequest struct {
	// Array of available pack sizes
	// Required: true
	// Example: [250,500,1000,2000,5000]
	PackSizes []int `json:"pack_sizes" binding:"required"`
}

// @Summary Update pack sizes
// @Description Update the available pack sizes configuration
// @Tags Packs
// @Accept json
// @Produce json
// @Param request body UpdatePackSizesRequest true "Pack sizes input"
// @Success 200 {object} types.SuccessResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /packs [post]
func UpdatePackSizesHandler(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req UpdatePackSizesRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
			return
		}

		if len(req.PackSizes) == 0 {
			c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "Pack sizes cannot be empty."})
			return
		}

		cfg.PackSizes = req.PackSizes
		c.JSON(http.StatusOK, types.SuccessResponse{Message: "Pack sizes updated successfully"})
	}
}
