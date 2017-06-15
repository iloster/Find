package main

import (
	"github.com/golang/glog"
	"flag"
	"spider/db"
)

func main(){
	flag.Parse()    // 1
	db.GetDB().Init()
	//zhihu.Start(fmt.Sprintf("http://www.zhihu.com/api/v4/members/jixin/activities?after_id=%d&limit=20&desktop=True",time.Now().Unix()))
	//blog.Start("http://coolshell.cn/rss")
	//jianshu.Start("http://www.jianshu.com/u/c22ccc510fb9")
	getFamousInfo()
	glog.Flush()

}

//1.获取他们的博客，微博，知乎，简书，推特地址
//2.爬取内容，存到数据库中
func getFamousInfo(){
	ret := []db.Famous{}
	ret = db.GetDB().GetFamousInfo()
	glog.Info("ret:",ret)
}