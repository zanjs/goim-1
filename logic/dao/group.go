package dao

import (
	"goim/logic/entity"
	"goim/logic/lib/session"
	"log"
)

type GroupDao struct {
	session.Sessioner
}

// Insert 插入一条群组信息
func (d *GroupDao) Insert(group entity.Group) (int, error) {
	result, err := d.Exec("insert into t_group(name) value(?)", group.Name)
	if err != nil {
		log.Println(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return int(id), nil
}

// Get 获取群组信息
func (d *GroupDao) Get(id int) (*entity.Group, error) {
	row := d.QueryRow("select name from t_group where id = ?", id)
	group := new(entity.Group)
	err := row.Scan(&group.Name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return group, nil
}

// GetGroupUser 获取群组人员信息
func (d *GroupDao) GetGroupUser(id int) ([]*entity.User, error) {
	sql := `select u.number,u.name,u.sex,u.img from t_group g left join t_user u on g.user_id = u.id where id = ?`
	rows, err := d.Query(sql, id)
	if err != nil {
		return nil, err
	}
	users := make([]*entity.User, 0, 5)
	for rows.Next() {
		user := new(entity.User)
		err := rows.Scan(&user.Number, &user.Name, &user.Sex, &user.Img)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
