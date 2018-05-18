package utils

import (
	"time"
)

//这方式比较特别，按照123456来记忆吧：01月02号 下午3点04分05秒 2006年
func GetTimesString( ) string{
	return time.Now().Format("2006-01-02 15:04:05")
}
func GetTimes() int64  {
	return time.Now().Unix()
}
func GetTimesNano() int64  {
	return time.Now().UnixNano()
}