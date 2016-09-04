//go:generate go-extpoints
package main

import (
    "fmt"
    "os"
    "github.com/mingderwang/go-ext/extpoints"
)

var subcommands = extpoints.Subcommands

func usage() {
    fmt.Println("Available commands:\n")
    for name, _ := range subcommands.All() {
        fmt.Println(" - ", name)
    }
    os.Exit(2)
}

func main() {
    if len(os.Args) < 2 {
        usage()
    }
    cmd := subcommands.Lookup(os.Args[1])
    if cmd == nil {
        usage()
    }
    cmd.Run(os.Args[2:])
}
