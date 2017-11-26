package notification

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"database/sql/driver"
	"reflect"
	//"database/sql"
	//"fmt"
	"testing"
	"time"

	"github.com/melodiez14/meiko/src/util/conn"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGet(t *testing.T) {
	type args struct {
		userID  int64
		page 	uint16
		limit	uint8
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
		want    []Notification
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				userID:	140810140060,
				page:	3,
				limit:	1,
			},
			mock: mock{
				query:  `^(\s*)SELECT(\s*)id,(\s*)name,(\s*)descriptions,(\s*)read_at,(\s*)table_id,(\s*)table_name,(\s*)created_at(\s*)FROM(\s*)notifications(\s*)WHERE(\s*)users_id(\s*)=(\s*)(.+)(\s*)ORDER BY(\s*)created_at(\s*)DESC(\s*)LIMIT(\s*)(.+)(\s*)\,(\s*)(.+)(\s*)`,
				column: []string{"id", "name", "description", "read_at", "table_id", "table_name", "created_at"},
				result: [][]driver.Value{
					[]driver.Value{
						"1", "Task 1", "notification for task 1", time.Now(), "1", "task", time.Now(),
					},
					[]driver.Value{
						"2", "Task 2", "notification for task 2", time.Now(), "1", "task", time.Now(),
					},
					[]driver.Value{
						"3", "Task 3", "notification for task 3", time.Now(), "1", "task", time.Now(),
					},
				},
				err:    nil,
			},
			want: []Notification{
				Notification{
					ID:           1,
					Name:         "Task 1",
					Description:  "notification for task 1",
					ReadAt:       mysql.NullTime{},
					TableID:      "1",
					TableName:    "task",
					CreatedAt:	time.Now(),
				},
				Notification{
					ID:           2,
					Name:         "Task 2",
					Description:  "notification for task 2",
					ReadAt:       mysql.NullTime{},
					TableID:      "1",
					TableName:    "task",
					CreatedAt:	time.Now(),
				},
				Notification{
					ID:           2,
					Name:         "Task 2",
					Description:  "notification for task 2",
					ReadAt:       mysql.NullTime{},
					TableID:      "1",
					TableName:    "task",
					CreatedAt:	time.Now(),
				},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				userID:	140810140060,
				page:	3,
				limit:	1,
			},
			mock: mock{
				query:  `^(\s*)SELECT(\s*)id,(\s*)name,(\s*)descriptions,(\s*)read_at,(\s*)table_id,(\s*)table_name,(\s*)created_at(\s*)FROM(\s*)notifications(\s*)WHERE(\s*)users_id(\s*)=(\s*)(.+)(\s*)ORDER BY(\s*)created_at(\s*)DESC(\s*)LIMIT(\s*)(.+)(\s*)\,(\s*)(.+)(\s*)`,
				column: []string{"id", "name", "description", "read_at", "table_id", "table_name", "created_at"},
				result: [][]driver.Value{},
				err:    sql.ErrNoRows,
			},
			want: []Notification{},
			wantErr: false,
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

		// db.Expected
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.userID, tt.args.page,tt.args.limit)
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