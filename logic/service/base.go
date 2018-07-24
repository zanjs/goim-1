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

type baseService struct {
	session *session.Session
}

func (s *baseService) setSession(session ...*session.Session) {
	if len(session) > 0 {
		s.session = session[0]
	} else {
		s.session = sf.GetSession()
	}
}
