package main

import (
    "io/ioutil"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func getHandler1(c *gin.Context) {
    name := c.Query("name")
    age := c.DefaultQuery("age", "18")
    if v, err := strconv.Atoi(age); err != nil {
        c.String(http.StatusBadRequest, "%s\n", err.Error())
    } else {
        c.JSON(http.StatusOK, gin.H{"name": name, "age": v})
        c.String(http.StatusOK, "\n")
        // c.Abort()
    }
}

func getHandler2(c *gin.Context) {
    bytes, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
        c.String(http.StatusBadRequest, "%s\n", err.Error())
    } else {
        c.String(http.StatusOK, "getHandler2: %s\n %s\n %s\n",
                    c.Query("name"), c.Query("age"), string(bytes))
    }
}

func main() {

    r := gin.Default()

    r.GET("/get", getHandler1, getHandler2)

    r.Run(":8080")

}

// Test it with:
// $ curl -X GET 'http://localhost:8080/get?name=wang&age=123'
// $ curl -X GET 'http://localhost:8080/get?name=wang'
// $ curl -X GET 'http://localhost:8080/get?name=wang&age=qwer'
