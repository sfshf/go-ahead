package main

import (
    "html/template"
    "log"

    "github.com/gin-gonic/gin"
)

/*

    http.Pusher is supported only go1.8+.

*/
func main() {

    r := gin.Default()
    r.Static("/assets", "./assets")
    r.SetHTMLTemplate(html)

    r.GET("/", func (c *gin.Context) {
        if pusher := c.Writer.Pusher(); pusher != nil {
            // use pusher.Push() to do server push
            if err := pusher.Push("/assets/app.js", nil); err != nil {
                log.Printf("Failed to push: %v", err)
            }
        }
        c.HTML(200, "https", gin.H{
            "status": "success",
        })
    })
    r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")

}

var html = template.Must(template.New("https").Parse(`
<html>
<head>
    <title>Https Test</title>
    <script src="assets/app.js"></script>
</head>
<body>
    <h1 style="color:red;" >Welcome, Ginner!</h1>
</body>
</html>
`))
