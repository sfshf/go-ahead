package main

import (
    "github.com/gin-gonic/gin"
)

/*
By default, logs output on console should be colorized depending on the detected TTY.
*/

func main() {

    main1()
    // main2()

}

// Never colorize logs:
func main1() {

    // Disable log's color
    gin.DisableConsoleColor()

    // Creates a gin router with default middleware:
    // logger and recovery (crash-free) middleware.
    router := gin.Default()

    router.GET("ping", func (c *gin.Context) {
        c.String(200, "pong")
    })

    router.Run()

}

// Always colorize logs:
func main2() {

    // Force log's color
    gin.ForceConsoleColor()

    // Creates a gin router with default middleware:
    // logger and recovery (crash-free) middleware.
    router := gin.Default()

    router.GET("/ping", func (c *gin.Context) {
        c.String(200, "pong")
    })

    router.Run()

}
