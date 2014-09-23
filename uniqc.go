package main

import (
    "bufio"
    "fmt"
    "os"
    )

func main() {
    var fp *os.File
    var err error

    if len(os.Args) < 2 {
        fp = os.Stdin
    } else {
        fp, err = os.Open(os.Args[1])
        if err != nil {
            panic(err)
        }
        defer fp.Close()
    }

    var val string
    var old_val string
    var cnt int = 0
    scanner := bufio.NewScanner(fp)
    for scanner.Scan() {
        val = scanner.Text()
        old_val = val
        cnt++
        break
    }
    for scanner.Scan() {
        val = scanner.Text()
        if val != old_val {
            fmt.Printf("%s\t%d\n", old_val,cnt)
            cnt = 0
        }
        old_val = val
        cnt++
    }
    fmt.Printf("%s\t%d\n", old_val,cnt)

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}

