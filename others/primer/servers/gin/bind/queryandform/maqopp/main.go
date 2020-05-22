package main

import (
    "fmt"

    "github.com/gin-gonic/gin"
)

/*

POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
Content-Type: application/x-www-form-urlencoded

names[first]=thinkerou&names[second]=tianou

*/
func main() {

    router := gin.Default()

    router.POST("/post", func (c *gin.Context) {

        ids := c.QueryMap("ids")
        names := c.PostFormMap("names")
        fmt.Printf("ids: %v; names: %v", ids, names)

    })

    router.Run()

}

// Test it with:
// curl -X POST -H "Content-Type=application/x-www.form-urlencoded" -d "names[first]=thinkerou&names[second]=tianou" localhost:8080/post?ids[a]=1234\&ids[b]=hello
