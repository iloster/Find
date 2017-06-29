package utils

import (
	"time"
	"strings"
	"net/url"
	"github.com/golang/glog"
	"strconv"
	"fmt"
)

func ParseTime(formatted string) (time.Time, error) {
	if strings.Contains(formatted,"/"){
		splitStr := strings.Split(formatted," ")
		if len(splitStr)!=2{
			glog.Error("ParseTime Error str:",formatted)
			return time.Now(),nil
		}

		splitStr1 := strings.Split(splitStr[0],"/")
		year,_ := strconv.Atoi(splitStr1[0])
		mon,_ := strconv.Atoi(splitStr1[1])
		day,_:=strconv.Atoi(splitStr1[2])
		splitStr2 := strings.Split(splitStr[1],":")
		hour,_ := strconv.Atoi(splitStr2[0])
		minute,_ := strconv.Atoi(splitStr2[1])
		second,_ := strconv.Atoi(splitStr2[2])
		return time.Date(year, time.Month(mon), day, hour, minute, second, 0, time.Local),nil
	}
	var layouts = [...]string{
		"Mon, _2 Jan 2006 15:04:05 CCT",
		"Mon, _2 Jan 2006 15:04:05 +0000",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		"Mon, 2, Jan 2006 15:4",
		"02 Jan 2006 15:04:05 CCT",
	}
	var t time.Time
	var err error
	formatted = strings.TrimSpace(formatted)
	loc, _ := time.LoadLocation("Local")
	for _, layout := range layouts {
		t, err = time.ParseInLocation(layout, formatted,loc)
		if !t.IsZero() {
			break
		}
	}
	return t, err
}
func SubString(str string,begin int,length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}

func UrlDecode(urlstr string)(ret string){
	ret1,_ := url.QueryUnescape(urlstr)
	return ret1
}

func PushToWeChat(name string,num1 int,num2 int,cost int64){
	text:=name+"运行成功"
	desp :=fmt.Sprintf("更新成功%d条记录,失败%d记录,总共用时%ds",num1,num2,cost)
	url := fmt.Sprintf("https://sc.ftqq.com/SCU9659T32a012053f440b7d103c7df59301b1de59550945a4fa9.send?text=%q&&desp=%q",text,desp)
	glog.Info(url)
	HttpGet(url)
}