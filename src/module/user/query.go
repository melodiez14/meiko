package user

const getUserByIDQuery = `
	SELECT
		*
	FROM
		users
	WHERE
		id = (%d)
`

const getUserByEmailQuery = `
	SELECT
		*
	FROM
		users
	WHERE
		email = ('%s') AND
		password = (md5('%s'))
`
