package dao

import (
	"goim/logic/entity"
	"goim/logic/lib/session"
	"log"
)

type GroupUserDao struct {
	base
}

func NewGroupUserDao(session *session.Session) *GroupUserDao {
	return &GroupUserDao{base{session}}
}

func (d *GroupUserDao) Get(id int) (*entity.Group, error) {
	row := d.session.QueryRow("select id,name from t_user where id = ?", id)
	var group entity.Group
	err := row.Scan(&group.Id, &group.Name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &group, nil
}

// ListGroupUser 获取群组用户信息
func (d *GroupUserDao) ListGroupUser(id int) ([]entity.GroupUser, error) {
	sql := `select g.label,u.number,u.name,u.sex,u.img from t_group_user g left join t_user u on g.user_id = u.id where group_id = ?`
	rows, err := d.session.Query(sql, id)
	if err != nil {
		return nil, err
	}
	groupUsers := make([]entity.GroupUser, 0, 5)
	for rows.Next() {
		var groupUser entity.GroupUser
		err := rows.Scan(&groupUser.Label, &groupUser.Number, &groupUser.Name, &groupUser.Sex, &groupUser.Img)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		groupUsers = append(groupUsers, groupUser)
	}
	return groupUsers, nil
}

// ListGroupUserId 获取群组用户id列表
func (d *GroupUserDao) ListGroupUserId(id int) ([]int, error) {
	rows, err := d.session.Query("select user_id t_group_user where group_id = ?", id)
	if err != nil {
		return nil, err
	}
	userIds := make([]int, 0, 5)
	for rows.Next() {
		var userId int
		err := rows.Scan(&userId)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		userIds = append(userIds, userId)
	}
	return userIds, nil
}

// ListByUser 获取用户群组id列表
func (d *GroupUserDao) ListbyUserId(userId int) ([]int, error) {
	rows, err := d.session.Query("select group_id from t_group_user where user_id = ?", userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var ids []int
	var id int
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

// Add 将用户添加到群组
func (d *GroupUserDao) Add(groupId int, userId int) error {
	_, err := d.session.Exec("insert ignore into t_group_user(group_id,user_id) values(?,?)", groupId, userId)
	if err != nil {
		log.Println(err)
	}
	return err
}

// Delete 将用户从群组删除
func (d *GroupUserDao) Delete(groupId int, userId int) error {
	_, err := d.session.Exec("delete from t_group_user where group_id = ? and user_id = ?", groupId, userId)
	if err != nil {
		log.Println(err)
	}
	return err
}

// UpdateLabel 更新用户群组备注
func (d *GroupUserDao) UpdateLabel(groupId int, userId int, label string) error {
	_, err := d.session.Exec("update t_group_user set label = ? where group_id = ? and user_id = ?", label, groupId, userId)
	if err != nil {
		log.Println(err)
	}
	return err
}
