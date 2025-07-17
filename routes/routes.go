package routes

import (
	"finanzapp/controllers"
	"finanzapp/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
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

	// TODO controller contact
	router.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", nil)
	})

	// Ruta protegida
	protected := router.Group("/")
	protected.Use(middlewares.AuthMiddleware())

	protected.GET("/dashboard", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", nil)
	})
	protected.GET("/mysql", func(c *gin.Context) {
		c.HTML(http.StatusOK, "mysql.html", nil)
	})
}
