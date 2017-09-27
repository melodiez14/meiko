package module

import (
	"fmt"
	"log"

	"github.com/melodiez14/meiko/src/util/conn"
)

func SelectByPage(page, offset uint16) ([]Module, error) {
	modules := []Module{}
	query := fmt.Sprintf(getQuery, page, offset)
	err := conn.DB.Select(&modules, query)
	if err != nil {
		return nil, err
	}

	return modules, nil
}

func GetPriviegeByRoleGroupID(rolegroupID int64) map[string][]string {

	var module string
	var ability string

	privilege := make(map[string][]string)
	query := fmt.Sprintf(getPrivilegeQuery, rolegroupID)
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
