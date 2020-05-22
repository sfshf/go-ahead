package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/contrib/sessions"
)

func main() {

    r := gin.Default()

    // 设置Session的存储方式
    store := sessions.NewCookieStore([]byte("secret"))
    // 设置Cookie的选项参数
    opts := sessions.Options{
        MaxAge: 0,
    }
    store.Options(opts)

    // 创建Session中间件，并使用
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/incr", func(c *gin.Context) {

        // 获取请求中的session
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count += 1
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.Run(":8080")

}
