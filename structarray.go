package main

import (
    "fmt"
)

const (
    cnum = 10
)

type Node struct {
    msg chan string
    port int
    rcv chan bool
}

func disp(n *Node) {
     m := <- n.msg
     fmt.Println(m)
     n.port += 10
     n.rcv <- true
}

func main() {

    var nodes []*Node = make([]*Node, cnum)
    mport := 55550
    for i, _ := range nodes {
        p := new(Node)
        p.msg = make(chan string)
        p.port = mport + i
        p.rcv = make(chan bool)
        nodes[i] = p
        //nodes[i].msg = make(chan string) 
        //nodes[i].port = mport + i
        //nodes[i].rcv = make(chan bool)
    }
    fmt.Println("disp")
    for _, n := range nodes {
        go disp(n)
    }

    for _, n := range nodes {
       str := fmt.Sprintf("msg:%d",n.port)
       n.msg <- str
       <- n.rcv 
    }
    for i, _ := range nodes {
        fmt.Println(nodes[i].port)
    }
}
