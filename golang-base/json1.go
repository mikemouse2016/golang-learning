package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson" // for json get
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	//"regexp"
	//"strconv"
)

func Get(url string) (content string, statusCode int) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		statusCode = -100
		return
	}
	defer resp.Body.Close()
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		statusCode = -200
		return
	}
	js, err := simplejson.NewJson(data)
	if err != nil {
		panic(err.Error())
	}
	//weilist := []weidata{}[:]
	for _, v := range js.Get("jokes").MustArray() {
		prase(v)
		//append(weilist, prase(v))
		//fmt.Println(i, v)
	}
	//fmt.Print(weilist)
	statusCode = resp.StatusCode
	content = string(data)
	return
}

/*
{"video_uri": "",
"user_cover_url_100x100": "/img/user/cover/377313-815105_100x100.jpg",
"user_cover_url": "/img/user/cover/377313-815105_25x25.jpg",
"tags": [],
"Jokeid": 211421,
"user_name": "\u590f\u6d45\u82cd",
"created": "2015-10-10 10:26:03",
"Userid": 377313,
"next_joke_id": 211427,
"uri": "",
"content": "\u95ee\uff1a\u54ea\u4e00\u4ef6\u4e8b\uff0c\u8ba9\u4f60\u5bf9\u81ea\u5df1\u7684\u65e0\u77e5\u611f\u5230\u9707\u61be\uff1f\r\n\u7b54\uff1a\u5c0f\u7684\u65f6\u5019\u5199\u4f5c\u6587\u8981\u5199\u7b14\u540d\uff0c\u6211\u5199\u7684\u4e00\u76f4\u662f\u201c\u4e2d\u534e\u7ed8\u56fe\u94c5\u7b14\u201d\u3002\u771f\u7684\u662f\u7b14\u540d\u554a\u2026\u2026",
"unlike_count": 0,
"comment_count": 3,
"like_count": 1,
"pre_joke_id": 211420,
"user_id": 377313,
"type": "text",
"created_cn": "3\u5c0f\u65f6\u524d"
*/
/**interface的强制类型转换**/
func prase(m1 interface{}) (wei weidata) {
	var dat map[string]interface{}
	dat = m1.(map[string]interface{})
	wei = weidata{dat["uri"].(string), dat["user_name"].(string), dat["content"].(string), string(dat["Jokeid"].(json.Number))}
	fmt.Println(wei)
	return
}

type weidata struct {
	url     string
	title   string
	content string
	key     string
}

func main() {
	fmt.Println(`Get index ...`)
	_, statusCode := Get("http://lengxiaohua.com/lengxiaohuaapi/joke?action=getJokes&interval=weekly&sort=popular&type=text%7Cimage&start=0&limit=20")
	if statusCode != 200 {
		return
	}

}
