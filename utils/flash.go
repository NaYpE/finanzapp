package utils

import "github.com/gin-gonic/gin"

// SetFlash establece dos cookies: clase CSS y mensaje
func SetFlash(c *gin.Context, cssClass, message string) {
	c.SetCookie("flash_css", cssClass, 1, "/", "", false, false)
	c.SetCookie("flash_msg", message, 1, "/", "", false, false)
}

// GetFlash obtiene y elimina las cookies flash
func GetFlash(c *gin.Context) (string, string) {
	css, err1 := c.Cookie("flash_css")
	msg, err2 := c.Cookie("flash_msg")

	if err1 == nil {
		c.SetCookie("flash_css", "", -1, "/", "", false, false)
	}
	if err2 == nil {
		c.SetCookie("flash_msg", "", -1, "/", "", false, false)
	}

	if err1 != nil || err2 != nil {
		return "", ""
	}
	return css, msg
}
