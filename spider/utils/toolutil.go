package utils

import (
	"time"
	"strings"
	"net/url"
	"github.com/golang/glog"
	"strconv"
	"fmt"
	"crypto/md5"
	"path/filepath"
	"os"
	"crypto/sha1"
	"encoding/base64"
	"crypto/hmac"
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
		"Fri Jun 16 2017 16:24:00 GMT+0800 (CST)",
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
	if t.IsZero(){
		return time.Now(),nil
	}
	return t, err
}

//2017年1月1日
func ParseTime2(formatted string) (time.Time, error) {
	if strings.Contains(formatted,"年"){
		splitStr := strings.Split(formatted,"年")
		year,_ := strconv.Atoi(splitStr[0])
		splitStr2 := strings.Split(splitStr[1],"月")
		month,_ := strconv.Atoi(splitStr2[0])
		splitStr3 := strings.Split(splitStr2[1],"日")
		day,_ := strconv.Atoi(splitStr3[0])
		return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local),nil
	}
	return time.Now(),nil
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

func PushToWeChat(name string,num1 int,num2 int){
	text:=name+"运行成功,"
	desp :=fmt.Sprintf("更新成功%d条记录,失败%d记录,%s",num1,num2,time.Now().Format("2006-01-02 15:04:05"))
	url := fmt.Sprintf("https://sc.ftqq.com/SCU9659T32a012053f440b7d103c7df59301b1de59550945a4fa9.send?text=%q&&desp=%q",text,desp)
	glog.Info(url)
	HttpGet(url)
}

func Md5(str string)string{
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		glog.Info("获取路径失败")
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func Sha1(str string)string{
	h := sha1.New()
	//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(str))
	//这个用来得到最终的散列值的字符切片。Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func Base64Encode(str string) string {
	return  base64.StdEncoding.EncodeToString([]byte(str))
}

func Hmac(str string,key string) string{
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(str))
	return fmt.Sprintf("%x", mac.Sum(nil))
}