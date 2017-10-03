package user

const (
	queryGet    = "SELECT %s FROM users"
	querySelect = "SELECT %s FROM users"
	queryInsert = "INSERT INTO users (%s) VALUES (%s)"
	queryUpdate = "UPDATE users SET %s"

	generateVerificationQuery = `
		UPDATE
			users
		SET
			email_verification_code = (%d),
			email_verification_expire_date = (DATE_ADD(NOW(), INTERVAL 30 MINUTE)),
			email_verification_attempt = 0,
			updated_at = NOW()
		WHERE
			identity_code = (%d)
	`

	getConfirmationQuery = `
		SELECT
			id,
			email_verification_attempt,
			email_verification_code
		FROM
			users
		WHERE
			email = ('%s') AND
			NOW() < email_verification_expire_date
	`

	attemptIncrementQuery = `
		UPDATE
			users
		SET
			email_verification_attempt = email_verification_attempt + 1,
			updated_at = NOW()
		WHERE
			id = (%d)
	`

	updateNewPasswordQuery = `
		UPDATE
			users
		SET
			password = md5('%s'),
			email_verification_code = NULL,
			email_verification_expire_date = NULL,
			email_verification_attempt = NULL,
			updated_at = NOW()
		WHERE
			email = ('%s')
	`
)
