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

// Insert 插入朋友一条朋友关系
func (d *FriendDao) Add(friend entity.Friend) error {
	_, err := d.session.Exec("insert into t_friend(user_id,friend,label) values(?,?,?)", friend.UserId, friend.Friend, friend.Label)
	if err != nil {
		log.Println(err)
	}
	return err
}

// GetFriends 获取用户的朋友列表
func (d *FriendDao) ListFriends(id int) ([]*entity.User, error) {
	rows, err := d.session.Query("select f.lable,u.id,u.number,u.name,u.sex,u.img from t_friend f left join t_user u on f.friend = u.id where f.id = ?", id)
	if err != nil {
		log.Println(err)
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
