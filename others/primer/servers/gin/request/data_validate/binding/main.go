package main

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

type Person struct {
    Name string         `form:"name" uri:"name" binding:"required"`
    Age int             `form:"age" uri:"age" binding:"required,gt=10"`
    Address string      `form:"address" uri:"address" binding:"required"`
    Birthday time.Time  `form:"birthday" time_format:"2006-01-02"`
}

func main() {

    r := gin.Default()

    r.GET("/get", func (c *gin.Context) {
        var person Person
        if err := c.ShouldBind(&person); err != nil {
            c.String(http.StatusBadRequest, "%s\n", err.Error())
        } else {
            c.JSON(http.StatusOK, person)
            c.String(http.StatusOK, "\n")
        }
    })

    r.GET("/get/:name/:age/:address", func(c *gin.Context) {
        var person Person
        if err := c.ShouldBindUri(&person); err != nil {
            c.String(http.StatusBadRequest, "%s\n", err.Error())
        } else {
            c.JSON(http.StatusOK, person)
            c.String(http.StatusOK, "\n")
        }
    })

    r.Run(":8080")

}

// Test it with:
// $ curl -X GET 'http://localhost:8080/get?name=wang&address=Shanghai'
// $ curl -X GET 'http://localhost:8080/get?name=wang&age=10&address=Shanghai'
// $ curl -X GET 'http://localhost:8080/get?name=wang&age=11&address=Shanghai'
// $ curl -X GET 'http://localhost:8080/get?name=wang&age=11&address=Shanghai&birthday=2017-05-03'
// $ curl -X GET 'http://localhost:8080/get/wang/10/Shanghai'
// $ curl -X GET 'http://localhost:8080/get/wang/11/Shanghai'
