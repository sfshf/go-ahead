package main

import (
    "log"

    "github.com/gin-gonic/gin"
)

type Person struct {
    Name string `form:"name"`
    Address string `form:"address"`
}

func startPage(c *gin.Context) {
    var person Person
    if c.ShouldBindQuery(&person) == nil {
        log.Println("====== Only Bind By Query String ======")
        log.Println(person.Name)
        log.Println(person.Address)
    }
    c.String(200, "Success")
}

/*

    `ShouldBindQuery` function only binds the query params and not the post data.

*/
func main() {

    route := gin.Default()
    route.Any("/testing", startPage)
    route.Run()

}

// Test it with:
// $ curl -X GET localhost:8080/testing?name=zhangsan\&address=beijing
// $ curl -X POST -F name=lisi -F address=shanghai localhost:8080/testing
