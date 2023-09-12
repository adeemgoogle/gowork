package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllWeatherData - ендпойнт для получения всех данных о погоде
func (h *Handler) GetAllWeatherData(ctx *gin.Context) {
	location := ctx.DefaultQuery("location", "Almaty")
	resp, err := h.weatherService.GetAllWeatherData(ctx, location)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
	return
}
