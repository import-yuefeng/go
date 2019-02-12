package main

import (
	"fmt"
	// "strconv"
	"errors"
)

func main(){
	_, err := my_sqrt(-1)
	if err != nil{
		fmt.Println(err)
	}
	items := [...]int{0, 10, 20, 30, 40}
	for index, _ := range items{
		items[index] *= 2
	}
	fmt.Println(items)
	

}


func my_sqrt(value float64) (result float64, err error){
	if value >= 0{
		fmt.Printf("Input value is %f\n", value)
		// valueString := strconv.FormatFloat(value, 'f', 6, 64)
		// fmt.Println("Input value is " + valueString)
		return value*value, nil
	}
	return -1, errors.New("negative Error")
}