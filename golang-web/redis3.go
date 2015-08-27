package main

import (
    "log"
    "time"
    "github.com/garyburd/redigo/redis"
)

func main() {
    c, _ := redis.Dial("tcp", "203.166.162.94:8888")
    n, err := c.Do("INFO")
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("info=%s", n)
    /*
    log.Println(time.Now())
    for i :=0;i<100000;i++{
        c.Do("sadd", "demo3", i)
    }
    */
    c.Do("set", "a", "123")
    n, err = c.Do("get", "a")
    if err != nil {
         log.Fatal(err)
    }
    log.Println(time.Now())
    for i :=0;i<10000;i++{
        c.Do("hset", "demo4", i, i)
    }
    log.Println(time.Now())
    for i :=0;i<10000;i++{
        c.Do("hget", "demo4", i)
    }
    log.Println(time.Now())
}
