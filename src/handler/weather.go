package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// GetCurrentData -
func (h *Handler) GetCurrentData(ctx *gin.Context) {
	location := ctx.DefaultQuery("location", "Almaty")
	err := h.weatherService.GetCurrentData(ctx, h.config, location)
	if err != nil {
		fmt.Println("Errors: ", err.Error())
		return
	}
}
