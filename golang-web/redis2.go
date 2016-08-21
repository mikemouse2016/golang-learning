package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func main() {
	conn, err := redis.DialTimeout("tcp", ":6380", 0, 1*time.Second, 1*time.Second)
	if err != nil {
		fmt.Println("error")
		return
	}
	size, err := conn.Do("DBSIZE")
	fmt.Printf("size is %d \n", size)
	_, err = conn.Do("SET", "user:user0", 123)
	_, err = conn.Do("SET", "user:user1", 456)
	_, err = conn.Do("APPEND", "user:user0", 87)

	user0, err := redis.Int(conn.Do("GET", "user:user0"))
	user1, err := redis.Int(conn.Do("GET", "user:user1"))
	user3, err := redis.Int(conn.Do("GET", "user:user3"))
	if err != nil {
		fmt.Println("1111", err)
	}
	fmt.Println(user3)
	likeList := fmt.Sprintf("ledis:like:list:%d", 1)
	likeListSet, err := redis.Ints(conn.Do("ZRANGE", likeList, 0, -1))
	if err != nil {
		fmt.Println("error like:%d", 1)
		fmt.Println(err.Error())
	}
	fmt.Println("likeList", likeListSet)

	fmt.Printf("user0 is %d , user1 is %d \n", user0, user1)

	size, _ = conn.Do("DBSIZE")
	fmt.Printf("size is %d \n", size)

	var value1 int
	var value2 string
	if _, err := redis.Scan(0, &value1, &value2); err != nil {
		// handle error
	}
	fmt.Println(value1, value2)
	conn.Close()
}
