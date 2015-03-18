package main

import (
    "github.com/iyoo14/go_work/lib"
    "fmt"
    )

func main() {
    fmt.Println("start")
    ins := lib.GetInstance()
    ins.M.Disp()
    ins.M.Back()
    ins.M.Back()
    ins.M.Disp()
    ins = lib.GetInstance()
    ins.M.Disp()
    ins.M.Back()
    ins.M.Disp()
}
