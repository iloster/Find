package db

import (
	"database/sql"
	"fmt"
	"github.com/golang/glog"
	_ "github.com/go-sql-driver/mysql"
	"spider/utils"
)
type Famous struct {
	Id int  	`json:"id"`
	Name string     `json:"name"`
	Blog string	`json:"blog"`
	ZhiHu string	`json:"zhihu"`
	JianShu string	`json:"jianshu"`
	Avater string   `json:"avater"`
}
type MysqlDB struct {
	DB *sql.DB
}

var db *MysqlDB

const (
	Table_ZhiHu = "timeline_zhihu"
	Table_JianShu = "timeline_jianshu"
	Table_Blog = "timeline_blog"
)
func init() {
	db = &MysqlDB{}
}

func GetDB() *MysqlDB {
	return db
}

func (this *MysqlDB)Init(){
	glog.Info("wait init db...")
	defer glog.Info("init db ok!")
	db,err := sql.Open("mysql","root:licheng@tcp(127.0.0.1:3307)/waste?charset=utf8")
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
	str := "insert into time_line (`userid`,`title`,`description`,`link`,`source`,`pub_data`) values (%d,%q,%q,%q,%d,%s)"
	sql := fmt.Sprintf(str,userid,title,utils.SubString(description,0,1500),link,source,pub_data)
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
	str := "insert into timeline_zhihu (`userid`,`title`,`description`,`link`,`verb`,`pub_date`) values (%d,%q,%q,%q,%q,%s)"
	sql := fmt.Sprintf(str,userid,title,utils.SubString(description,0,1500),link,verb,pub_date)
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
	str := "insert into timeline_jianshu (`userid`,`title`,`description`,`link`,`pub_date`) values (%d,%q,%q,%q,%s)"
	sql := fmt.Sprintf(str,userid,title,utils.SubString(description,0,1500),link,pub_date)
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
	str := "insert into timeline_blog (`userid`,`title`,`description`,`link`,`pub_date`) values (%d,%q,%q,%q,%s)"
	sql := fmt.Sprintf(str,userid,title,utils.SubString(description,0,1500),link,pub_date)
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
	str := "select * from %s where `link`= %q"
	sql := fmt.Sprintf(str,table,link)
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
	sql := "SELECT * FROM `famous`"
	rows,err := this.DB.Query(sql)
	defer rows.Close()
	if err!=nil{
		panic(err.Error())
	}
	ret := []Famous{}
	for rows.Next() {
		item := Famous{}
		err = rows.Scan(&item.Id,&item.Name,&item.Blog,&item.ZhiHu,&item.JianShu,&item.Avater)
		glog.Info(item)
		ret = append(ret,item)
	}

	return ret
}
