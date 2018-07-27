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
func (d *UserDao) Add(user entity.User) (int, error) {
	result, err := d.session.Exec("insert ignore into t_user(number,name,sex,img,password) values(?,?,?,?,?)",
		user.Number, user.Name, user.Sex, user.Img, user.Password)
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

// Get 获取用户信息
func (d *UserDao) GetPassword(id int) (string, error) {
	row := d.session.QueryRow("select password from t_user where id = ?", id)
	var password string
	err := row.Scan(&password)
	if err != nil {
		log.Println(err)
	}
	return password, err
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
