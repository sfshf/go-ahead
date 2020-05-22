package main

import (
    "net/http"
    "html/template"
    "fmt"
    "time"

    "github.com/gin-gonic/gin"
)

func main() {

    main5()
    // main2()
    // main1()

}

/*

    使用 LoadHTMLGlob() 或 LoadHTMLFiles()

*/
func main1() {

    router := gin.Default()
    router.LoadHTMLGlob("templates/*")
    // router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
    router.GET("/index", func (c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "title": "Main website",
        })
    })
    router.Run()

}

/*

    使用不同目录里的同名模板文件。

*/
func main2() {

    router := gin.Default()
    router.LoadHTMLGlob("templates/**/*")
    router.GET("posts/index", func (c *gin.Context) {
        c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
            "title": "Posts",
        })
    })
    router.GET("/users/index", func (c *gin.Context) {
        c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
            "title": "Users",
        })
    })
    router.Run()

}

/*

    自定义模板渲染器

*/
func main3(){

    router := gin.Default()
    html := template.Must(template.ParseFiles("file1", "file2"))
    router.SetHTMLTemplate(html)
    router.Run()

}

/*

    自定义限定器

*/
func main4() {

    r := gin.Default()
    r.Delims("{[{", "}]")
    r.LoadHTMLGlob("path/to/templates")

}

/*

    自定义模板函数

*/
func main5() {

    router := gin.Default()
    router.Delims("{[{", "}]}")
    router.SetFuncMap(template.FuncMap{
        "formatAsDate": formatAsDate,
    })
    router.LoadHTMLFiles("templates/raw.tmpl")

    router.GET("/raw", func (c *gin.Context) {
        c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{} {
            "now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
        })
    })
    router.Run()

}

func formatAsDate(t time.Time) string {
    year, month, day := t.Date()
    return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}
