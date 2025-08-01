package middlewares

import (
	"login-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("Authorization")
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		claims, err := utils.ParseJWT(token)
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		// Guardar el ID del usuario en el contexto
		c.Set("claims", claims)
		c.Next()
	}
}

func AlreadyLoggedInMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("Authorization")
		if err == nil {
			// Si el token existe y es válido, redirige al dashboard
			if _, err := utils.ParseJWT(token); err == nil {
				c.Redirect(http.StatusFound, "/dashboard")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
