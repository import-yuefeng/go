package main

import (
	"fmt"

)


func f() (res int){
	fmt.Println(res)
	defer func(){
		fmt.Println(res)
		res++

	}()
	return 1


}

func main(){

	fmt.Println(f())
}