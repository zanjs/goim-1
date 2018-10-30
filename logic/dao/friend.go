package dao

import (
	"goim/logic/model"
	"goim/public/ctx"
	"goim/public/logger"
)

type friendDao struct{}

var FriendDao = new(friendDao)

// Get 获取一个朋友关系
func (*friendDao) Get(ctx *ctx.Context, userId int64, friendId int64) (*model.Friend, error) {
	var friend model.Friend
	row := ctx.Session.QueryRow(`select id,user_id,friend_id,lable,create_time,update_time 
		from t_friend where user_id = ? and friend_id = ?)`,
		userId, friendId)
	err := row.Scan(&friend.Id, &friend.UserId, &friend.Id, &friend.Label, &friend.CreateTime, &friend.UpdateTime)
	if err != nil {
		logger.Sugaer.Error(err)
		return nil, err
	}
	return &friend, nil
}

// Add 插入一条朋友关系
func (*friendDao) Add(ctx *ctx.Context, friend model.Friend) error {
	_, err := ctx.Session.Exec("insert ignore into t_friend(user_id,friend_id,label) values(?,?,?)",
		friend.UserId, friend.FriendId, friend.Label)
	if err != nil {
		logger.Sugaer.Error(err)
	}
	return err
}

// Delete 删除一条朋友关系
func (*friendDao) Delete(ctx *ctx.Context, userId, friend int) error {
	_, err := ctx.Session.Exec("delete from t_friend where user_id = ? and friend_id = ? ",
		userId, friend)
	if err != nil {
		logger.Sugaer.Error(err)
	}
	return err
}

// ListFriends 获取用户的朋友列表
func (*friendDao) ListUserFriend(ctx *ctx.Context, id int) ([]model.FriendUser, error) {
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
