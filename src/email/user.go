package email

// SendEmailValidation is used for sending an email validation
func SendEmailValidation(name, email string, code uint16) {

	data := map[string]interface{}{
		"code": code,
	}

	NewRequest(email, "Email Validation").
		SetTemplate("files/var/www/meiko/email/email_validation.html", data).
		Deliver()
}

// SendForgotPassword is used for sending an email validation
func SendForgotPassword(name, email string, code uint16) {

	data := map[string]interface{}{
		"code": code,
	}

	NewRequest(email, "Forgot Password").
		SetTemplate("files/var/www/meiko/email/forgot_password.html", data).
		Deliver()
}

func SendAccountCreated(name, email string, code uint16) {
	data := map[string]interface{}{
		"code": code,
	}

	NewRequest(email, "Your account has been created").
		SetTemplate("files/var/www/meiko/email/account_created.html", data).
		Deliver()
}
