package handler

import (
	"gin-web-template/internal/errors"
	"gin-web-template/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// loginRequest defines the required fields for user login
type loginRequest struct {
	Username string `json:"username" example:"john_doe"`
	Password string `json:"password" example:"password123"`
}

// LoginHandler godoc
// @Summary 用户登录
// @Description 用户登录接口
// @Tags 认证
// @Accept json
// @Produce json
// @Param request body loginRequest true "登录请求"
// @Success 200 {object} handler.Response
// @Failure 400 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Router /sessions [post]
func LoginHandler(c *gin.Context) {
	var req loginRequest
	// Bind JSON request data
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(errors.ErrParseRequest))
		return
	}

	// Call the login service with provided credentials
	data, err := service.LoginService(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse(data))
}
