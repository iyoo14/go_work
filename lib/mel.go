package lib

import (
    "fmt"
    )

type Ml struct {
   id int
}

func NewMl()(*Ml){
    return new(Ml) 
}

func (m *Ml) Back() {
    m.id--
    fmt.Println(m.id)
}

func (m *Ml) Disp() {
    fmt.Printf("now id is %d\n", m.id)
}

