package main

import (
	"fmt"
)

func main() {
	a := []int{}
	fmt.Println(len(a), cap(a), a)
	a = append(a, 1)
	fmt.Println(len(a), cap(a), a)
	a = append(a, 10)
	fmt.Println(len(a), cap(a), a)
}

curl -XPOST localhost:9200/test -d '{
  "settings":{
      "number_of_shards":1
  },
  "mappings":{
      "type1":{
          "_source":{"enabled":false},
          "preperties":{
              "field1":{"type":"string",
                      "index":"not_analyzed"
              }
          }
      }
  }
}'

curl -XPUT 'http://localhost:9200/twitter/tweet/1'-d '{
  "user":"kimchy",
  "post_date":"2012-12-12",
  "message":"trying out ElasticSearch!"
}'
