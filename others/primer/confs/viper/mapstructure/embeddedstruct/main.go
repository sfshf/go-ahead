package main

import (
    "fmt"

    "github.com/mitchellh/mapstructure"
)

/*
    Squashing multiple embedded structs is allowed using the squash tag.
    This is demonstrated by creating a composite struct of multiple types and decoding into it.
    In this case, a person can carry with it both a Family and a Location, as well as their own FirstName.
*/

type Family struct {
    LastName string
}

type Location struct {
    City string
}

type Person struct {
    Family      `mapstructure:",squash"`
    Location    `mapstructure:",squash"`
    FirstName string
}

var input = map[string]interface{}{
    "FirstName": "Mitchell",
    "LastName": "Hashimoto",
    "City": "San Francisco",
}

func main() {

    var result Person

    err := mapstructure.Decode(input, &result)
    if err != nil { panic(err) }

    fmt.Printf("%s %s, %s\n", result.FirstName, result.LastName, result.City)

}
