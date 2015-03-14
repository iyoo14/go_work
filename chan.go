package main

import "fmt"
import "runtime"

const msize int = 100

func sum(a []int, c chan int, v chan int) {
    total := 0
    for _, v := range a {
        for i := 0; i < msize; i++ {
            total += v
        }
    }
    c <- total
    v <- 1
}

func main() {

    cpunum := runtime.NumCPU()
    fmt.Printf("cpunum : %v\n", cpunum)
    runtime.GOMAXPROCS(runtime.NumCPU())
    var a [msize]int
    l := len(a)
    for i := 0; i < l; i++ {
        a[i] = 1;
    }
    c := make(chan int, cpunum)
    v := make(chan int)
    ofs := len(a)/cpunum

    s, e := 0, 0
    for i := 0; i < cpunum; i++ {
        s = ofs * i
        e = s + ofs
        //fmt.Println(s, e)
        go sum(a[s:e], c, v)
    }
    vcnt := 0
    total := 0
    //for vcnt != cpunum {
    for i := 0; i < cpunum; i++ {
        fmt.Println(vcnt)
        total += <-c
        vcnt += <-v
    }

    fmt.Println(total)
}
