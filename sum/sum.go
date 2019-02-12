package main

import (
    "fmt"
)


func sum(a int, b int) int{
    c := a + b
    fmt.Println(c)
    return c
}


func main(){
    sum(1, 2)
    a_test()
}

func a_test() {
    a := make([]int, 10)
    b := a[:]
    fmt.Printf("%T %T\n", a, b)
    fmt.Println(a)

}