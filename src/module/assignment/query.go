package assignment

const queryGetIncompleteAssignment = `
	SELECT
		id,
		name,
		status,
		upload_date,
		due_date
	FROM
		assigments
	WHERE
		EXISTS (
			SELECT 
				id
			FROM
				grade_parameters
			WHERE
				EXISTS (
					SELECT
						courses_id
					FROM
						p_users_courses
					WHERE
						users_id = (%d)
				)
		) AND id NOT IN (
			SELECT
				assigments_id
			FROM
				p_users_assignments
			WHERE
				users_id = (%d)
		);
`
