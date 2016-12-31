package main

import ( 
    "fmt" 
    "os" 
    "os/signal" 
)

func main() { 
    c := make(chan os.Signal) 
    signal.Notify(c)

    //signal.Notify(c, syscall.SIGHUP, syscall.SIGUSR2)  //监听指定信号

    s := <-c //阻塞直至有信号传入 
    fmt.Println("get signal:", s) 
}