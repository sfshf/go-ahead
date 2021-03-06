package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

/*

    Using `JSONP` to request data from a server in a different domain. Add `callback` to response body if the query parameter `callback` exists.

*/
func main() {

    r := gin.Default()

    r.HTML()

    r.GET("/JSONP?callback=x", func (c *gin.Context) {
        data := map[string]interface{} {
            "foo": "bar",
        }
        // callback is x
        // Will output: x({\"foo\":\"bar\"})
        c.JSONP(http.StatusOK, data)
    })

    // Listen and serve on 0.0.0.0:8080
    r.Run()

}
