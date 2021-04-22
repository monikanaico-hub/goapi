package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/monikanaico-hub/goapi/configaration"
)

func dsn(db string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", configaration.Username, configaration.Password, configaration.Hostname, configaration.Dbname)
}

func Dberror(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func DbConnection() *sql.DB {
	db, err := sql.Open("mysql", dsn(configaration.Dbname))
	Dberror(err)
	return db
}
