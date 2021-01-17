package mysqldb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	user     string = "root"
	password string = "s0meG0od@ndStr0ngP@ssWird"
	driver   string = "mysql"
)

func InitDBConnection(dbName string) (*sql.DB, error) {
	db, err := sql.Open(driver, getDSN(dbName))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getDSN(dbName string) string {
	return fmt.Sprintf("%s:%s@/%s", user, password, dbName)
}
