package dao

import (
	"goim/logic/model"
	"goim/public/ctx"
	"goim/public/logger"
)

type deviceDao struct{}

var DeviceDao = new(deviceDao)

// Insert 插入一条设备信息
func (*deviceDao) Add(ctx *ctx.Context, device model.Device) (int64, error) {
	result, err := ctx.Session.Exec("insert into t_device(token,type,model,version) values(?,?,?,?)",
		device.Token, device.Type, device.Model, device.Version)
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

// Get 获取设备
func (*deviceDao) Get(ctx *ctx.Context, id int64) (*model.Device, error) {
	device := model.Device{Id: id}
	row := ctx.Session.QueryRow("select user_id,token,type,model,version,status,create_time,update_time "+
		"from t_device where id = ? ", id)
	err := row.Scan(&device.UserId, &device.Token, &device.Type, &device.Model, &device.Version,
		&device.Status, &device.CreateTime, &device.UpdateTime)
	if err != nil {
		logger.Sugaer.Error(err)
	}
	return &device, err
}

// UpdateUserId 更新设备绑定用户
func (*deviceDao) UpdateUserId(ctx *ctx.Context, id, userId int64) error {
	_, err := ctx.Session.Exec("update t_device set user_id = ? where id = ? ", userId, id)
	if err != nil {
		logger.Sugaer.Error(err)
		return err
	}
	return nil
}

// UpdateStatus 更新设备的在线状态
func (*deviceDao) UpdateStatus(ctx *ctx.Context, id int64, status int) error {
	_, err := ctx.Session.Exec("update t_device set status = ? where id = ? ", status, id)
	if err != nil {
		logger.Sugaer.Error(err)
		return err
	}
	return nil
}

// ListUserOnline 查询用户所有的在线设备
func (*deviceDao) ListOnlineByUserId(ctx *ctx.Context, userId int64) ([]*model.Device, error) {
	rows, err := ctx.Session.Query("select id,type,model,version from t_device where user_id = ? and status = 1",
		userId)
	if err != nil {
		logger.Sugaer.Error(err)
		return nil, err
	}

	devices := make([]*model.Device, 0, 5)
	for rows.Next() {
		device := new(model.Device)
		err = rows.Scan(&device.Id, &device.Type, &device.Model, &device.Version)
		if err != nil {
			logger.Sugaer.Error(err)
			return nil, err
		}
		devices = append(devices, device)
	}
	return devices, nil
}
