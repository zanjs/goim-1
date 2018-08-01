package dao

import (
	"goim/logic/entity"
	"goim/logic/lib/session"
	"log"
)

type GroupDao struct {
	base
}

func NewGroupDao(session *session.Session) *GroupDao {
	return &GroupDao{base{session}}
}

// Get 获取群组信息
func (d *GroupDao) Get(id int) (*entity.Group, error) {
	row := d.session.QueryRow("select name from t_group where id = ?", id)
	group := new(entity.Group)
	err := row.Scan(&group.Name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return group, nil
}

// Insert 插入一条群组信息
func (d *GroupDao) Add(name string) (int, error) {
	result, err := d.session.Exec("insert into t_group(name) value(?)", name)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return int(id), nil
}
