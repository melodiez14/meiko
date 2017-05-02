package conn

import (
	"testing"

	"github.com/jmoiron/sqlx"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func initDBMock() (*sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, err
	}

	DB = sqlx.NewDb(db, "sqlmock")

	return &mock, nil
}

func TestInitDB(t *testing.T) {
	_, err := initDBMock()
	if err != nil {
		t.Errorf("Failed to InitMockDB")
	}

	cases := []struct {
		db       *sqlx.DB
		expected error
	}{
		{
			db:       DB,
			expected: nil,
		},
	}
	for _, val := range cases {
		if got := val.db.Ping(); got != val.expected {
			t.Errorf("Get() = %v, expected %v", got, val.expected)
		}
	}
}
