package main

import (
	"github.com/golang/glog"
	"flag"
	"spider/utils"
	"spider/structs"
	"encoding/xml"
	"spider/db"
	"fmt"
)

const (
	Source_Blog = 1
	Source_Weibo = 2
	Source_Zhihu = 3
	Source_Jianshu = 4
	Source_Twitter = 5
)
func main(){
	flag.Parse()    // 1
	db.GetDB().Init()
	start()
	glog.Flush()
}

//1.获取他们的博客，微博，知乎，简书，推特地址
//2.爬取内容，存到数据库中
func start(){
	url := "http://coolshell.cn/rss"
	html :=utils.HttpGet(url)
	//glog.Info("html:",html)
	ret := structs.Feed{}
	err := xml.Unmarshal([]byte(html),&ret)
	if err != nil {
		glog.Info("error: %v", err)
		return
	}
	for _,item := range ret.Channel.Items{
		if !db.GetDB().IsExistTimeLineByLink(item.Link){
			glog.Info("[success] title:",item.Title,"| description:",item.Description,"| link:",item.Link,"| pubData:",item.PubDate)
			tm,_ := utils.ParseTime(item.PubDate)
			db.GetDB().InsertTimeLine(1,item.Title,item.Description,item.Link,Source_Blog,fmt.Sprintf("%d",tm.Unix()))
		}else{
			glog.Info("[exist] title:",item.Title,"| description:",item.Description,"| link:",item.Link,"| pubData:",item.PubDate)
		}
	}

}

