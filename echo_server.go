package main

import (
    "log"
    //"math/rand"
    "net"
    "runtime"
    //"time"
    "strings"
)

func main() {
    cpunum := runtime.NumCPU()
    log.Printf("cpu : %v\n", cpunum)
    runtime.GOMAXPROCS(cpunum)

    service := ":12345"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service);
    listener, err := net.ListenTCP("tcp", tcpAddr);

    if err != nil {
        log.Fatalln(err)
    }

    for {
        conn, err := listener.AcceptTCP()

        if err != nil {
            log.Printf("Accept Error:%v\n", err)
            continue
        }

        log.Printf("Accept[%v]\n", conn.RemoteAddr())

        go doProcess(conn)
    }
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

