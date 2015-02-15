package main

import (
   //"bufio"
    "fmt"
    //"io/ioutil"
    "net"
    "os"
    //"log"
    //"flag"
)

const (
    cnum = 2
)

type Node struct {
    msg chan string
    conn *net.TCPConn
    rcv chan bool
}

func disp(n Node) {
   val := <- n.msg
   conn := n.conn
   _, err := conn.Write([]byte(val))
   checkError(err)

   buf := make([]byte, 1024)
   //rlen := 0
   _, err = conn.Read(buf)
   //result, err := ioutil.ReadAll(conn)
   checkError(err)
   //fmt.Println(string(buf[:rlen]))
   fmt.Printf("> %s", string(buf))

   n.rcv <- true
}

func main() {
    var nodes [cnum]Node
    mport := 55550
    for i, _ := range nodes {
        service := fmt.Sprintf("%s:%d", "localhost", mport+i)
        tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
        checkError(err)
        conn, err := net.DialTCP("tcp", nil, tcpAddr)
        checkError(err)
        nodes[i].conn = conn
        nodes[i].msg = make(chan string) 
        nodes[i].rcv = make(chan bool)
    }
    for _, n := range nodes {
        go disp(n)
    }
    for i, n := range nodes {
       str := fmt.Sprintf("%s:%d", "hello! localhost", mport+i)
       n.msg <- str
       <- n.rcv 
    }
    os.Exit(0)
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
        os.Exit(1)
    }
}
