package blog

import (
	"spider/utils"
	"encoding/xml"
	"github.com/golang/glog"
	"spider/db"
	"fmt"
	"strings"
	"bytes"
	"golang.org/x/net/html/charset"
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


type Atom struct {
	Feed xml.Name `xml:"feed"`
	Entrys []Entry `xml:"entry"`
}

type Entry struct {
	Title string `xml:"title"`
	AtomLink AtomLink `xml:"link"`
	Summary string `xml:"summary"`
	PubDate string `xml:"published"`
	UpdateDate string `xml:"updated"`
}

type AtomLink struct {
	Link xml.Name `xml:"link"`
	Href string `xml:"href,attr"`
}

type TimeLine struct {
	Id  int `field:"id"`
	UserId int `field:"userid"`
	Title string `field:"title"`
	Description string `field:"description"`
	Link string `field:"link"`
	PubData string `field:"pub_data"`
}


func Start(id int,url string,blogurl string) (int,int){
	glog.Info("blog url:",url,"| id:",id)
	successNum := 0
	failedNum := 0
	html :=utils.HttpGet(url)
	//glog.Info(strings.Index(html,"<feed"))
	if id == 31 {
		return MeiTuanGet(id,url,blogurl)
	}
	if strings.Index(html,"<feed") < 100&&strings.Index(html,"<feed")>0{
		ret := Atom{}
		reader := bytes.NewReader([]byte(html))
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = charset.NewReaderLabel
		err := decoder.Decode(&ret)
		//err := xml.Unmarshal([]byte(html),&ret)
		if err != nil{
			glog.Info("error:%v",err.Error())
			return 0,0
		}
		//glog.Info(ret)
		for _,entry := range ret.Entrys{
			glog.Info("link:",entry.AtomLink.Href)
			if strings.Index(entry.AtomLink.Href,"/") == 0{
				//相对路径的情况
				entry.AtomLink.Href = blogurl + entry.AtomLink.Href
			}
			if !db.GetDB().IsExistTimeLineByLink(db.Table_Blog,entry.AtomLink.Href){
				if entry.PubDate == "" {
					entry.PubDate = entry.UpdateDate
				}
				entry.AtomLink.Href = utils.UrlDecode(entry.AtomLink.Href)
				tm, _ := utils.ParseTime(entry.PubDate)
				_, err = db.GetDB().InsertTimeLineBlog(id, entry.Title, utils.SubString(strings.Replace(entry.Summary,"\n","",-1),0,500), entry.AtomLink.Href, fmt.Sprintf("%d", tm.Unix()))
				if err == nil {
					//glog.Info("[Success] blog title:", entry.Title, "| description:", entry.Summary, "| link:", entry.AtomLink.Href, "| pubData:", entry.PubDate)
					successNum++
				} else {
					failedNum++
					glog.Info("[Error] blog title:", entry.Title, "| description:", utils.SubString(strings.Replace(entry.Summary,"\n","",-1),0,500), "| link:", entry.AtomLink.Href, "| pubData:", entry.PubDate, "|err:", err.Error())
				}
			}else{
				//glog.Info("[Exist] title:", entry.Title, "| description:", entry.Summary, "| link:", entry.AtomLink.Href, "| pubData:", entry.PubDate)
			}
		}
	}else {
		ret := Feed{}
		err := xml.Unmarshal([]byte(html), &ret)
		if err != nil {
			glog.Info("error: %v", err.Error())
			return 0,0
		}
		for _, item := range ret.Channel.Items {
			if strings.Index(item.Link,"/") == 0{
				//相对路径的情况
				item.Link = blogurl + item.Link
			}
			item.Link = utils.UrlDecode(item.Link)
			if !db.GetDB().IsExistTimeLineByLink(db.Table_Blog,item.Link) {

				tm, _ := utils.ParseTime(item.PubDate)
				_, err = db.GetDB().InsertTimeLineBlog(id, item.Title, utils.SubString(strings.Replace(item.Description,"\n","",-1),0,500), item.Link, fmt.Sprintf("%d", tm.Unix()))
				if err == nil {
					//glog.Info("[Success] blog title:", item.Title, "| description:", item.Description, "| link:", item.Link, "| pubData:", item.PubDate)
					successNum++
				} else {
					failedNum++
					glog.Info("[Error] blog title:", item.Title, "| description:",utils.SubString(strings.Replace(item.Description,"\n","",-1),0,500), "| link:", item.Link, "| pubData:", item.PubDate, "|err:", err.Error())
				}
			} else {
				//glog.Info("[Exist] title:", item.Title, "| description:", item.Description, "| link:", item.Link, "| pubData:", item.PubDate)
			}
		}
	}
	return successNum,failedNum
}
