package main

import (
    "github.com/robfig/cron"
    "log"
    "time"
)

type ff func()

func (f ff) Run(){
    f()
}

func (f func()) Run(){
    f()
}

func main() {
    log.Println(
    struct {
    a int
    b int}{a:1,b:2})
    log.Println(time.Now().Local())
    i := 0
    c := cron.New()
    spec := "0 */2 * * * ?"
    spec2 := "2 * * * * *"
    c.AddFunc(spec, func() {
        i++
        log.Println("cron running:", i)
        log.Println(time.Now())
    })
    c.AddFunc(spec2, func() {
        log.Println("deme2", time.Now())
    })
    c.Start()
    ff(func(){
        log.Println("demo")
    }).Run()
    select{}
}
