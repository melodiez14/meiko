package course

// query constant to insert update select data from database
const (
	queryGet    = "SELECT %s FROM courses"
	querySelect = "SELECT %s FROM courses"
	queryInsert = "INSERT INTO courses (%s) VALUES (%s)"
	queryUpdate = "UPDATE courses SET %s"
)

// constanta query to show user course in database
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

// constanta query to get course by by user id in coure in database
const queryGetIDByUserID = `
	SELECT
		courses_id
	FROM
		p_users_courses
	WHERE
		users_id = (%d)
		(%s)
`
