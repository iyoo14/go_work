package main

import (
    "fmt"
    "reflect"
)

type Hoge struct {
    N struct{ s string; v bool}
    M int
}

func main() {
    h := &Hoge{struct{s string; v bool}{"aaa", true},10}
    v := reflect.ValueOf(*h)
    t := v.Type()
    fmt.Println("Name of h: v=reflect.ValueOf(h), t=v.Type(), t.Name()")
    fmt.Printf("%s\n", t.Name())

    fmt.Println("Num of Field: v=reflect.ValueOf(h), v.NumField()")
    fmt.Println(v.NumField())
    //name := v.FieldByIndex([]int{0,1})
    fmt.Println("Field Name: v=reflect.ValueOf(h), t=v.Type(), t.Field(0).Name")
    name := t.Field(0).Name
    fmt.Println(name)
    fmt.Println("Value : v=reflect.ValueOf(h), v.FieldByName(name).Interface()")
    val := v.FieldByName(name).Interface()
    setval(&val)
    fmt.Printf("%d\n", val)
}

func setval(val *interface{}) {
    *val = 200
}
