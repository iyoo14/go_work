package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var val string
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        val += fmt.Sprintf("%s\n", scanner.Text())
    }
    fmt.Printf("%s", val)
}
