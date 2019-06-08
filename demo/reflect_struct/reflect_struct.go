package main

import (
    "fmt"
    "reflect"
)

type Student struct {
    Name string
    Sex int
    Age int
    Score float32
    //xxx int
}
func main() {
    //testStruct()
    testStructPoint()
}

func testStructPoint() {
    var s Student
    v := reflect.ValueOf(&s)
    // *v
    v.Elem().Field(0).SetString("stu01")
    v.Elem().FieldByName("Sex").SetInt(2)
    v.Elem().FieldByName("Age").SetInt(18)
    v.Elem().FieldByName("Score").SetFloat(99.5)

    fmt.Printf("s: %#v\n", s)
}

func testStruct() {
    var s Student
    v := reflect.ValueOf(s)
    t := v.Type()
    //t := reflect.TypeOf(s)

    kind := t.Kind()
    switch kind {
    case reflect.Int64:
        fmt.Printf("s is int64\n")
    case reflect.Float32:
        fmt.Printf("s is float32\n")
    case reflect.Struct:
        fmt.Printf("s is struct\n")

        fmt.Printf("field num of s is %d\n", v.NumField())
        for i:=0; i<v.NumField();i++ {
            field := v.Field(i)
            //fmt.Printf("name:%s type:%v value:%v\n",
            //    t.Field(i).Name, field.Type(), field.Interface())
            fmt.Printf("name:%s type:%v value:%v\n",
                t.Field(i).Name, field.Type().Kind(), field.Interface())

        }
    default:
        fmt.Printf("default\n")
    }
}
