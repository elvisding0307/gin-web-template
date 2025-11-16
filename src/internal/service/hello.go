package service

import (
	"strings"
)

func HelloService() string {
	return "Hello there!"
}

func HelloUserService(username string) string {
	// 验证和处理用户名
	username = strings.TrimSpace(username)

	return "Hello " + username + "!"
}
