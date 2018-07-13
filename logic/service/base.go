package service

import (
	"goim/logic/lib/session"

	_ "github.com/go-sql-driver/mysql"
)

var sf *session.SessionFactory

func init() {
	var err error
	sf, err = session.NewSessionFactory("mysql", "root:Liu123456@tcp(localhost:3306)/im?charset=utf8")
	if err != nil {
		panic(err)
	}
}

func SetSession(service *session.Sessioner, sessioner ...session.Sessioner) {
	if len(sessioner) > 0 {
		service = &sessioner[0]
	} else {
		*service = sf.GetSession()
	}
}
