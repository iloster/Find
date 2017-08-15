package utils

import (
	"github.com/huichen/sego"
	"github.com/golang/glog"
	"math"
	"regexp"
	"os"
	"bufio"
	"io"
	"strings"
)


func SegWords()  {
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("../vendor/github.com/huichen/sego/data/dictionary.txt")

	// 分词
	text := []byte("情景分析 通常我们在破解apk的时候，第一步肯定先反编译apk文件，然后开始修改代码和资源文件，最后回编译")
	segments := segmenter.Segment(text)
	out:=sego.SegmentsToSlice(segments, false)
	//
	glog.Info("list:",out)
	list := []string{}
	for _,key := range out{
		if !isFilter(key){
			list = append(list,key)
		}
	}
	glog.Info("list1:",list)
	getRank(list)
}

func isFilter(key string) bool{
	if strings.TrimSpace(key)==""{
		return true
	}
	f,err :=os.Open("../vendor/github.com/huichen/sego/data/stopword.txt")
	if err!=nil{
		panic(err.Error())
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	stopword := make(map[string]int)
	for{
		line,err := rd.ReadString('\n')
		if err==nil{
			stopword[strings.TrimSpace(line)] = 1
		}else if err==io.EOF{
			break
		}
	}
	reg := regexp.MustCompile(`[\pP]+`)
	ret := reg.FindAllString(key, -1)
	if len(ret)>0 {
		return true
	}else {
		_,exist := stopword[key]
		if exist{
			return true
		}else {
			return false
		}
	}
}

func getRank(list []string){
	length:=len(list)
	score := make(map[string]float64)
	worlds := make(map[string]map[string]int)
	for i:= 0;i<length;i++{
		worlds[list[i]] = make(map[string]int)
		for j := 1;j<=5;j++ {
			if (i+j)<length {
				//去重
				worlds[list[i]][list[i + j]] = 1
			}
			if (i-j)>=0{
				//去重
				worlds[list[i]][list[i - j]] = 1
			}
			//for k,v :=range worlds{
			//	glog.Info("j:",j,"|key:",k,"|value:",v)
			//}
		}
	}
	for k,v :=range worlds{
		glog.Info("key:",k,"|value:",v)
	}
	for i := 0;i<=200;i++ {
		m := make(map[string]float64)
		max_diff := float64(0)
		for k, items := range worlds {
			m[k] = float64(1 - 0.85)
			for v, count:= range items {

				size := len(worlds[k])*count
				if k == v || size == 0 {
					continue
				}

				_, exist := score[v]
				if exist {
					m[v] = m[v] + float64(0.85) / float64(size) * score[k]
				} else {
					m[v] = m[v]
				}
			}
			_, exist := score[k]
			if exist {
				if max_diff<math.Abs(float64(m[k])-float64(score[k])){
					max_diff = math.Abs(m[k]-score[k])
				}
			} else {
				if max_diff<m[k]{
					max_diff = m[k]
				}
			}
		}
		score = m;
		if max_diff <= 0.001 {
			break
		}
	}
	glog.Info("worlds:",score)

}