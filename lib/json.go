package lib

import (
    "fmt"
    "github.com/bitly/go-simplejson"
    "io/ioutil"
    "log"
)

type Jsong struct {
    Js *simplejson.Json
}

func NewJsong()(Jsong) {
    rf, err := ioutil.ReadFile("./conf/db.json")
    if err != nil {
        log.Fatal(err)
    }

    js, err := simplejson.NewJson(rf)
    if err != nil {
        log.Fatal(err)
    }
    return Jsong{js}
}

func (j Jsong)Get(key string) {
    b, err := j.Js.Get(key).String()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(b)
}

