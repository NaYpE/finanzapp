package validations

import (
	"regexp"
	"strings"
)

// Valida si el email tiene un formato válido
func IsValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}

// Valida si la contraseña tiene al menos 8 caracteres, una mayúscula, una minúscula y un número
func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	return hasUpper && hasLower && hasNumber
}

// Limpia espacios en blanco
func SanitizeInput(input string) string {
	return strings.TrimSpace(input)
}
