package lid

import (
	"database/sql"
	"errors"
	"goim/public/logger"
	"sync"
	"time"
)

var (
	ErrLidUnavailable    = errors.New("lid unavailable")    // lid不可用
	ErrBufferUnavailable = errors.New("buffer unavailable") // buffer不可用
)

type Lid struct {
	db         *sql.DB    // 数据库连接
	businessId string     // 业务id
	buffers    [2]buffer  // 自增健缓存
	being_used int        // 正在使用的buffer
	lock       sync.Mutex // 互斥锁
	available  bool       // 是否可用
}

type buffer struct {
	min       int64 // 最小值
	max       int64 // 最大值
	median    int64 // 中间值
	available bool  // 是否可用
}

// NewLid 创建一个lid
func NewLid(db *sql.DB, businessId string) (*Lid, error) {
	lid := Lid{
		db:         db,
		businessId: businessId,
		being_used: 1,
		available:  true,
	}
	err := lid.getFromDB()
	lid.being_used = 0
	if err != nil {
		logger.Sugaer.Error(err)
		return nil, err
	}
	return &lid, nil
}

// Get 获取自增id
func (l *Lid) Get() (int64, error) {
	for {
		key, err := l.getKey()
		if err == ErrBufferUnavailable {
			time.Sleep(time.Microsecond * 200)
			continue
		}
		if err == ErrLidUnavailable {
			return 0, ErrLidUnavailable

		}
		return key, nil
	}
}

// GetKey 获取自增id
func (l *Lid) getKey() (int64, error) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.available == false {
		return 0, ErrLidUnavailable
	}

	if l.buffers[l.being_used].available == false {
		return 0, ErrBufferUnavailable
	}

	// 如果buffer已经达到中间值，从数据库初始化另一个buffer
	if l.buffers[l.being_used].min == l.buffers[l.being_used].median {
		go l.getFromDB()
	}

	// 如果buffer已经到最大值，切换到另一个buffer
	if l.buffers[l.being_used].min >= l.buffers[l.being_used].max {
		if l.buffers[1-l.being_used].available == true {
			l.buffers[l.being_used].available = false
			l.being_used = 1 - l.being_used
		} else {
			return 0, ErrBufferUnavailable
		}
	}
	l.buffers[l.being_used].min++
	return l.buffers[l.being_used].min, nil
}

func (l *Lid) initAnotherBuffer() {
	l.lock.Lock()
	defer l.lock.Unlock()

	// 重试5次，如果5次都失败，将lid改为不可用
	for i := 0; i < 5; i++ {
		err := l.getFromDB()
		if err == nil {
			return
		}
		continue
	}
	l.available = false
}

func (l *Lid) getFromDB() error {
	var (
		maxId int64
		step  int64
	)

	tx, err := l.db.Begin()
	if err != nil {
		logger.Sugaer.Error(err)
		return err
	}
	defer tx.Rollback()

	row := tx.QueryRow("select max_id,step from t_lid where business_id = ? for update", l.businessId)
	err = row.Scan(&maxId, &step)
	if err != nil {
		logger.Sugaer.Error(err)
		return err
	}

	_, err = tx.Exec("update t_lid set max_id = ? where business_id = ?", maxId+step, l.businessId)
	if err != nil {
		logger.Sugaer.Error(err)
		return err
	}
	err = tx.Commit()
	if err != nil {
		logger.Sugaer.Error(err)
		return err
	}

	unused := 1 - l.being_used
	l.buffers[unused].min = maxId
	l.buffers[unused].max = maxId + step
	l.buffers[unused].median = (maxId + maxId + step) / 2
	l.buffers[unused].available = true
	return nil
}
