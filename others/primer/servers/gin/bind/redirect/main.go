package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {

    // main1()
    main2()

}

/*
    Issuing a HTTP redirect is easy. Both internal and external locations are supported.
*/
func main1() {

    r := gin.Default()

    r.GET("/test", func (c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
    })

    r.Run()

}

/*
    Issuing a Router redirect, use HandleContext like below.
*/
func main2() {

    r := gin.Default()

    r.GET("/test", func (c *gin.Context) {
        c.Request.URL.Path = "/test2"
        r.HandleContext(c)
    })

    r.GET("test2", func (c *gin.Context) {
        c.JSON(200, gin.H { "hello": "world" })
    })

    r.Run()

}
