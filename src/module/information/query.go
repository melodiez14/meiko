package information

const (
	queryGet    = "SELECT %s FROM informations"
	querySelect = "SELECT %s FROM informations"
	queryInsert = "INSERT INTO informations (%s) VALUES (%s)"
	queryUpdate = "UPDATE files SET %s"

	querySelectAuthorized = `
		SELECT
			id,
			title,
			description,
			type,
			courses_id,
			created_at
		FROM
			informations
		WHERE
			type = (%d)
		OR
			courses_id IN (%s)
		ORDER BY created_at DESC
		LIMIT 100
	`
)
