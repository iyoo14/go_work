package lib

import (
    "fmt"
)

type Bony struct {
    name string
}

func NewBony() *Bony {
    b := new(Bony)
    b.name = "bony"
    return b
}

func (b Bony)Say() {
    fmt.Println(b.name)
}
