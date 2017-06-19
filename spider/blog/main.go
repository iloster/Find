package blog

import (
	"spider/utils"
	"encoding/xml"
	"github.com/golang/glog"
	"spider/db"
	"fmt"
)

var Source_Blog = 1


type Feed struct {
	Rss xml.Name `xml:"rss"`
	Version string `xml:"version,attr"`
	Channel Channel `xml:"channel"`
}
type Channel struct {
	Title string `xml:"title"`
	Items [] Item `xml:"item"`
}
type Item struct {
	Title string `xml:"title"`
	Link string `xml:"link"`
	PubDate string `xml:"pubDate"`
	Description string `xml:"description"`
}

type Sql struct {
	Table string
	Data  interface{}
}

type TimeLine struct {
	Id  int `field:"id"`
	UserId int `field:"userid"`
	Title string `field:"title"`
	Description string `field:"description"`
	Link string `field:"link"`
	PubData string `field:"pub_data"`
}


func Start(id int,url string){
	glog.Info("blog url:",url,"| id:",id)
	html :=utils.HttpGet(url)
	ret := Feed{}
	err := xml.Unmarshal([]byte(html),&ret)
	if err != nil {
		glog.Info("error: %v", err.Error())
		return
	}
	for _,item := range ret.Channel.Items{
		if !db.GetDB().IsExistTimeLineByLink(item.Link){

			tm,_ := utils.ParseTime(item.PubDate)
			_,err = db.GetDB().InsertTimeLine(id,item.Title,item.Description,item.Link,Source_Blog,fmt.Sprintf("%d",tm.Unix()))
			if err == nil {
				glog.Info("[Success] blog title:", item.Title, "| description:", item.Description, "| link:", item.Link, "| pubData:", item.PubDate)
			}else{
				glog.Info("[Error] blog title:", item.Title, "| description:", item.Description, "| link:", item.Link, "| pubData:", item.PubDate, "|err:", err.Error())
			}
		}else{
			glog.Info("[Exist] title:",item.Title,"| description:",item.Description,"| link:",item.Link,"| pubData:",item.PubDate)
		}
	}

}
