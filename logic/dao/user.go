package dao

import (
	"goim/logic/entity"
	"goim/logic/lib/session"
	"log"
)

type UserDao struct {
	base
}

func NewUserDao(session *session.Session) *UserDao {
	return &UserDao{base{session}}
}

// Insert 插入一条用户信息
func (d *UserDao) Insert(user entity.User) (int, error) {
	result, err := d.session.Exec("insert into t_user(number,name,password,sex，img) valus(?,?,?,?)",
		user.Number, user.Name, user.Password, user.Sex, user.Img)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return int(id), nil
}

// Get 获取用户信息
func (d *UserDao) Get(id int) (*entity.User, error) {
	row := d.session.QueryRow("select number,name,password,sex,img from t_user where id = ?", id)
	user := new(entity.User)
	err := row.Scan(&user.Number, &user.Name, &user.Password, &user.Sex, &user.Img)
	if err != nil {
		log.Println(err)
	}
	return user, err
}

// UpdatePassword 更新用户密码
func (d *UserDao) UpdatePassword(id int) (*entity.User, error) {
	row := d.session.QueryRow("select number,name,password,sex from t_user where id = ?", id)
	user := new(entity.User)
	err := row.Scan(&user.Number, &user.Name, &user.Password, &user.Sex)
	if err != nil {
		log.Println(err)
	}
	return user, err
}
