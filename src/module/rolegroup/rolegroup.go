package rolegroup

import (
	"fmt"
	"log"

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

func Insert(name string) error {
	query := fmt.Sprintf(insertQuery, name)
	_, err := conn.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
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
	}
}

func GetRoleList() []string {
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

func GetModuleAccess(id int64) map[string][]string {

	var module string
	var ability string

	privilege := make(map[string][]string)
	query := fmt.Sprintf(queryGetModuleAccess, id)
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
