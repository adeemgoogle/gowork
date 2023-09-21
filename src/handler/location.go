package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetLocations - ендпойт для получения списка всех локации
func (h *Handler) GetLocations(ctx *gin.Context) {
	resp, err := h.locationService.GetLocations(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
	return
}
