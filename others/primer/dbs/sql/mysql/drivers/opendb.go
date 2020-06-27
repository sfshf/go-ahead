package main

import (
    "database/sql"
    "fmt"
    "os"

    "github.com/go-sql-driver/mysql"
)

func Exit1IfHasError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        os.Exit(1)
    }
}

func main() {

    /*
        A word on sql.Open
    */

    driverName := "mysql"  //一般由配置文件给出值
    dataSourceName := "root:000000@/crashcourse?charset=utf8"  //一般由配置文件给出值

    config, err := mysql.ParseDSN(dataSourceName)
    Exit1IfHasError(err)
    db, err := sql.Open(driverName, config.FormatDSN())
    Exit1IfHasError(err)
    defer db.Close()

    // Open doesn't open a connection. Validate DSN data:
    err = db.Ping()
    Exit1IfHasError(err)

    fmt.Fprintln(os.Stdout, "Opening database correctly.")

    fmt.Fprintln(os.Stdout)

}
