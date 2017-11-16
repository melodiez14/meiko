package rolegroup

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/melodiez14/meiko/src/util/conn"
)

func GetByPage(page, offset uint16) ([]RoleGroup, error) {
	rolegroups := []RoleGroup{}
	query := fmt.Sprintf(getQuery, page, offset)
	err := conn.DB.Select(&rolegroups, query)
	if err != nil {
		return nil, err
	}

	return rolegroups, nil
}

func Update(id int64, name string) error {
	query := fmt.Sprintf(updateQuery, name, id)
	_, err := conn.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func GetModuleList() []string {
	return []string{
		ModuleUser,
		ModuleCourse,
		ModuleRole,
		ModuleAttendance,
		ModuleSchedule,
		ModuleAssignment,
		ModuleInformation,
	}
}

func GetAbilityList() []string {
	return []string{
		RoleRead,
		RoleCreate,
		RoleUpdate,
		RoleDelete,
		RoleXRead,
		RoleXCreate,
		RoleXUpdate,
		RoleXDelete,
	}
}

func SelectModuleAccess(id int64) map[string][]string {

	var module string
	var ability string

	privilege := make(map[string][]string)
	query := fmt.Sprintf(`
		SELECT
			module,
			ability
		FROM
			rolegroups_modules
		WHERE
			rolegroups_id = (%d)
	`, id)
	rows, err := conn.DB.Query(query)
	if err != nil {
		return privilege
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&module, &ability); err != nil {
			log.Fatalln(err.Error())
		}
		privilege[module] = append(privilege[module], ability)
	}

	return privilege
}

// IsExistName ...
func IsExistName(name string) bool {

	var x string
	query := fmt.Sprintf(`
		SELECT
			'x'
		FROM
			rolegroups
		WHERE
			name = ('%s')
		LIMIT 1;
	`, name)

	err := conn.DB.Get(&x, query)
	if err != nil {
		return false
	}

	return true
}

// Insert ...
func Insert(name string, tx *sqlx.Tx) (int64, error) {

	query := fmt.Sprintf(`
		INSERT INTO
			rolegroups (
				name,
				created_at,
				updated_at
			)
			VALUES (
				('%s'),
				NOW(),
				NOW()
			);
	`, name)

	var result sql.Result
	var err error
	if tx != nil {
		result, err = tx.Exec(query)
	} else {
		result, err = conn.DB.Exec(query)
	}
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

// InsertModuleAccess ...
func InsertModuleAccess(rolegroupsID int64, privileges map[string][]string, tx *sqlx.Tx) error {

	var value []string
	for module, abilities := range privileges {
		for _, ability := range abilities {
			value = append(value, fmt.Sprintf("(%d, '%s', '%s', NOW(), NOW())", rolegroupsID, module, ability))
		}
	}

	queryValue := strings.Join(value, ", ")
	query := fmt.Sprintf(`
		INSERT INTO
			rolegroups_modules (
				rolegroups_id,
				module,
				ability,
				created_at,
				updated_at
			)
			VALUES %s;
	`, queryValue)

	var err error
	if tx != nil {
		_, err = tx.Exec(query)
	} else {
		_, err = conn.DB.Exec(query)
	}
	if err != nil {
		return err
	}

	return nil
}
