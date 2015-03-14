package main

import (
    "fmt"
)

type Hoge struct {
    //N struct{ s string; v bool}
   N  struct { s string}
   M int
}

func main() {
    h := Hoge{struct{s string}{"aaa"}, 10}
    fmt.Printf("%v\n", h)
}

