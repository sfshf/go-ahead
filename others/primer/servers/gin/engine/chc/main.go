package main

import (
    "github.com/gin-gonic/gin"
)


// Custom HTTP configuration
func main() {

    main1()
    main2()

}

// Use http.ListenAndServe() directly, like this:
func main1() {

    router := gin.Default()
    http.ListenAndServe(":8080", router)

}

// or
func main2() {

    router := gin.Default()
    s := &http.Server{
        Addr: ":8080",
        Handler: router,
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServe()

}
