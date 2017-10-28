package conn

import (
	"fmt"
	"log"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	// mysql module used for sqlx purpose
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DatabaseConfig is used for the setting of database
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Driver   string `json:"driver"`
	Database string `json:"database"`
}

// DB is connection to database
var DB *sqlx.DB

// InitDB is used for Initialize setting of database
func InitDB(cfg DatabaseConfig) {
	log.Println("Initializing database")

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := sqlx.Connect(cfg.Driver, connString)
	if err != nil {
		log.Fatalln("Error to connect database")
		return
	}

	DB = db
	log.Println("Database successfully connected")
}

// InitDBMock used for initialize database connection mocking
func InitDBMock() (sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, err
	}

	DB = sqlx.NewDb(db, "sqlmock")

	return mock, nil
}
