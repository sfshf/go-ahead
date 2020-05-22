package main

import (
    "fmt"
    "os"

    "github.com/pelletier/go-toml"
)

func main() {

    config, err := toml.Load(`
        [postgres]
        user = "pelletier"
        password = "mypassword"`)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        os.Exit(1)
    }
    // retrieve data directly
    user := config.Get("postgres.user").(string)

    // or using an intermediate object
    postgresConfig := config.Get("postgres").(*toml.Tree)
    password := postgresConfig.Get("password").(string)

    fmt.Fprintf(os.Stdout, "User: %s\n", user)
    fmt.Fprintf(os.Stdout, "Password: %s\n", password)

}
