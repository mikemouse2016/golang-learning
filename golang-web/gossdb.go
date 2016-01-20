package main

import(
    "github.com/seefan/gossdb"
)

func main(){
pool, err := gossdb.NewPool(&gossdb.Config{
    Host:             "127.0.0.1",
    Port:             6380,
    MinPoolSize:      5,
    MaxPoolSize:      50,
    AcquireIncrement: 5,
})
if err != nil {
    log.Fatal(err)
    return
}


c, err := pool.NewClient()
if err != nil {
    log.Println(err.Error())
    return
}
defer c.Close()
c.Set("test","hello world.")
re, err := c.Get("test")
if err != nil {
    log.Println(err)
} else {
    log.Println(re, "is get")
}
//设置10 秒过期
c.Set("test1",1225,10)
//取出数据，并指定类型为 int
re, err = c.Get("test1")
if err != nil {
    log.Println(err)
} else {
    log.Println(re.Int(), "is get")
}
}
