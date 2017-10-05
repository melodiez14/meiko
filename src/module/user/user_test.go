package user

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/melodiez14/meiko/src/util/conn"
)

func TestGetByEmail(t *testing.T) {
	type args struct {
		email  string
		column []string
	}
	type mock struct {
		query  string
		column []string
		result []driver.Value
		err    error
	}
	tests := []struct {
		name    string
		args    args
		mock    mock
		want    User
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				email:  "risal@live.com",
				column: []string{},
			},
			mock: mock{
				query:  `SELECT(\s*)(.+)(\s*)FROM(\s*)users(\s*)WHERE(\s*)email(\s*)=(\s*)(.+)(\s*)LIMIT(\s*)1`,
				column: []string{"id", "name", "email", "gender", "note", "status", "identity_code", "line_id", "phone", "rolegroups_id"},
				result: []driver.Value{"1", "Risal Falah", "risal@live.com", "1", "", "2", "140810140016", nil, nil, nil},
				err:    nil,
			},
			want: User{
				ID:           1,
				Name:         "Risal Falah",
				Email:        "risal@live.com",
				Gender:       1,
				Note:         "",
				Status:       2,
				IdentityCode: 140810140016,
				LineID:       sql.NullString{},
				Phone:        sql.NullString{},
				RoleGroupsID: sql.NullInt64{},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				email:  "risal@live.com",
				column: []string{ColID, ColName, ColEmail},
			},
			mock: mock{
				query:  `SELECT(\s*)(.+)(\s*)FROM(\s*)users(\s*)WHERE(\s*)email(\s*)=(\s*)(.+)(\s*)LIMIT(\s*)1`,
				column: []string{"id", "name", "email"},
				result: []driver.Value{"1", "Risal Falah", "risal@live.com"},
				err:    nil,
			},
			want: User{
				ID:    1,
				Name:  "Risal Falah",
				Email: "risal@live.com",
			},
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				email:  "risal@live.com",
				column: []string{ColID, ColName, ColEmail},
			},
			mock: mock{
				query:  `SELECT(\s*)(.+)(\s*)FROM(\s*)users(\s*)WHERE(\s*)email(\s*)=(\s*)(.+)(\s*)LIMIT(\s*)1`,
				column: []string{"id", "name", "xxx"},
				result: []driver.Value{},
				err:    fmt.Errorf("Invalid column"),
			},
			want:    User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		db, _ := conn.InitDBMock()
		q := db.ExpectQuery(tt.mock.query)
		if tt.mock.err == nil {
			q.WillReturnRows(sqlmock.NewRows(tt.mock.column).
				AddRow(tt.mock.result...))
		} else {
			q.WillReturnError(tt.mock.err)
		}

		// db.Expected
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetByEmail(tt.args.email, tt.args.column...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
