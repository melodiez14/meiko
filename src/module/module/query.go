package module

const (
	getQuery = `
		SELECT
			id,
			name
		FROM
			modules
		LIMIT %d
		OFFSET %d
	`

	getPrivilegeQuery = `
		SELECT
			modules.name,
			p_rolegroups_modules.ability
		FROM
			p_rolegroups_modules
		LEFT JOIN
			modules
		ON
			p_rolegroups_modules.modules_id = modules_id
		WHERE
			p_rolegroups_modules.rolegroups_id = (%d)
	`
)
