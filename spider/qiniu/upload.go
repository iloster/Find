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
	"spider/cfg"
)


//首先下载每个人的头像
func Download(fid int,suffix string,url string,path string){
	res,err := http.Get(url)
	if err!=nil{
		panic(err)
	}
	filename := fmt.Sprintf("%d.%s",fid,suffix)
	if createDir(path){
		f,err := os.Create(path+filename)
		if err!=nil{
			panic(err)
		}
		glog.Info("下载成功","|fid:",fid,"|url:",url)
		io.Copy(f,res.Body)
	}else{
		glog.Info("path路径错误",path)
	}

}

func createDir(path string) bool{
	fi,err:=os.Stat(path)
	if err!= nil{
		return false
	}
	if fi.IsDir(){
		err:= os.MkdirAll(path,0666)
		if err==nil{
			return true
		}else{
			return false
		}
	}
	return false
}

func Upload(fid int,suffix string,path string){

	bucket := "trace"
	key := fmt.Sprintf("%d.%s",fid,suffix)
	localFile := path+key
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac:=qbox.NewMac(cfg.GetCfg().GetQiniuCfg().AccessKey,cfg.GetCfg().GetQiniuCfg().SecretKey)
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
	glog.Info("上传成功","|fid:",fid,"|key:",ret.Key,"|hash:",ret.Hash)

}
