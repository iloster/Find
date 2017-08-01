package zhihu

import (
	"encoding/json"
	"fmt"
	"spider/db"
	"strconv"
	"github.com/golang/glog"
	"io/ioutil"
	"spider/utils"
)
var Source_Zhihu = 3
//ANSWER_CREATE  回答问题
//ANSWER_VOTE_UP 赞同问题
//QUESTION_FOLLOW 关注了问题
//TOPIC_FOLLOW  关注了话题
//MEMBER_FOLLOW_COLLECTION 关注收藏夹
//QUESTION_CREATE 提出问题
//MEMBER_VOTEUP_ARTICLE 赞了文章
//MEMBER_FOLLOW_ROUNDTABLE  关注了圆桌
//MEMBER_CREATE_ARTICLE  发表了文章
//知乎动态的结构

const (
	Verb_ANSWER_CREATE = "ANSWER_CREATE"
	Verb_ANSWER_VOTE_UP = "ANSWER_VOTE_UP"
	Verb_QUESTION_FOLLOW = "QUESTION_FOLLOW"
	Verb_TOPIC_FOLLOW = "TOPIC_FOLLOW"
	Verb_MEMBER_FOLLOW_COLLECTION = "MEMBER_FOLLOW_COLLECTION"
	Verb_QUESTION_CREATE = "QUESTION_CREATE"
	Verb_MEMBER_VOTEUP_ARTICLE = "MEMBER_VOTEUP_ARTICLE"
	Verb_MEMBER_FOLLOW_ROUNDTABLE = "MEMBER_FOLLOW_ROUNDTABLE"
	Verb_MEMBER_CREATE_ARTICLE = "MEMBER_CREATE_ARTICLE"
)
type ZhihuActivity struct {
	Paging  ZhihuPaging  `json:"paging"`
	Data    []ZhihuData	`json:"data"`
}
type ZhihuPaging struct {
	Next  string `json:"next"`
	Previous string `json:"previous"`
}
type ZhihuData struct {
	Verb   		string  	`json:"verb"`
	Target 		ZhihuTarget 	`json:"target"`
	ActionType 	string 		`json:"action_text"`
	CreateTime 	int 		`json:"created_time"`
}
type ZhihuTarget struct {
	Author     ZhihuAuthor    `json:"author"`
	Question   ZhihuQuestion  `json:"question"`
	Excerpt    string         `json:"excerpt"`
	Name       string         `json:"name"`          //圆桌会议的标题
	Description string        `json:"description"`   //圆桌会议的描述
	Url        string         `json:"url"`
	Title      string         `json:"title"`     //文章的标题
	Id         json.Number           `json:"id,Number"`
}
type ZhihuQuestion struct {
	Title string 	`json:"title"`
	Id     int       `json:"id"`
}
type ZhihuAuthor struct {
	Name   string  `json:"name"`    //回答者昵称
	Avatar string `json:"avatar_url"` //回答者头像
}

func Start(id int,url string)(int,int){
	res,err:=GetSession().Get(url)
	glog.Info("zhihu url:",url,"| id:",id)
	successNum := 0
	failedNum := 0
	if err == nil{
		bodyByte, _ := ioutil.ReadAll(res.Body)
		resStr := string(bodyByte)
		//glog.Info(resStr)
		ret := ZhihuActivity{}
		err := json.Unmarshal([]byte(resStr),&ret)
		if err!=nil{
			glog.Info("err",err)
		}
		if err == nil{
			for _,item := range ret.Data {
				var title = ""
				var desc = ""
				var link = ""
				switch item.Verb {
				case Verb_ANSWER_CREATE:{
					//	回答了问题
					title ="回答了："+item.Target.Question.Title + "的问题"
					desc = item.Target.Excerpt
					link = fmt.Sprintf("https://www.zhihu.com/question/%d/answer/%s",item.Target.Question.Id,string(item.Target.Id))
				}
				case Verb_ANSWER_VOTE_UP:{
					//赞同问题
					title ="赞同了:" + item.Target.Question.Title + "--"+item.Target.Author.Name+"的回答"
					desc  = item.Target.Excerpt
					link = fmt.Sprintf("https://www.zhihu.com/question/%d/answer/%s",item.Target.Question.Id,string(item.Target.Id))
				}
				case Verb_MEMBER_FOLLOW_ROUNDTABLE:{
					title = "关注了圆桌:" + item.Target.Name
					desc = item.Target.Description
					link = fmt.Sprintf("https://www.zhihu.com/roundtable/%s",string(item.Target.Id))
				}
				case Verb_MEMBER_VOTEUP_ARTICLE:{
					title = "赞了文章:" + item.Target.Title + "--" + item.Target.Author.Name
					desc = item.Target.Excerpt
					link = fmt.Sprintf("https://zhuanlan.zhihu.com/p/%s",string(item.Target.Id))
				}
				case Verb_QUESTION_FOLLOW:{
					//	关注了问题
					title = "关注了问题:"+item.Target.Title
					desc = ""
					link = fmt.Sprintf("https://www.zhihu.com/question/%s",string(item.Target.Id))
				}
				case Verb_MEMBER_FOLLOW_COLLECTION:{
					//关注收藏夹
					title = "关注了收藏夹:"+item.Target.Title
					desc = ""
					link = fmt.Sprintf("https://www.zhihu.com/collection/%s",string(item.Target.Id))
				}
				case Verb_TOPIC_FOLLOW:{
					//关注了话题
					title = "关注了话题:"+item.Target.Name
					desc = ""
					link = fmt.Sprintf("https://www.zhihu.com/topic/%s/hot",string(item.Target.Id))
				}
				case Verb_QUESTION_CREATE:{
					//提出问题
					title = "提出了问题:" + item.Target.Title
					desc = ""
					link = fmt.Sprintf("https://www.zhihu.com/question/%s",string(item.Target.Id))
				}
				case Verb_MEMBER_CREATE_ARTICLE:{
					//发表了文章
					title = "发表了文章:"+item.Target.Title
					desc = item.Target.Excerpt
					link = fmt.Sprintf("https://zhuanlan.zhihu.com/p/%s",string(item.Target.Id))
				}


				}
				if !db.GetDB().IsExistTimeLineByLink(db.Table_ZhiHu,link){
					_, err = db.GetDB().InsertTimeLineZhiHu(id, title, utils.SubString(desc,0,1000), link, item.Verb, strconv.Itoa(item.CreateTime))
					if err == nil {
						//glog.Info("[Success] zhihu title:", title, "| description:", desc, "| link:", link, "| pub_data:", item.CreateTime)
						successNum++
					} else {
						failedNum++
						glog.Info("[Error] zhihu title:", title, "| description:", desc, "| link:", link, "| pub_data:", item.CreateTime, "|error:", err.Error())
					}
				}else{
					//glog.Info("[Exist] zhihu title:", title, "| description:", desc, "| link:", link, "| pub_data:", item.CreateTime)
				}
			}
		}
	}else{
		glog.Info(err.Error())

	}
	return successNum,failedNum
}