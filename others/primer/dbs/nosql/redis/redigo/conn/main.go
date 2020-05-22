package main

import (
    "github.com/gomodule/redigo/redis"
    "fmt"
    "os"
    "encoding/json"
)

func checkerr(what string, err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s: %v\n", what, err)
        os.Exit(1)
    }
}

type Person struct {
    Name string
    Age int
}

func main() {
    conn, err := redis.Dial("tcp", ":6379")
    checkerr("redis.Dial()", err)
    defer conn.Close()

    reply, err := conn.Do("APPEND", "name", "张三")
    checkerr("redis.Conn.Do()", err)
    fmt.Printf("%T\t%[1]v\n", reply)
    reply, err = conn.Do("DEL", "name")
    checkerr("redis.Conn.Do()", err)
    fmt.Printf("%T\t%[1]v\n", reply)

    reply, err = conn.Do("APPEND", 1, Person{"张三", 23})
    checkerr("redis.Conn.Do()", err)
    fmt.Printf("%T\t%[1]v\n", reply)
    reply, err = conn.Do("GET", 1)
    checkerr("redis.Conn.Do()", err)
    fmt.Printf("%T\t%[1]v\n", reply)
    reply, err = conn.Do("DEL", 1)
    checkerr("redis.Conn.Do()", err)
    fmt.Printf("%T\t%[1]v\n", reply)

    p := Person{"李四", 21}
    bytes, err := json.Marshal(&p)
    checkerr("json.Marshal()", err)
    reply, err = conn.Do("APPEND", p.Name, bytes)
    checkerr("redis.Conn.Do()", err)
    fmt.Printf("%T\t%[1]v\n", reply)

    reply, err = conn.Do("GET", p.Name)
    checkerr("redis.Conn.Do()", err)
    fmt.Printf("%T\t%[1]v\n", reply)
    p2 := Person{}
    if bytes, ok := reply.([]byte); ok {
        err = json.Unmarshal(bytes, &p2)
        checkerr("json.Unmarshal()", err)
    }
    fmt.Println(p2)
    fmt.Println(p==p2)

    reply, err = conn.Do("DEL", p.Name)
    checkerr("redis.Conn.Do()", err)
    fmt.Printf("%T\t%[1]v\n", reply)

}
