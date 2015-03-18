package main

import (
    "fmt"
    "github.com/iyoo14/go_work/model"
    )

func main() {
    m := model.NewModel()
    m.Do()
    // export ok
    fmt.Println(m.Gname)
    // export ng
    //fmt.Println(m.name)
}
