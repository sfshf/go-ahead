package main

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func main() {

    r := gin.Default()

    r.GET("/url2params/:name/:age", func(c *gin.Context) {
        name := c.Param("name")
        agestr := c.Param("age")
        if age, err := strconv.Atoi(agestr); err != nil {
            c.String(http.StatusBadRequest, "Invalid value of 'age' param -- '%s'\n", agestr)
        } else {
            c.JSON(http.StatusOK, gin.H{"name": name, "age": age})
        }

    })

    r.Run()

}

// Test it with:
// $ curl -X GET 'http://localhost:8080/url2params/jiang/asdf'
// $ curl -X GET 'http://localhost:8080/url2params/zhang/28'
