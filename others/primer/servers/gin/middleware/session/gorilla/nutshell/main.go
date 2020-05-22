package main

import (
    "net/http"
    "os"

    "github.com/gorilla/sessions"
)

/*
    Note: Don't store your key in your source code.
    Pass it via an environmental variable, or flag (or both),
    and don't accidentally commit it alongside your code.
    Ensure your key is sufficiently random - i.e. use Go's crypto/rand
    or securecookie.GenerateRandomKey(32) and persist the result.

    首先，我们调用`NewCookieStore()`初始化一个的`session存储`，`并传入一个用于认证session的密钥`。
    在`处理器`内部，我们调用`store.Get()`来检索现有session或创建一个新session。
    然后，我们在`session.Values`（这是一个`map[interface{}]interface{}`类型）中设置一些`session值`。
    最后，我们调用`session.Save()`将session保存到`响应`中。
*/
var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func MyHandler(w http.ResponseWriter, r *http.Request) {
    /*
        Get a session. We're ignoring the error resulted from decoding an existing session: Get() always returns a session, even if empty.
    */
    session, _ := store.Get(r, "session-name")
    // Set some session values.
    session.Values["foo"] = "bar"
    session.Values[42] = 43
    // Save it before we write to the response/return from the handler.
    err := session.Save(r, w)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func main() {

    http.HandleFunc("/session", MyHandler)
    http.ListenAndServe(":8080", nil)

}
