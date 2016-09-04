package main

import (
    "fmt"
)

func init() {
    subcommands.Register(new(HelloComponent), "hello")
}

type HelloComponent struct {}

func (p *HelloComponent) Run(args []string) {
    fmt.Println("Hello world!")
}
