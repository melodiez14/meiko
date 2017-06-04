package assignment

import (
	"reflect"
	"testing"

	"github.com/melodiez14/meiko/src/util/conn"

	"time"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetByCourseID(t *testing.T) {

	type args struct {
		courseID int64
	}
	type mock struct {
		query string
		rows  *sqlmock.Rows
		err   error
	}

	now := time.Now()
	mockDB, _ := conn.InitDBMock()
	m := *mockDB

	tests := []struct {
		name    string
		args    args
		mock    mock
		want    []Assignment
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				courseID: 1,
			},
			mock: mock{
				query: `SELECT (.+) FROM assignments WHERE EXISTS \( SELECT (.+) FROM grade_parameters WHERE courses_id = (.+) \);`,
				rows: sqlmock.NewRows([]string{"id", "name", "status", "upload_date", "due_date"}).
					AddRow(1, "Object Oriented Programming", 0, now, now),
			},
			want: []Assignment{
				Assignment{
					ID:         1,
					Name:       "Object Oriented Programming",
					Status:     0,
					UploadDate: now,
					DueDate:    now,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.ExpectQuery(tt.mock.query).WillReturnRows(tt.mock.rows).WillReturnError(tt.mock.err)
			got, err := GetByCourseID(tt.args.courseID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByCourseID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByCourseID() = %v, want %v", got, tt.want)
			}
		})
	}
}
