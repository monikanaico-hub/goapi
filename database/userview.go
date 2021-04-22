package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/monikanaico-hub/goapi/models"
)

func CreateUser(user models.User) (int64, error) {
	db, err := DbConnection()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	query := "INSERT INTO user (username, password, firstname) VALUES(?, ?, ?);"
	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(user.Username, user.Password, user.Firstname)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, err
}

func GetallUser() ([]models.User, error) {
	db, err := DbConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	query := "Select * from user;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var userrslt []models.User
	for rows.Next() {
		user := models.User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Firstname)
		userrslt = append(userrslt, user)
	}
	err = rows.Err()
	if err != nil {
		return userrslt, err
	}
	return userrslt, err

}

func FindById(id int64) (models.User, error) {
	var userList models.User
	db, err := DbConnection()
	if err != nil {
		return userList, err
	}
	defer db.Close()

	query := "SELECT iduser, username, password, firstname FROM user WHERE iduser = ?;"

	row := db.QueryRow(query, id)
	row.Scan(&userList.Id, &userList.Username, &userList.Password, &userList.Firstname)

	return userList, nil
}

func UpdateUser(id int64, userList models.User) (models.User, error) {
	db, err := DbConnection()
	if err != nil {
		return userList, err
	}
	defer db.Close()

	query := "UPDATE user SET username = ?, password = ?, firstname = ? WHERE iduser = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return userList, err
	}

	_, queryErr := stmt.Exec(userList.Username, userList.Password, userList.Firstname, id)
	if queryErr != nil {
		return userList, queryErr
	}

	userList.Id = id
	return userList, queryErr
}

func DeleteUser(userList models.User) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "DELETE FROM user WHERE iduser = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	_, queryErr := stmt.Exec(userList.Id)

	return queryErr
}
