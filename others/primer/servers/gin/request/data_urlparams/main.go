package main

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func testing(c *gin.Context) {

    name := c.Param("name")
    age := c.Param("age")
    if v, err := strconv.Atoi(age); err != nil {
        c.String(http.StatusBadRequest, "%s\n", err.Error())
    } else {
        c.String(http.StatusOK, "name: %s\n", name)
        c.String(http.StatusOK, "age: %d\n", v)
    }

}

func main() {

    r := gin.Default()

    r.GET("/get/:name/:age", testing)
    r.POST("/post/:name/:age", testing)

    r.Run(":8080")

}

// Test it with:
// curl -X GET 'http://localhost:8080/get/wang/123'
// curl -X GET 'http://localhost:8080/get/wang/asdf'
// curl -X POST 'http://localhost:8080/post/zhang/123' -d 'name=wang&age=321'
// curl -X POST 'http://localhost:8080/post/zhang/asdf' -d 'name=wang&age=321'
