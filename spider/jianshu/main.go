package jianshu

import (
	"github.com/golang/glog"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"strings"
	"spider/utils"
	"spider/db"
)
var Source_JianShu = 4
type JianShuData struct {
	Title    string
	Href     string
	Abstract string
	PubDate string
	Img      string
}
func Start(id int,url string){
	glog.Info("jianshu url:",url,"| id:",id)
	ret:=getData(url)
	for _,item :=range ret{
		item.Href = "http://www.jianshu.com/" + item.Href
		if !db.GetDB().IsExistTimeLineByLink(item.Href){
			tm,_ := utils.ParseTime(item.PubDate)
			_,err := db.GetDB().InsertTimeLine(id,item.Title,item.Abstract,item.Href,Source_JianShu,fmt.Sprintf("%d",tm.Unix()))
			if err == nil {
				glog.Info("[Success] title:", item.Title, "| description:", item.Abstract, "| link:", item.Href, "| pubData:", item.PubDate)
			}else{
				glog.Info("[Error] title:", item.Title, "| description:", item.Abstract, "| link:", item.Href, "| pubData:", item.PubDate, "|err:", err.Error())
			}
		}else{
			glog.Info("[Exist] title:",item.Title,"| description:",item.Abstract,"| link:",item.Href,"| pubData:",item.PubDate)
		}
	}
}

func getData(url string) []JianShuData{
	ret := []JianShuData{}
	for i:=1;i<5;i++ {
		pageurl := fmt.Sprintf(url + "?order_by=shared_at&page=%d",i)
		glog.Info("jianshu Start pageurl:",pageurl)
		doc, err := goquery.NewDocument(pageurl)
		if err == nil {
			doc.Find(".content").Each(func(i int, s *goquery.Selection) {
				item := JianShuData{}
				item.Title = s.Find(".title").Text()
				item.Href, _ = s.Find(".title").Attr("href")
				item.Abstract = strings.TrimSpace(s.Find(".abstract").Text())
				item.PubDate, _ = s.Find(".time").Attr("data-shared-at")
				item.Img,_ = s.Find(".img-blur-done").Attr("src")

				ret = append(ret,item)
			})

		}else{
			glog.Info("jianshu url:",pageurl,"|err:",err.Error())
			break;
		}

	}
	return ret
}


