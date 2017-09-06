package rolegroup

import (
	"fmt"

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
