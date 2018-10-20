package dao

import (
	"goim/public/context"
	"log"
)

type deviceSendSequenceDao struct{}

var DeviceSendSequenceDao = new(deviceSendSequenceDao)

// Add 添加设备发送序列号
func (*deviceSendSequenceDao) Add(ctx *context.Context, deviceId int64, sendSequence int64) error {
	_, err := ctx.Session.Exec("insert into t_device_send_sequence(device_id,send_sequence) values(?,?)",
		deviceId, sendSequence)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Get 获取设备已经发送的同步序列号
func (*deviceSendSequenceDao) Get(ctx *context.Context, id int64) (int64, error) {
	row := ctx.Session.QueryRow("select send_sequence from t_device_send_sequence where device_id = ?", id)
	var syncSeq int64
	err := row.Scan(&syncSeq)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return syncSeq, nil
}

// UpdateSequence 更新设备已经发送的消息序列号
func (*deviceSendSequenceDao) UpdateSequence(ctx *context.Context, deviceId int64, sequence int64) error {
	_, err := ctx.Session.Exec("update t_device_send_sequence set send_sequence = ? where device_id = ?",
		sequence, deviceId)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
