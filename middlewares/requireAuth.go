package middlewares

import (
	"ThaiLy/configs"
	"ThaiLy/graph/model"
	helper "ThaiLy/helpers"
	"strings"

	"github.com/gin-gonic/gin"
)

// Define a custom type for the context key
type ContextKey string

// Define a constant for your context key (optional, but makes it easier to use and prevents string literal typos)
const AccountKey ContextKey = "account"

func RequireAuth(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "" {
		if len(token) > 7 && strings.HasPrefix(token, "Bearer ") {
			token = token[7:] // Lấy token sau "Bearer "
			Claims, _ := helper.ParseJWT(token)
			if Claims != nil {
				var account = model.AccountDB{}
				result := configs.GetDB().Where("id = ?", Claims.ID).First(&account)
				if result.Error != nil {
					c.JSON(401, gin.H{
						"code":    "error",
						"massage": "Người dùng cố tình hack hệ thống mình à",
					})
					c.Abort()
					return
				}
				c.Set("account", account)
			}
		}
	}
	c.Next()
}
