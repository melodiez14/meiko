package user

const (
	insertNewUserQuery = `
		INSERT INTO
			users(
				id,
				name,
				email,
				password,
				created_at,
				updated_at
			)
		VALUES (
			(%d),
			('%s'),
			('%s'),
			(md5('%s')),
			NOW(),
			NOW()
		)
	`

	getUserByIDQuery = `
		SELECT
			id,
			name,
			gender,
			college,
			note,
			rolegroups_id,
			status,
			phone,
			line_id
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
			email_verification_attempt = 0,
			updated_at = NOW()
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
			email_verification_attempt = email_verification_attempt + 1,
			updated_at = NOW()
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
			email_verification_attempt = NULL,
			updated_at = NOW()
		WHERE
			email = ('%s')
	`

	getUserByStatusQuery = `
		SELECT
			id,
			name,
			email
		FROM
			users
		WHERE
			status = (%d)
	`

	setStatusUserQuery = `
		UPDATE 
			users
		SET 
			status = (%d),
			updated_at = NOW()
		WHERE
			email = ('%s')
	`
	updateUserAccountQuery = `
		UPDATE
			users
		SET
			name = ('%s'),
			gender = (%d),
			phone = ('%s'),
			line_id = ('%s'),
			college = ('%s'),
			note = ('%s'), 
			Updated_at = NOW()
		WHERE
			id = (%d)
	`
	setChangePasswordQuery = `
		UPDATE
			users
		SET
			password = md5('%s'),
			updated_at = NOW()
		WHERE
			id = (%d)
`

	getUserByIDStatusQuery = `
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
			id = (%d) AND
			status = (%d)
		LIMIT 1
	`
)
