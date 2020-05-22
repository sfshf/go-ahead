package main

import (
    "fmt"
    "net/http"
    "log"
    
    "github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}


// We need an object that implements the http.Handler interface.
// Therefore we need a type for which we implement the ServeHTTP method.
// We just use a map here, in which we map host names (with port) to http.Handlers.
type HostSwitch map[string]http.Handler

// Implement the ServeHTTP method on our new type.
func (hs HostSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // Check if a http.Handler is registered for the given host.
    // If yes, use it to handle the request.
    if handler := hs[r.Host]; handler != nil {
        handler.ServeHTTP(w, r)
    } else {
        // Handle host names for which no handler is registered.
        http.Error(w, "Forbidden", 403)  // Or Redirect?
    }
}

// 如果在Linux系统上测试代码时，需要在`/etc/hosts`文件中添加`127.0.0.1 example.com`。
// 然后运行本服务器程序，再在浏览器访问`http://example.com:12345/`地址。
func main() {

    // Initialize a router as usual.
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/hello/:name", Hello)

    // Make a new HostSwitch and insert the router (our http handler)
    // for example.com and port 12345.
    hs := make(HostSwitch)
    hs["example.com:12345"] = router

    // Use the HostSwitch to listen and serve on port 12345
    log.Fatal(http.ListenAndServe(":12345", hs))

}
