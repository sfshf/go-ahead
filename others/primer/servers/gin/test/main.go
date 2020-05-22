package main

import (
    "github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

    r := gin.Default()
    r.GET("/ping", func (c *gin.Context) {

        c.String(200, "pong")

    })

    return r

}

/*
    The `net/http/httptest` package is preferable way for HTTP testing.
*/
func main() {

    r := setupRouter()
    r.Run()

}
