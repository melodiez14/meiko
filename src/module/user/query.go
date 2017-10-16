package user

const (
	queryGet        = "SELECT %s FROM users"
	querySelect     = "SELECT %s FROM users"
	queryInsert     = "INSERT INTO users (%s) VALUES (%s)"
	queryUpdate     = "UPDATE users SET %s"
	querySelectByID = `
		SELECT
			%s
		FROM
			users
		WHERE
			id IN (%s);
	`
	queryGetByEmail = `
		SELECT
			%s
		FROM
			users
		WHERE
			email = ('%s')
		LIMIT 1;
	`
	queryGetByIdentityCode = `
		SELECT
			%s
		FROM
			users
		WHERE
			identity_code = (%d)
		LIMIT 1;
	`
	querySignIn = `
		SELECT
			id,
			name,
			email,
			gender,
			note,
			status,
			identity_code,
			line_id,
			phone,
			rolegroups_id
		FROM
			users
		WHERE
			email = ('%s') AND
			password = ('%s')
		LIMIT 1;
	`
	querySignUp = `
		INSERT INTO
			users (
				name,
				email,
				password,
				identity_code,
				created_at,
				updated_at
			) VALUES (
				('%s'),
				('%s'),
				('%s'),
				(%d),
				NOW(),
				NOW()
			);
	`
	queryUpdateToVerified = `
		UPDATE
			users
		SET
			status = (%d),
			email_verification_code = NULL,
			email_verification_expire_date = NULL,
			email_verification_attempt = NULL,
			updated_at = NOW()
		WHERE
			identity_code = (%d);
	`
	queryUpdateStatus = `
		UPDATE
			users
		SET
			status = (%d),
			updated_at = NOW()
		WHERE
			identity_code = (%d);
	`
	querySelectDashboard = `
		SELECT
			identity_code,
			name,
			email,
			status
		FROM
			users
		WHERE
			(status = (%d) OR status = (%d)) AND
			id != (%d)
		LIMIT %d
		OFFSET %d;
	`

	generateVerificationQuery = `
		UPDATE
			users
		SET
			email_verification_code = (%d),
			email_verification_expire_date = (DATE_ADD(NOW(), INTERVAL 30 MINUTE)),
			email_verification_attempt = NULL,
			updated_at = NOW()
		WHERE
			identity_code = (%d);
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
		LIMIT 1;
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

	queryForgotNewPassword = `
		UPDATE
			users
		SET
			password = ('%s'),
			email_verification_code = NULL,
			email_verification_expire_date = NULL,
			email_verification_attempt = NULL,
			updated_at = NOW()
		WHERE
			email = ('%s');
	`
)
