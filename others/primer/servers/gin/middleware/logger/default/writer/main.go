package main

import (
    "os"
    "io"

    "github.com/gin-gonic/gin"
)

func main() {

    // Disable Console Color, you don't need console color when writing the logs to file.
    gin.DisableConsoleColor()

    // Logging to a file.
    f, _ := os.Create("./gin_deprecated.log")
    gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

    // Use the following code if you need to write the logs to file and console at the same time.
    // gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
    router := gin.Default()
    router.GET("ping", func (c *gin.Context) {
        c.String(200, "pong\n")
    })

    router.Run()

}
