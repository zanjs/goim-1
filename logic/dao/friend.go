package dao

import (
	"goim/logic/entity"
	"goim/logic/lib/session"
	"log"
)

type FriendDao struct {
	base
}

func NewFriendDao(session *session.Session) *FriendDao {
	return &FriendDao{base{session}}
}

// Add 插入一条朋友关系
func (d *FriendDao) Add(friend entity.Friend) error {
	_, err := d.session.Exec("insert ignore into t_friend(user_id,friend,label) values(?,?,?)", friend.UserId, friend.Friend, friend.Label)
	if err != nil {
		log.Println(err)
	}
	return err
}

// Delete 删除一条朋友关系
func (d *FriendDao) Delete(userId, friend int) error {
	_, err := d.session.Exec("delete from t_friend where user_id = ? and friend = ? ", userId, friend)
	if err != nil {
		log.Println(err)
	}
	return err
}

// ListFriends 获取用户的朋友列表
func (d *FriendDao) ListUserFriend(id int) ([]entity.FriendUser, error) {
	rows, err := d.session.Query("select f.label,u.id,u.number,u.name,u.sex,u.img from t_friend f left join t_user u on f.friend = u.id where f.user_id = ?", id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	users := make([]entity.FriendUser, 0, 5)
	for rows.Next() {
		var user entity.FriendUser
		err := rows.Scan(&user.Label, &user.UserId, &user.Number, &user.Name, &user.Sex, &user.Img)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
