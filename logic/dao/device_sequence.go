package dao

import (
	"goim/public/context"
	"log"
)

type deviceSequenceDao struct{}

var DeviceSequenceDao = new(deviceSequenceDao)

// GetSeq 获取设备已经同步的消息序列号
func (*deviceSequenceDao) Add(ctx *context.Context, deviceId int64) error {
	_, err := ctx.Session.Exec("insert into t_device_seq(device_id) values(?)", deviceId)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// GetSeq 获取设备已经同步的消息序列号
func (*deviceSequenceDao) GetSeq(ctx *context.Context, id int) (int, error) {
	row := ctx.Session.QueryRow("select sync_seq from t_device_seq where device_id = ?", id)
	var syncSeq int
	err := row.Scan(&syncSeq)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return syncSeq, nil
}

// UpdateSeq 更新设备已经同步的消息序列号
func (*deviceSequenceDao) UpdateSeq(ctx *context.Context, deviceId int, syncSeq int) error {
	_, err := ctx.Session.Exec("update t_device_seq set sync_seq = ? where device_id = ?", syncSeq, deviceId)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
