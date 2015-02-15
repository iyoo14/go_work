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

func disp(n Node) {
   m := <- n.msg
   fmt.Println(m)
   n.rcv <- true
}

func main() {

    var nodes [cnum]Node
    mport := 55550
    for i, _ := range nodes {
        nodes[i].msg = make(chan string) 
        nodes[i].port = mport + i
        nodes[i].rcv = make(chan bool)
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
}
