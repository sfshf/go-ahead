package main

import (
    "fmt"

    "github.com/mitchellh/mapstructure"
)

type Person struct {
    Name string
    Age int
    Emails []string
    Extra map[string]string
}

/*
    This input can come from anywhere,
    but typically comes from something like decoding JSON where we're not quite sure of the struct initially.
*/
var input = map[string]interface{}{
    "name": 123,
    "age": "bad value",
    "emails": []int{1, 2, 3},
}

func main() {

    var result Person

    err := mapstructure.Decode(input, &result)
    if err == nil { panic("should have an error") }

    fmt.Println(err.Error())

}
