package course

const queryGetCourseByUserID = `
	SELECT
		id,
		name,
		ucu,
		semester,
		status
	FROM
		courses
	WHERE
		EXISTS (
			SELECT
				courses_id
			FROM
				p_users_courses
			WHERE
				users_id = (%d)
		);
`

const queryGetIDByUserID = `
	SELECT
		courses_id
	FROM
		p_users_courses
	WHERE
		users_id = (%d)
`
