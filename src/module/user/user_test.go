package user

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/melodiez14/meiko/src/util/conn"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
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
				query:  `^\s*SELECT(\s*)(.+)(\s*)FROM(\s*)users(\s*)WHERE(\s*)email(\s*)=(\s*)(.+)(\s*)LIMIT(\s*)1`,
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
				query:  `^\s*SELECT(\s*)(.+)(\s*)FROM(\s*)users(\s*)WHERE(\s*)email(\s*)=(\s*)(.+)(\s*)LIMIT(\s*)1`,
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
				column: []string{ColID, ColName, "xxx"},
			},
			mock: mock{
				query:  `^\s*SELECT(\s*)(.+)(\s*)FROM(\s*)users(\s*)WHERE(\s*)email(\s*)=(\s*)(.+)(\s*)LIMIT(\s*)1`,
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

func TestGetByIdentityCode(t *testing.T) {
	type args struct {
		identityCode int64
		column       []string
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
				identityCode: 140810140016,
				column:       []string{},
			},
			mock: mock{
				query:  `^\s*SELECT(\s*)(.+)(\s*)FROM(\s*)users(\s*)WHERE(\s*)identity_code(\s*)=(\s*)(.+)(\s*)LIMIT(\s*)1`,
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
				identityCode: 140810140016,
				column:       []string{ColID, ColName, ColEmail},
			},
			mock: mock{
				query:  `^\s*SELECT(\s*)(.+)(\s*)FROM(\s*)users(\s*)WHERE(\s*)identity_code(\s*)=(\s*)(.+)(\s*)LIMIT(\s*)1`,
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
				identityCode: 140810140016,
				column:       []string{ColID, ColName, "xxx"},
			},
			mock: mock{
				query:  `^\s*SELECT(\s*)(.+)(\s*)FROM(\s*)users(\s*)WHERE(\s*)identity_code(\s*)=(\s*)(.+)(\s*)LIMIT(\s*)1`,
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

		t.Run(tt.name, func(t *testing.T) {
			got, err := GetByIdentityCode(tt.args.identityCode, tt.args.column...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByIdentityCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByIdentityCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignIn(t *testing.T) {
	type args struct {
		email    string
		password string
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
				email:    "risal@live.com",
				password: "2af9b1ba42dc5eb01743e6b3759b6e4b",
			},
			mock: mock{
				query:  `^\s*SELECT(\s*)id,(\s*)name,(\s*)email,(\s*)gender,(\s*)note,(\s*)status,(\s*)identity_code,(\s*)line_id,(\s*)phone,(\s*)rolegroups_id(\s*)FROM(\s*)users(\s*)WHERE(\s*)email(\s*)=(\s*)(.+)AND(\s*)password(\s*)=(\s*)(.+)(\s*)LIMIT(\s*)1;$`,
				column: []string{"id", "name", "email", "gender", "note", "status", "identity_code", "line_id", "phone", "rolegroups_id"},
				result: []driver.Value{"1", "Risal Falah", "risal@live.com", "1", "", "2", "140810140016", nil, nil, "1"},
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
				RoleGroupsID: sql.NullInt64{Valid: true, Int64: 1},
			},
			wantErr: false,
		},
		{
			name: "Test Case 1",
			args: args{
				email:    "viavallen@metal.com",
				password: "2af9b1ba42dc5eb01743e6b3759b6e4b",
			},
			mock: mock{
				query:  `^\s*SELECT(\s*)id,(\s*)name,(\s*)email,(\s*)gender,(\s*)note,(\s*)status,(\s*)identity_code,(\s*)line_id,(\s*)phone,(\s*)rolegroups_id(\s*)FROM(\s*)users(\s*)WHERE(\s*)email(\s*)=(\s*)(.+)AND(\s*)password(\s*)=(\s*)(.+)(\s*)LIMIT(\s*)1;$`,
				column: []string{"id", "name", "email", "gender", "note", "status", "identity_code", "line_id", "phone", "rolegroups_id"},
				result: []driver.Value{},
				err:    sql.ErrNoRows,
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

		t.Run(tt.name, func(t *testing.T) {
			got, err := SignIn(tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPhoneExist(t *testing.T) {
	type args struct {
		identityCode int64
		phone        string
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
				identityCode: 140810140016,
				phone:        "085860141146",
			},
			mock: mock{
				query:  `^\s*SELECT(\s*)phone(\s*)FROM(\s*)users(\s*)WHERE(\s*)phone(\s*)=(\s*)(.+)AND(\s*)identity_code(\s*)!=(\s*)(.+)(\s*)LIMIT(\s*)1`,
				column: []string{"phone"},
				result: []driver.Value{"085860141146"},
				err:    nil,
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				identityCode: 140810140016,
				phone:        "081231231234",
			},
			mock: mock{
				query:  `^\s*SELECT(\s*)phone(\s*)FROM(\s*)users(\s*)WHERE(\s*)phone(\s*)=(\s*)(.+)AND(\s*)identity_code(\s*)!=(\s*)(.+)(\s*)LIMIT(\s*)1`,
				column: []string{"phone"},
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
			if got := IsPhoneExist(tt.args.identityCode, tt.args.phone); got != tt.want {
				t.Errorf("IsPhoneExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLineIDExist(t *testing.T) {
	type args struct {
		identityCode int64
		lineID       string
	}
	type mock struct {
		query  string
		column []string
		result []driver.Value
		err    error
	}
	tests := []struct {
		name string
		mock mock
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				identityCode: 140810140016,
				lineID:       "risalfa",
			},
			mock: mock{
				query:  `^\s*SELECT(\s*)line_id(\s*)FROM(\s*)users(\s*)WHERE(\s*)line_id(\s*)=(\s*)(.+)AND(\s*)identity_code(\s*)!=(\s*)(.+)(\s*)LIMIT(\s*)1`,
				column: []string{"line_id"},
				result: []driver.Value{"085860141146"},
				err:    nil,
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				identityCode: 140810140016,
				lineID:       "viavallen",
			},
			mock: mock{
				query:  `^\s*SELECT(\s*)line_id(\s*)FROM(\s*)users(\s*)WHERE(\s*)line_id(\s*)=(\s*)(.+)AND(\s*)identity_code(\s*)!=(\s*)(.+)(\s*)LIMIT(\s*)1`,
				column: []string{"line_id"},
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
			if got := IsLineIDExist(tt.args.identityCode, tt.args.lineID); got != tt.want {
				t.Errorf("IsLineIDExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidConfirmationCode(t *testing.T) {
	type args struct {
		email string
		code  uint16
	}
	type mockSelect struct {
		query  string
		column []string
		result []driver.Value
		err    error
	}
	type mockUpdate struct {
		query        string
		lastInsertID int64
		rowsAffected int64
		err          error
	}
	tests := []struct {
		name       string
		args       args
		mockSelect mockSelect
		mockUpdate mockUpdate
		want       bool
	}{
		{
			name: "Test Case 1",
			args: args{
				email: "risal@live.com",
				code:  1234,
			},
			mockSelect: mockSelect{
				query:  `^\s*SELECT(\s*)email_verification_attempt,(\s*)email_verification_code(\s*)FROM(\s*)users(\s*)WHERE(\s*)email(\s*)=(\s*)(.+)AND(\s*)NOW\(\)(\s*)<(\s*)email_verification_expire_date(\s*)LIMIT(\s*)1`,
				column: []string{"email_verification_attempt", "email_verification_code"},
				result: []driver.Value{"0", "1234"},
				err:    nil,
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				email: "risal@live.com",
				code:  1234,
			},
			mockSelect: mockSelect{
				query:  `^\s*SELECT(\s*)email_verification_attempt,(\s*)email_verification_code(\s*)FROM(\s*)users(\s*)WHERE(\s*)email(\s*)=(\s*)(.+)AND(\s*)NOW\(\)(\s*)<(\s*)email_verification_expire_date(\s*)LIMIT(\s*)1`,
				column: []string{"email_verification_attempt", "email_verification_code"},
				result: []driver.Value{"4", "1234"},
				err:    nil,
			},
			want: false,
		},
		{
			name: "Test Case 3",
			args: args{
				email: "risal@live.com",
				code:  1234,
			},
			mockSelect: mockSelect{
				query:  `^\s*SELECT(\s*)email_verification_attempt,(\s*)email_verification_code(\s*)FROM(\s*)users(\s*)WHERE(\s*)email(\s*)=(\s*)(.+)AND(\s*)NOW\(\)(\s*)<(\s*)email_verification_expire_date(\s*)LIMIT(\s*)1`,
				column: []string{"email_verification_attempt", "email_verification_code"},
				result: []driver.Value{"1", "3213"},
				err:    nil,
			},
			want: false,
		},
		{
			name: "Test Case 4",
			args: args{
				email: "risal@live.com",
				code:  1234,
			},
			mockSelect: mockSelect{
				query:  `^\s*SELECT(\s*)email_verification_attempt,(\s*)email_verification_code(\s*)FROM(\s*)users(\s*)WHERE(\s*)email(\s*)=(\s*)(.+)AND(\s*)NOW\(\)(\s*)<(\s*)email_verification_expire_date(\s*)LIMIT(\s*)1`,
				column: []string{},
				result: []driver.Value{},
				err:    fmt.Errorf("Error connection"),
			},
			want: false,
		},
		{
			name: "Test Case 5",
			args: args{
				email: "risal@live.com",
				code:  1234,
			},
			mockSelect: mockSelect{
				query:  `^\s*SELECT(\s*)email_verification_attempt,(\s*)email_verification_code(\s*)FROM(\s*)users(\s*)WHERE(\s*)email(\s*)=(\s*)(.+)AND(\s*)NOW\(\)(\s*)<(\s*)email_verification_expire_date(\s*)LIMIT(\s*)1`,
				column: []string{"email_verification_attempt", "email_verification_code"},
				result: []driver.Value{"1", "4321"},
				err:    nil,
			},
			mockUpdate: mockUpdate{
				query:        `UPDATE(\s*)users(\s*)SET(\s*)email_verification_attempt(\s*)=(\s*)email_verification_attempt \+ 1,(\s*)updated_at = NOW\(\)(\s*)WHERE(\s*)id(\s*)=(\s*)(.+)`,
				lastInsertID: 1,
				rowsAffected: 1,
				err:          nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		db, _ := conn.InitDBMock()
		q := db.ExpectQuery(tt.mockSelect.query)
		// QuerySelect Mock
		if tt.mockSelect.err == nil {
			q.WillReturnRows(sqlmock.NewRows(tt.mockSelect.column).
				AddRow(tt.mockSelect.result...))
		} else {
			q.WillReturnError(tt.mockSelect.err)
		}
		// QueryUpdate Mock
		u := db.ExpectExec(tt.mockUpdate.query)
		if tt.mockUpdate.err == nil {
			u.WillReturnResult(sqlmock.NewResult(tt.mockUpdate.lastInsertID, tt.mockUpdate.rowsAffected))
		} else {
			u.WillReturnError(tt.mockUpdate.err)
		}

		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidConfirmationCode(tt.args.email, tt.args.code); got != tt.want {
				t.Errorf("IsValidConfirmationCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectDashboard(t *testing.T) {
	type args struct {
		id     int64
		limit  uint16
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
		want    []User
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				id:     1,
				limit:  10,
				offset: 10,
			},
			mock: mock{
				query:  `^\s*SELECT(\s*)identity_code,(\s*)name,(\s*)email,(\s*)status(\s*)FROM(\s*)users(\s*)WHERE(\s*)\(status(\s*)=(.+)OR(\s*)status(.+)=(\s*)(.+)\)(\s*)AND(\s*)id(\s*)!=(\s*)(.+)(\s*)LIMIT(\s*)(.+)(\s*)OFFSET(\s*)(.+)`,
				column: []string{"identity_code", "name", "email", "status"},
				result: [][]driver.Value{
					[]driver.Value{
						"140810140016", "Risal Falah", "risal@live.com", "2",
					},
					[]driver.Value{
						"140810140020", "Rifki Muhammad", "rifkirifkigue@gmail.com", "1",
					},
					[]driver.Value{
						"140810140070", "Asep Nur Muhammad Iskandar Yusuf", "asepasepgue@gmail.com", "2",
					},
				},
				err: nil,
			},
			want: []User{
				User{
					IdentityCode: 140810140016,
					Name:         "Risal Falah",
					Email:        "risal@live.com",
					Status:       2,
				},
				User{
					IdentityCode: 140810140020,
					Name:         "Rifki Muhammad",
					Email:        "rifkirifkigue@gmail.com",
					Status:       1,
				},
				User{
					IdentityCode: 140810140070,
					Name:         "Asep Nur Muhammad Iskandar Yusuf",
					Email:        "asepasepgue@gmail.com",
					Status:       2,
				},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				id:     1,
				limit:  10,
				offset: 10,
			},
			mock: mock{
				query:  `^\s*SELECT(\s*)identity_code,(\s*)name,(\s*)email,(\s*)status(\s*)FROM(\s*)users(\s*)WHERE(\s*)\(status(\s*)=(.+)OR(\s*)status(.+)=(\s*)(.+)\)(\s*)AND(\s*)id(\s*)!=(\s*)(.+)(\s*)LIMIT(\s*)(.+)(\s*)OFFSET(\s*)(.+)`,
				column: []string{"identity_code", "name", "email", "status"},
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
			got, err := SelectDashboard(tt.args.id, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectDashboard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectDashboard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateStatus(t *testing.T) {
	type args struct {
		identityCode int64
		status       int8
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
				identityCode: 140810140016,
				status:       1,
			},
			mock: mock{
				query:        `^\s*UPDATE\s*users\s*SET\s*status\s*=\s*\(\d\),\s*updated_at\s*=\s*NOW\(\)\s*WHERE\s*identity_code\s*=\s*\(\d*\);$`,
				lastInsertID: 1,
				rowsAffected: 1,
				err:          nil,
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				identityCode: 140810140016,
				status:       1,
			},
			mock: mock{
				query:        `^\s*UPDATE\s*users\s*SET\s*status\s*=\s*\(\d\),\s*updated_at\s*=\s*NOW\(\)\s*WHERE\s*identity_code\s*=\s*\(\d*\);$`,
				lastInsertID: 0,
				rowsAffected: 0,
				err:          fmt.Errorf("Error connection"),
			},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			args: args{
				identityCode: 140810140016,
				status:       1,
			},
			mock: mock{
				query:        `^\s*UPDATE\s*users\s*SET\s*status\s*=\s*\(\d\),\s*updated_at\s*=\s*NOW\(\)\s*WHERE\s*identity_code\s*=\s*\(\d*\);$`,
				lastInsertID: 0,
				rowsAffected: 0,
				err:          nil,
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
			if err := UpdateStatus(tt.args.identityCode, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("UpdateStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateToVerified(t *testing.T) {
	type args struct {
		identityCode int64
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
				identityCode: 140810140016,
			},
			mock: mock{
				query:        `^\s*UPDATE(\s*)users(\s*)SET(\s*)status(\s*)=(\s*)(.+)(\s*)email_verification_code(\s*)=(\s*)NULL,(\s*)email_verification_expire_date(\s*)=(\s*)NULL,(\s*)email_verification_attempt(\s*)=(\s*)NULL,(\s*)updated_at(\s*)=(\s*)NOW\(\)(\s*)WHERE(\s*)identity_code(\s*)=(\s*)(.+)`,
				lastInsertID: 1,
				rowsAffected: 1,
				err:          nil,
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				identityCode: 140810140016,
			},
			mock: mock{
				query:        `^\s*UPDATE(\s*)users(\s*)SET(\s*)status(\s*)=(\s*)(.+)(\s*)email_verification_code(\s*)=(\s*)NULL,(\s*)email_verification_expire_date(\s*)=(\s*)NULL,(\s*)email_verification_attempt(\s*)=(\s*)NULL,(\s*)updated_at(\s*)=(\s*)NOW\(\)(\s*)WHERE(\s*)identity_code(\s*)=(\s*)(.+)`,
				lastInsertID: 0,
				rowsAffected: 0,
				err:          nil,
			},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			args: args{
				identityCode: 140810140016,
			},
			mock: mock{
				query:        `^\s*UPDATE(\s*)users(\s*)SET(\s*)status(\s*)=(\s*)(.+)(\s*)email_verification_code(\s*)=(\s*)NULL,(\s*)email_verification_expire_date(\s*)=(\s*)NULL,(\s*)email_verification_attempt(\s*)=(\s*)NULL,(\s*)updated_at(\s*)=(\s*)NOW\(\)(\s*)WHERE(\s*)identity_code(\s*)=(\s*)(.+)`,
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
			if err := UpdateToVerified(tt.args.identityCode); (err != nil) != tt.wantErr {
				t.Errorf("UpdateToVerified() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChangePassword(t *testing.T) {
	type args struct {
		identityCode int64
		password     string
		oldPassword  string
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
				identityCode: 140810140016,
				password:     "f1cf8402f0fb0511a8054c697fc4bee1",
				oldPassword:  "f6ec409a28c6d93c11c056f1409ed887",
			},
			mock: mock{
				query:        `^\s*UPDATE\s*users\s*SET\s*password\s=\s\('([\w]*)'\)\s*WHERE\s*identity_code\s=\s\((\d+)\)\sAND\s*password\s=\s\('(\w*)'\);$`,
				lastInsertID: 1,
				rowsAffected: 1,
				err:          nil,
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				identityCode: 140810140016,
				password:     "f1cf8402f0fb0511a8054c697fc4bee1",
				oldPassword:  "f6ec409a28c6d93c11c056f1409ed887",
			},
			mock: mock{
				query:        `^\s*UPDATE\s*users\s*SET\s*password\s=\s\('([\w]*)'\)\s*WHERE\s*identity_code\s=\s\((\d+)\)\sAND\s*password\s=\s\('(\w*)'\);$`,
				lastInsertID: 0,
				rowsAffected: 0,
				err:          nil,
			},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			args: args{
				identityCode: 140810140016,
				password:     "f1cf8402f0fb0511a8054c697fc4bee1",
				oldPassword:  "f6ec409a28c6d93c11c056f1409ed887",
			},
			mock: mock{
				query:        `^\s*UPDATE\s*users\s*SET\s*password\s=\s\('([\w]*)'\)\s*WHERE\s*identity_code\s=\s\((\d+)\)\sAND\s*password\s=\s\('(\w*)'\);$`,
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
			if err := ChangePassword(tt.args.identityCode, tt.args.password, tt.args.oldPassword); (err != nil) != tt.wantErr {
				t.Errorf("ChangePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestForgotNewPassword(t *testing.T) {
	type args struct {
		email    string
		password string
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
				email:    "risal@live.com",
				password: "f1cf8402f0fb0511a8054c697fc4bee1",
			},
			mock: mock{
				query:        `^\s*UPDATE\s*users\s*SET\s*password\s=\s\('([\w]*)'\),\s*email_verification_code\s=\sNULL,\s*email_verification_expire_date\s=\sNULL,\s*email_verification_attempt\s=\sNULL,\s*updated_at\s=\sNOW\(\)\s*WHERE\s*email\s=\s\('([\w@.]*)'\);$`,
				lastInsertID: 1,
				rowsAffected: 1,
				err:          nil,
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				email:    "risal@live.com",
				password: "f1cf8402f0fb0511a8054c697fc4bee1",
			},
			mock: mock{
				query:        `^\s*UPDATE\s*users\s*SET\s*password\s=\s\('([\w]*)'\),\s*email_verification_code\s=\sNULL,\s*email_verification_expire_date\s=\sNULL,\s*email_verification_attempt\s=\sNULL,\s*updated_at\s=\sNOW\(\)\s*WHERE\s*email\s=\s\('([\w@.]*)'\);$`,
				lastInsertID: 0,
				rowsAffected: 0,
				err:          nil,
			},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			args: args{
				email:    "risal@live.com",
				password: "f1cf8402f0fb0511a8054c697fc4bee1",
			},
			mock: mock{
				query:        `^\s*UPDATE\s*users\s*SET\s*password\s=\s\('([\w]*)'\),\s*email_verification_code\s=\sNULL,\s*email_verification_expire_date\s=\sNULL,\s*email_verification_attempt\s=\sNULL,\s*updated_at\s=\sNOW\(\)\s*WHERE\s*email\s=\s\('([\w@.]*)'\);$`,
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
			if err := ForgotNewPassword(tt.args.email, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("ForgotNewPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerateVerification(t *testing.T) {
	type args struct {
		identity int64
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
		want    Verification
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				identity: 140810140016,
			},
			mock: mock{
				query:        `^\s*UPDATE\s*users\s*SET\s*email_verification_code\s=\s\(\d+\),\s*email_verification_expire_date\s=\s\(DATE_ADD\(NOW\(\), INTERVAL 30 MINUTE\)\),\s*email_verification_attempt\s=\sNULL,\s*updated_at\s=\sNOW\(\)\s*WHERE\s*identity_code\s=\s\(\d+\);`,
				lastInsertID: 1,
				rowsAffected: 1,
				err:          nil,
			},
			want: Verification{
				Attempt:        0,
				Code:           0,
				ExpireDate:     time.Now(),
				ExpireDuration: "30 Minutes",
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				identity: 140810140016,
			},
			mock: mock{
				query:        `^\s*UPDATE\s*users\s*SET\s*email_verification_code\s=\s\(\d+\),\s*email_verification_expire_date\s=\s\(DATE_ADD\(NOW\(\), INTERVAL 30 MINUTE\)\),\s*email_verification_attempt\s=\sNULL,\s*updated_at\s=\sNOW\(\)\s*WHERE\s*identity_code\s=\s\(\d+\);`,
				lastInsertID: 0,
				rowsAffected: 0,
				err:          nil,
			},
			want: Verification{
				Attempt:        0,
				Code:           0,
				ExpireDate:     time.Now(),
				ExpireDuration: "30 Minutes",
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
			got, err := GenerateVerification(tt.args.identity)
			tt.want.Code = got.Code
			tt.want.ExpireDate = got.ExpireDate
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateVerification() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateVerification() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateProfile(t *testing.T) {
	type args struct {
		identityCode int64
		name         string
		note         string
		phone        sql.NullString
		lineID       sql.NullString
		gender       int8
	}
	type mock struct {
		query        string
		lastInsertID int64
		rowsAffected int64
		err          error
	}
	tests := []struct {
		name    string
		mock    mock
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				identityCode: 140810140016,
				name:         "Risal Falah",
				note:         "Hello my name is risal!!@#$%^&*&^%$##$%^&*(",
				phone:        sql.NullString{Valid: true, String: "085860141146"},
				lineID:       sql.NullString{Valid: true, String: "risalfa"},
				gender:       1,
			},
			mock: mock{
				query:        `^\s*UPDATE\s*users\s*SET\s*name\s=\s\('([\w ]*)'\),\s*phone\s=\s\('\d+'\),\s*line_id\s=\s\('(.+)'\),\s*note\s=\s\('(.+),\s*gender\s=\s\(\d\),\s*updated_at\s=\sNOW\(\)\s*WHERE\s*identity_code\s=\s\(\d+\);$`,
				lastInsertID: 1,
				rowsAffected: 1,
				err:          nil,
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				identityCode: 140810140016,
				name:         "Risal Falah",
				note:         "Hello my name is risal!!@#$%^&*&^%$##$%^&*(",
				phone:        sql.NullString{Valid: true, String: "085860141146"},
				lineID:       sql.NullString{Valid: true, String: "risalfa"},
				gender:       1,
			},
			mock: mock{
				query:        `^\s*UPDATE\s*users\s*SET\s*name\s=\s\('([\w ]*)'\),\s*phone\s=\s\('\d+'\),\s*line_id\s=\s\('(.+)'\),\s*note\s=\s\('(.+),\s*gender\s=\s\(\d\),\s*updated_at\s=\sNOW\(\)\s*WHERE\s*identity_code\s=\s\(\d+\);$`,
				lastInsertID: 0,
				rowsAffected: 0,
				err:          nil,
			},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			args: args{
				identityCode: 140810140016,
				name:         "Risal Falah",
				note:         "Hello my name is risal!!@#$%^&*&^%$##$%^&*(",
				phone:        sql.NullString{Valid: true, String: "085860141146"},
				lineID:       sql.NullString{Valid: true, String: "risalfa"},
				gender:       1,
			},
			mock: mock{
				query:        `^\s*UPDATE\s*users\s*SET\s*name\s=\s\('([\w ]*)'\),\s*phone\s=\s\('\d+'\),\s*line_id\s=\s\('(.+)'\),\s*note\s=\s\('(.+),\s*gender\s=\s\(\d\),\s*updated_at\s=\sNOW\(\)\s*WHERE\s*identity_code\s=\s\(\d+\);$`,
				lastInsertID: 0,
				rowsAffected: 0,
				err:          fmt.Errorf("Error connection"),
			},
			wantErr: true,
		},
		{
			name: "Test Case 4",
			args: args{
				identityCode: 140810140016,
				name:         "Risal Falah",
				note:         "Hello my name is risal!!@#$%^&*&^%$##$%^&*(",
				phone:        sql.NullString{Valid: true, String: "085860141146"},
				lineID:       sql.NullString{Valid: true, String: "risalfa"},
				gender:       3,
			},
			mock: mock{
				query:        `^\s*UPDATE\s*users\s*SET\s*name\s=\s\('([\w ]*)'\),\s*phone\s=\s\('\d+'\),\s*line_id\s=\s\('(.+)'\),\s*note\s=\s\('(.+),\s*gender\s=\s\(\d\),\s*updated_at\s=\sNOW\(\)\s*WHERE\s*identity_code\s=\s\(\d+\);$`,
				lastInsertID: 1,
				rowsAffected: 1,
				err:          nil,
			},
			wantErr: false,
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
			if err := UpdateProfile(tt.args.identityCode, tt.args.name, tt.args.note, tt.args.phone, tt.args.lineID, tt.args.gender); (err != nil) != tt.wantErr {
				t.Errorf("UpdateProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSignUp(t *testing.T) {
	type args struct {
		identityCode int64
		name         string
		email        string
		password     string
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
				identityCode: 140810140016,
				name:         "Risal Falah",
				email:        "risal@live.com",
				password:     "2af9b1ba42dc5eb01743e6b3759b6e4b",
			},
			mock: mock{
				query:        `^\s*INSERT\sINTO\s*users\s\(\s*name,\s*email,\s*password,\s*identity_code,\s*created_at,\s*updated_at\s*\)\sVALUES\s\(\s*\('([\w ]*)'\),\s*\('([\w@.]+)'\),\s*\('([\w]+)'\),\s*\(\d+\)\s*,\s*NOW\(\),\s*NOW\(\)\s*\);$`,
				lastInsertID: 1,
				rowsAffected: 1,
				err:          nil,
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				identityCode: 140810140016,
				name:         "Risal Falah",
				email:        "risal@live.com",
				password:     "2af9b1ba42dc5eb01743e6b3759b6e4b",
			},
			mock: mock{
				query:        `^\s*INSERT\sINTO\s*users\s\(\s*name,\s*email,\s*password,\s*identity_code,\s*created_at,\s*updated_at\s*\)\sVALUES\s\(\s*\('([\w ]*)'\),\s*\('([\w@.]+)'\),\s*\('([\w]+)'\),\s*\(\d+\)\s*,\s*NOW\(\),\s*NOW\(\)\s*\);$`,
				lastInsertID: 0,
				rowsAffected: 0,
				err:          nil,
			},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			args: args{
				identityCode: 140810140016,
				name:         "Risal Falah",
				email:        "risal@live.com",
				password:     "2af9b1ba42dc5eb01743e6b3759b6e4b",
			},
			mock: mock{
				query:        `^\s*INSERT\sINTO\s*users\s\(\s*name,\s*email,\s*password,\s*identity_code,\s*created_at,\s*updated_at\s*\)\sVALUES\s\(\s*\('([\w ]*)'\),\s*\('([\w@.]+)'\),\s*\('([\w]+)'\),\s*\(\d+\)\s*,\s*NOW\(\),\s*NOW\(\)\s*\);$`,
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
			if err := SignUp(tt.args.identityCode, tt.args.name, tt.args.email, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSelectByID(t *testing.T) {
	type args struct {
		id     []int64
		column []string
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
		want    []User
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				id:     []int64{1, 2, 3},
				column: []string{ColID, ColName, ColEmail, ColPhone},
			},
			mock: mock{
				query:  `^\s*SELECT\s*(.+)\s*FROM\s*users\s*WHERE\s*id\s*IN\s*\(([\d, ]*)\);$`,
				column: []string{"id", "name", "email", "phone"},
				result: [][]driver.Value{
					[]driver.Value{
						"1", "Risal Falah", "risal@live.com", "085860141146",
					},
					[]driver.Value{
						"2", "Rifki Muhammad", "rifkirifkigue@gmail.com", "085860141146",
					},
					[]driver.Value{
						"3", "Asep Nur Muhammad Iskandar Yusuf", "asepasepgue@gmail.com", "085860141146",
					},
				},
				err: nil,
			},
			want: []User{
				User{
					ID:    1,
					Name:  "Risal Falah",
					Email: "risal@live.com",
					Phone: sql.NullString{Valid: true, String: "085860141146"},
				},
				User{
					ID:    2,
					Name:  "Rifki Muhammad",
					Email: "rifkirifkigue@gmail.com",
					Phone: sql.NullString{Valid: true, String: "085860141146"},
				},
				User{
					ID:    3,
					Name:  "Asep Nur Muhammad Iskandar Yusuf",
					Email: "asepasepgue@gmail.com",
					Phone: sql.NullString{Valid: true, String: "085860141146"},
				},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				id:     []int64{1, 2},
				column: []string{},
			},
			mock: mock{
				query:  `^\s*SELECT\s*(.+)\s*FROM\s*users\s*WHERE\s*id\s*IN\s*\(([\d, ]*)\);$`,
				column: []string{"id", "name", "email", "gender", "note", "status", "identity_code", "line_id", "phone", "rolegroups_id"},
				result: [][]driver.Value{
					[]driver.Value{"1", "Risal Falah", "risal@live.com", "1", "", "2", "140810140016", nil, nil, nil},
				},
				err: nil,
			},
			want: []User{
				User{
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
			},
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				id:     []int64{1, 2},
				column: []string{},
			},
			mock: mock{
				query:  `^\s*SELECT\s*(.+)\s*FROM\s*users\s*WHERE\s*id\s*IN\s*\(([\d, ]*)\);$`,
				column: []string{"id", "name", "email", "gender", "note", "status", "identity_code", "line_id", "phone", "rolegroups_id"},
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
			rows := sqlmock.NewRows(tt.mock.column)
			for _, val := range tt.mock.result {
				rows.AddRow(val...)
			}
			q.WillReturnRows(rows)
		} else {
			q.WillReturnError(tt.mock.err)
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectByID(tt.args.id, tt.args.column...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
