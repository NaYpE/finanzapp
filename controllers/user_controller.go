package controllers

import (
	"finanzapp/config"
	"finanzapp/dto"
	"finanzapp/models"
	"finanzapp/utils"
	"finanzapp/validations"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input dto.RegisterInput
	if err := c.ShouldBind(&input); err != nil {
		utils.SetFlash(c, "alert-warning", "Datos inválidos")
		c.HTML(http.StatusSeeOther, "signup.html", gin.H{"error": "Datos inválidos"})
		return
	}

	if !validations.IsValidEmail(input.Email) {
		utils.SetFlash(c, "alert-warning", "Correo inválido")
		c.Redirect(http.StatusSeeOther, "/signup")
		return
	}

	var existingUser models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		utils.SetFlash(c, "alert-warning", "El correo ya está registrado")
		c.Redirect(http.StatusSeeOther, "/signup")
		return
	}

	if !validations.IsValidPassword(input.Password) {
		utils.SetFlash(c, "alert-warning", "Password inválido")
		c.Redirect(http.StatusSeeOther, "/signup")
		return
	}

	if input.Password != input.ConfirmPassword {
		utils.SetFlash(c, "alert-warning", "Las contraseñas no coinciden")
		c.Redirect(http.StatusSeeOther, "/signup")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	if err != nil {
		utils.SetFlash(c, "alert-danger", "Error al procesar la contraseña")
		c.Redirect(http.StatusSeeOther, "/signup")
		return
	}

	user := models.User{Name: input.Name, Email: input.Email, Password: string(hashedPassword)}

	if err := config.DB.Create(&user).Error; err != nil {
		utils.SetFlash(c, "alert-danger", "No se pudo crear el usuario")
		c.HTML(http.StatusSeeOther, "signup.html", gin.H{"error": "No se pudo crear el usuario"})
		return
	}

	utils.SetFlash(c, "alert-success", "Usuario creado exitosamente!")
	c.Redirect(http.StatusFound, "/login")
}

func Login(c *gin.Context) {
	var input dto.LoginInput
	if err := c.ShouldBind(&input); err != nil {
		utils.SetFlash(c, "alert-danger", "Datos inválidos")
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	var user models.User

	// Validacion de correo
	if !validations.IsValidEmail(input.Email) {
		utils.SetFlash(c, "alert-warning", "Correo inválido")
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
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
		utils.SetFlash(c, "alert-danger", "Error generando token")
		c.Redirect(http.StatusFound, "/login")
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
