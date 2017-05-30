package notification

const queryGet = `
	SELECT
		id,
		name,
		descriptions,
		read_at,
		table_id,
		table_name,
		created_at
	FROM
		notifications
	WHERE
		users_id = (%d)
	ORDER BY
		created_at DESC
	LIMIT %d, %d
`
