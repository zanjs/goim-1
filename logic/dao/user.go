package dao

import (
	"goim/lib/context"
	"goim/logic/entity"
	"log"
)

type userDao struct{}

var UserDao = new(userDao)

// Add 插入一条用户信息
func (*userDao) Add(ctx *context.Context, user entity.User) (int64, error) {
	result, err := ctx.Session.Exec("insert ignore into t_user(number,name,sex,img,password) values(?,?,?,?,?)",
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
	return id, nil
}

// Get 获取用户信息
func (*userDao) Get(ctx *context.Context, id int) (*entity.User, error) {
	row := ctx.Session.QueryRow("select number,name,password,sex,img from t_user where id = ?", id)
	user := new(entity.User)
	err := row.Scan(&user.Number, &user.Name, &user.Password, &user.Sex, &user.Img)
	if err != nil {
		log.Println(err)
	}
	return user, err
}

// Get 获取用户信息
func (*userDao) GetPassword(ctx *context.Context, id int64) (string, error) {
	row := ctx.Session.QueryRow("select password from t_user where id = ?", id)
	var password string
	err := row.Scan(&password)
	if err != nil {
		log.Println(err)
	}
	return password, err
}

// UpdatePassword 更新用户密码
func (*userDao) UpdatePassword(ctx *context.Context, id int) (*entity.User, error) {
	row := ctx.Session.QueryRow("select number,name,password,sex from t_user where id = ?", id)
	user := new(entity.User)
	err := row.Scan(&user.Number, &user.Name, &user.Password, &user.Sex)
	if err != nil {
		log.Println(err)
	}
	return user, err
}
