package model

import (
    "github.com/iyoo14/go_work/lib"
    "github.com/iyoo14/go_work/common"
    )

type mo struct {
    id int
    name string
    Gname string
}

func NewModel()(mo) {
    return mo{1, "stanza", "Mr.stanza"}
}

func (m mo)Do() {
    var mml lib.Ml = lib.NewMl()
    common.Err(mml)
}
