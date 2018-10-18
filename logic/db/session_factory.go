package db

import (
	"goim/public/session"

	_ "github.com/go-sql-driver/mysql"
)

var Factoty *session.SessionFactory

func init() {
	var err error
	Factoty, err = session.NewSessionFactory("mysql", "root:Liu123456@tcp(localhost:3306)/im?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
}
