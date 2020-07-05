package main

import (
    "fmt"
    "regexp"
)

func main() {

    re := regexp.MustCompile(`(\d)abc1`)
    fmt.Println(re.MatchString("1abc1"))

}
