package main

import (
    "fmt"
    r "reflect"
    )

type member struct {
    Id int
    Name string
    Flg bool
}
type S struct {
    I int
    Str string
}

type error string

func (e error) String() string { return string(e) }

func main() {
    s := member{1, "Bunnies", true}
    //s := S{5, "Bunnies"}
    m, _ := StructToMap(&s)
    fmt.Println(m)
    //MapToStructN(s)
    //m := make(map[string]interface{})
    s1 := member{}
    MapToStruct(m, &s1)
    fmt.Println(s1)

    id := m["Id"]
    recss := []interface{}{&id, m["Name"], m["Flg"]}
    recs := []interface{}{}
    for k, v := range m {
        fmt.Printf("%v - %v\n", k, v)
        recs = append(recs, &v)
    }
    fmt.Println(recss)

    /*
    m := make(map[string]interface{})
    m["I"] = [1]string{"123"}
    m["S"] = [1]string{"hello"}
    fmt.Println(m)
    s1 := S{}
    fmt.Println(s1)
    MapToStruct(m, &s1)

    fmt.Printf("%+v\n%v\n%+v\n", s, m, s1)
    */
}

func StructToMap(val interface{}) (mapVal map[string]interface{}, ok bool) {
    structVal := r.Indirect(r.ValueOf(val))
    typ := structVal.Type()

    mapVal = make(map[string]interface{})

    for i := 0; i < typ.NumField(); i++ {
        field := structVal.Field(i)

        if field.CanSet() {
            mapVal[typ.Field(i).Name] = field.Interface()
        }
    }
    return
}

func MapToStruct(mapVal map[string]interface{}, val interface{}) (ok bool) {
    structVal := r.Indirect(r.ValueOf(val))
    for name, elem := range mapVal {
        v := r.ValueOf(elem)
        fmt.Println(v)
        fmt.Println(name)
        structVal.FieldByName(name).Set(v)
    }
    return
}

func MapToStructN(h interface{}) {
    v := r.ValueOf(h)
    t := v.Type()
    fmt.Printf("val name: %s\n", t.Name())

    fnum := v.NumField()
    fmt.Printf("Num of Field: %d\n" , fnum)

    for i := 0; i < fnum; i++ {
        name := t.Field(i).Name
        val := v.FieldByName(name).Interface()
        fmt.Printf("name-value: %v - %v\n", name, val)
    }
}
