package main

import (
	"strings"
	"qiniu/db"
	"qiniu/cfg"
	"flag"
	"github.com/golang/glog"
)

func main(){
	flag.Parse()    // 1
	cfg.GetCfg().LoadCfg("config.json")
	db.GetDB().Init()
	ret := []db.Famous{}
	ret = db.GetDB().GetFamousInfo()
	//for _,item :=range ret {
	//	if strings.TrimSpace(item.Avater)!="" && !strings.Contains(item.Avater,"http://ou08bmaya") {
	//		path := cfg.GetCfg().GetQiniuCfg().Path
	//		Download(item.Id, "jpg",item.Avater,path)
	//		Upload(item.Id, "jpg",path)
	//	}
	//}
	for _,item :=range ret {
		if strings.TrimSpace(item.Avater)!="" && !strings.Contains(item.Avater,"http://ou08bmaya"){
			db.GetDB().UpdateAvater(item.Id)
		}
	}

	glog.Flush()
}
