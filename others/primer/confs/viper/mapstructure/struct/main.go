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
    "name": "Mitchell",
    "age": 91,
    "emails": []string{"one", "two", "three"},
    "extra": map[string]string{
        "twitter": "mitchellh",
    },
}

func main() {

    var result Person

    err := mapstructure.Decode(input, &result)
    if err != nil { panic(err) }

    fmt.Printf("%#v\n", result)

}
