package main

import (
    "fmt"

    "github.com/mitchellh/mapstructure"
)

/*
    Note that the mapstructure tags defined in the struct type can indicate which fields the values are mapped to.
*/
type Person struct {
    Name string
    Age int
    Other map[string]interface{} `mapstructure:",remain"`
}

var input = map[string]interface{}{
    "name": "Mitchellh",
    "age": 91,
    "email": "mitchell@example.com",
}

func main() {

    var result Person

    err := mapstructure.Decode(input, &result)
    if err != nil { panic(err) }

    fmt.Printf("%#v\n", result)

}
