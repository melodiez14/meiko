package rolegroup

import (
	"reflect"
	"database/sql/driver"
	"testing"
	"fmt"
	"github.com/melodiez14/meiko/src/util/conn"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"

)

func TestGetByPage(t *testing.T) {
	type args struct {
		page   uint16
		offset uint16
	}
	type mock struct {
		query  string
		column []string
		result [][]driver.Value
		err    error
	}
	tests := []struct {
		name    string
		args    args
		mock    mock
		want    []RoleGroup
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				page:     1,
				offset: 10,
			},
			mock: mock{
				query:  `^(\s*)SELECT(\s*)id,(\s*)name(\s*)FROM(\s*)rolegroups(\s*)LIMIT(\s*)(.+)(\s*)OFFSET(\s*)(.+)`,
				column: []string{"id", "name"},
				result: [][]driver.Value{
					[]driver.Value{
						"1", "Master",
					},
					[]driver.Value{
						"2", "Assistant",
					},
					[]driver.Value{
						"3", "User",
					},
				},
				err: nil,
			},
			want: []RoleGroup{
				RoleGroup{
					ID: 	1,
					Name:   "Master",
				},
				RoleGroup{
					ID: 	2,
					Name:   "Assistant",
				},
				RoleGroup{
					ID: 	3,
					Name:   "User",
				},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				page:     1,
				offset: 10,
			},
			mock: mock{
				query:  `^(\s*)SELECT(\s*)id,(\s*)name(\s*)FROM(\s*)rolegroups(\s*)LIMIT(\s*)(.+)(\s*)OFFSET(\s*)(.+)`,
				column: []string{"id", "name"},
				result: [][]driver.Value{},
				err:    fmt.Errorf("Error connection"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		db, _ := conn.InitDBMock()
		q := db.ExpectQuery(tt.mock.query)
		if tt.mock.err == nil {
			rows := sqlmock.NewRows(tt.mock.column)
			for _, val := range tt.mock.result {
				rows.AddRow(val...)
			}
			q.WillReturnRows(rows)
		} else {
			q.WillReturnError(tt.mock.err)
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetByPage(tt.args.page, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	type args struct {
		name    string
	}
	type mock struct {
		query        string
		lastInsertID int64
		rowsAffected int64
		err          error
	}
	tests := []struct {
		name    string
		args    args
		mock    mock
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				name:    "master",
			},
			mock: mock{
				query:	`^(\s*)INSERT\sINTO(\s*)rolegroups\s\((\s*)name,(\s*)created_at,(\s*)updated_at(\s)\)(\s*)VALUES\s\(\s*\('([\w ]*)'\),\s*NOW\(\),\s*NOW\(\)\s*\);$`,
				lastInsertID: 1,
				rowsAffected: 1,
				err:          nil,
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				name:    "assistant",
			},
			mock: mock{
				query:	`^(\s*)INSERT\sINTO(\s*)rolegroups\s\((\s*)name,(\s*)created_at,(\s*)updated_at(\s)\)(\s*)VALUES\s\(\s*\('([\w ]*)'\),\s*NOW\(\),\s*NOW\(\)\s*\);$`,
				lastInsertID: 0,
				rowsAffected: 0,
				err:          nil,
			},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			args: args{
				name:    "assistant",
			},
			mock: mock{
				query:	`^(\s*)INSERT\sINTO(\s*)rolegroups\s\((\s*)name,(\s*)created_at,(\s*)updated_at(\s)\)(\s*)VALUES\s\(\s*\('([\w ]*)'\),\s*NOW\(\),\s*NOW\(\)\s*\);$`,
				lastInsertID: 0,
				rowsAffected: 0,
				err:          fmt.Errorf("Error connection"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		db, _ := conn.InitDBMock()
		q := db.ExpectExec(tt.mock.query)
		if tt.mock.err == nil {
			q.WillReturnResult(sqlmock.NewResult(tt.mock.lastInsertID, tt.mock.rowsAffected))
		} else {
			q.WillReturnError(tt.mock.err)
		}
		t.Run(tt.name, func(t *testing.T) {
			if err := Insert(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}



func TestUpdate(t *testing.T) {
	type args struct {
		id 		int64
		name    string
	}
	type mock struct {
		query        string
		lastInsertID int64
		rowsAffected int64
		err          error
	}
	tests := []struct {
		name    string
		args    args
		mock    mock
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				id:		 1,
				name:    "master",
			},
			mock: mock{
				query:        `^(\s*)UPDATE(\s*)rolegroups(\s*)SET(\s*)name(\s*)=(\s*)(.+)WHERE(\s*)id(\s*)=(\s*)(.+)(\s*)`,
				lastInsertID: 1,
				rowsAffected: 1,
				err:          nil,
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				id:		 2,
				name:    "assistant",

			},
			mock: mock{
				query:        `^(\s*)UPDATE(\s*)rolegroups(\s*)SET(\s*)name(\s*)=(\s*)(.+)WHERE(\s*)id(\s*)=(\s*)(.+)(\s*)`,
				lastInsertID: 0,
				rowsAffected: 0,
				err:          nil,
			},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			args: args{
				id:		 2,
				name:    "assistant",
			},
			mock: mock{
				query:        `^(\s*)UPDATE(\s*)rolegroups(\s*)SET(\s*)name(\s*)=(\s*)(.+)WHERE(\s*)id(\s*)=(\s*)(.+)(\s*)`,
				lastInsertID: 0,
				rowsAffected: 0,
				err:          fmt.Errorf("Error connection"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		db, _ := conn.InitDBMock()
		q := db.ExpectExec(tt.mock.query)
		if tt.mock.err == nil {
			q.WillReturnResult(sqlmock.NewResult(tt.mock.lastInsertID, tt.mock.rowsAffected))
		} else {
			q.WillReturnError(tt.mock.err)
		}
		t.Run(tt.name, func(t *testing.T) {
			if err := Update(tt.args.id, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
