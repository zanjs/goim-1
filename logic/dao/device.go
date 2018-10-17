package dao

import (
	"goim/lib/context"
	"goim/logic/entity"
	"log"
)

type deviceDao struct{}

var DeviceDao = new(deviceDao)

// Insert 插入一条设备信息
func (*deviceDao) Add(ctx *context.Context, device entity.Device) (int64, error) {
	result, err := ctx.Session.Exec("insert into t_device(token,type,model,version) values(?,?,?,?)", device.Token, device.Type, device.Model, device.Version)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetToken 获取设备的token
func (*deviceDao) GetToken(ctx *context.Context, id int64) (string, error) {
	var token string
	row := ctx.Session.QueryRow("select token from t_device where id = ? ", id)
	err := row.Scan(&token)
	if err != nil {
		log.Println(err)
	}
	return token, err
}

// UpdateUserId 更新设备绑定用户
func (*deviceDao) UpdateUserId(ctx *context.Context, id, userId int64) error {
	_, err := ctx.Session.Exec("update t_device set user_id = ? where id = ? ", userId, id)
	if err != nil {
		log.Println(err)
	}
	return nil
}

// UpdateUserId 更新设备绑定用户
func (*deviceDao) UpdateStatus(ctx context.Context, id, status int) error {
	_, err := ctx.Session.Exec("update t_device set status = ? where id = ? ", status, id)
	if err != nil {
		log.Println(err)
	}
	return nil
}

// ListUserOnline 查询用户所有的在线设备
func (*deviceDao) ListOnlineByUserId(ctx *context.Context, userId int) ([]*entity.Device, error) {
	rows, err := ctx.Session.Query("select id,type,model,version from t_device where user_id = ? and status = 1", userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	devices := make([]*entity.Device, 0, 5)
	for rows.Next() {
		device := new(entity.Device)
		err = rows.Scan(&device.Id, &device.Type, &device.Model, &device.Version)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		devices = append(devices, device)
	}
	return devices, nil

}
