package user

const (
	getUserByIDQuery = `
		SELECT
			id,
			name,
			gender,
			college,
			note,
			rolegroups_id,
			status
		FROM
			users
		WHERE
			id = (%d)
	`

	getUserEmailQuery = `
		SELECT
			id,
			name,
			gender,
			college
		FROM
			users
		WHERE
			email = ('%s')
	`

	getUserLoginQuery = `
		SELECT
			id,
			name,
			gender,
			college,
			note,
			rolegroups_id,
			status
		FROM
			users
		WHERE
			email = ('%s') AND
			password = (md5('%s'))
	`

	generateVerificationQuery = `
		UPDATE
			users
		SET
			email_verification_code = (%d),
			email_verification_expire_date = (DATE_ADD(NOW(), INTERVAL 30 MINUTE)),
			email_verification_attempt = 0
		WHERE
			id = (%d)
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
			email_verification_attempt = email_verification_attempt + 1
		WHERE
			id = (%d)
	`

	setNewPasswordQuery = `
		UPDATE
			users
		SET
			password = md5('%s'),
			email_verification_code = NULL,
			email_verification_expire_date = NULL,
			email_verification_attempt = NULL
		WHERE
			email = ('%s')
	`
)
