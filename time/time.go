package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println(time.Now())
	t := time.Now()
	fmt.Println(t.Day())
	fmt.Println(t.Format("02 Jan 2006 15:04"))

}