package handler

import (
	"gin-web-template/internal/dao"
	"gin-web-template/internal/model"
	"gin-web-template/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloHandler godoc
// @Summary 普通问候
// @Description 获取普通问候语，不需登录
// @Tags 问候
// @Accept json
// @Produce json
// @Success 200 {object} handler.Response
// @Router /hello [get]
func HelloHandler(c *gin.Context) {
	data := service.HelloService()
	c.JSON(http.StatusOK, NewSuccessResponse(data))
}

// HelloUserHandler godoc
// @Summary 用户问候
// @Description 获取带用户名的问候语，需登录
// @Tags 问候
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} handler.Response
// @Router /hello-user [get]
func HelloUserHandler(c *gin.Context) {
	// 从上下文中获取用户ID
	userID := c.MustGet("user_id").(uint64)

	// 查询用户名
	username := GetUsernameById(userID)

	data := service.HelloUserService(username)
	c.JSON(http.StatusOK, NewSuccessResponse(data))
}

func GetUsernameById(userId uint64) string {
	// 实现通过用户ID查询用户名的逻辑
	db, err := dao.GetMysqlInstance()
	if err != nil {
		return "Unknown User"
	}

	var user model.User
	if err := db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		return "Unknown User"
	}

	return user.Username
}
