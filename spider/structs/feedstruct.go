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