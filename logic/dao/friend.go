package dao

import (
	"goim/logic/model"
	"goim/public/context"
	"goim/public/logger"
)

type friendDao struct{}

var FriendDao = new(friendDao)

// Add 插入一条朋友关系
func (*friendDao) Add(ctx *context.Context, friend model.Friend) error {
	_, err := ctx.Session.Exec("insert ignore into t_friend(user_id,friend_id,label) values(?,?,?)",
		friend.UserId, friend.FriendId, friend.Label)
	if err != nil {
		logger.Sugaer.Error(err)
	}
	return err
}

// Delete 删除一条朋友关系
func (*friendDao) Delete(ctx *context.Context, userId, friend int) error {
	_, err := ctx.Session.Exec("delete from t_friend where user_id = ? and friend_id = ? ",
		userId, friend)
	if err != nil {
		logger.Sugaer.Error(err)
	}
	return err
}

// ListFriends 获取用户的朋友列表
func (*friendDao) ListUserFriend(ctx *context.Context, id int) ([]model.FriendUser, error) {
	rows, err := ctx.Session.Query("select f.label,u.id,u.number,u.name,u.sex,u.img from t_friend f left join "+
		"t_user u on f.friend = u.id where f.user_id = ?", id)
	if err != nil {
		logger.Sugaer.Error(err)
		return nil, err
	}

	users := make([]model.FriendUser, 0, 5)
	for rows.Next() {
		var user model.FriendUser
		err := rows.Scan(&user.Label, &user.UserId, &user.Number, &user.Name, &user.Sex, &user.Img)
		if err != nil {
			logger.Sugaer.Error(err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
