package middlewares

import (
	"gin-web-template/internal/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Authorization header 获取 token
		authHeader := c.GetHeader("Authorization")
		log.Println("authHeader: ", authHeader)
		if authHeader == "" {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		// 解析和验证 token
		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			// 确保使用的是正确的签名方法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Println("Invalid signing method")
				return nil, jwt.ErrSignatureInvalid
			}
			// 注意：实际使用时应该从配置文件读取密钥
			cfg, err := config.ServerConfig()
			if err != nil {
				log.Println("Failed to get server config: ", err)
				return nil, err
			}

			return []byte(cfg.ServerSecretKey), nil
		})

		if err != nil || !token.Valid {
			log.Println("Invalid token: ", err)
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		// 将 token 中的信息存储到上下文中
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			user_id := uint64(claims["user_id"].(float64))
			c.Set("user_id", user_id)
		}

		c.Next()
	}
}
