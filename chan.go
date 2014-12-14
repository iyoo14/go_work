package main

import "fmt"
import "runtime"

const msize int = 100000

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

    cpu := runtime.NumCPU()
    fmt.Printf("cpu : %v\n", cpu)
    runtime.GOMAXPROCS(runtime.NumCPU())
    var a [msize]int
    l := len(a)
    for i := 0; i < l; i++ {
        a[i] = 1;
    }
    c := make(chan int, 5)
    v := make(chan int)
    ofs := len(a)/cpu

    s, e := 0, 0
    for i := 0; i < cpu; i++ {
        s = ofs * i
        e = s + ofs
        fmt.Println(s, e)
        go sum(a[s:e], c, v)
    }
    vcnt := 0
    total := 0
    for vcnt != cpu {
        total += <-c
        vcnt += <-v
    }

    fmt.Println(total)
}
