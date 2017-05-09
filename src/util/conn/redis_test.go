package conn

import (
	"reflect"
	"testing"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/rafaeljusto/redigomock"
)

func initRedisMock() *redigomock.Conn {
	conn := redigomock.NewConn()
	Redis = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 10 * time.Second,
		Dial:        func() (redis.Conn, error) { return conn, nil },
	}
	return conn
}

func TestInitRedisMock(t *testing.T) {
	cases := []struct {
		command        string
		key            string
		expectedResult string
		expectedErr1   error
		expectedErr2   error
	}{
		{
			command:        "GET",
			key:            "test",
			expectedResult: "success",
			expectedErr1:   nil,
			expectedErr2:   nil,
		},
	}

	modConn := initRedisMock()

	for _, val := range cases {
		cmd := modConn.Command(val.command, val.key).Expect(val.expectedResult)

		r, err := redis.String(modConn.Do(val.command, val.key))
		if !reflect.DeepEqual(err, val.expectedErr1) {
			t.Errorf("expect: %v got: %v", val.expectedErr1, err)
			t.FailNow()
		}

		if r != val.expectedResult {
			t.Errorf("expect %v got %v", val.expectedResult, r)
			t.FailNow()
		}

		if modConn.Stats(cmd) != 1 {
			t.Errorf("Command was not used")
			t.FailNow()
		}
	}

}
