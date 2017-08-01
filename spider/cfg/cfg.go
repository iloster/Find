package cfg

import (
	"os"
	"encoding/json"
	"github.com/golang/glog"
)

type MysqlCfg struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Port 	 int `json:"port"`
	DataBase string `json:"database"`
}

type ZhihuCfg struct {
	Authorization string `json:"authorization"`
	Open  bool `json:"open"`
}

type JianshuCfg struct {
	Open bool `json:"open"`
}

type BlogCfg struct {
	Open bool `json:"open"`
}
type Cfg struct {
	MysqlCfg 	*MysqlCfg 	`json:"mysql"`
	Push    	bool       	`json:"push"`
	ZhihuCfg   	*ZhihuCfg  	`json:"zhihu"`
	JianshuCfg	*JianshuCfg 	`json:"jianshu"`
	BlogCfg		*BlogCfg	`json:"blog"`
}
var cfg *Cfg

func init(){
	cfg = &Cfg{}
}

func GetCfg() *Cfg  {
	return cfg
}

func (this *Cfg)LoadCfg(path string){
	fd, err := os.Open(path)
	if err != nil {
		panic("无法打开配置文件 config.json: " + err.Error())
	}
	defer fd.Close()
	err = json.NewDecoder(fd).Decode(cfg)
	if err != nil {
		panic("解析配置文件出错: " + err.Error())
	}
	glog.Info("配置文件加载成功")
}

func (this *Cfg)GetMysqlCfg() *MysqlCfg{
	return this.MysqlCfg
}

func (this *Cfg)GetPushStatus() bool{
	return this.Push
}

func (this *Cfg)GetZhihuCfg() *ZhihuCfg {
	return this.ZhihuCfg
}

func (this *Cfg)GetBlogCfg() *BlogCfg{
	return this.BlogCfg
}

func (this *Cfg)GetJianshuCfg() *JianshuCfg{
	return this.JianshuCfg
}
