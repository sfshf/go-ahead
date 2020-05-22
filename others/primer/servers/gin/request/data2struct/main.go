package main

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

type Person struct {
    Name string         `form:"name" uri:"name"`
    // 此处重点记忆
    Birthday time.Time  `form:"birthday" uri:"birthday" time_format:"2006-01-02"`
    Address string      `form:"address" uri:"address"`
}

func testing(c *gin.Context) {

    contentType := c.GetHeader("Content-Type")
    var person Person
    if err := c.ShouldBind(&person); err != nil {
        c.String(http.StatusBadRequest, "%s\n", err.Error())
    } else {
        c.String(http.StatusOK, "%s\n", contentType)
        c.String(http.StatusOK, "%v\n", person)
    }

}

func urltesting(c *gin.Context) {
    var person Person
    if err := c.ShouldBindUri(&person); err != nil {
        c.String(http.StatusBadRequest, "%s\n", err.Error())
    } else {
        c.String(http.StatusOK, "%v\n", person)
    }
}

func main() {

    r := gin.Default()

    r.GET("/get", testing)
    r.GET("/url/:name/:birthday/:address", urltesting)
    r.POST("/post", testing)

    r.Run(":8080")

}

// Test it with:
// $ curl -X GET 'http://localhost:8080/get?name=wang&birthday=1995-06-23&address=Shanghai'
// $ curl -X POST 'http://localhost:8080/post' -d 'name=zhang&birthday=2001-11-03&address=Beijing'
// $ curl -X POST 'http://localhost:8080/post' -H 'Content-Type:application/json' -d '{"name":"zhang","address":"Beijing"}'
// $ curl -X GET 'http://localhost:8080/url/wang/2019-05-05/Shanghai'

// 此处测试，访问失败，返回结果有报错，具体原因不详。
// $ curl -X POST 'http://localhost:8080/post' -H 'Content-Type:application/json' -d '{"name":"zhang","birthday":"2001-11-03","address":"Beijing"}'
