package utils

import (
	"encoding/json"
	"github.com/golang/glog"
	"strings"
)

type Ret struct {
	Code int `json:"code"`
	Tag  string `json:"tag"`
}

func GetTag(title string) string {
	url:="http://localhost:5000/title/"+strings.Replace(title,"/","",-1)
	body:=HttpGet(url)
	ret:=&Ret{}
	err:=json.Unmarshal([]byte(body),ret)
	if err!=nil {
		glog.Info("[tagutils] Error:",err.Error()," url:",url)
	}
	if ret.Code == 1 {
		return ret.Tag
	}else{
		return ""
	}
}

