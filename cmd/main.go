package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"finanzapp/config"
	"finanzapp/middlewares"
	"finanzapp/routes"
	"finanzapp/utils"
)

func main() {

	// Cargar .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar .env")
	}

	// Conexión a BD
	config.ConnectDatabase()
	//database.Connection()

	router := gin.Default()

	// Servir archivos estáticos (opcional)
	//router.Static("/static", "./static") //public en el otro

	// Cargar templates
	router.LoadHTMLGlob("templates/*.html")

	// Recupera panics de cualquier handler
	router.Use(middlewares.ErrorHandler())

	// Rutas
	routes.RegisterWebRoutes(router)

	// 404
	router.NoRoute(func(c *gin.Context) {
		token, err := c.Cookie("Authorization")
		loggedIn := false
		if err == nil {
			if _, err := utils.ParseJWT(token); err == nil {
				loggedIn = true
			}
		}
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"loggedIn": loggedIn,
		})
	})

	// Puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8282"
	}

	log.Printf("Servidor corriendo en http://localhost:%s", port)
	router.Run(":" + port)
}
