package handler

import (
	"github.com/adeemgoogle/gowork/src/model/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUser - ендпойт для создания профиля пользователя
func (h *Handler) CreateUser(ctx *gin.Context) {
	deviceId := ctx.GetHeader("device-id")
	if deviceId == "" {
		ctx.JSON(http.StatusBadRequest, "device-id is not specified in the header")
		return
	}

	var rqUser request.RqUser
	if err := ctx.ShouldBindJSON(&rqUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"request error": err.Error()})
		return
	}

	resp, err := h.userService.CreateUser(ctx, deviceId, rqUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
	return
}

// GetUser - ендпойт для получения профиля пользователя
func (h *Handler) GetUser(ctx *gin.Context) {
	deviceId := ctx.GetHeader("device-id")
	if deviceId == "" {
		ctx.JSON(http.StatusBadRequest, "device-id is not specified in the header")
		return
	}

	resp, err := h.userService.GetUser(ctx, deviceId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
	return
}
