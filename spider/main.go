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
	"spider/wechat"
)

func main(){
	flag.Parse()    // 1
	cfg.GetCfg().LoadCfg(fmt.Sprintf("%s/%s",utils.GetCurrentDirectory(),"config.json"))
	db.GetDB().Init()
	db.GetRedidDB().Init()
	spider()
	wechat.Start()
	db.GetRedidDB().Del("latest*")
	glog.Flush()

}


func spider(){
	ret := []db.Famous{}
	ret = db.GetDB().GetFamousInfo()

	blogSuccTotal:=0
	blogFailedTotal:=0

	zhihuSuccTotal:=0
	zhihuFailedTotal:=0

	jianshuSuccTotal:=0
	jianshuFailedTotal:=0
	for _,item := range ret{
		if cfg.GetCfg().BlogCfg.Open && item.BlogSpider != "" {
			num1, num2:= blog.Start(item.Id, strings.TrimSpace(item.BlogSpider),item.Blog)
			blogSuccTotal+=num1
			blogFailedTotal +=num2
		}
		if cfg.GetCfg().BlogCfg.Open && item.HexoSpider!=""{
			num1, num2:= blog.HexoStart(item.Id, strings.TrimSpace(item.HexoSpider),item.Blog)
			blogSuccTotal+=num1
			blogFailedTotal +=num2
		}

		if cfg.GetCfg().JianshuCfg.Open && item.JianShuSpider != "" {
			num1,num2:=jianshu.Start(item.Id, item.JianShuSpider)
			zhihuSuccTotal+=num1
			zhihuFailedTotal +=num2

		}
		if cfg.GetCfg().ZhihuCfg.Open && item.ZhiHuSpider != "" {
			num1,num2 := zhihu.Start(item.Id, fmt.Sprintf("https://www.zhihu.com/api/v4/members/%s/activities?after_id=%d&limit=20&desktop=True", strings.TrimSpace(item.ZhiHuSpider), time.Now().Unix()))
			jianshuSuccTotal+=num1
			jianshuFailedTotal +=num2
		}
		if cfg.GetCfg().JuejinCfg.Open && item.JuejinSpider != ""{
			juejin.Start(item.Id,item.JuejinSpider)
		}
	}
	if cfg.GetCfg().GetPushStatus(){
		utils.PushToWeChat("blog",blogSuccTotal,blogFailedTotal)
		utils.PushToWeChat("zhihu",zhihuSuccTotal,zhihuFailedTotal)
		utils.PushToWeChat("jianshu",jianshuSuccTotal,jianshuFailedTotal)
	}
	//utils.SegWords();
}


