package main

import (
    "fmt"
    "os"

    "xorm.io/xorm"
    _ "github.com/go-sql-driver/mysql"
)

func main() {

    // Create Engine
    driverName := "mysql"
    dataSourceName := "root:000000@/crashcourse?charset=utf8"
    engine, err := xorm.NewEngine(driverName, dataSourceName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        return
    }
    defer engine.Close()

    results, err := engine.Query("SELECT prod_name FROM products;")
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        return
    }

    fmt.Fprintf(os.Stdout, "%T\n", results)
    fmt.Fprintf(os.Stdout, "%v\n", results)

}
