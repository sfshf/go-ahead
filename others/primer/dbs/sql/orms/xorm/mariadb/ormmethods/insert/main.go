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


    user1 := User {
        Name: "王小琴",
        Addr: "上海",
        Age: 21,
        Birth: "2010-04-05",
        UpdateAt: time.Now(),
        CreateAt: time.Now(),
    }

    // 1. Insert one or multiple records to database;
    affected, err := engine.Insert(&user1)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        return
    }
    fmt.Fprintf(os.Stdout, "INSERT INTO struct () VALUES (); \n-- %d\n", affected)


    user2 := *&user1
    user2.Name = "王小名"
    user2.Addr = "西安"
    user2.Age = 14
    user2.Birth = "2001-09-21"
    user2.Sex = 1

    user3 := *&user1
    user3.Name = "王小海"
    user3.Addr = "北京"
    user3.Age = 13
    user3.Birth = "2009-09-21"
    user3.Sex = 1

    affected, err = engine.Insert(&user2, &user3)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        return
    }
    fmt.Fprintf(os.Stdout, "INSERT INTO struct () VALUES (); \nINSERT INTO struct () VALUES (); \n-- %d\n", affected)


    user4 := *&user1
    user4.Name = "王小米"
    user4.Addr = "天津"
    user4.Age = 21
    user4.Birth = "1995-10-23"

    user5 := *&user1
    user5.Name = "王小倩"
    user5.Addr = "广州"
    user5.Age = 27
    user5.Birth = "1994-03-20"

    users := []User{ user4, user5}
    affected, err = engine.Insert(&users)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        return
    }
    fmt.Fprintf(os.Stdout, "INSERT INTO struct () VALUES (), (), (); \n-- %d\n", affected)

}
