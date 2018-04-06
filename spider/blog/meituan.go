package blog

import (
	"spider/utils"
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"bytes"
	"github.com/golang/glog"
	"strings"
	"fmt"
	"spider/db"
)

type MeiTuanAtom struct{
	Feed xml.Name `xml:"feed"`
	Entrys []MeiTuanEntry `xml:"entry"`
}

type MeiTuanEntry struct {
	Title string `xml:"title"`
	Content string `xml:"content"`
	PubDate string `xml:"updated"`
	Link   string `xml:"id"`
}

func MeiTuanGet(id int,url string,blogurl string)(int,int){
	successNum := 0
	failedNum := 0
	html :=utils.HttpGet(url)
	ret:= MeiTuanAtom{}
	reader := bytes.NewReader([]byte(html))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&ret)
	if err != nil{
		glog.Info("error:%v",err.Error())
		return 0,0
	}
	for _,entry := range ret.Entrys{
		if strings.Index(entry.Link,"/") == 0{
			//相对路径的情况
			entry.Link = blogurl + entry.Link
		}
		if !db.GetDB().IsExistTimeLineByLink(db.Table_Blog,entry.Link){
			entry.Link = utils.UrlDecode(entry.Link)
			tm, _ := utils.ParseTime(entry.PubDate)
			_, err = db.GetDB().InsertTimeLineBlog(id, entry.Title,strings.Replace(entry.Content,"\n","",-1), entry.Link, fmt.Sprintf("%d", tm.Unix()))
			if err == nil {
				//glog.Info("[Success] blog title:", entry.Title, "| description:", entry.Summary, "| link:", entry.AtomLink.Href, "| pubData:", entry.PubDate)
				successNum++
			} else {
				failedNum++
				glog.Info("[Error] blog title:", entry.Title, "| description:", utils.SubString(strings.Replace(entry.Content,"\n","",-1),0,500), "| link:", entry.Link, "| pubData:", entry.PubDate, "|err:", err.Error())
			}
		}else{
			glog.Info("[Exist] title:", entry.Title, "| description:", entry.Content, "| link:", entry.Link, "| pubData:", entry.PubDate)
		}
	}
	return successNum,failedNum
}

