package dao

import (
	"goim/logic/model"
	"goim/public/ctx"
	"goim/public/logger"
)

type userDao struct{}

var UserDao = new(userDao)

// Add 插入一条用户信息
func (*userDao) Add(ctx *ctx.Context, user model.User) (int64, error) {
	result, err := ctx.Session.Exec("insert ignore into t_user(number,name,sex,img,password) values(?,?,?,?,?)",
		user.Number, user.Name, user.Sex, user.Avatar, user.Password)
	if err != nil {
		logger.Sugaer.Error(err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Sugaer.Error(err)
		return 0, err
	}
	return id, nil
}

// Get 获取用户信息
func (*userDao) Get(ctx *ctx.Context, id int) (*model.User, error) {
	row := ctx.Session.QueryRow("select number,name,password,sex,img from t_user where id = ?", id)
	user := new(model.User)
	err := row.Scan(&user.Number, &user.Name, &user.Password, &user.Sex, &user.Avatar)
	if err != nil {
		logger.Sugaer.Error(err)
		return nil, err
	}
	return user, err
}

// Get 获取用户信息
func (*userDao) GetPassword(ctx *ctx.Context, id int64) (string, error) {
	row := ctx.Session.QueryRow("select password from t_user where id = ?", id)
	var password string
	err := row.Scan(&password)
	if err != nil {
		logger.Sugaer.Error(err)
		return "", err
	}
	return password, nil
}

// UpdatePassword 更新用户密码
func (*userDao) UpdatePassword(ctx *ctx.Context, id int) (*model.User, error) {
	row := ctx.Session.QueryRow("select number,name,password,sex from t_user where id = ?", id)
	user := new(model.User)
	err := row.Scan(&user.Number, &user.Name, &user.Password, &user.Sex)
	if err != nil {
		logger.Sugaer.Error(err)
		return nil, err
	}
	return user, nil
}
