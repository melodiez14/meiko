package assignment

import (
	"time"
	"database/sql/driver"
	"database/sql"
	"fmt"
	"reflect"
	"testing"
	

	"github.com/melodiez14/meiko/src/util/conn"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetByCourseID(t *testing.T) {
	type args struct {
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
		want    []Assignment
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				courseID:	1,
			},
			mock: mock{
				query:	`^\s*SELECT(\s*)id,(\s*)name,(\s*)status,(\s*)upload_date,(\s*)due_date(\s*)FROM(\s*)assignments(\s*)WHERE(\s*)EXISTS\s\(\s*SELECT(\s*)id(\s*)FROM(\s*)grade_parameters(\s*)WHERE(\s*)courses_id(\s*)=(\s*)(.+)\s*\);$`,
				column: []string{"id", "name", "status", "upload_date", "due_date"},
				result: [][]driver.Value{
					[]driver.Value{
						"1", "Basis Data", "1", time.Now(),time.Now(),
					},
					[]driver.Value{
						"1", "Basis Data", "2", time.Now(),time.Now(),
					},
					[]driver.Value{
						"1", "Basis Data", "3", time.Now(),time.Now(),
					},
				},					
				err: nil,
			},
			want: []Assignment{
				Assignment{
					ID:    1,
					Name:  "Basis Data",
					Status: 1,
					UploadDate: time.Now(),
					DueDate: time.Now(),
				},
				Assignment{
					ID:    1,
					Name:  "Basis Data",
					Status: 2,
					UploadDate: time.Now(),
					DueDate: time.Now(),
				},
				Assignment{
					ID:    1,
					Name:  "Basis Data",
					Status: 3,
					UploadDate: time.Now(),
					DueDate: time.Now(),
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
				query:	`^\s*SELECT(\s*)id,(\s*)name,(\s*)status,(\s*)upload_date,(\s*)due_date(\s*)FROM(\s*)assignments(\s*)WHERE(\s*)EXISTS\s\(\s*SELECT(\s*)id(\s*)FROM(\s*)grade_parameters(\s*)WHERE(\s*)courses_id(\s*)=(\s*)(d.+)\s*\);$`,
				column: []string{"id", "name", "status", "upload_date", "due_date"},
				result: [][]driver.Value{},
				err: sql.ErrNoRows,
			},
			want: []Assignment{},
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				courseID:	1,
			},
			mock: mock{
				query:	`^\s*SELECT(\s*)id,(\s*)name,(\s*)status,(\s*)upload_date,(\s*)due_date(\s*)FROM(\s*)assignments(\s*)WHERE(\s*)EXISTS\s\(\s*SELECT(\s*)id(\s*)FROM(\s*)grade_parameters(\s*)WHERE(\s*)courses_id(\s*)=(\s*)(d.+)\s*\);$`,
				column: []string{"id", "name", "status", "upload_date", "due_date"},
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
			got, err := GetByCourseID(tt.args.courseID)
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

func TestGetIncompleteByUserID(t *testing.T) {
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
		want    []Assignment
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				userID:	140810140060,
			},
			mock: mock{
				query:	`^(\s*)SELECT(\s*)id,(\s*)name,(\s*)status,(\s*)upload_date,(\s*)due_date(\s*)FROM(\s*)assigments(\s*)WHERE(\s*)EXISTS(\s*)\((\s*)SELECT(\s*)id(\s*)FROM(\s*)grade_parameters(\s*)WHERE(\s*)EXISTS(\s*)\((\s*)SELECT(\s*)courses_id(\s*)FROM(\s*)p_users_courses(\s*)WHERE(\s*)users_id(\s*)=(\s*)(/+)(\s*)\)\)(\s*)AND(\s*)id(\s*)NOT(\s*)IN(\s*)\((\s*)SELECT(\s*)assignments_id(\s*)FROM(\s*)p_users_assignments(\s*)WHERE(\s*)users_id(\s*)=(\s*)(.+)(\s*)\);$`,
				column: []string{"id", "name", "status", "upload_date", "due_date"},
				result: [][]driver.Value{
					[]driver.Value{
						"1", "Basis Data", "1", time.Now(),time.Now(),
					},
					[]driver.Value{
						"2", "Kriptografi", "1", time.Now(),time.Now(),
					},
					[]driver.Value{
						"3", "Pemrograman Web", "1", time.Now(),time.Now(),
					},
				},					
				err: nil,
			},
			want: []Assignment{
				Assignment{
					ID:    1,
					Name:  "Basis Data",
					Status: 1,
					UploadDate: time.Now(),
					DueDate: time.Now(),
				},
				Assignment{
					ID:    2,
					Name:  "Kriptografi",
					Status: 1,
					UploadDate: time.Now(),
					DueDate: time.Now(),
				},
				Assignment{
					ID:    3,
					Name:  "Pemrograman Web",
					Status: 1,
					UploadDate: time.Now(),
					DueDate: time.Now(),
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
				query:	`^\s*SELECT(\s*)id,(\s*)name,(\s*)status,(\s*)upload_date,(\s*)due_date(\s*)FROM(\s*)assignments(\s*)WHERE(\s*)EXISTS\s\(\s*SELECT(\s*)id(\s*)FROM(\s*)grade_parameters(\s*)WHERE(\s*)courses_id(\s*)=(\s*)(d.+)\s*\);$`,
				column: []string{"id", "name", "status", "upload_date", "due_date"},
				result: [][]driver.Value{},
				err: sql.ErrNoRows,
			},
			want: []Assignment{},
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				userID:	140810140060,
			},
			mock: mock{
				query:	`^\s*SELECT(\s*)id,(\s*)name,(\s*)status,(\s*)upload_date,(\s*)due_date(\s*)FROM(\s*)assignments(\s*)WHERE(\s*)EXISTS\s\(\s*SELECT(\s*)id(\s*)FROM(\s*)grade_parameters(\s*)WHERE(\s*)courses_id(\s*)=(\s*)(d.+)\s*\);$`,
				column: []string{"id", "name", "status", "upload_date", "due_date"},
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
			got, err := GetIncompleteByUserID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetcompleteByUserID(t *testing.T) {
	type args struct {
		userID	int64
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
				userID:	140810140060,
			},
			mock: mock{
				query:	`^(\s*)SELECT(\s*)assignments_id(\s*)FROM(\s*)p_users_assignments(\s*)WHERE(\s*)users_id(\s*)=(\s*)(.+)(\s*);$`,
				column: []string{"assignments_id"},
				result: []driver.Value{"1", "2", "3",},
				err: nil,
			},
			want: []int64{1,2,3},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				userID:	140810140060,
			},
			mock: mock{
				query:	`^(\s*)SELECT(\s*)assignments_id(\s*)FROM(\s*)p_users_assignments(\s*)WHERE(\s*)users_id(\s*)=(\s*)(.+)(\s*);$`,
				column: []string{"assignments_id"},
				result: []driver.Value{},
				err: sql.ErrNoRows,
			},
			want: []int64{},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			args: args{
				userID:	140810140060,
			},
			mock: mock{
				query:	`^(\s*)SELECT(\s*)assignments_id(\s*)FROM(\s*)p_users_assignments(\s*)WHERE(\s*)users_id(\s*)=(\s*)(.+)(\s*);$`,
				column: []string{"assignments_id"},
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
			got, err := GetCompleteByUserID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}
