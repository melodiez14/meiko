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

	updateQuery = `
		UPDATE
			rolegroups
		SET
			name = ('%s')
		WHERE
			id = (%d)
	`
)
