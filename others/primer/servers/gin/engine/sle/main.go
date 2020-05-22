package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/autotls"
    "golang.org/x/crypto/acme/autocert"
)

func main() {

    // main1()
    main2()

}

/*
    example for 1-line LetsEncrypt HTTPS servers.
*/
func main1() {

    r := gin.Default()

    // Ping handler
    r.GET("/ping", func (c *gin.Context) {
        c.String(200, "pong")
    })

    log.Fatal(autotls.Run(r, "www.example1.com", "www.example2.com"))

}

/*
    example for custom autocert manager.
*/
func main2() {

    r := gin.Default()

    // Ping handler
    r.GET("/ping", func (c *gin.Context) {
        c.String(200, "pong")
    })

    m := autocert.Manager{
        Prompt: autocert.AcceptTOS,
        HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
        Cache: autocert.DirCache("/var/www/.cache"),
    }

    log.Fatal(autotls.RunWithManager(r, &m))

}
