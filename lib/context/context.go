package context

import "goim/lib/session"

type Context struct {
	Session *session.Session
}

func NewContext(Session *session.Session) *Context {
	return &Context{Session: Session}
}
