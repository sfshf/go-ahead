package main

import (
    "fmt"
    "os"

    "github.com/pelletier/go-toml"
    "github.com/pelletier/go-toml/query"
)

func main() {

    config, err := toml.Load(`
        [Postgres]
        User = "pelletier"
        Password = "mypassword"`)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        os.Exit(1)
    }

    // Use a query to gather elements without walking the tree.
    q, err := query.Compile("$..[User, Password]")
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        os.Exit(1)
    }
    results := q.Execute(config)
    for li, item := range results.Values() {
        fmt.Fprintf(os.Stdout, "Query result %d: %v\n", li, item)
    }

}
