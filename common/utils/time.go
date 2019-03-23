package utils

import (
	"time"
)

//这方式比较特别，按照123456来记忆吧：01月02号 下午3点04分05秒 2006年
func GetTimesString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//秒
func GetTimes() int64 {
	return time.Now().Unix()
}

//纳秒
func GetTimesNano() int64 {
	return time.Now().UnixNano()
}

//日期转时间戳
func TransformUnix(str string) (int64, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return 0, err
	}
	tm, err := time.ParseInLocation("2006-01-02", str, loc)
	if err != nil {
		return 0, err
	}
	return tm.Unix(), nil
}

//时间戳转日期
func TimestampToDate(tm int64) string {
	return time.Unix(tm, 0).Format("2006-01-02")
}

//时间戳转日期
func TimestampToStr(tm int64) string {
	return time.Unix(tm, 0).Format("2006-01-02 15:04:05")
}

//时间戳转日期
func TimestampToSqlDate(tm int64) string {
	return time.Unix(tm, 0).Format("20060102")
}

//时间戳转日期
func TimestampToSqlDateTime(tm int64) string {
	return time.Unix(tm, 0).Format("20060102150405")
}
