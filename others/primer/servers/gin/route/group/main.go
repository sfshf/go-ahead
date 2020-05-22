package main

import (
    "log"

    "github.com/gin-gonic/gin"
)

func loginEndpoint(c *gin.Context) {
    log.Println(c.Request.URL.Path)
}

func submitEndpoint(c *gin.Context) {
    log.Println(c.Request.URL.Path)
}

func readEndpoint(c *gin.Context) {
    log.Println(c.Request.URL.Path)
}

func main() {

    router := gin.Default()

    // Simple group: v1
    v1 := router.Group("/v1")
    {
        v1.POST("/login", loginEndpoint)
        v1.POST("/submit", submitEndpoint)
        v1.POST("/read", readEndpoint)
    }

    // Simple group: v2
    v2 := router.Group("/v2")
    {
        v2.POST("/login", loginEndpoint)
        v2.POST("/submit", submitEndpoint)
        v2.POST("/read", readEndpoint)
    }

    router.Run()

}

// Test it with:
// $ curl -X POST 'http://localhost:8080/v1/login'
// $ curl -X POST 'http://localhost:8080/v1/submit'
// $ curl -X POST 'http://localhost:8080/v1/read'

// $ curl -X POST 'http://localhost:8080/v2/login'
// $ curl -X POST 'http://localhost:8080/v2/submit'
// $ curl -X POST 'http://localhost:8080/v2/read'
