package main

import (
	"github.com/golang/glog"
	"flag"
	"spider/utils"
	"spider/structs"
	"encoding/xml"
	"spider/db"
	"fmt"
	"spider/zhihu"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"time"
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
	zhiHu()
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

func zhiHu(){
	//zhihu.Init()
	url := fmt.Sprintf("http://www.zhihu.com/api/v4/members/jixin/activities?after_id=%d&limit=20&desktop=True",time.Now().Unix())
	res,err:=zhihu.GetSession().Get(url)
	glog.Info("url:",url)
	if err == nil{
		bodyByte, _ := ioutil.ReadAll(res.Body)
		resStr := string(bodyByte)
		glog.Info(resStr)
		ret := zhihu.ZhihuActivity{}
		err := json.Unmarshal([]byte(resStr),&ret)
		glog.Info("err",err)
		if err == nil{
			for _,item := range ret.Data {
				var title = ""
				var desc = ""
				var link = ""
				switch item.Verb {
					case zhihu.Verb_ANSWER_CREATE:{
						//	回答了问题

					}
					case zhihu.Verb_ANSWER_VOTE_UP:{
						//赞同问题
						title ="赞同了:" + item.Target.Question.Title + "--"+item.Target.Author.Name+"的回答"
						desc  = item.Target.Excerpt
						link = fmt.Sprintf("https://www.zhihu.com/question/%s/answer/%s",item.Target.Question.Id,string(item.Target.Id))
					}
					case zhihu.Verb_MEMBER_FOLLOW_ROUNDTABLE:{
						title = "关注了圆桌:" + item.Target.Name
						desc = item.Target.Description
						link = fmt.Sprintf("https://www.zhihu.com/roundtable/%s",string(item.Target.Id))
					}
					case zhihu.Verb_MEMBER_VOTEUP_ARTICLE:{
						title = "赞了文章:" + item.Target.Title + "--" + item.Target.Author.Name
						desc = item.Target.Excerpt
						link = fmt.Sprintf("https://zhuanlan.zhihu.com/p/%s",string(item.Target.Id))
					}

				}

				_, err = db.GetDB().InsertTimeLine(1, title, desc, link, Source_Zhihu, strconv.Itoa(item.CreateTime))
				if err == nil {
					glog.Info("[Success] title:", title, "| description:", desc, "| link:", link, "| pub_data:", item.CreateTime)
				} else {
					glog.Info("[Error] title:", title, "| description:", desc, "| link:", link, "| pub_data:", item.CreateTime, "|error:", err.Error())
				}
			}
		}
	}else{
		glog.Info(err.Error())
	}
}