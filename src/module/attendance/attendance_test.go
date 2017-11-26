package attendance

import (
	"time"
	"database/sql/driver"
	"database/sql"
	"fmt"
	"testing"
	"reflect"

	"github.com/melodiez14/meiko/src/util/conn"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"	
)

func TestGetByUserCourseID(t *testing.T) {
	type args struct {
		userID		int64
		courseID	int64
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
		want    []Attendance
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				userID:		140810140060,
				courseID:	1,
			},
			mock: mock{
				query:	`^(\s*)SELECT(\s*)id,(\s*)meeting_number,(\s*)status,(\s*)meeting_date(\s*)FROM(\s*)attendances(\s*)WHERE(\s*)p_users_courses_users_id(\s*)=(\s*)(.+)(\s*)AND(\s*)p_users_courses_courses_id(\s*)=(\s*)(.+)(\s*);$`,
				column: []string{"id", "meeting_number", "status", "meeting_date"},
				result: [][]driver.Value{
					[]driver.Value{
						"1", "1", "1", time.Now(),
					},
					[]driver.Value{
						"1", "2", "1", time.Now(),
					},
					[]driver.Value{
						"1", "3", "1", time.Now(),
					},
				},					
				err: nil,
			},
			want: []Attendance{
				Attendance{
					ID:    1,
					MeetingNumber:  1,
					Status: 1,
					Date: time.Now(),
				},
				Attendance{
					ID:    1,
					MeetingNumber:  2,
					Status: 1,
					Date: time.Now(),
				},
				Attendance{
					ID:    1,
					MeetingNumber:  3,
					Status: 1,
					Date: time.Now(),
				},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				courseID:	1,
			},
			mock: mock{
				query:	`^(\s*)SELECT(\s*)id,(\s*)meeting_number,(\s*)status,(\s*)meeting_date(\s*)FROM(\s*)attendances(\s*)WHERE(\s*)p_users_courses_users_id(\s*)=(\s*)(.+)(\s*)AND(\s*)p_users_courses_courses_id(\s*)=(\s*)(.+)(\s*);$`,
				column: []string{"id", "meeting_number", "status", "meeting_date"},
				result: [][]driver.Value{},
				err: sql.ErrNoRows,
			},
			want: []Attendance{},
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				courseID:	1,
			},
			mock: mock{
				query:	`^(\s*)SELECT(\s*)id,(\s*)meeting_number,(\s*)status,(\s*)meeting_date(\s*)FROM(\s*)attendances(\s*)WHERE(\s*)p_users_courses_users_id(\s*)=(\s*)(.+)(\s*)AND(\s*)p_users_courses_courses_id(\s*)=(\s*)(.+)(\s*);$`,
				column: []string{"id", "meeting_number", "status", "meeting_date"},
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
			rows := sqlmock.NewRows(tt.mock.column)
			for _, val := range tt.mock.result {
				rows.AddRow(val...)
			}
			q.WillReturnRows(rows)
		} else {
			q.WillReturnError(tt.mock.err)
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := GetByUserCourseID(tt.args.userID,tt.args.courseID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ByUserCourseID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByCourseID() = %v, want %v", got, tt.want)
			}
		})
	}
}
