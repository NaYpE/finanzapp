package controllers

import (
	"finanzapp/config"
	"finanzapp/dto"
	"finanzapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input dto.RegisterInput
	if err := c.ShouldBind(&input); err != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"error": "Datos inválidos"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	user := models.User{Name: input.Name, Email: input.Email, Password: string(hashedPassword)}

	if err := config.DB.Create(&user).Error; err != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"error": "No se pudo crear el usuario"})
		return
	}

	c.Redirect(http.StatusFound, "/login")
}

func Login(c *gin.Context) {
	var input dto.LoginInput
	if err := c.ShouldBind(&input); err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Datos inválidos"})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Correo o contraseña inválidos"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Correo o contraseña inválidos"})
		return
	}

	// Aquí podrías guardar sesión o JWT. Por ahora, redireccionamos:
	c.Redirect(http.StatusFound, "/dashboard")
}
