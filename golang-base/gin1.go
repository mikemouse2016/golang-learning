package main

import "github.com/gin-gonic/gin"
import "fmt"

func main() {
    r := gin.Default()
    r.POST("/v2/favorite/list", func(c *gin.Context) {
        c.MustGet("ddd")
        fmt.Println(c.Params)
        fmt.Println(c.Keys)
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run() // listen and server on 0.0.0.0:8080
}
