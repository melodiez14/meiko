package place

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"testing"
	"reflect"

	"github.com/jmoiron/sqlx"
	"github.com/melodiez14/meiko/src/util/conn"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestSearch(t *testing.T) {
	type args struct {
		id string
	}
	type mock struct {
		query  string
		column []string
		result []driver.Value
		err    error
	}
	tests := []struct {
		name 	string
		mock	 mock
		args 	args
		want 	[]string
		wantErr	bool
	}{
		{
			name: "Test Case 1",
			args: args{
				id: "453",
			},
			mock: mock{
				query:  `^(\s*)SELECT(\s*)id(\s*)FROM(\s*)places(\s*)WHERE(\s*)id(\s*)LIKE(\s*)\(\d*\)(\s*)`,
				column: []string{"id"},
				result: []driver.Value{"45311","45312"},
				err:    nil,
			},
			want: []string{"45311","45312"},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				id: "453",
			},
			mock: mock{
				query:  `^(\s*)SELECT(\s*)id(\s*)FROM(\s*)places(\s*)WHERE(\s*)id(\s*)LIKE(\s*)\(\d*\)(\s*)`,
				column: []string{"id"},
				result: []driver.Value{},
				err:    sql.ErrNoRows,
			},
			want: []string{},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			args: args{
				id: "453",
			},
			mock: mock{
				query:  `^(\s*)SELECT(\s*)id(\s*)FROM(\s*)places(\s*)WHERE(\s*)id(\s*)LIKE(\s*)\(\d*\)(\s*)`,
				column: []string{"id"},
				result: nil,				
				err: fmt.Errorf("Error Connection"),
			},
			want: nil,
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

		t.Run(tt.name, func(t *testing.T) {
			got, err := Search(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsID(t *testing.T) {
	type args struct {
		id string
	}
	type mock struct {
		query  string
		column []string
		result []driver.Value
		err    error
	}
	tests := []struct {
		name string
		args args
		mock mock
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				id: "453",
			},
			mock: mock{
				query:  `^(\s*)SELECT(\s*)id(\s*)FROM(\s*)places(\s*)HERE(\s*)id(\s*)=(\s*)(.+)(\s*)LIMIT(\s*)1(\s*)`,
				column: []string{"id"},
				result: []driver.Value{"453"},
				err:    nil,
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				id: "453",
			},
			mock: mock{
				query:  `^(\s*)SELECT(\s*)id(\s*)FROM(\s*)places(\s*)HERE(\s*)id(\s*)=(\s*)(.+)(\s*)LIMIT(\s*)1(\s*)`,
				column: []string{"id"},
				result: []driver.Value{},
				err:    sql.ErrNoRows,
			},
			want: false,
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

		t.Run(tt.name, func(t *testing.T) {
			if got := IsExistID(tt.args.id); got != tt.want {
				t.Errorf("IsIDExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
