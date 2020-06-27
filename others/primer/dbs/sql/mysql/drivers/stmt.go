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

    config, err := mysql.ParseDSN("root:000000@/crashcourse?charset=utf8")
    Exit1IfHasError(err)
    db, err := sql.Open("mysql", config.FormatDSN())
    Exit1IfHasError(err)
    defer db.Close()

    result, err := db.Exec(`CREATE TABLE IF NOT EXISTS squares(
                                number INT(11) NOT NULL,
                                squareNumber INT(11) NOT NULL,
                                PRIMARY KEY(number)
                            ) ENGINE=InnoDB;`)
    Exit1IfHasError(err)
    n, err := result.RowsAffected()
    Exit1IfHasError(err)
    fmt.Fprintf(os.Stdout, "RowsAffected: %d\n", n)

    /*
        Prepared Statements
    */

    // Prepare statement for inserting data.
    stmtIns, err := db.Prepare("INSERT INTO squares VALUES(?, ?);")
    Exit1IfHasError(err)
    defer stmtIns.Close()

    // Prepare statement for reading data.
    stmtOut, err := db.Prepare("SELECT squareNumber FROM squares WHERE number = ?;")
    Exit1IfHasError(err)
    defer stmtOut.Close()

    // Insert square numbers for 0-24 in the database.
    for i := 0; i < 25; i ++ {
        result, err :=  stmtIns.Exec(i, (i * i))
        Exit1IfHasError(err)
        n, err := result.RowsAffected()
        Exit1IfHasError(err)
        fmt.Fprintf(os.Stdout, "RowsAffected: %d\n", n)
    }

    var squareNum int  // We "scan" the result in here.

    // Query the square-number of 13
    err = stmtOut.QueryRow(13).Scan(&squareNum)
    Exit1IfHasError(err)
    fmt.Fprintf(os.Stdout, "The square number of 13 is: %d\n", squareNum)

    // Query another number.. 1 maybe?
    err = stmtOut.QueryRow(1).Scan(&squareNum)
    Exit1IfHasError(err)
    fmt.Fprintf(os.Stdout, "The square number of 1 is: %d\n", squareNum)

    // 删除表
    result, err = db.Exec(`DROP TABLE IF EXISTS squares;`)
    Exit1IfHasError(err)
    n, err = result.RowsAffected()
    Exit1IfHasError(err)
    fmt.Fprintf(os.Stdout, "Dropping the table successfully.")
    fmt.Fprintf(os.Stdout, "RowsAffected: %d\n", n)

}
