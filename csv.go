package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "os"
)

func main() {
    var fp *os.File
    if len(os.Args) < 2 {
        fp = os.Stdin
    } else {
        var err error
        fp, err = os.Open(os.Args[1])
        if err != nil {
            panic(err)
        }
        defer fp.Close()
    }

    reader := csv.NewReader(fp)
    reader.Comma = '\t'
    reader.LazyQuotes = true
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            panic(err)
        }
        for _, v := range record {
            fmt.Printf("%v\n", v)
        }
    }
}
