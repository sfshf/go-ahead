package main

import (
    "bytes"
    "net/http"
    "strconv"
    "io/ioutil"

    "github.com/gin-gonic/gin"
)

func main() {

    r := gin.Default()

    r.POST("/post1", func(c *gin.Context) {

        name := c.PostForm("name")
        age := c.DefaultPostForm("age", "18")
        if v, err := strconv.Atoi(age); err != nil {
            c.String(http.StatusBadRequest, "%s\n", err.Error())
        } else {
            c.JSON(http.StatusOK, gin.H{"name": name, "age": v})
            c.String(http.StatusOK, "\n")
        }

    })

    r.POST("/post2", func(c *gin.Context) {
        if bytes, err := ioutil.ReadAll(c.Request.Body); err != nil {
            c.String(http.StatusBadRequest, "%s\n", err.Error())
        } else {
            c.String(http.StatusOK, "%s\n", string(bytes))
        }
    })

    r.POST("/post3", func(c *gin.Context) {
        bytes, err := ioutil.ReadAll(c.Request.Body)
        if err != nil {
            c.String(http.StatusBadRequest, "%s\n", err.Error())
        } else {
            c.String(http.StatusOK, "%s\n", string(bytes))
        }
        c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bytes))
        name := c.PostForm("name")
        age := c.DefaultPostForm("age", "18")
        if v, err := strconv.Atoi(age); err != nil {
            c.String(http.StatusBadRequest, "%s\n", err.Error())
        } else {
            c.JSON(http.StatusOK, gin.H{"name": name, "age": age})
            c.String(http.Status, "\n")
        }
    })

    r.Run(":8080")

}

// Test it with:
// $ curl -X POST 'http://localhost:8080/post1' -d 'name=wang'
// $ curl -X POST 'http://localhost:8080/post1' -d 'name=wang&age=123'
// $ curl -X POST 'http://localhost:8080/post1' -d 'name=wang&age=qwer'
// $ curl -X POST 'http://localhost:8080/post2' -d 'name=wang&age=qwer'
// $ curl -X POST 'http://localhost:8080/post3' -d 'name=wang&age=123'
// $ curl -X POST 'http://localhost:8080/post3' -d 'name=wang&age=qwer'
