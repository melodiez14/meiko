package attendance

const queryGetByUserCourse = `
	SELECT
		id,
		meeting_number,
		status,
		meeting_date
	FROM
		attendances
	WHERE
		p_users_courses_users_id = (%d) AND
		p_users_courses_courses_id = (%d);
`
