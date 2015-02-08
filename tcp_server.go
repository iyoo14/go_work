package main

import (
    "log"
    "math/rand"
    "net"
    "runtime"
    "time"
    "strings"
    "flag"
    "fmt"
    "os"
)

const (
    num_child = 10 
)

func main() {
    var p = flag.Int("p", 12345, "port number.")
    flag.Parse()

    // argc
    if flag.NArg() > 0 {
        fmt.Fprintf(os.Stderr,"error:illegale args.\n")
        os.Exit(1)
    }
    port := *p 

    cpunum := runtime.NumCPU()
    log.Printf("cpu : %v\n", cpunum)
    log.Printf("port : %d\n", port)
    runtime.GOMAXPROCS(cpunum)

    service := fmt.Sprintf(":%d", port)
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service);
    listener, err := net.ListenTCP("tcp", tcpAddr);
    checkError(err)

    for i := 0; i < num_child; i++ {
        go accept(listener)
    }
    for {
        time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
    }
}

func accept(listener *net.TCPListener) {
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("Accept Error:%v\n", err)
            return
        }
        go doProcess(conn)
    }
    return
}

func doProcess(conn net.Conn) {
    var rlen int
    var err error

    tcpConn := conn.(*net.TCPConn)
    //tcpConn := conn

    defer tcpConn.Close()

/*
    err = tcpConn.SetDeadline(time.Now().Add(5*time.Second))

    if err != nil {
        log.Printf("[%v]: %v\n", tcpConn.RemoteAddr(), err)
        return
    }
*/
    //time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

    for {
        buf := make([]byte, 1024)

        rlen, err = tcpConn.Read(buf)

        if err != nil {
            log.Printf("Receive Error[%v]: %v\n", tcpConn.RemoteAddr(), err)
            return
        }

        s := string(buf[:rlen])

        log.Printf("Receive[%v]: %v\n", tcpConn.RemoteAddr(), s)

        s = strings.TrimRight(s, "\n")
        s = s + ":OK\r\n"
        rlen, err = tcpConn.Write([]byte(s))

        if err != nil {
            log.Printf("Send Error[%v]\n", tcpConn.RemoteAddr(), err)
            return
        }
    }

}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}

