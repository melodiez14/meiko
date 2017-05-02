package conn

import (
	"fmt"
	"log"

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
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := sqlx.Connect(cfg.Driver, connString)
	if err != nil {
		log.Fatalln("Error to connect database")
		return
	}
	DB = db
}
