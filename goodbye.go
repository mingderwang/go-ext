package main

import (
    "fmt"
)

func init() {
    subcommands.Register(new(GoodbyeComponent), "goodbye")
}

type GoodbyeComponent struct {}

func (p *GoodbyeComponent) Run(args []string) {
    fmt.Println("Goodbye world!")
}
