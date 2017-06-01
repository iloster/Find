package zhihu

import "encoding/json"

//ANSWER_CREATE  回答问题
//ANSWER_VOTE_UP 赞同问题
//QUESTION_FOLLOW 关注了问题
//TOPIC_FOLLOW  关注了话题
//MEMBER_FOLLOW_COLLECTION 关注收藏夹
//QUESTION_CREATE 提出问题
//MEMBER_VOTEUP_ARTICLE 赞了文章
//MEMBER_FOLLOW_ROUNDTABLE  关注了圆桌
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