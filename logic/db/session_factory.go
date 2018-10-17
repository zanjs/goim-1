package db

import "goim/lib/session"

var Factoty *session.SessionFactory

func init() {
	var err error
	Factoty, err = session.NewSessionFactory("mysql", "root:Liu123456@tcp(localhost:3306)/im?charset=utf8")
	if err != nil {
		panic(err)
	}
}
