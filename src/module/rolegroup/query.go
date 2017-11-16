package rolegroup

const (
	updateQuery = `
		UPDATE
			rolegroups
		SET
			name = ('%s')
		WHERE
			id = (%d)
	`
)
