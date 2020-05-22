package main

import (
    "fmt"

    "github.com/gin-gonic/gin"
)

/*

POST /post?id=1234&page=1 HTTP/1.1
Content-Type: multipart/form-data

name=manu&message=this_is_great

*/
func main() {

    router := gin.Default()

    router.POST("/post", func (c *gin.Context) {

        id := c.Query("id")
        page := c.DefaultQuery("page", "0")
        name := c.PostForm("name")
        message := c.PostForm("message")
        fmt.Printf("id: %s; page: %s; name: %s; message: %s", id ,page, name, message)

    })

    router.Run()

}

// Test it with:
// $ curl -v -X POST -F 'name=manu' -F 'message=this_is_great' localhost:8080/post?id=1234\&page=1
