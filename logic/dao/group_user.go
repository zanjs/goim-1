package dao

import (
	"goim/logic/model"
	"goim/public/context"
	"log"
)

type groupUserDao struct{}

var GroupUserDao = new(groupUserDao)

func (*groupUserDao) Get(ctx *context.Context, id int64) (*model.Group, error) {
	row := ctx.Session.QueryRow("select id,name from t_group where id = ?", id)
	var group model.Group
	err := row.Scan(&group.Id, &group.Name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &group, nil
}

// ListGroupUser 获取群组用户信息
func (*groupUserDao) ListGroupUser(ctx *context.Context, id int64) ([]model.GroupUser, error) {
	sql := `select g.label,u.id,u.number,u.name,u.sex,u.avatar from t_group_user g left join t_user u on g.user_id = u.id where group_id = ?`
	rows, err := ctx.Session.Query(sql, id)
	if err != nil {
		return nil, err
	}
	groupUsers := make([]model.GroupUser, 0, 5)
	for rows.Next() {
		var groupUser model.GroupUser
		err := rows.Scan(&groupUser.Label, &groupUser.UserId, &groupUser.Number, &groupUser.Name, &groupUser.Sex, &groupUser.Img)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		groupUsers = append(groupUsers, groupUser)
	}
	return groupUsers, nil
}

// ListGroupUserId 获取群组用户id列表
func (*groupUserDao) ListGroupUserId(ctx *context.Context, id int) ([]int, error) {
	rows, err := ctx.Session.Query("select user_id t_group_user where group_id = ?", id)
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
func (*groupUserDao) ListbyUserId(ctx *context.Context, userId int) ([]int64, error) {
	rows, err := ctx.Session.Query("select group_id from t_group_user where user_id = ?", userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var ids []int64
	var id int64
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
func (*groupUserDao) Add(ctx *context.Context, groupId int64, userId int64) error {
	_, err := ctx.Session.Exec("insert ignore into t_group_user(group_id,user_id) values(?,?)", groupId, userId)
	if err != nil {
		log.Println(err)
	}
	return err
}

// Delete 将用户从群组删除
func (d *groupUserDao) Delete(ctx *context.Context, groupId int64, userId int64) error {
	_, err := ctx.Session.Exec("delete from t_group_user where group_id = ? and user_id = ?", groupId, userId)
	if err != nil {
		log.Println(err)
	}
	return err
}

// UpdateLabel 更新用户群组备注
func (*groupUserDao) UpdateLabel(ctx *context.Context, groupId int, userId int, label string) error {
	_, err := ctx.Session.Exec("update t_group_user set label = ? where group_id = ? and user_id = ?", label, groupId, userId)
	if err != nil {
		log.Println(err)
	}
	return err
}
