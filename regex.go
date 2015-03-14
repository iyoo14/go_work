package main

import (
    "fmt"
    "regexp"
    )

func main() {
    fmt.Println("---MatchString---")
    v := regexp.MustCompile("^[a-zA-Z0-9]+$")
    fmt.Println(v.MatchString("0123"))
    fmt.Println(v.MatchString("-we"))
    fmt.Println("---ReplaceAllString---")
    fmt.Println(regexp.MustCompile("hello").ReplaceAllString("hello world", "HELLO"))
    fmt.Println("---FindString---")
    fmt.Println(regexp.MustCompile("([a-z]+)@gmail.com").FindString("mail: yourname@gmail.com"))
    fmt.Println("---FindStringSubmatch---")
    fmt.Println(regexp.MustCompile("([a-z]+)@(.*?).com").FindStringSubmatch("mail: yourname@gmail.com"))


}
