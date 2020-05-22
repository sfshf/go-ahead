package main

import (
    "fmt"

    "github.com/mitchellh/mapstructure"
)

/*
    Note that the mapstructure tags define in the struct type can indicate which fields the values are mapped to.
*/
type Person struct {
    Name string     `mapstructure:"person_name"`
    Age int         `mapstructure:"person_age"`
}

var input = map[string]interface{}{
    "person_name": "Mitchell",
    "person_age": 91,
}

func main() {

    var result Person

    err := mapstructure.Decode(input, &result)
    if err != nil { panic(err) }

    fmt.Printf("%#v\n", result)

}
