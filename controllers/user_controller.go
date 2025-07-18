package controllers

import (
	"finanzapp/config"
	"finanzapp/dto"
	"finanzapp/models"
	"finanzapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input dto.RegisterInput
	if err := c.ShouldBind(&input); err != nil {
		utils.SetFlash(c, "alert-warning", "Datos inválidos")
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"error": "Datos inválidos"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	user := models.User{Name: input.Name, Email: input.Email, Password: string(hashedPassword)}

	if err := config.DB.Create(&user).Error; err != nil {
		utils.SetFlash(c, "alert-danger", "No se pudo crear el usuario")
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"error": "No se pudo crear el usuario"})
		return
	}

	utils.SetFlash(c, "alert-success", "Usuario creado exitosamente!")
	c.Redirect(http.StatusFound, "/login")
}

func Login(c *gin.Context) {
	var input dto.LoginInput
	if err := c.ShouldBind(&input); err != nil {
		utils.SetFlash(c, "alert-danger", "Datos inválidos")
		c.Redirect(http.StatusBadRequest, "/login")
		return
	}

	var user models.User
	// Buscar el usuario por correo
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		// Si no se encuentra el usuario, establecer mensaje flash y redirigir
		utils.SetFlash(c, "alert-warning", "Correo inválido")
		c.Redirect(http.StatusFound, "/login")
		return
	}

	// Verificar la contraseña
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		// Si la contraseña no coincide, establecer mensaje flash y redirigir
		utils.SetFlash(c, "alert-warning", "Contraseña inválida")
		c.Redirect(http.StatusFound, "/login")
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Name)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"error": "Error generando token"})
		return
	}

	// Guardar token en cookie segura
	c.SetCookie("Authorization", token, 3600*24, "/", "", false, true)
	utils.SetFlash(c, "alert-success", "Inicio de sesión exitoso")
	c.Redirect(http.StatusFound, "/dashboard")
}

func Logout(c *gin.Context) {
	// Eliminar la cookie del token
	c.SetCookie("Authorization", "", -1, "/", "", false, true)
	utils.SetFlash(c, "alert-success", "Sesión cerrada")
	c.Redirect(http.StatusFound, "/login")
}
