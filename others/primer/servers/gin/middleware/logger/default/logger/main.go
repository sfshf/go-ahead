package main

import (
    "time"
    "log"

    "github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {

    return func (c *gin.Context) {

        t := time.Now()

        // Set example variable
        c.Set("example", "12345")

        // before request
        log.Printf("1:%v\n", "before request")

        c.Next()

        // after request
        latency := time.Since(t)
        log.Printf("2:%v\n", latency)

        // access the status we are sending
        status := c.Writer.Status()
        log.Printf("3:%v\n", status)

    }

}

func main() {

    r := gin.Default()
    r.Use(Logger())

    r.GET("/", func (c *gin.Context) {
        example := c.MustGet("example").(string)

        // it would print "12345"
        log.Printf("4:%v\n", example)
    })

    r.GET("/test/", func (c *gin.Context) {
        log.Println("route '/test/'")
    })

    // Listen and serve on 0.0.0.0:8080
    r.Run(":8080")

}

// Test it with:
// $ curl localhost:8080/
// $ curl localhost:8080/test
// $ curl localhost:8080/test/
