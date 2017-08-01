package juejin

var BaseUrl = "https://timeline-merger-ms.juejin.im/v1/get_entry_by_self?src=web&targetUid=%s&type=post&before&limit=20&order=createdAt"

type JuejinD struct {
	Total  int  `json:"total"`
	EntryList []JuejinEntry `json:"entrylist"` 
}

type JuejinEntry struct {

} 

type JuejinData struct {
	S  int `json:"s"`
	M  string `json:"m"`
	D  JuejinD `json:"d"`
}

func Start(id int,url string){

}