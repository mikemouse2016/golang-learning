package main

import(
    "gopkg.in/redis.v3"
    "fmt"
)

var client = GetRedis()

func GetRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	PONG, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(PONG)
	return client
}

func main(){
     v:=[]string{}

     v=append(v, "a:3278")
     v=append(v, "a")
     b := [][]byte{[]byte("a"), []byte("a:3278")}
//     fmt.Println(client.Get("a:3278").Result())
     fmt.Println(client.MGet(v...).Result())
}
