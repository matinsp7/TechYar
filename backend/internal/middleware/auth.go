package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/romina-kh/NetEng-Project/backend/internal/dto"
	utility "github.com/romina-kh/NetEng-Project/backend/internal/utils"
)

func JWTAuth(jwtManager *utility.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("auth_token")
		if err != nil || token == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				dto.ErrorResponse{
					Error: "authentication cookie required",
				},
			)
			return
		}

		claims, err := jwtManager.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				dto.ErrorResponse{
					Error: "invalid or expired token",
				},
			)
			return
		}

		c.Set("user_id", claims.UserID)

		c.Next()
	}
}
