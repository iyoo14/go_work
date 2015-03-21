package lib

import (
    "fmt"
)

type Pony struct {
    name string
}

func NewPony() *Pony {
    p := new(Pony)
    p.name = "pony"
    return p
}

func (p Pony)Say() {
    fmt.Println(p.name)
}
