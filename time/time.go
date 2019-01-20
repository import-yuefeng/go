package main

import (
	"fmt"
	"time"
	"runtime"
	"log"
)

func where() {
    _, file, line, _ := runtime.Caller(1)
    log.Printf("%s:%d", file, line)
}


func main(){

	fmt.Println(time.Now())
	where()

	
	t := time.Now()
	fmt.Println(t.Day())
	// fmt.Println(t.Format("02 Jan 2006 15:04"))
	where()

	
}