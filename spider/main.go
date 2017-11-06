package main

import (
	"github.com/golang/glog"
	"flag"
	"spider/db"
	"spider/jianshu"
	"spider/cfg"
	"spider/utils"
	"spider/blog"
	"strings"
	"spider/zhihu"
	"fmt"
	"time"
	"spider/juejin"
)

func main(){
	flag.Parse()    // 1
	cfg.GetCfg().LoadCfg(fmt.Sprintf("%s/%s",utils.GetCurrentDirectory(),"config.json"))
	db.GetDB().Init()
	db.GetRedidDB().Init()
	//spider()
	blog.ZBlogGet(1,"http://rarnu.com","http://rarnu.com")
	db.GetRedidDB().Del("latest*")
	glog.Flush()

}


func spider(){
	ret := []db.Famous{}
	ret = db.GetDB().GetFamousInfo()

	for _,item := range ret{
		if cfg.GetCfg().BlogCfg.Open && item.BlogSpider != "" {
			blog.Start(item.Id, strings.TrimSpace(item.BlogSpider),item.Blog)
		}
		if cfg.GetCfg().BlogCfg.Open && item.HexoSpider!=""{
			blog.HexoStart(item.Id, strings.TrimSpace(item.HexoSpider),item.Blog)
		}

		if cfg.GetCfg().JianshuCfg.Open && item.JianShuSpider != "" {
			jianshu.Start(item.Id, item.JianShuSpider)

		}
		if cfg.GetCfg().ZhihuCfg.Open && item.ZhiHuSpider != "" {
			zhihu.Start(item.Id, fmt.Sprintf("https://www.zhihu.com/api/v4/members/%s/activities?after_id=%d&limit=20&desktop=True", strings.TrimSpace(item.ZhiHuSpider), time.Now().Unix()))

		}
		if cfg.GetCfg().JuejinCfg.Open && item.JuejinSpider != ""{
			juejin.Start(item.Id,item.JuejinSpider)
		}
	}
	if cfg.GetCfg().GetPushStatus(){
		utils.PushToWeChat("blog",0,0)
		utils.PushToWeChat("zhihu",0,0)
		utils.PushToWeChat("jianshu",0,0)
	}
	//utils.SegWords();
}


