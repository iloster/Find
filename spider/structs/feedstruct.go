package structs

import "encoding/xml"

type Feed struct {
	Rss xml.Name `xml:"rss"`
	Version string `xml:"version,attr"`
	Channel Channel `xml:"channel"`
}
type Channel struct {
	Title string `xml:"title"`
	Items [] Item `xml:"item"`
}
type Item struct {
	Title string `xml:"title"`
	Link string `xml:"link"`
	PubDate string `xml:"pubDate"`
	Description string `xml:"description"`
}

type Sql struct {
	Table string
	Data  interface{}
}

type TimeLine struct {
	Id  int `field:"id"`
	UserId int `field:"userid"`
	Title string `field:"title"`
	Description string `field:"description"`
	Link string `field:"link"`
	PubData string `field:"pub_data"`
}

//知乎动态的结构
type ZhihuActivity struct {
	Paging  ZhihuPaging  `json:"paging"`
	Data    []ZhihuData	`json:"data"`
}
type ZhihuPaging struct {
	Next  string `json:"next"`
	Previous string `json:"previous"`
}
type ZhihuData struct {
	Target ZhihuTarget `json:"target"`

	ActionType string `json:"action_text"`
	CreateTime int `json:"created_time"`
}
type ZhihuTarget struct {
	Excerpt string `json:"excerpt"`
	Name string `json:"name"`
	Url  string `json:"url"`
}