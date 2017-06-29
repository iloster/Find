package main

import (
	"github.com/golang/glog"
	"flag"
	"spider/db"

	"spider/blog"
	"strings"
	"spider/jianshu"
	"spider/zhihu"
	"fmt"
	"time"
	"spider/utils"
)

func main(){
	flag.Parse()    // 1
	db.GetDB().Init()
	ret := []db.Famous{}
	ret = db.GetDB().GetFamousInfo()
	start_time := time.Now().Unix()

	zhihu_success_num := 0
	zhihu_failed_num := 0
	jianshu_success_num := 0
	jianshu_failed_num := 0
	blog_success_num := 0
	blog_failed_num := 0
	for _,item := range ret{
		//if item.Id != 29 {
		//	continue
		//}
		if item.Blog != "" {
			num1,num2:=blog.Start(item.Id, strings.TrimSpace(item.Blog))
			blog_success_num+=num1
			blog_failed_num+=num2
		}
		if item.JianShu != "" {
			num1,num2:=jianshu.Start(item.Id, item.JianShu)
			jianshu_success_num+=num1
			jianshu_failed_num+=num2
		}
		if item.ZhiHu != "" {
			num1,num2:=zhihu.Start(item.Id, fmt.Sprintf("https://www.zhihu.com/api/v4/members/%s/activities?after_id=%d&limit=20&desktop=True", strings.TrimSpace(item.ZhiHu), time.Now().Unix()))
			zhihu_success_num+=num1
			zhihu_failed_num+=num2
		}
	}
	utils.PushToWeChat("博客",blog_success_num,blog_failed_num,time.Now().Unix() - start_time)
	utils.PushToWeChat("简书",jianshu_success_num,jianshu_failed_num,time.Now().Unix() - start_time)
	utils.PushToWeChat("知乎",zhihu_success_num,zhihu_failed_num,time.Now().Unix() - start_time)
	glog.Flush()

}

//1.获取他们的博客，微博，知乎，简书，推特地址
//2.爬取内容，存到数据库中
func getFamousInfo(){
	ret := []db.Famous{}
	ret = db.GetDB().GetFamousInfo()
	glog.Info("ret:",ret)
}
