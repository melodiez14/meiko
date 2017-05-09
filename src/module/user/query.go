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

	getUserByEmailQuery = `
		SELECT
			*
		FROM
			users
		WHERE
			email = ('%s') AND
			password = (md5('%s'))
	`

	insertUser = `
		INSERT INTO
			users (
				name,
				email,
				password,
				gender,
				college,
				note,
				rolegroups_id,
				status
			)
			VALUES (
				(%s),
				(%s),
				(%s),
				(%s),
				(%s),
				(%d),
				(%d)
			)
	`
)
