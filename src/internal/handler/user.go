package handler

import (
	"gin-web-template/internal/errors"
	"gin-web-template/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// registerRequest defines the required fields for user registration
type registerRequest struct {
	Username string `json:"username" example:"john_doe"`
	Password string `json:"password" example:"password123"`
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
		c.JSON(http.StatusBadRequest, NewErrorResponse(errors.ErrInvalidCredentials))
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

// changePasswordRequest defines the required fields for changing password
type changePasswordRequest struct {
	UserID      uint64 `json:"user_id" example:"1"`
	OldPassword string `json:"old_password" example:"old_password123"`
	NewPassword string `json:"new_password" example:"new_password456"`
}

// ChangePasswordHandler godoc
// @Summary 修改密码
// @Description 用户修改密码接口
// @Tags 用户
// @Accept json
// @Produce json
// @Param request body changePasswordRequest true "修改密码请求"
// @Success 200 {object} handler.Response
// @Failure 400 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /user/password [post]
func ChangePasswordHandler(c *gin.Context) {
	var req changePasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(errors.ErrParseRequest))
		return
	}

	err := service.ChangePasswordService(req.UserID, req.OldPassword, req.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse(nil))
}
