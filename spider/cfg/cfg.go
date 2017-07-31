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

type Cfg struct {
	MysqlCfg *MysqlCfg `json:"mysql"`
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
	return this.MysqlCfg;
}

