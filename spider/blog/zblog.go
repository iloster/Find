package blog

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
	"strings"
	"spider/utils"
	"fmt"
	"spider/db"
)
type ZBlogPost struct {
	Title  		string
	Content 	string
	PubDate  	string
	Link 		string
}

func ZBlogGet(id int,url string,blogurl string){
	glog.Info("[ZBlog]"," id:",id," url:",url," blogurl:",blogurl)
	doc, err := goquery.NewDocument(blogurl)
	if err!=nil{
		glog.Info("[ZBlog] error:",err.Error()," | blogurl:",blogurl)
		return
	}
	successNum := 0
	failedNum := 0
	items := []*ZBlogPost{}
	doc.Find(".post").Each(func(i int, s *goquery.Selection) {
		item := &ZBlogPost{}
		item.Title = strings.TrimSpace(s.Find(".post-title").Text())
		tag := utils.GetTag(item.Title)
		glog.Info("[ZBlog] ZBlogGet:",tag)
		item.Content = strings.TrimSpace(s.Find(".post-body").Text())
		item.PubDate = strings.TrimSpace(s.Find(".post-date").Text())
		item.Link,_ = s.Find(".post-title").Find("a").Eq(0).Attr("href")
		items = append(items,item)
	})
	for _,item :=range items{
		if !db.GetDB().IsExistTimeLineByLink(db.Table_Blog,item.Link){
			tm,_ := utils.ParseTime(item.PubDate)
			_,err := db.GetDB().InsertTimeLineBlog(id,item.Title,strings.Replace(item.Content,"\n","",-1),item.Link,fmt.Sprintf("%d",tm.Unix()))
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
	//return successNum,failedNum

}





