package main

import (
    "github.com/gin-gonic/gin"
)

func main() {

    r := gin.Default()

    r.Static("/static", "./static")
    r.StaticFS("/assets", gin.Dir("./assets", true))
    r.StaticFile("/favicon.ico", "./static/img/favicon.ico")

    r.Run(":8080")

}

// Test it with:
// $ curl -X GET 'http://localhost:8080/static/images/favicon.ico'
// $ curl -X GET 'http://localhost:8080/assets'
// $ curl -X GET 'http://localhost:8080/favicon.ico'
