package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {

    r := gin.Default()

    r.GET("/get", func(c *gin.Context) {
        c.String(http.StatusOK, "GET: ok!\n")
    })

    r.GET("/get/params", func(c *gin.Context) {
        c.String(http.StatusOK, "GET: name -- %s\n", c.Query("name"))
        c.String(http.StatusOK, "GET: password -- %s\n", c.Query("password"))
        c.String(http.StatusOK, "GET: birthday -- %s\n", c.Query("birthday"))
        c.String(http.StatusOK, "GET: ok!\n")
    })


    r.POST("/post", func(c *gin.Context) {
        c.String(http.StatusOK, "POST: ok!\n")
    })

    r.POST("/post/params", func(c *gin.Context) {
        c.String(http.StatusOK, "POST: name -- %s\n", c.PostForm("name"))
        c.String(http.StatusOK, "POST: password -- %s\n", c.PostForm("password"))
        c.String(http.StatusOK, "POST: birthday -- %s\n", c.PostForm("birthday"))
        c.String(http.StatusOK, "POST: ok!\n")
    })

    r.PUT("/put", func(c *gin.Context) {
        c.String(http.StatusOK, "PUT: ok!\n")
    })

    r.HEAD("/head", func(c *gin.Context) {
        c.String(http.StatusOK, "HEAD: ok!\n")
    })

    r.DELETE("/delete", func(c *gin.Context) {
        c.String(http.StatusOK, "DELETE: ok!\n")
    })

    r.PATCH("/patch", func(c *gin.Context) {
        c.String(http.StatusOK, "PATCH: ok!\n")
    })

    r.OPTIONS("/options", func(c *gin.Context) {
        c.String(http.StatusOK, "OPTIONS: ok!\n")
    })

    r.ANY("/any", func(c *gin.Context) {
        c.String(http.StatusOK, "ANY: ok!\n")
    })

    r.Run(":8080")

}

// Test it with:
// $ curl -X GET 'http://localhost:8080/get'
// $ curl -X GET 'http://localhost:8080/get/params?name=wang&password=000000&birthday=1994-07-12'
// $ curl -X POST 'http://localhost:8080/post'
// $ curl -X POST 'http://localhost:8080/post/params' -d 'name=jiang' -d 'password=123456' -d 'birthday=1997-05-21'
// ... ...
