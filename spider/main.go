package main

import (
	"github.com/golang/glog"
	"flag"
	"spider/db"

	"spider/blog"
	"strings"
	"spider/jianshu"
)

func main(){
	flag.Parse()    // 1
	db.GetDB().Init()
	ret := []db.Famous{}
	ret = db.GetDB().GetFamousInfo()
	glog.Info("ret:",ret)
	for _,item := range ret{
		if item.Blog!="" {
			blog.Start(item.Id, strings.TrimSpace(item.Blog))
		}
		if item.JianShu != "" {
			jianshu.Start(item.Id, item.JianShu)
		}
		//if item.ZhiHu != "" {
		//	zhihu.Start(item.Id, fmt.Sprintf("https://www.zhihu.com/api/v4/members/%s/activities?after_id=%d&limit=20&desktop=True", "jixin", time.Now().Unix()))
		//}
	}

	glog.Flush()

}

//1.获取他们的博客，微博，知乎，简书，推特地址
//2.爬取内容，存到数据库中
func getFamousInfo(){
	ret := []db.Famous{}
	ret = db.GetDB().GetFamousInfo()
	glog.Info("ret:",ret)
}