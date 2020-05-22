package main

import (
    "fmt"
    "os"
    "time"

    _ "github.com/go-sql-driver/mysql"
    "xorm.io/xorm"
)

type User struct {
    Id int
    Name string
    Addr string
    Age int
    Birth string
    Sex int
    UpdateAt time.Time
    CreateAt time.Time
}

func main() {

    driverName := "mysql"
    dataSourceName := "root:000000@/test_new_db?charset=utf8"
    engine, err := xorm.NewEngine(driverName, dataSourceName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        return
    }
    defer engine.Close()

    //
    var user User
    // Query one record or one variable from datebase.
    has, err := engine.Get(&user)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        return
    }
    fmt.Fprintf(os.Stdout, "SELECT * FROM user LIMIT 1; \n-- %t --%v\n", has, user)

    // 查询单列的值
    var id int64
    has, err = engine.Table("user").Where("name = ?", "王小海").Cols("id").Get(&id)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        return
    }
    fmt.Fprintf(os.Stdout, "SELECT id FROM user WHERE name = ? LIMIT 1; \n-- %t --%v\n", has, id)

    var users1 []User
    err = engine.Find(&users1)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        return
    }
    fmt.Fprintf(os.Stdout, "SELECT * FROM user;\n-- %v\n", users1)

    // users2 := make(map[int64]User)
    // err = engine.Find(&users2)
    // if err != nil {
    //     fmt.Fprintf(os.Stderr, "%s\n", err.Error())
    //     return
    // }
    // fmt.Fprintf(os.Stdout, "SELECT * FROM user;\n%v\n", users2)

    var userids []int64
    err = engine.Table("user").Cols("id").Find(&userids)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        return
    }
    fmt.Fprintf(os.Stdout, "SELECT id FROM user;\n-- %v\n", userids)

}
