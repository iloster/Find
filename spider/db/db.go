package db

import (
	"database/sql"
	"fmt"
	"github.com/golang/glog"
	_ "github.com/go-sql-driver/mysql"
	"spider/utils"

	"spider/cfg"
)
type Famous struct {
	Id 		int  		`json:"id"`
	Name 		string     	`json:"name"`
	BlogSpider 	string 		`json:"blogspider"`
	HexoSpider	string		`json:"hexoSpider"`
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
	Brief 		string		`json:"brief"`

}
type MysqlDB struct {
	DB *sql.DB
}

var db *MysqlDB

const (
	Table_ZhiHu = "tb_zhihu"
	Table_JianShu = "tb_jianshu"
	Table_Blog = "tb_blog"
	Table_JueJing = "tb_JueJing"
)
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

//func Connect(driveName string) *MysqlDB{
//	glog.Info("wait init db...")
//	defer glog.Info("init db ok!")
//	db,err := sql.Open("mysql",driveName)
//	//defer db.Close()
//	if err!=nil{
//		panic(err.Error())
//	}else{
//		this := new(MysqlDB)
//		this.DB = db
//		return this
//	}
//}

func (this *MysqlDB)InsertTimeLine(userid int,title string,description string,link string,source int,pub_data string) (int ,error){
	//insert into time_line (`id`,`userid`,`title`,`description`,`link`,`pub_data`) values ('0','1','test title','test Description','test Link','1496031517')
	str := "insert into time_line (`fid`,`title`,`description`,`link`,`source`,`pub_data`) values (%d,%q,%q,%q,%d,%s)"
	sql := fmt.Sprintf(str,userid,title,description,link,source,pub_data)
	//glog.Info("sql:",sql)
	res, err := this.DB.Exec(sql)
	if err != nil{
		return 0,err
	}
	row ,err := res.LastInsertId()
	if err != nil{
		return 0,err
	}
	return int(row),nil
}

//func (this *MysqlDB)IsExistTimeLineByLink(link string) bool{
//	str := "select * from time_line where `link`= %q"
//	sql := fmt.Sprintf(str,link)
//	rows, err := this.DB.Query(sql)
//	defer rows.Close()
//	if err != nil{
//		//return err
//		panic(err.Error())
//	}
//	if rows.Next(){
//		return true
//	}else{
//		return false
//	}
//
//}

func (this *MysqlDB)InsertTimeLineZhiHu(userid int,title string,description string,link string,verb string,pub_date string) (int ,error){
	//insert into time_line (`id`,`userid`,`title`,`description`,`link`,`pub_data`) values ('0','1','test title','test Description','test Link','1496031517')
	tag := utils.GetTag(title)
	str := "insert into tb_zhihu (`fid`,`title`,`description`,`link`,`linkid`,`verb`,`pub_date`,`tag`) values (%d,%q,%q,%q,%q,%q,%s,%q)"
	sql := fmt.Sprintf(str,userid,title,description,link,utils.Md5(link),verb,pub_date,tag)
	//glog.Info("sql:",sql)
	res, err := this.DB.Exec(sql)
	if err != nil{
		return 0,err
	}
	row ,err := res.LastInsertId()
	if err != nil{
		return 0,err
	}
	return int(row),nil
}

func (this *MysqlDB)InsertTimeLineJianShu(userid int,title string,description string,link string,pub_date string) (int ,error){
	//insert into time_line (`id`,`userid`,`title`,`description`,`link`,`pub_data`) values ('0','1','test title','test Description','test Link','1496031517')
	tag := utils.GetTag(title)
	str := "insert into tb_jianshu (`fid`,`title`,`description`,`link`,`linkid`,`pub_date`,`tag`) values (%d,%q,%q,%q,%q,%s,%q)"
	sql := fmt.Sprintf(str,userid,title,description,link,utils.Md5(link),pub_date,tag)
	//glog.Info("sql:",sql)
	res, err := this.DB.Exec(sql)
	if err != nil{
		return 0,err
	}
	row ,err := res.LastInsertId()
	if err != nil{
		return 0,err
	}
	return int(row),nil
}

func (this *MysqlDB)InsertTimeLineBlog(userid int,title string,description string,link string,pub_date string) (int ,error){
	//insert into time_line (`id`,`userid`,`title`,`description`,`link`,`pub_data`) values ('0','1','test title','test Description','test Link','1496031517')
	tag := utils.GetTag(title)
	str := "insert into tb_blog (`fid`,`title`,`description`,`link`,`linkid`,`pub_date`,`tag`) values (%d,%q,%q,%q,%q,%s,%q)"
	sql := fmt.Sprintf(str,userid,title,description,link,utils.Md5(link),pub_date,tag)
	//glog.Info("sql:",sql)
	res, err := this.DB.Exec(sql)
	if err != nil{
		return 0,err
	}
	row ,err := res.LastInsertId()
	if err != nil{
		return 0,err
	}
	return int(row),nil
}


func (this *MysqlDB)InsertTimeLineJuejin(userid int,title string,description string,link string,pub_date string)(int,error){
	str := "insert into tb_juejin (`fid`,`title`,`description`,`link`,`linkid`,`pub_date`) values (%d,%q,%q,%q,%q,%s)"
	sql := fmt.Sprintf(str,userid,title,description,0,link,utils.Md5(link),pub_date)
	//glog.Info("sql:",sql)
	res, err := this.DB.Exec(sql)
	if err != nil{
		return 0,err
	}
	row ,err := res.LastInsertId()
	if err != nil{
		return 0,err
	}
	return int(row),nil
}

func (this *MysqlDB)IsExistTimeLineByLink(table string,link string) bool{
	str := "select * from %s where `linkid`= %q"
	sql := fmt.Sprintf(str,table,utils.Md5(link))
	rows, err := this.DB.Query(sql)
	defer rows.Close()
	if err != nil{
		//return err
		panic(err.Error())
	}
	if rows.Next(){
		return true
	}else{
		return false
	}

}
func (this *MysqlDB)GetFamousInfo()[]Famous {
	sql := "SELECT * FROM `tb_famous`"
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
		err = rows.Scan(&item.Id,&item.Name,&item.BlogSpider,&item.HexoSpider,&item.ZhiHuSpider,&item.JianShuSpider,&item.JuejinSpider,&item.Blog,&item.ZhiHu,&item.JianShu,&item.Weibo,&item.Juejin,&item.Github,&item.Avater,&item.Brief)
		glog.Info(err,item)
		ret = append(ret,item)
	}
	return ret
}

func (this *MysqlDB)UpdateAvater(fid int) int{
	str:="update `tb_famous` set `avater`=%q where `id`=%d"
	url := fmt.Sprintf("http://ou08bmaya.bkt.clouddn.com/%d.jpg",fid)
	sql := fmt.Sprintf(str,url,fid)
	_, err := this.DB.Exec(sql)
	if err != nil{
		glog.Info("err:",err)
		return 0
	}

	return 1
}