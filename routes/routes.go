package routes

import (
	"finanzapp/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterWebRoutes(router *gin.Engine) {
	//paginas estaticas - vistas
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	router.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", nil)
	})

	router.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", nil)
	})

	router.GET("/mysql", func(c *gin.Context) {
		c.HTML(http.StatusOK, "mysql.html", nil)
	})

	router.GET("/dashboard", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", nil)
	})

	//endpoints - acciones
	router.POST("/signup", controllers.Register)
	router.POST("/login", controllers.Login)
}
