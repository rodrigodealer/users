package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLConn interface {
	Connect()
	Ping() (bool, error)
}

type MySQLConnection struct {
	Conn *sql.DB
}

func (m *MySQLConnection) Connect() {
	var connectionString = "root:root@/nutripad-rest-api_production?charset=utf8"
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Printf("Failed to connect to MySQL: (%s)\n %s", connectionString, err.Error())
	}
	m.Conn = db
}

func (m *MySQLConnection) Ping() (bool, error) {
	var err = m.Conn.Ping()
	var working = true
	if err != nil {
		working = false
		log.Printf("Failed to ping MySQL: \n %s", err.Error())
	}
	return working, err
}
