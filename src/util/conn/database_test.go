package conn

import (
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestInitDB(t *testing.T) {
	_, err := InitDBMock()
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

func TestInitDBMock(t *testing.T) {
	_, err := InitDBMock()
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
