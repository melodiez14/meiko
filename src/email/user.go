package email

import (
	"fmt"
)

// SendEmailValidation is used for sending an email validation
func SendEmailValidation(name, email string, code uint16) {
	bg, err := encodeImage("files/var/www/meiko/static/bg.jpg")
	if err != nil {
		fmt.Println("Error load bg.jpg")
		return
	}

	logo, err := encodeImage("files/var/www/meiko/static/logo.png")
	if err != nil {
		fmt.Println("Error load logo.png")
		return
	}

	data := map[string]interface{}{
		"code":       code,
		"background": bg,
		"logo":       logo,
	}

	NewRequest(email, "Email Validation").
		SetTemplate("files/var/www/meiko/email/email_validation.html", data).
		Deliver()
}

// SendForgotPassword is used for sending an email validation
func SendForgotPassword(name, email string, code uint16) {
	bg, err := encodeImage("files/var/www/meiko/static/bg.jpg")
	if err != nil {
		fmt.Println("Error load bg.jpg")
		return
	}

	logo, err := encodeImage("files/var/www/meiko/static/logo.png")
	if err != nil {
		fmt.Println("Error load logo.png")
		return
	}

	data := map[string]interface{}{
		"code":       code,
		"background": bg,
		"logo":       logo,
	}

	NewRequest(email, "Email Validation").
		SetTemplate("files/var/www/meiko/email/forgot_password.html", data).
		Deliver()
}
