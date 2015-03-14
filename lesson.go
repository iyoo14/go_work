package main

import (
    "fmt"
)

func main() {
    var i *int
    n := 3

    i = &n
    fmt.Println(*i)
    v := []interface{}{"a", "b", "c"}
    fmt.Printf("%s - %s -%s\n",v...) 
}

