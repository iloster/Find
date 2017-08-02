package qiniu

import (
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"io"
	"github.com/golang/glog"
)

const (
	ACCESS_KEY = "xxxx"
	SECRET_KEY = "xxxx"
)


func Start(fid int,url string){
	//tmp := strings.Split(url,".")
	//suffix := tmp[1]
	glog.Info("fid:",fid,"| url:",url)
	Download(fid,"jpg",url)
}

//首先下载每个人的头像
func Download(fid int,suffix string,url string){
	res,err := http.Get(url)
	if err!=nil{
		panic(err)
	}
	filename := fmt.Sprintf("%d.%s",fid,suffix)
	f,err := os.Create("/Users/dev/Desktop/trace/"+filename)
	if err!=nil{
		panic(err)
	}
	io.Copy(f,res.Body)
}

func Upload(fid int,suffix string,path string){

	bucket := "trace"
	key := fmt.Sprintf("%d.%s",fid,suffix)
	localFile := path+key
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac:=qbox.NewMac(ACCESS_KEY,SECRET_KEY)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		panic(err)
		return
	}
	//fmt.Println(ret.Key,ret.Hash)
	glog.Info("fid:",fid,"|key:",ret.Key,"|hash:",ret.Hash)

}
