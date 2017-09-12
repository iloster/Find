package blog

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
	"strings"
	"fmt"
	"spider/utils"
	"spider/db"
)

type HexoPost struct {
	Title  		string
	Content 	string
	PubDate  	string
	Link 		string
}



func HexoStart(id int,url string,blogurl string)(int,int){
	glog.Info("id:",id,"|url:",url,"|blogurl:",blogurl)
	successNum := 0
	failedNum := 0
	doc, err := goquery.NewDocument(blogurl)
	if err!=nil{
		glog.Info("[hexo] error:",err.Error()," | blogurl:",blogurl)
	}
	items := []*HexoPost{}
	doc.Find(".post-block").Each(func(i int, s *goquery.Selection) {
		item := &HexoPost{}
		item.Title = s.Find(".post-title-link").Text()
		item.Link,_ = s.Find(".post-title-link").Attr("href")
		item.Link = fmt.Sprintf("%s/%s",blogurl,item.Link)
		item.Content,_ = s.Find(".post-body").Html()
		item.Content = strings.TrimSpace(item.Content)
		item.PubDate,_ = s.Find("time").Attr("datetime")
		items = append(items,item)
	})
	for _,item :=range items{
		if !db.GetDB().IsExistTimeLineByLink(db.Table_Blog,item.Link){
			tm,_ := utils.ParseTime(item.PubDate)
			_,err := db.GetDB().InsertTimeLineBlog(id,item.Title,utils.SubString(strings.Replace(item.Content,"\n","",-1),0,500),item.Link,fmt.Sprintf("%d",tm.Unix()))
			if err == nil {
				//glog.Info("[Success] title:", item.Title, "| description:", item.Abstract, "| link:", item.Href, "| pubData:", item.PubDate)
				successNum++
			}else{
				failedNum++
				glog.Info("[Error] title:", item.Title, "| description:", item.Content, "| link:", item.Link, "| pubData:", item.PubDate, "|err:", err.Error())
			}
		}else{
			//glog.Info("[Exist] title:",item.Title,"| description:",item.Content,"| link:",item.Link,"| pubData:",item.PubDate)
		}
	}
	return successNum,failedNum

}
