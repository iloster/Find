package db

import (
	"database/sql"
	"fmt"
	"github.com/golang/glog"
	_ "github.com/go-sql-driver/mysql"
	"qiniu/cfg"
)
type Famous struct {
	Id 		int  		`json:"id"`
	Name 		string     	`json:"name"`
	BlogSpider 	string 		`json:"blogspider"`
	ZhiHuSpider 	string 		`json:"zhihuspider"`
	JianShuSpider 	string 		`json:"jianshuspider"`
	JuejinSpider  	string		`json:"juejinspider"`
	Blog 		string		`json:"blog"`
	ZhiHu 		string		`json:"zhihu"`
	JianShu 	string		`json:"jianshu"`
	Weibo		string		`json:"weibo"`
	Juejin		string		`json:"juejin"`
	Github		string 		`json:"github"`
	Avater 		string   	`json:"avater"`
	Brief 		string
}
type MysqlDB struct {
	DB *sql.DB
}

var db *MysqlDB

func init() {
	db = &MysqlDB{}
}

func GetDB() *MysqlDB {
	return db
}

func (this *MysqlDB)Init(){
	mysqlcfg := cfg.GetCfg().GetMysqlCfg()
	glog.Info(mysqlcfg)
	path:=fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/%s?charset=utf8mb4",mysqlcfg.UserName,mysqlcfg.Password,mysqlcfg.Port,mysqlcfg.DataBase)
	glog.Info("wait init db...",path)
	defer glog.Info("init db ok!")
	db,err := sql.Open("mysql",path)
 	if err!=nil{
		panic(err.Error())
	}else{
		this.DB = db
	}

}

func (this *MysqlDB)GetFamousInfo()[]Famous {
	sql := "SELECT * FROM `famous`"
	rows,err := this.DB.Query(sql)
	defer rows.Close()
	if err!=nil{
		panic(err.Error())
	}
	if rows.Err()!=nil{
		panic(rows.Err().Error())
	}
	ret := []Famous{}

	for rows.Next() {
		item := Famous{}
		err = rows.Scan(&item.Id,&item.Name,&item.BlogSpider,&item.ZhiHuSpider,&item.JianShuSpider,&item.JuejinSpider,&item.Blog,&item.ZhiHu,&item.JianShu,&item.Weibo,&item.Juejin,&item.Github,&item.Avater,&item.Brief)
		glog.Info(err,item)
		ret = append(ret,item)
	}
	return ret
}

func (this *MysqlDB)UpdateAvater(fid int) int{
	str:="update `famous` set `avater`=%q where `id`=%d"
	url := fmt.Sprintf("http://ou08bmaya.bkt.clouddn.com/%d.jpg",fid)
	sql := fmt.Sprintf(str,url,fid)
	_, err := this.DB.Exec(sql)
	if err != nil{
		glog.Info("err:",err)
		return 0
	}

	return 1
}