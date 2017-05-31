package db

import (
	"database/sql"
	"fmt"
	"github.com/golang/glog"
)

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
	sql := fmt.Sprintf(str,userid,title,description,link,source,pub_data)
	glog.Info("sql:",sql)
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

func (this *MysqlDB)IsExistTimeLineByLink(link string) bool{
	str := "select * from time_line where `link`= %q"
	sql := fmt.Sprintf(str,link)
	rows, err := this.DB.Query(sql)
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
