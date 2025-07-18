package routes

import (
	"finanzapp/controllers"
	"finanzapp/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RegisterWebRoutes(router *gin.Engine) {
	// Rutas públicas
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/login", controllers.Login)
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	router.POST("/signup", controllers.Register)
	router.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", nil)
	})

	router.GET("/logout", controllers.Logout)

	// TODO controller contact
	router.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", nil)
	})

	// Ruta protegida
	protected := router.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{

		protected.GET("/dashboard", func(c *gin.Context) {

			claimsInterface, exists := c.Get("claims")
			if !exists {
				c.Redirect(http.StatusFound, "/login")
				return
			}

			claims := claimsInterface.(jwt.MapClaims)

			name, ok := claims["name"].(string)
			if !ok {
				name = "Usuario" // Valor por defecto
			}

			c.HTML(http.StatusOK, "dashboard.html", gin.H{
				"Name": name,
			})
		})

		protected.GET("/mysql", func(c *gin.Context) {
			c.HTML(http.StatusOK, "mysql.html", nil)
		})
	}

}
