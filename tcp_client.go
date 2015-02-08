package main

import (
    "fmt"
    //"io/ioutil"
    "net"
    "os"
    "log"
    "flag"
    )

func main() {
    var p = flag.Int("p", 12345, "port number.")
    var h = flag.String("h", "localhost", "host name.")
    flag.Parse()

    // argc
    if flag.NArg() > 0 {
        fmt.Fprintf(os.Stderr,"error:illegale args.\n")
        os.Exit(1)
    }
    port := *p 
    host := *h
    log.Printf("host: %s\n", host)
    log.Printf("port: %d\n", port)
    service := fmt.Sprintf("%s:%d", host, port)
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    checkError(err)
    _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
    checkError(err)

    buf := make([]byte, 1024)
    var rlen int
    rlen, err = conn.Read(buf)
    //result, err := ioutil.ReadAll(conn)
    checkError(err)
    //fmt.Println(string(buf[:rlen]))
    fmt.Println(rlen)
    fmt.Println(string(buf))
    os.Exit(0)
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
        os.Exit(1)
    }
}
