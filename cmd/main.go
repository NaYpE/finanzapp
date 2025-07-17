package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"finanzapp/config"
	"finanzapp/routes"
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

	// Rutas
	routes.RegisterWebRoutes(router)

	// Puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8282"
	}

	log.Printf("Servidor corriendo en http://localhost:%s", port)
	router.Run(":" + port)
}
