package handler

import (
	"gin-web-template/internal/errors"
	"gin-web-template/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// registerRequest defines the required fields for user registration
type registerRequest struct {
	Username string `json:"username" binding:"required" example:"john_doe"`
	Password string `json:"password" binding:"required" example:"password123"`
}

// RegisterHandler godoc
// @Summary 用户注册
// @Description 用户注册接口
// @Tags 用户
// @Accept json
// @Produce json
// @Param request body registerRequest true "注册请求"
// @Success 200 {object} handler.Response
// @Failure 400 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /user [post]
func RegisterHandler(c *gin.Context) {
	var req registerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(errors.InvalidRequestError))
		return
	}

	userId, err := service.RegisterService(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse(map[string]interface{}{
		"user_id":  userId,
		"username": req.Username,
	}))
}
