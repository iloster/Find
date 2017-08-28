package utils


type MapSorter []Item

type Item struct {
	Key string
	Val float64
}

func NewMapSorter(m map[string]float64) MapSorter {
	ms := make(MapSorter, 0, len(m))

	for k, v := range m {
		ms = append(ms, Item{k, v})
	}

	return ms
}


func (ms MapSorter) Len() int {
	return len(ms)
}

func (ms MapSorter) Less(i, j int) bool {
	return ms[i].Val > ms[j].Val // 按值排序
	//return ms[i].Key < ms[j].Key // 按键排序
}

func (ms MapSorter) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}


//func SegWords()  {
//	var segmenter sego.Segmenter
//	segmenter.LoadDictionary("../vendor/github.com/huichen/sego/data/dictionary.txt")
//
//	// 分词
//	text := []byte("这跟非阻塞发送一点关系都没有，就算你用阻塞 IO ，对方一样可能分多次收到数据。这是 TCP 这个字节流协议的固有性质。")
//	segments := segmenter.Segment(text)
//	out:=sego.SegmentsToSlice(segments, false)
//	//
//	glog.Info("list:",out)
//	list := []string{}
//	for _,key := range out{
//		if !isFilter(key){
//			list = append(list,key)
//		}
//	}
//	glog.Info("list1:",list)
//	getTopWords(getRank(list),3)
//
//}
//
//func isFilter(key string) bool{
//	if strings.TrimSpace(key)==""{
//		return true
//	}
//	if len(key) == 3 {
//		return true
//	}
//	f,err :=os.Open("../vendor/github.com/huichen/sego/data/stopword.txt")
//	if err!=nil{
//		panic(err.Error())
//	}
//	defer f.Close()
//	rd := bufio.NewReader(f)
//	stopword := make(map[string]int)
//	for{
//		line,err := rd.ReadString('\n')
//		if err==nil{
//			stopword[strings.TrimSpace(line)] = 1
//		}else if err==io.EOF{
//			break
//		}
//	}
//	reg := regexp.MustCompile(`[\pP]+`)
//	ret := reg.FindAllString(key, -1)
//	if len(ret)>0 {
//		return true
//	}else {
//		_,exist := stopword[key]
//		if exist{
//			return true
//		}else {
//			return false
//		}
//	}
//}
//
//func getRank(list []string) map[string]float64{
//	length:=len(list)
//	score := make(map[string]float64)
//	worlds := make(map[string]map[string]int)
//	for i:= 0;i<length;i++{
//		worlds[list[i]] = make(map[string]int)
//		for j := 1;j<=5;j++ {
//			if (i+j)<length {
//				//去重
//				worlds[list[i]][list[i + j]] = 1
//			}
//			if (i-j)>=0{
//				//去重
//				worlds[list[i]][list[i - j]] = 1
//			}
//			//for k,v :=range worlds{
//			//	glog.Info("j:",j,"|key:",k,"|value:",v)
//			//}
//		}
//	}
//	for k,v :=range worlds{
//		glog.Info("key:",k,"|value:",v)
//	}
//
//	for i := 0;i<=200;i++ {
//		m := make(map[string]float64)
//		max_diff := float64(0)
//		for k, items := range worlds {
//			m[k] = float64(1 - 0.85)
//			for v, count := range items {
//
//				size := len(worlds[v])*count
//				if k == v || size == 0 {
//					continue
//				}
//
//				_, exist := score[v]
//				if exist {
//					m[k] = m[k] + float64(0.85) / float64(size) * score[v]
//				} else {
//					m[k] = m[k]
//				}
//			}
//			_, exist := score[k]
//			if exist {
//				if max_diff<math.Abs(float64(m[k])-float64(score[k])){
//					max_diff = math.Abs(m[k]-score[k])
//				}
//			} else {
//				if max_diff<m[k]{
//					max_diff = m[k]
//				}
//			}
//		}
//		score = m;
//		if max_diff <= 0.001 {
//			break
//		}
//	}
//	return score
//}
//
////获取分数最高的几个数字
//func getTopWords(score map[string]float64,n int){
//	ms := NewMapSorter(score)
//	sort.Sort(ms)
//	for k,v :=range ms{
//		glog.Info("k:",k,"|v:",v)
//	}
//}