package dao

import (
	"goim/logic/entity"
	"goim/logic/lib/session"
	"log"
)

type DeviceDao struct {
	session.Sessioner
}

// Insert 插入一条设备信息
func (d *DeviceDao) Insert(device entity.Device) (int, error) {
	result, err := d.Exec("insert into t_device(token,type,model,version) values(?,?,?,?)", device.Token, device.Model, device.Model, device.Version)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// UpdateUserIdAndStatus 更新设备绑定用户和在线状态
func (d *DeviceDao) UpdateUserIdAndStatus(id, userId, status int) error {
	_, err := d.Exec("update t_device set user_id = ?,status = ? where id = ? ", userId, status, id)
	if err != nil {
		log.Println(err)
	}
	return nil
}

// ListUserOnline 查询用户所有的在线设备
func (d *DeviceDao) ListOnlineByUserId(userId int) ([]*entity.Device, error) {
	rows, err := d.Query("select id,type,model,version from t_device where user_id = ? and status = 1", userId)
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
