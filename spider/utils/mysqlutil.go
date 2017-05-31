package utils

import (
	"database/sql"
	"github.com/golang/glog"
	"strconv"
	"reflect"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"spider/structs"
)
type DBMysql struct {
	DB *sql.DB
}

func Connect(driveName string) *DBMysql{
	db,err := sql.Open("mysql",driveName)
	//defer db.Close()
	if err!=nil{
		panic(err.Error())
	}else{
		this := new(DBMysql)
		this.DB = db
		return this
	}
}

func (this *DBMysql)Insert(data interface{})(int ,error){

	sql := this.getInsertSql(data)
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

func (this *DBMysql)getInsertSql(data interface{}) string{
	ret,ok :=data.(*structs.Sql)
	glog.Info(data)
	if ok {
		args := toString(ret.Data)
		fields := []string{}
		values := []string{}
		for k, v := range (args) {
			fields = append(fields, k)
			values = append(values, "'"+v+"'")
		}
		sql := "insert into " + ret.Table + " (`" + strings.Join(fields, "`,`") + "`) values (" + strings.Join(values, ",") + ")"
		return sql
	}
	return ""
}

func (this *DBMysql)Query(data interface{},result []structs.Sql) error {
	sql := this.getQuerySql(data)
	glog.Info("sql:",sql)
	rows, err := this.DB.Query(sql)
	if err != nil{
		//return err
		panic(err.Error())
	}

	colums,err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	ret := []map[string]string{}
	values := make([][]byte,len(colums))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err!=nil{
			panic(err.Error())
		}
		var value string
		item := map[string]string{}
		for i,col := range values{
			if col== nil {
				value ="NULL"
			}else{
				value = string(col)
			}
			item[colums[i]] = value
		}
		ret = append(ret,item)
		//glog.Info(ret)
	}
	return nil
}


func (this *DBMysql)getQuerySql(data interface{}) string{
	ret,ok := data.(*structs.Sql)
	if ok {
		args := toString(ret.Data)
		str :=[]string{}
		for k, v := range (args) {
			if (k == "id" && len(v) <=1){
				continue
			}
			if len(v) == 0{
				continue
			}
			str = append(str, k + "= '" + v + "'")
		}

		sql := "select * from "+ ret.Table +" where "+strings.Join(str," and ")
		return sql
	}
	return ""
}


func toString(data interface{}) map[string]string {
	val := map[string]string{}
	var f = func(v interface{}) string {
		s := ""
		switch v.(type) {
		case string:
			s = v.(string)
		case int:
			s = strconv.Itoa(v.(int))
		case int8:
			s = strconv.Itoa(int(v.(int8)))
		case int16:
			s = strconv.Itoa(int(v.(int16)))
		case int32:
			s = strconv.Itoa(int(v.(int32)))
		case int64:
			s = strconv.FormatInt(v.(int64), 10)
		}
		return s
	}
	if tmp, ok := data.(map[string]string); ok {
		return tmp
	} else if tmp, ok := data.(map[string]interface{}); ok {
		for k, v := range tmp {
			val[k] = f(v)
		}
	} else {
		v := reflect.ValueOf(data)
		if !v.IsValid() {
			return val
		}
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if !v.IsValid() {
			return val
		}
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			fv := v.Field(i)
			ft := t.Field(i)
			field := ft.Tag.Get("field")
			if field != "" && field != "-" {
				switch ft.Type.Kind() {
				case reflect.Int, reflect.Int64, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					v := fv.Int()
					val[field] = strconv.FormatInt(v, 10)
				case reflect.Float64, reflect.Float32:
					v := fv.Float()
					val[field] = strconv.FormatFloat(v, 'f', -1, 64)
				case reflect.String:
					val[field] = fv.String()
				case reflect.Bool:
					if fv.Bool() {
						val[field] = "1"
					} else {
						val[field] = "0"
					}
				case reflect.Interface:
					v := fv.Interface()
					val[field] = f(v)
				}
			}
		}

	}
	return val
}