package main

import (
    "github.com/iyoo14/go_work/lib"
)

func main() {
    var myif lib.Myif
    myif = lib.NewPony()
    myif.Say()
    myif = lib.NewBony()
    myif.Say()
}
