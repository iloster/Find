package juejin

import (
	"github.com/golang/glog"
	"spider/utils"
	"encoding/json"
	"fmt"
	"spider/db"
)

var BaseUrl = "https://timeline-merger-ms.juejin.im/v1/get_entry_by_self?src=web&targetUid=%s&type=post&before&limit=20&order=createdAt"

type JuejinD struct {
	Total  int  `json:"total"`
	EntryList []JuejinEntry `json:"entrylist"` 
}

type JuejinEntry struct {
	Title    string  `json:"title"`    //标题
	Content  string  `json:"content"`  //内容
	Link     string `json:"originalUrl"` //链接
	Pub_Date string `json:"createdAt"`  //创建时间
} 

type JuejinData struct {
	S  int `json:"s"`
	M  string `json:"m"`
	D  JuejinD `json:"d"`
}

func Start(id int,jujinId string){
	url := fmt.Sprintf(BaseUrl,jujinId);
	glog.Info("juejin url:",url,"| id:",id)
	html := utils.HttpGet(url)
	ret := JuejinData{}
	err := json.Unmarshal([]byte(html),&ret)
	if err==nil{
		for _,item := range ret.D.EntryList {
			if !db.GetDB().IsExistTimeLineByLink("timeline_juejin",item.Link) {
				tm, _ := utils.ParseTime(item.Pub_Date)
				_, err = db.GetDB().InsertTimeLineJuejin(id,item.Title, item.Content, item.Link, fmt.Sprintf("%d", tm.Unix()))
				if err==nil{
					glog.Info("[Success] title:", item.Title, "| description:", item.Content, "| link:", item.Link, "| pubData:", item.Pub_Date)
				}else{
					glog.Info("[Error] title:", item.Title, "| description:", item.Content, "| link:", item.Link, "| pubData:", item.Pub_Date, "|err:", err.Error())
				}
			}else{
				glog.Info("[Exist] title:",item.Title,"| description:",item.Content,"| link:",item.Link,"| pubData:",item.Pub_Date)
			}
		}
	}else{
		panic(err.Error())
	}
}