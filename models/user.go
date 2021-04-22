package models

type User struct {
	Id        int64  `db:"iduser" json:"id"`
	Username  string `db:"username" json:"username"`
	Password  string `db:"password" json:"password"`
	Firstname string `db:"firstname" json:"firstname"`
}
