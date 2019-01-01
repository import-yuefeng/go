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
}