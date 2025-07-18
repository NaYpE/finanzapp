package routes

import (
	"finanzapp/controllers"
	"finanzapp/middlewares"
	"finanzapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RegisterWebRoutes(router *gin.Engine) {
	// Rutas públicas
	router.GET("/", middlewares.AlreadyLoggedInMiddleware(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/login", controllers.Login)
	router.GET("/login", middlewares.AlreadyLoggedInMiddleware(), func(c *gin.Context) {
		css, msg := utils.GetFlash(c)
		c.HTML(http.StatusOK, "login.html", gin.H{
			"flash_css": css,
			"flash_msg": msg,
		})
	})

	router.POST("/signup", controllers.Register)
	router.GET("/signup", middlewares.AlreadyLoggedInMiddleware(), func(c *gin.Context) {
		css, msg := utils.GetFlash(c)
		c.HTML(http.StatusOK, "signup.html", gin.H{
			"flash_css": css,
			"flash_msg": msg,
		})
	})

	router.GET("/logout", controllers.Logout)

	// TODO controller contact
	router.GET("/contact", middlewares.AlreadyLoggedInMiddleware(), func(c *gin.Context) {
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
			// Mensaje flash
			css, msg := utils.GetFlash(c)

			c.HTML(http.StatusOK, "dashboard.html", gin.H{
				"Name":      name,
				"flash_css": css,
				"flash_msg": msg,
			})
		})

		protected.GET("/mysql", func(c *gin.Context) {
			c.HTML(http.StatusOK, "mysql.html", nil)
		})
	}

}
