package utils

import (
	"minercms/consts"
	"time"
)

func String2Time(str, tmpl string) time.Time {
	//获取本地location
	loc, _ := time.LoadLocation(consts.Local)          //重要：获取时区
	theTime, _ := time.ParseInLocation(tmpl, str, loc) //使用模板在对应时区转化为time.time类型
	return theTime
}

//str必须为 yyyy-MM-dd hh:mm:ss
func String2TimeNyrSfm(str string) time.Time {
	return String2Time(str, consts.NyrSfm)
}

//str必须为 yyyy-MM-dd
func String2TimeNyr(str string) time.Time {
	return String2Time(str, consts.Nyr)
}
func Time2String(t time.Time, templ string) string {
	sr := t.Unix() //转化为时间戳 类型是int64
	//时间戳转日期
	dataTimeStr := time.Unix(sr, 0).Format(templ) //设置时间戳 使用模板格式化为日期字符串
	return dataTimeStr
}
func Time2StringNyr(t time.Time) string {
	ret := Time2String(t, consts.Nyr)
	return ret
}
func Time2StringNyrSfm(t time.Time) string {
	ret := Time2String(t, consts.NyrSfm)
	return ret
}
func Time2StringNyrSfmXhx(t time.Time) string {
	ret := Time2String(t, consts.NyrSfmXhx)
	return ret
}
func Time2StringNyrXhx(t time.Time) string {
	ret := Time2String(t, consts.NyrXhx)
	return ret
}
func Time2StringNyrWu(t time.Time) string {
	ret := Time2String(t, consts.NyrWu)
	return ret
}
