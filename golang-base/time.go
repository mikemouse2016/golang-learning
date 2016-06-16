package main

import (
	"fmt"
	"strconv"
	"time"
)

const FORMAT_DATE_TIME string = "2006-01-02 15:04:05"

//获取本地时间戳纳秒，以字符串形式返回
func UnixNano() string {

	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

//获取本地时间戳秒，以字符串形式返回
func Unix() string {

	return strconv.FormatInt(time.Now().Unix(), 10)
}

//获取当前系统时间
func GetTime() string {
	return time.Now().Format(FORMAT_DATE_TIME)
}

//返回时间对象
func ParseDate(dstr string) time.Time {
	t, _ := time.Parse(FORMAT_DATE_TIME, dstr)
	return t
}

func Utc() string {
	return strconv.FormatInt(time.Now().UTC().Unix(), 10)

}

func Date() {
	fmt.Println(time.Unix(time.Now().UTC().Unix(), 0))
	fmt.Println(time.Unix(time.Now().Unix(), 0))
}
func main() {
	name, off := time.Now().Zone()
	fmt.Println(name, off)
	fmt.Println(UnixNano())
	fmt.Println(Unix())
        fmt.Println(Utc())
	fmt.Println(GetTime())
	t := time.Now()
	fmt.Println(t)
	t = t.Add(1*time.Second - time.Duration(t.Nanosecond())*time.Nanosecond)
	fmt.Println(t)
        Date()
}
