package dao

import (
	"goim/logic/lib/session"
	"log"
)

type DeviceSeqDao struct {
	base
}

func NewDeviceSeqDao(session *session.Session) *DeviceSeqDao {
	return &DeviceSeqDao{base{session}}
}

// GetSeq 获取设备已经同步的消息序列号
func (d *DeviceSeqDao) Insert(deviceId int) error {
	_, err := d.session.Exec("insert into t_device_seq(device_id) values(?)", deviceId)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// GetSeq 获取设备已经同步的消息序列号
func (d *DeviceSeqDao) GetSeq(id int) (int, error) {
	row := d.session.QueryRow("select sync_seq from t_device_seq where device_id = ?", id)
	var syncSeq int
	err := row.Scan(&syncSeq)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return syncSeq, nil
}

// UpdateSeq 更新设备已经同步的消息序列号
func (d *DeviceSeqDao) UpdateSeq(deviceId int, syncSeq int) error {
	_, err := d.session.Exec("update t_device_seq set sync_seq = ? where device_id = ?", syncSeq, deviceId)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
