package views

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/monikanaico-hub/goapi/database"
	"github.com/monikanaico-hub/goapi/models"
)

func CreateUser(user models.User) (int64, error) {
	db := database.DbConnection()
	defer db.Close()

	query := "INSERT INTO user (username, password, firstname) VALUES(?, ?, ?);"
	stmt, stmtErr := db.Prepare(query)
	database.Dberror(stmtErr)

	res, queryErr := stmt.Exec(user.Username, user.Password, user.Firstname)
	database.Dberror(queryErr)

	id, getLastInsertIdErr := res.LastInsertId()
	database.Dberror(getLastInsertIdErr)

	return id, queryErr
}

func GetallUser() ([]models.User, error) {
	db := database.DbConnection()
	defer db.Close()
	query := "Select * from user;"
	rows, queryErr := db.Query(query)
	database.Dberror(queryErr)

	defer rows.Close()
	var userrslt []models.User
	for rows.Next() {
		user := models.User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Firstname)
		userrslt = append(userrslt, user)
	}
	queryErr = rows.Err()
	database.Dberror(queryErr)

	return userrslt, queryErr

}
