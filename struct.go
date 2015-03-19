package main

import (
    "fmt"
    )

type parent struct {
    pid string
    id int
}

type child struct {
    *parent
    cid string
}

func (p parent) common() {
    fmt.Println(p.pid)
}

func (c *child) doSomething() {
    c.common()
}

func (c *child)disp() {
    fmt.Println(c.id)
}

func (c *child)cal() {
    c.id++
    fmt.Println(c.id)
}

func main() {
    //ob := new(child)
    p := new(parent)
    p.id = 1;
    p.pid = "parent"
    //ob.parent = p
    //ob. cid = "child"
    ob := &child{p, "child"}
    ob.doSomething()
    ob.disp()
    ob.cal()
    ob.disp()
    ob.cal()
    ob.disp()
    fmt.Println(ob.cid)
}

