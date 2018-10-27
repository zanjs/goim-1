package lib

import "time"

// UnixTime 将时间转化为毫秒数
func UnixTime(t time.Time) int64 {
	return t.UnixNano() / 1000000
}

// UnunixTime 将毫秒数转为为时间
func UnunixTime(unix int64) time.Time {
	return time.Unix(0, unix*1000000)
}
