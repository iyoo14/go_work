package main

import (
    "fmt"
)

func main() {
    var i *int
    n := 3

    i = &n
    fmt.Println(*i)
}

