package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

/*
使用 AsciiJSON来生成转义了非Ascii字符的纯ASCII编码的JSON对象。
*/
func main() {

    r := gin.Default()
    r.GET("/someJSON", func (c *gin.Context) {
        data := map[string]interface{} {
            "lang": "GO语言",
            "tag": "<br>",
        }

        c.AsciiJSON(http.StatusOK, data)

    })

    r.Run()

}
