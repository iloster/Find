package main

import (
	"github.com/golang/glog"
	"flag"
	"spider/utils"
	"spider/structs"
	"encoding/xml"
	"spider/db"
	"fmt"
	"regexp"
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
	//start()
	zhihu()
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

			tm,_ := utils.ParseTime(item.PubDate)
			_,err = db.GetDB().InsertTimeLine(1,item.Title,item.Description,item.Link,Source_Blog,fmt.Sprintf("%d",tm.Unix()))
			if err == nil {
				glog.Info("[Success] title:", item.Title, "| description:", item.Description, "| link:", item.Link, "| pubData:", item.PubDate)
			}else{
				glog.Info("[Error] title:", item.Title, "| description:", item.Description, "| link:", item.Link, "| pubData:", item.PubDate, "|err:", err.Error())
			}
		}else{
			glog.Info("[Exist] title:",item.Title,"| description:",item.Description,"| link:",item.Link,"| pubData:",item.PubDate)
		}
	}

}

func zhihu(){
	url := "https://www.zhihu.com/people/jixin/answers"
	html := utils.HttpGet(url)
	//glog.Info("html:"+html)
	reg := regexp.MustCompile(`<span class="ActivityItem-metaTitle">(.*?)</span>`)
	glog.Info(reg.FindAll([]byte(html),-1))
}