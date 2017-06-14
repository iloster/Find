package jianshu

import (
	"github.com/golang/glog"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"strings"
)

type JianShuData struct {
	Title    string
	Href     string
	Abstract string
	Pub_date string
	Img      string
}
func Start(url string){
	ret:=getData(url)
	for _,item :=range ret{
		glog.Info("jianshu item:", item)
	}
}

func getData(url string) []JianShuData{
	ret := []JianShuData{}
	for i:=1;i<5;i++ {
		pageurl := fmt.Sprintf(url + "?order_by=shared_at&page=%d",i)
		glog.Info("jianshu Start pageurl:",pageurl)
		doc, err := goquery.NewDocument(pageurl)
		if err == nil {
			doc.Find(".content").Each(func(i int, s *goquery.Selection) {
				item := JianShuData{}
				item.Title = s.Find(".title").Text()
				item.Href, _ = s.Find(".title").Attr("href")
				item.Abstract = strings.TrimSpace(s.Find(".abstract").Text())
				item.Pub_date, _ = s.Find(".time").Attr("data-shared-at")
				item.Img,_ = s.Find(".img-blur-done").Attr("src")

				ret = append(ret,item)
			})

		}else{
			glog.Info("jianshu url:",pageurl,"|err:",err.Error())
			break;
		}

	}
	return ret
}


