package course

import (
	"fmt"
	"testing"
	"database/sql/driver"
	"reflect"
	//"strings"

	//"github.com/jmoiron/sqlx"

	"database/sql"

	"github.com/melodiez14/meiko/src/util/conn"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetByUserID(t *testing.T) {
	type args struct {
		userID	int64
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
		want    []Course
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				userID:	140810140060,
			},
			mock: mock{
				query:	`^(\s*)SELECT(\s*)id,(\s*)name,(\s*)ucu,(\s*)semester,(\s*)status(\s*)FROM(\s*)courses(\s*)WHERE(\s*)EXISTS(\s*)\((\s*)SELECT(\s*)courses_id(\s*)FROM(\s*)p_users_courses(\s*)WHERE(\s*)users_id(\s*)=(\s*)(.+)(\s*)\);$`,
				column: []string{"id", "name", "ucu", "semester", "status"},
				result: [][]driver.Value{
					[]driver.Value{
						"1", "Basis Data", "1", "1", "1",
					},
					[]driver.Value{
						"2", "Praktikum", "1", "1", "1",
					},
					[]driver.Value{
						"3", "Pemrograman Web", "1", "1", "1",
					},
				},					
				err: nil,
			},
			want: []Course{
				Course{
					ID:    		1,
					Name:  		"Basis Data",
					UCU:		1,
					Semester:	1,
					Status: 	1,
				},
				Course{
					ID:    		2,
					Name:  		"Praktikum",
					UCU:		1,
					Semester:	1,
					Status: 	1,
				},
				Course{
					ID:    		3,
					Name:  		"Pemrograman Web",
					UCU:		1,
					Semester:	1,
					Status: 	1,
				},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				userID:	140810140060,
			},
			mock: mock{
				query:	`^(\s*)SELECT(\s*)id,(\s*)name,(\s*)ucu,(\s*)semester,(\s*)status(\s*)FROM(\s*)courses(\s*)WHERE(\s*)EXISTS(\s*)\((\s*)SELECT(\s*)courses_id(\s*)FROM(\s*)p_users_courses(\s*)WHERE(\s*)users_id(\s*)=(\s*)(.+)(\s*)\);$`,
				column: []string{"id", "name", "ucu", "semester", "status"},
				result: [][]driver.Value{},
				err: sql.ErrNoRows,
			},
			want: []Course{},
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				userID:	140810140060,
			},
			mock: mock{
				query:	`^(\s*)SELECT(\s*)id,(\s*)name,(\s*)ucu,(\s*)semester,(\s*)status(\s*)FROM(\s*)courses(\s*)WHERE(\s*)EXISTS(\s*)\((\s*)SELECT(\s*)courses_id(\s*)FROM(\s*)p_users_courses(\s*)WHERE(\s*)users_id(\s*)=(\s*)(.+)(\s*)\);$`,
				column: []string{"id", "name", "ucu", "semester", "status"},
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
			got, err := GetByUserID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ByCourseID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByCourseID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectIDByUserID(t *testing.T) {
	type args struct {
		userID     int64
		status	 []string
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
		want    []int64
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				userID:     140810140060,
				status: 	[]string{},
			},
			mock: mock{
				query:  `^(\s*)SELECT(\s*)courses_id(\s*)FROM(\s*)p_users_courses(\s*)WHERE(\s*)users_id(\s*)=(\s*)(.+)(\s*);$`,
				column: []string{"courses_id"},
				result: []driver.Value{"1","2","3"},
				err: nil,
			},
			want: []int64{1,2,3},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				userID:     140810140060,
				status: []string{},
			},
			mock: mock{
				query:  `^(\s*)SELECT(\s*)courses_id(\s*)FROM(\s*)p_users_courses(\s*)WHERE(\s*)users_id(\s*)=(\s*)(.+)(\s*);$`,
				column: []string{"courses_id"},
				result: []driver.Value{1},
				err: nil,
			},
			want: []int64{1},
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				userID:     140810140060,
				status: []string{},
			},
			mock: mock{
				query:  `^(\s*)SELECT(\s*)courses_id(\s*)FROM(\s*)p_users_courses(\s*)WHERE(\s*)users_id(\s*)=(\s*)(.+)(\s*);$`,
				column: []string{"courses_id"},
				result: nil,
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
			q.WillReturnRows(sqlmock.NewRows(tt.mock.column).
				AddRow(tt.mock.result...))
		} else {
			q.WillReturnError(tt.mock.err)
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := GetByUserID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestIsEnrolled(t *testing.T) {
	type args struct {
		userID 		int64
		courseID    int64
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
				userID: 	140810140016,
				courseID:   1,
			},
			mock: mock{
				query:  `^(\s)SELECT(\s)users_id(\s)FROM(\s)p_users_courses(\s)WHERE(\s)users_id(\s)=(\s)(.+)(\s)AND(\s)courses_id(\s)=(\s)(.+)(\s)LIMIT(\s)1`,
				column: []string{"users_id"},
				result: []driver.Value{"140810140060"},
				err:    nil,
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				userID: 	140810140016,
				courseID:   1,
			},
			mock: mock{
				query:  `^(\s)SELECT(\s)users_id(\s)FROM(\s)p_users_courses(\s)WHERE(\s)users_id(\s)=(\s)(.+)(\s)AND(\s)courses_id(\s)=(\s)(.+)(\s)LIMIT(\s)1`,
				column: []string{"users_id"},
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
			if got := IsEnrolled(tt.args.userID, tt.args.courseID); got != tt.want {
				t.Errorf("IsEnrolled() = %v, want %v", got, tt.want)
			}
		})
	}
}

// SelectAssistantID Query to select Assisten by its iD
/*
	@params:
		courseID	= int64
	@example:
		courseID	= 1
	@return:
		userIDs	= userID{}
*/
func TestSelectAssistantID(t *testing.T) {
	type args struct {
		courseID     int64
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
		want    []int64
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				courseID:     1,
			},
			mock: mock{
				query:  `^(\s*)SELECT(\s*)users_id(\s*)FROM(\s*)p_users_courses(\s*)WHERE(\s*)courses_id(\s*)=(\s*)(.+)(\s*)AND(\s*)status(\s*)=(\s*)(.+)(\s*)`,
				column: []string{"users_id"},
				result: []driver.Value{"140810140060","140810140061","140810140062"},
				err: nil,
			},
			want: []int64{140810140060,140810140061,140810140062},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				courseID:     1,
			},
			mock: mock{
				query:  `^(\s*)SELECT(\s*)users_id(\s*)FROM(\s*)p_users_courses(\s*)WHERE(\s*)courses_id(\s*)=(\s*)(.+)(\s*)AND(\s*)status(\s*)=(\s*)(.+)(\s*)`,
				column: []string{"users_id"},
				result: []driver.Value{"140810140060"},
				err: nil,
			},
			want: []int64{140810140060},
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				courseID:     1,
			},
			mock: mock{
				query:  `^(\s*)SELECT(\s*)users_id(\s*)FROM(\s*)p_users_courses(\s*)WHERE(\s*)courses_id(\s*)=(\s*)(.+)(\s*)AND(\s*)status(\s*)=(\s*)(.+)(\s*)`,
				column: []string{"users_id"},
				result: nil,
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
			q.WillReturnRows(sqlmock.NewRows(tt.mock.column).
				AddRow(tt.mock.result...))
		} else {
			q.WillReturnError(tt.mock.err)
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectAssistantID(tt.args.courseID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}
