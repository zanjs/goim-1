package dao

import "goim/logic/lib/session"

type MessageDao struct {
	session.Sessioner
}

func (d *MessageDao) List(userId int) {

}
