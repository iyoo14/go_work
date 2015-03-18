package lib

import (
    "fmt"
)

type sng struct{
    M  *Ml
    name string
}

var instance *sng

func GetInstance() (*sng){
    if instance == nil {
        fmt.Println("nil")
        m := new(Ml)
        m.id = 5
        instance = &sng{m, "rat"}
    }
    instance.M.Disp()
    return instance
}
