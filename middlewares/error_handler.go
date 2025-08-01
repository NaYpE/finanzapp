package middlewares

import (
	"login-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				// Aquí capturamos cualquier panic
				utils.SetFlash(c, "alert-danger", "Ocurrió un error inesperado. Intenta nuevamente.")
				c.Redirect(http.StatusFound, "/")
				c.Abort()
			}
		}()

		// Continuar con el siguiente middleware o handler
		c.Next()

		// También podemos manejar errores HTTP aquí
		status := c.Writer.Status()
		if status >= 500 {
			utils.SetFlash(c, "alert-danger", "Error interno del servidor.")
			c.Redirect(http.StatusFound, "/")
			c.Abort()
		}
	}
}
