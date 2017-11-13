package utils

import (
	"net/http"
	"io/ioutil"
	"time"
	"github.com/golang/glog"
	"crypto/tls"
	"math/rand"
)

func HttpGet(url string) string{
	req, err := http.NewRequest("GET", url, nil)
	if err!=nil{
		glog.Info("HttpGet error: ",err.Error())
		return ""
	}
	req.Header.Set("User-Agent", GetUserAgent())
	client := http.DefaultClient
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client.Timeout = 10*time.Second
	client.Transport = tr
	res, e := client.Do(req)
	if e != nil {
		//log4go.INFO("Get请求%s返回错误:%s", url, e)
		glog.Info("url:",url,"| error:",e)
		return ""
	}
	if res.StatusCode == 200 {
		body := res.Body
		defer body.Close()
		bodyByte, _ := ioutil.ReadAll(body)
		resStr := string(bodyByte)
		return resStr
	}
	return ""
}


func GetUserAgent() string {
	var userAgent = [...]string{
		//"Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)",
		//"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
		//"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
		//"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0,",
		//"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
		//"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
		//"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		//"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
		//"Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		//"Mozilla/5.0 (Linux, U, Android 6.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
		//"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
		//"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
		//"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
		//"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
	 }

	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	return userAgent[r.Intn(len(userAgent))]
	//return userAgent[0]
}
