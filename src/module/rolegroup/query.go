package rolegroup

const (
	getQuery = `
		SELECT
			id,
			name
		FROM
			rolegroups
		LIMIT %d
		OFFSET %d
	`

	insertQuery = `
		INSERT INTO
			rolegroups(
				name,
				created_at,
				updated_at
			)
		VALUES (
			('%s'),
			NOW(),
			NOW()
		)
	`

	updateQuery = `
		UPDATE
			rolegroups
		SET
			name = ('%s')
		WHERE
			id = (%d)
	`
)
