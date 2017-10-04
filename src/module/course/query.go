package course

const (
	queryGet    = "SELECT %s FROM courses"
	querySelect = "SELECT %s FROM courses"
	queryInsert = "INSERT INTO courses (%s) VALUES (%s)"
	queryUpdate = "UPDATE courses SET %s"
)

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
		(%s)
`
