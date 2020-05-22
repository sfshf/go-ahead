package main

import (
    "fmt"
    "os"

    "github.com/pelletier/go-toml"
)

func main() {

    type Postgres struct {

        User string
        Password string

    }

    type Config struct {
        Postgres Postgres
    }

    doc := []byte(`
        [Postgres]
        User = "pelletier"
        Password = "mypassword"`)

    config := Config{}
    err := toml.Unmarshal(doc, &config)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        os.Exit(1)
    }
    fmt.Fprintln(os.Stdout, "User:", config.Postgres.User)

}
