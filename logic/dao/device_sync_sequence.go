package dao

import (
	"goim/public/ctx"
	"goim/public/logger"
)

type deviceSyncSequenceDao struct{}

var DeviceSyncSequenceDao = new(deviceSyncSequenceDao)

// GetSeq 获取设备已经同步的消息序列号
func (*deviceSyncSequenceDao) Add(ctx *ctx.Context, deviceId int64, syncSequence int64) error {
	_, err := ctx.Session.Exec("insert into t_device_seq(device_id,sync_sequence) values(?,?)",
		deviceId)
	if err != nil {
		logger.Sugaer.Error(err)
		return err
	}
	return nil
}

// Get 获取设备已经同步的消息序列号
func (*deviceSyncSequenceDao) Get(ctx *ctx.Context, id int) (int, error) {
	row := ctx.Session.QueryRow("select sync_sequence from t_device_sync_sequence where device_id = ?", id)
	var syncSeq int
	err := row.Scan(&syncSeq)
	if err != nil {
		logger.Sugaer.Error(err)
		return 0, err
	}
	return syncSeq, nil
}

// UpdateSequence 更新设备已经同步的消息序列号
func (*deviceSyncSequenceDao) UpdateSequence(ctx *ctx.Context, deviceId int64, sequence int64) error {
	_, err := ctx.Session.Exec("update t_device_sync_sequence set sync_sequence = ? where device_id = ?",
		sequence, deviceId)
	if err != nil {
		logger.Sugaer.Error(err)
		return err
	}
	return nil
}
