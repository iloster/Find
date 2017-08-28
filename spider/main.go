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
	"spider/qiniu"
)

func main(){
	flag.Parse()    // 1
	cfg.GetCfg().LoadCfg("config.json")
	db.GetDB().Init()

	spider()
	qn()
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

func qn(){
	ret := []db.Famous{}
	ret = db.GetDB().GetFamousInfo()
	for _,item :=range ret {
		if strings.TrimSpace(item.Avater)!="" && !strings.Contains(item.Avater,"http://ou08bmaya") {
			path := cfg.GetCfg().GetQiniuCfg().Path
			qiniu.Download(item.Id, "jpg",item.Avater,path)
			qiniu.Upload(item.Id, "jpg",path)
		}
	}
	for _,item :=range ret {
		if strings.TrimSpace(item.Avater)!="" && !strings.Contains(item.Avater,"http://ou08bmaya"){
			db.GetDB().UpdateAvater(item.Id)
		}
	}
}
//1.获取他们的博客，微博，知乎，简书，推特地址
//2.爬取内容，存到数据库中
func getFamousInfo(){
	ret := []db.Famous{}
	ret = db.GetDB().GetFamousInfo()
	glog.Info("ret:",ret)
}
