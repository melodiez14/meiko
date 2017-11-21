package email

import "github.com/melodiez14/meiko/src/util/alias"

// SendEmailValidation is used for sending an email validation
func SendEmailValidation(name, email string, code uint16) {

	data := map[string]interface{}{
		"code": code,
	}

	NewRequest(email, "Email Validation").
		SetTemplate(alias.Dir["email"]+"/email_validation.html", data).
		Deliver()
}

// SendForgotPassword is used for sending an email validation
func SendForgotPassword(name, email string, code uint16) {

	data := map[string]interface{}{
		"code": code,
	}

	NewRequest(email, "Forgot Password").
		SetTemplate(alias.Dir["email"]+"/forgot_password.html", data).
		Deliver()
}

func SendAccountCreated(name, email string, code uint16) {
	data := map[string]interface{}{
		"code": code,
	}

	NewRequest(email, "Your account has been created").
		SetTemplate(alias.Dir["email"]+"/account_created.html", data).
		Deliver()
}
