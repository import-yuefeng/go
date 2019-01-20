package main

import (
	"fmt"
)

func main(){

	for a := 0; a < 5; a++{
		fmt.Println(a)
	}

	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}

	// for i := 0; i < 3; {
    // 	fmt.Println("Value of i:", i)
	// }
	for i, j, s := 0, 5, "a"; i < 3 && j < 10 && s != "aaaaa"; i, j,
		s = i+1, j+1, s + "a" {
			fmt.Println("Value of i, j, s:", i, j, s)
	}

	

}







